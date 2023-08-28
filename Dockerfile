FROM golang:1.16-alpine AS builder

RUN apk add --no-cache upx make gcc musl-dev linux-headers git

ADD . /workspace

WORKDIR /workspace

RUN go build -ldflags "-s -w" -o /app ./cmd/server
RUN upx /app

FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=builder /app /app

ENTRYPOINT ["/app"]
