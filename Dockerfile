# built via:
#  docker build -t matthiaswinzeler/dynatrace-counters:latest .
#  docker push matthiaswinzeler/dynatrace-counters:latest

FROM golang:1.17 as builder

WORKDIR /workspace

COPY go.mod .
COPY go.sum .
COPY main.go .

# Build
RUN go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o service main.go

# Use distroless as minimal base image to package the binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /workspace/service .
USER nonroot:nonroot

ENTRYPOINT ["/service"]
