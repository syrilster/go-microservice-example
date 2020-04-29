export GO111MODULE=on
export GOFLAGS=-mod=vendor
APP=currency-conversion-service
KUBE-PROJECT-ID=kube-go-exp
GCLOUD-REGION-PREFIX=asia.gcr.io

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

.PHONY: \
	clean \
	build \
	test \
	container \
	push \