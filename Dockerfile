# build the application
FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod .
RUN go mod download 

COPY . .
RUN CGO_ENABLED="0" go build -ldflags="-s -w" -o app main.go


# Copy go build file to alpine image

FROM alpine:3.19.1

WORKDIR /app

COPY --from=builder /app/static ./static
COPY --from=builder /app/app .

CMD [ "./app", "-listen", "0.0.0.0:80"]
