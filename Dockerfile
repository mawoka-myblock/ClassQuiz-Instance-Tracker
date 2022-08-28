FROM golang:1.19 as builder

WORKDIR /usr/local/go/src/InstanceTracker

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY src ./
RUN go build -o /instance_tracker


FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=builder /instance_tracker /instance_tracker
ENV DATABASE_PATH="/data/data.db"
ENV GIN_MODE=release
RUN mkdir /data
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/instance_tracker"]
