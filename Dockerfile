FROM golang:1.14.0 AS builder

ENV GOPATH=/go

COPY . /go/src/library
WORKDIR /go/src/library

RUN go mod download -x \
    && go build -a -installsuffix cgo main.go

EXPOSE 8080
CMD ["/go/src/library/main"]

#FROM alpine:3.11.3
#RUN adduser -D -h /app app
#WORKDIR /app
#COPY --from=builder /go/src/v2ray-ui/v2ray-ui .
#COPY --from=builder /go/src/v2ray-ui/conf ./conf
#COPY --from=builder /go/src/v2ray-ui/static ./static
#COPY --from=builder /go/src/v2ray-ui/views ./views
#EXPOSE 8080
#CMD ["/app/v2ray-ui"]