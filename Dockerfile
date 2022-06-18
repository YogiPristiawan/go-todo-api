FROM golang:alpine
RUN mkdir /go-todo-api
WORKDIR /go-todo-api
COPY . .
RUN go clean --modcache
RUN go mod tidy
RUN go build -o binary

EXPOSE 8080
ENTRYPOINT ["/go-todo-api/binary"]