FROM golang:1.15.2 as builder
COPY go.mod go.sum /go/src/github.com/jgleon/topsecret/
WORKDIR /go/src/github.com/jgleon/topsecret
RUN go mod download
COPY . /go/src/github.com/jgleon/topsecret
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/topsecret github.com/jgleon/topsecret

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src//github.com/jgleon/topsecret/build/topsecret /usr/bin/topsecret
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/topsecret"]