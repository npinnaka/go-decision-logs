# Use the official Go image as the base image
FROM golang:1.23.6-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download the dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o server .

# Expose the port the server will run on
EXPOSE 3000

# Specify the command to run the application
CMD ["./server"]