# Go Hello World HTTP server

- Install with `go get github.com/apxamccallum/go-hello-world-http/go-hello-world-http`
- Build with `docker build -t go-hello-world-http .`
- Run with `docker run --env-file default.env -p 8282:80 --rm go-hello-world-http`
- Or run with `docker run -e "WEBBODY=Hello World" -p 8282:80 --rm go-hello-world-http`
- Run and curl/open in a browser localhost:8282
- Profit! :)
