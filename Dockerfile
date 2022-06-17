FROM golang:alpine
RUN mkdir /go-todo-api
WORKDIR /go-todo-api
COPY go.mod .
COPY go.sum .
RUN go clean --modcache
RUN go mod tidy

COPY . .
RUN go build -o binary

EXPOSE 8080
ENTRYPOINT ["/go-todo-api/binary"]