# Use an official Golang runtime as an image
FROM golang:1.21.1-alpine3.17

# Set working directory
WORKDIR /go/src/app

# Copy the local package files to the container’s workspace.
COPY . .


# Install package
RUN go mod download
RUN go build -o main .

# Run the service
CMD ["./main"]

