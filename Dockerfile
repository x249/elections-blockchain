FROM golang:latest as builder

LABEL maintainer="Osama Adil <adilosama47@gmail.com>"

ADD https://github.com/golang/dep/releases/download/v0.5.4/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/phr3nzy/elections-blockchain/
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o main .
RUN pwd

FROM alpine:latest  

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/phr3nzy/elections-blockchain .
EXPOSE 8080
CMD ["./main"]