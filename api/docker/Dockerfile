FROM golang:latest as builder-setup
WORKDIR /workdir
COPY ./src ./
# COPY ./docker/app/config ./config
RUN go mod tidy

FROM builder-setup
# Set Environment Variable
ENV CGO_ENABLED=0
# Build
RUN go build -o app

FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=builder /workdir/app /app

ENTRYPOINT ["/app"]
