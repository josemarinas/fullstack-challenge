FROM golang:alpine3.14 AS build_base
ENV CGO_ENABLED=1
ENV GO111MODULE=on
WORKDIR /src
# We want to populate the module cache based on the go.{mod,sum} files.
COPY . .
RUN rm go.mod go.sum
RUN go mod init server
RUN go mod tidy
# Build the Go app
RUN go build -o ./bin/server ./main.go
# Start fresh from a smaller image
FROM alpine:3.14
WORKDIR /app/bin
COPY --from=build_base /src/bin/server /app/bin/
COPY --from=build_base /src/conf.toml /app/bin/
RUN chmod +x server
# This container exposes port 8080 to the outside world
EXPOSE 8000
# Run the binary program produced by `go install`
ENTRYPOINT ./server