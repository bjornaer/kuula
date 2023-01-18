import kserve
import argparse


from torchserver.model import PyTorchModel

DEFAULT_MODEL_NAME = "model"
DEFAULT_LOCAL_MODEL_DIR = "/tmp/model"
DEFAULT_MODEL_CLASS_NAME = "PyTorchModel"

parser = argparse.ArgumentParser(parents=[kserve.model_server.parser])
parser.add_argument('--model_dir', required=True,
                    help='A URI pointer to the model directory')
parser.add_argument('--model_name', default=DEFAULT_MODEL_NAME,
                    help='The name that the model is served under.')
parser.add_argument('--model_class_name', default=DEFAULT_MODEL_CLASS_NAME,
                    help='The class name for the model.')
args, _ = parser.parse_known_args()

if __name__ == "__main__":
    model = PyTorchModel(args.model_name, args.model_class_name, args.model_dir)
    model.load()
    kserve.ModelServer().start([model])