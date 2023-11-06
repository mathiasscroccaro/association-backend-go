test:
	go test ./... -v

build:
	go build cmd/webserver/main.go

run: build
	./main

clean:
	rm main
