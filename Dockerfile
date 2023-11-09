FROM golang:latest AS build-env
COPY . /src
WORKDIR /src
RUN ls -l
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o crash-test main.go

# final stage
FROM gcr.io/distroless/base
COPY --from=build-env /src/crash-test /
EXPOSE 8080
CMD ["/crash-test"]
