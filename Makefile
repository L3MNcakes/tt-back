.PHONY: build start test

build:
	docker build -t tt-back .

start:
	docker run -it --rm -p 8080:8080 --name tt-api tt-back

test:
	docker run --rm --name tt-api-test tt-back /bin/bash -c "go test; go test ./routes"
