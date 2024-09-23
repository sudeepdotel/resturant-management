# Dockerfile
FROM golang:alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o resturant-management .
CMD ["/app/resturant-management"]