FROM golang:1.25-alpine3.23 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# CGO_ENABLED=0 - tells Go compiler to build statically linked executable
RUN CGO_ENABLED=0 GOOS=linux    \
    go build -o /out/l2chat main.go

# ------------------------------------------------------
FROM scratch

WORKDIR /app
COPY --from=build /out/l2chat ./l2chat

ENTRYPOINT ["/app/l2chat"]
CMD ["list"]

