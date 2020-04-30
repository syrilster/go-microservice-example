export GO111MODULE=on
export GOFLAGS=-mod=vendor
PROJECT_NAME="go-microservice-example"
PKG= "github.com/syrilster/$(PROJECT_NAME)"
PKG_LIST=$(shell go list ${PKG}/... | grep -v /vendor/)
APP=currency-conversion-service
KUBE-PROJECT-ID=kube-go-exp
GCLOUD-REGION-PREFIX=asia.gcr.io

update-vendor:
	go mod tidy
	go mod vendor

clean:
	rm -f ${APP}

lint:
	golint -set_exit_status ${PKG_LIST}

build: clean
	go build -o ${APP}

container: build
	docker build . -t ${GCLOUD-REGION-PREFIX}/${KUBE-PROJECT-ID}/${APP}

push:
	docker push ${GCLOUD-REGION-PREFIX}/${KUBE-PROJECT-ID}/${APP}

test:
	go test -v ./... 2>&1 | tee test-output.txt

test-coverage:
	go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	cat cover.out >> test-output.txt

.PHONY: \
	clean \
	build \
	test \
	test-coverage \
	container \
	push \
	lint \