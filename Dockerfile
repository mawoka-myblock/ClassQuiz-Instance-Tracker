FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY src/ /app/src/
COPY main.go ./
RUN go build -o /instance_tracker


FROM gcr.io/distroless/base-debian11

WORKDIR /data

COPY --from=builder /instance_tracker /instance_tracker
ENV DATABASE_PATH="/data/data.db"
#ENV GIN_MODE=release
EXPOSE 8080
#ENV PORT=0.0.0.0:8080
ENTRYPOINT ["/instance_tracker"]
