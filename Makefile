.PHONY: build clean prepare update docker

GO = CGO_ENABLED=0 GO111MODULE=on go

MICROSERVICES=cmd/web-helloworld

.PHONY: $(MICROSERVICES)

DOCKERS=docker_web_helloworld
.PHONY: $(DOCKERS)

VERSION=$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)
BASE_VERSION=$(shell cat ./BASE_VERSION 2>/dev/null || echo 0.0.0)


GIT_SHA=$(shell git rev-parse HEAD)

build: $(MICROSERVICES)

cmd/web-helloworld:
	$(GO) build -mod=vendor -o $@ ./cmd


clean:
	rm -f $(MICROSERVICES)

docker: $(DOCKERS)

docker_web_helloworld:
	docker build -f ./Dockerfile \
		-t worldmaomao/web-helloworld:$(VERSION) \
		.

docker.push:
	docker push worldmaomao/web-helloworld:$(VERSION)