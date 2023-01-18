from torchserver import PyTorchModel
import torch
import torchvision
import torchvision.transforms as transforms
import os

model_dir = model_dir = os.path.join(os.path.dirname(__file__), "example_model", "model")


def test_model():
    server = PyTorchModel("model", "Net", model_dir)
    server.load()

    transform = transforms.Compose([transforms.ToTensor(),
                                    transforms.Normalize((0.5, 0.5, 0.5), (0.5, 0.5, 0.5))])
    testset = torchvision.datasets.CIFAR10(root='./data', train=False,
                                           download=True, transform=transform)
    testloader = torch.utils.data.DataLoader(testset, batch_size=4,
                                             shuffle=False, num_workers=2)
    dataiter = iter(testloader)
    images, _ = dataiter.next()

    request = {"instances": images[0:1].tolist()}
    response = server.predict(request)
    assert isinstance(response["predictions"][0], list)