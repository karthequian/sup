FROM golang:1.7.3 as builder
#FROM golang:1.7.3

WORKDIR /go/src/github.com/karthequian/sup/

RUN go get -d -v github.com/gorilla/mux
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

#COPY --from=0 /go/src/github.com/karthequian/sup/main .
COPY --from=builder /go/src/github.com/karthequian/sup/main .

CMD ["./main"]