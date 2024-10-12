FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app *.go
FROM alpine
WORKDIR /app
COPY --from=builder /app/app ./
EXPOSE 7259
CMD [ "./app" ]
