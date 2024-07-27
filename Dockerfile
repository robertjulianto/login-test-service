FROM golang:1.22-alpine3.20

WORKDIR ./app

COPY . .

RUN go build main.go

CMD "./main"