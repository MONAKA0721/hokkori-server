REVISION := $(shell git rev-parse --short HEAD)

default:
	go build -o bin/hokkori -ldflags "-X 'main.revision=$(REVISION)'" \
		main.go

build-tag-push: build tag push

# build:
# 	docker build -t mo7ka/hokkori:$(REVISION) .
build:
	docker build -t hokkori:$(REVISION) .

# tag:
# 	docker tag mo7ka/hokkori:$(REVISION) kix.ocir.io/nro0opflyunj/hokkori/hokkori:$(REVISION)
tag:
	docker tag hokkori:$(REVISION) 607142159265.dkr.ecr.ap-northeast-1.amazonaws.com/hokkori:$(REVISION)

# push:
# 	docker push kix.ocir.io/nro0opflyunj/hokkori/hokkori:$(REVISION)
push:
	docker push 607142159265.dkr.ecr.ap-northeast-1.amazonaws.com/hokkori:$(REVISION)

gen:
	go generate ./...
