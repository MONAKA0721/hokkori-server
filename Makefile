REVISION := $(shell git rev-parse --short HEAD)

default:
	go build -o bin/hokkori -ldflags "-X 'main.revision=$(REVISION)'" \
		main.go

build-tag-push: build tag push

build:
	docker build -t mo7ka/hokkori:$(REVISION) .

tag:
	docker tag mo7ka/hokkori:$(REVISION) kix.ocir.io/nro0opflyunj/hokkori/hokkori:$(REVISION)

push:
	docker push kix.ocir.io/nro0opflyunj/hokkori/hokkori:$(REVISION)
