export GO111MODULE=on
APP=currency-conversion

update-vendor:
	go mod tidy
	go mod vendor

clean:
	rm -f ${APP}

build: clean
	go build -o ${APP}

container: build
	docker build -t $(APP):latest .

test:
	go test

.PHONY: \
	clean \
	build \
	test \
	container \