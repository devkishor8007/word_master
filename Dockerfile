FROM golang:1.21.1

WORKDIR /app

RUN go clean -modcache
RUN go clean -cache


COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o myapp

EXPOSE 3002

CMD ["./myapp"]
