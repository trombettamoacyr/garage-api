FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV GARAGE_API_PORT=":8080"
ENV GARAGE_API_DB_HOST="localhost"
ENV GARAGE_API_DB_PORT="5432"
ENV GARAGE_API_DB_NAME="postgres"
ENV GARAGE_API_DB_USER="postgres"
ENV GARAGE_API_DB_PASSWORD="postgres"
ENV GOOGLE_APPLICATION_CREDENTIALS="/app/firestore/token.json"

RUN go build

CMD ["./garage-api"]