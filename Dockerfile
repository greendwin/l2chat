FROM golang:1.25-alpine3.23 AS build

WORKDIR /src

RUN apk add musl-dev gcc libpcap-dev

COPY go.mod go.sum ./
RUN go mod download

COPY *.go .
COPY cmd ./cmd
COPY methods ./methods
COPY proto ./proto
COPY server ./server

# CGO_ENABLED=0 - tells Go compiler to build statically linked executable
RUN CGO_ENABLED=1 GOOS=linux    \
    go build -o /out/l2chat main.go

ENTRYPOINT ["/out/l2chat"]
CMD ["list"]

# ---------------

FROM alpine:latest

RUN apk add libpcap-dev

WORKDIR /app
COPY --from=build /out/l2chat /app/l2chat

ENTRYPOINT ["/app/l2chat"]
CMD ["list"]

