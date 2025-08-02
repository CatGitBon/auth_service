FROM golang:1.24.4-alpine

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o auth_service ./cmd

EXPOSE 50051

CMD ["./auth_service"]