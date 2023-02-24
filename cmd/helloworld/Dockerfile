# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /go/bin/app -v ./cmd/helloworld

# final stage
FROM alpine:latest
ENV WEBBODY='Default Environment Varable'
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT ["/app"]
LABEL Name=gohelloworldhttp Version=0.0.1
EXPOSE 80