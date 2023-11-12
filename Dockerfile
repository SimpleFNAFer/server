FROM golang:1.21
LABEL authors="nasy"

WORKDIR /server
COPY go.mod ./
COPY go.sum ./

COPY app/ ./app/
COPY cmd/ ./cmd/
COPY template/ ./template/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd

EXPOSE 8080 8082

CMD ["./server"]