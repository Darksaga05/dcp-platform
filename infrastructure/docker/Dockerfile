FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o dcp-backend

EXPOSE 8080

CMD ["./dcp-backend"]
