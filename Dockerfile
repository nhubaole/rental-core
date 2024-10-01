FROM golang:1.23-alpine AS build

WORKDIR /build
COPY go.mod go.sum ./
# RUN go mod tidy
RUN go mod download && go mod verify
COPY . .

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o run cmd/server/main.go 

# Moving the binary to the 'final Image' to make it smaller
FROM alpine
WORKDIR /app
COPY --from=build /build/run .
COPY configs/local.yaml configs/local.yaml
ENV GIN_MODE=release
EXPOSE 8080
CMD ["/app/run"]