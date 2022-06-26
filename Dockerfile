FROM golang:alpine as build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go clean --modcache && go mod tidy
COPY . .
RUN go build -o binary

FROM alpine:3.16.0
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app ./
EXPOSE 8080
ENTRYPOINT ["/app/binary"]