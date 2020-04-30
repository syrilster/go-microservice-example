PROJECT_NAME := "go-microservice-example"
PKG := "github.com/syrilster/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
APP=currency-conversion-service
KUBE-PROJECT-ID=kube-go-exp
GCLOUD-REGION-PREFIX=asia.gcr.io
export GO111MODULE=on
export GOFLAGS=-mod=vendor

update-vendor:
	go mod tidy
	go mod vendor

clean:
	rm -f ${APP}

build: clean
	go build -o ${APP}

container: build
	docker build . -t ${GCLOUD-REGION-PREFIX}/${KUBE-PROJECT-ID}/${APP}

push:
	docker push ${GCLOUD-REGION-PREFIX}/${KUBE-PROJECT-ID}/${APP}

test:
	go test -v ./... 2>&1 | tee test-output.txt

test-coverage:
	@go test -short -coverprofile coverage.txt -covermode=atomic ${PKG_LIST}


.PHONY: \
	clean \
	build \
	test \
	test-coverage \
	container \
	push \