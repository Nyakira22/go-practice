FROM golang:1.21.5

WORKDIR /go/src

COPY go* ./
RUN go mod download

COPY . .