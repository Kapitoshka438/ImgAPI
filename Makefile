build:
	docker build -t imgapi-app .
	docker run --publish 8080:8000 --rm --name imgapi-app-run imgapi-app
	explorer "http://localhost:8080"

test:
	go test -v ./...

clean:
	rm -rf build

.PHONY: build test clean
