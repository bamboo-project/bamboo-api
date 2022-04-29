FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o main .
EXPOSE 8080

ENTRYPOINT ["/app/main"]