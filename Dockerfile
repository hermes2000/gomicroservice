FROM golang:1.18.4 as builder

COPY . /go/src/microservice

RUN cd /go/src/microservice && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/microservice


FROM gcr.io/distroless/static-debian11

ENV PORT 8080
EXPOSE 8080

# Import the user and group files from the builder.

COPY --from=builder /go/bin/microservice /go/bin/microservice

USER nonroot:nonroot

ENTRYPOINT ["/go/bin/microservice"]


