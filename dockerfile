FROM golang:1.23-alpine
WORKDIR /
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["/main"]
