# Go Hello World HTTP server
This is a simple HTTP server written in Go that responds with a message from a environment verriable.

## Installation
To install the server, run the following command:
```
go get github.com/apxamccallum/go-hello-world-http/go-hello-world-http
```

## Build with Docker
To build the server with Docker, run the following command:
```
docker build -t go-hello-world-http ./cmd/helloworld/
```

## Run with Docker
To run the server with Docker, use the following command:
```
docker run --env-file ./cmd/helloworld/default.env -p 8282:80 --rm go-hello-world-http
```
This command starts a container with the `go-hello-world-http` image, maps the container's port 80 to the host's port 8282, and uses an environment file to set `WEBBODY`.
Alternatively, you can run the server with a custom message by setting the `WEBBODY` environment variable:
```
docker run -e "WEBBODY=Hello World" -p 8282:80 --rm go-hello-world-http
```

## Usage
To test the server, open a web browser and navigate to `http://localhost:8282` or use the `curl` command:
```
curl http://localhost:8282
```
## Launch with Docker Compose
Alternatively, you can launch the server with Docker Compose by running the following command:
```
docker-compose up
```
This command starts two instances of the server in containers and exposes them on port 8082 and 8083.
