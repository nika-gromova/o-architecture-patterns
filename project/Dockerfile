FROM golang:1.23 as build

WORKDIR /app

# Copy the Go module files
COPY /go.mod .
COPY /go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o bin/app ./cmd

FROM ubuntu:22.04

COPY --from=build /app/bin/app /app
ADD ./config.yaml /config.yaml
ADD ./pkg/rules_v1/api.swagger.json /swagger.json

CMD ["/app"]