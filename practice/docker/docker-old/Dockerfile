# single stage
FROM golang:1.23.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /server

EXPOSE 8080

ENTRYPOINT [ "/server" ]
