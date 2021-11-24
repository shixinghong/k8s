
# Image URL to use all building/pushing image targets
IMG ?= hub.fastonetech.com/infra/webhook
TEMP_TAG ?=`date +%Y%m%d-%H%M%S`
# ENVTEST_K8S_VERSION refers to the version of kubebuilder assets to be downloaded by envtest binary.
ENVTEST_K8S_VERSION = 1.22

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build docker-build docker-push deploy

.PHONY: test
test: build
	sudo docker build -t ${IMG}:test .

.PHONY: build-image
build-image: build  docker-build

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o webhook main.go

.PHONY: docker-build
docker-build:
	sudo docker build -t ${IMG} .

.PHONY: docker-push
docker-push:
	sudo docker push ${IMG}

.PHONY: deploy
deploy:
	echo "deploy"
