# Use the official Golang image as the base image
FROM golang:1.22-alpine as build-stage
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go mod file
COPY go.mod ./
# Download all dependencies. Dependencies will be cached if the go.mod file is not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
RUN go build -o main consumer/...

FROM alpine:latest
# Command to run the executable
COPY --from=build-stage /app/main /bin
CMD ["/bin/main"]
