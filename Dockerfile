FROM golang:1.8.3-alpine

RUN apk add --no-cache git

WORKDIR /go/src/mini-url
COPY . .

RUN go-wrapper download
RUN go-wrapper install
CMD ["go-wrapper", "run"]