.PHONY: container

IMAGE_NAME=josefuentes/pm-ipsum
IMAGE_TAG=latest

container:
	podman build -t $(IMAGE_NAME):$(IMAGE_TAG) .
