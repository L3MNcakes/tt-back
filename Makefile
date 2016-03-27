.PHONY: build start

build:
	docker build -t tt-back .

start:
	docker run -it --rm -p 8080:8080 --name tt-api tt-back
