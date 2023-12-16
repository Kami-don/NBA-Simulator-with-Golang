# Step 1: Modules caching
FROM golang:1.21-alpine3.19 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.21-alpine3.19 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app ./cmd/app

# Step 3: Final
FROM scratch
COPY --from=builder /app/data.json ./
COPY --from=builder /app/config /config
COPY --from=builder /bin/app /app
COPY --from=builder /app/templates /templates 
CMD ["/app"]

