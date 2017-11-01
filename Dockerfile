FROM golang:1.7.3 AS builder

WORKDIR /go/src/github.com/karthequian/sup/

RUN go get -d -v github.com/gorilla/mux
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

TAG karthequian/sup:step1


CMD ["./main"]

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/karthequian/sup/main ./test/

CMD ["./test/main"]