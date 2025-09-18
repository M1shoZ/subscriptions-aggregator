FROM golang:1.24.5

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY . .
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go mod tidy

CMD ["air", "-c", ".air.toml"]