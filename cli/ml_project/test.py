"""Start a worker.

Typically one worker is run per 'resource'. A resource is a compute element
such as a CPU core or a GPU.

We register callbacks for some Celery signals so that we can write out and
clean up 'healthcheck' files. These allow Docker Compose and Kubernetes to
search the container for these files to determine readiness and liveness.
"""
import argparse
import datetime
import os
import random
import re
import string

import celery.signals
from loguru import logger

from pipeline_cluster_common.celery import app

from .tasks import ExecuteRunTask

# Queue names must match this regular expression
# Between 8 and 16 lowercase alphanumeric characters
QUEUE_NAME_PATTERN = re.compile("^[a-z0-9]{8,16}$")


def generate_random_queue_name():
    return "".join(random.choices(string.ascii_lowercase + string.digits, k=16))


def hostname(queue):
    """Construct a 'hostname' for this worker.

    We abuse the Celery worker hostname to describe the features this worker
    has, e.g. an accelerator card.
    """
    # TODO detect worker capabilities and construct hostname accordingly
    return f"{queue}-cpu@%h"


def queue_name_type(value):
    if not QUEUE_NAME_PATTERN.match(value):
        logger.error("queue must be between 8 and 16 lowercase alphanumeric characters")
        raise argparse.ArgumentTypeError
    return value


@celery.signals.worker_ready.connect
def on_worker_ready(**_):
    """Write a 'ready' file on receiving the ready signal."""
    with open("worker.ready", "w") as f:
        f.write(str(datetime.datetime.now().timestamp()))


@celery.signals.worker_shutdown.connect
def on_worker_shutdown(**_):
    """Remove 'ready' and 'heartbeat' files on receiving the shutdown signal."""
    for fname in ["worker.ready", "worker.heartbeat"]:
        try:
            os.remove(fname)
        except FileNotFoundError:
            pass


@celery.signals.heartbeat_sent.connect
def on_worker_heartbeat(**_):
    """Write a 'heartbeat' file on receiving the heartbeat signal."""
    with open("worker.heartbeat", "w") as f:
        f.write(str(datetime.datetime.now().timestamp()))


def run():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "queue",
        nargs="?",
        default=None,
        type=queue_name_type,
        help="Queue name to consume tasks from.",
    )
    args = parser.parse_args()
    q = args.queue or generate_random_queue_name()

    name = hostname(q)
    logger.info("Starting worker {name} watching queue {queue}", name=name, queue=q)

    app.register_task(ExecuteRunTask())

    app.worker_main(["worker", f"--queues={q}", f"--hostname={name}"])
