#
#  Input variables
#
#  - DISABLE_DOCKER
#
#  - DOCKERFILE
#
#  - DOCKER_REGISTRY
#  - DOCKER_ORGANISATION
#  - CONTAINER_NAME
#  - IMAGE_VERSION
#
#  - SHELL
#  - BIND_VOLUMES
#  - EXPOSE_PORTS
#  - EXTRA_RUN_OPTS
#

ifndef DISABLE_DOCKER

ifndef IMAGE_VERSION
GIT_TAG:=$(shell git describe --tags 2> /dev/null)
GIT_TAG_STRIPPED:=$(patsubst v%,%,$(GIT_TAG))
ifneq ("$(GIT_TAG_STRIPPED)_","_")
IMAGE_VERSION:=$(GIT_TAG_STRIPPED)
else
IMAGE_VERSION:=latest
endif
endif

ifndef SHELL
SHELL:=/bin/bash
endif

ifndef CONTAINER_NAME
DIR_NAME:=$(notdir $(realpath .))
CONTAINER_NAME:=$(subst -,,$(DIR_NAME))
endif

#
#  build up the IMAGE_NAME, optionally including REGISTRY and ORGANISATION
#
IMAGE_NAME:=$(CONTAINER_NAME):$(IMAGE_VERSION)
ifneq ("$(DOCKER_ORGANISATION)x","x")
IMAGE_NAME:=$(DOCKER_ORGANISATION)/$(IMAGE_NAME)
endif
ifneq ("$(DOCKER_REGISTRY)x","x")
IMAGE_NAME:=$(DOCKER_REGISTRY)/$(IMAGE_NAME)
endif

#
# also build image:latest
#
IMAGE_LATEST:=$(patsubst %:$(IMAGE_VERSION),%:latest,$(IMAGE_NAME))

ifeq ($(DOCKERFILE)x,x)
DOCKERFILE:=$(wildcard Dockerfile)
endif
ifneq ($(DOCKERFILE)x,x)
ifndef EXPOSE_PORTS
EXPOSE_PORTS:=$(shell gawk '/^EXPOSE/ { for(i=2;i<=NF;i++) printf("-p %s:%s ", $$i, $$i) }' Dockerfile)
endif
endif

ifndef BIND_VOLUMES
BIND_VOLUMES:=
endif

.PHONY: image
image:
	rm -f $(TAR_FILE)
	docker build --rm --force-rm -t $(IMAGE_LATEST) -f $(DOCKERFILE) .
	[ "$(IMAGE_VERSION)" == "latest" ] || docker tag $(IMAGE_LATEST) $(IMAGE_NAME)

.PHONY: clean
clean::
	rm -f $(TAR_FILE)
	-docker kill $(CONTAINER_NAME) 2> /dev/null
	-docker rm $(CONTAINER_NAME) 2> /dev/null

.PHONY: clobber
clobber:: clean
	-docker rmi -f $(IMAGE_NAME) $(IMAGE_LATEST) 2> /dev/null

TAR_FILE:=$(notdir $(subst :,_,$(IMAGE_NAME))).tar

.PHONY: docker-tar
docker-tar:
	docker save '$(IMAGE_NAME)' > $(TAR_FILE)
	gzip $(TAR_FILE)

#
# Push both the versioned and the "latest" image. They might be the same thing - but,
# if so, then the second one will complete very quickly.
#
.PHONY: docker-push
docker-push:
	docker push $(IMAGE_NAME)
	docker push $(IMAGE_LATEST)

.PHONY: docker-run
docker-run: clean
	docker run -d $(EXTRA_RUN_OPTS) $(EXPOSE_PORTS) $(BIND_VOLUMES) --name $(CONTAINER_NAME) $(IMAGE_NAME)

.PHONY: docker-exec
docker-exec:
	docker exec -ti $(CONTAINER_NAME) $(SHELL)

.PHONY: docker-entry
docker-entry:
	docker run -ti $(EXTRA_RUN_OPTS) $(EXPOSE_PORTS) $(BIND_VOLUMES) --entrypoint $(SHELL) $(IMAGE_NAME)

# Useful with a scratch container. Use in Dockerfile as:
#
#   ADD ca-bundle.crt /etc/pki/tls/certs/ca-bundle.crt
ca-bundle.crt:
	curl -sL https://mkcert.org/generate/ > $@

.PHONY: docker-clean
docker-clean:
	docker images | grep '<none>' | gawk '{ print $$3 }' | xargs docker rmi

endif
