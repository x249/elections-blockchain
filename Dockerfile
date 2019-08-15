FROM golang:latest as builder

LABEL maintainer="Osama Adil <adilosama47@gmail.com>"

WORKDIR /app

COPY . .

RUN go build -o main .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]