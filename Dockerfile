################
FROM golang:1.16 as builder
RUN go version

WORKDIR /go/src/github.com/noelruault/go-fibonacci/
COPY . .

RUN [ -d bin ] || mkdir bin
RUN GOOS=linux CGO_ENABLED=0 go build -o bin/go-fibonacci ./cmd/

################
FROM alpine
COPY --from=builder /go/src/github.com/noelruault/go-fibonacci/bin/ bin

RUN chmod +x /bin/go-fibonacci

ENTRYPOINT ./bin/go-fibonacci
