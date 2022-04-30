FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN apk add git
RUN go mod tidy
RUN go build -o main .
EXPOSE 8080

ENTRYPOINT ["/app/main"]