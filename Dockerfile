# Use the official Go image version 1.22.2 as the base image
FROM golang:1.22.2

# Set the working directory inside the container
WORKDIR /app

# Copy all application files from the host machine to the working directory in the container
COPY . .

# Download and tidy up Go module dependencies
RUN go mod tidy

# Build the Go application and output the binary named "main"
RUN go build -o main .

# Expose port 8080 for the application
EXPOSE 8080

# Specify the command to run the application
CMD ["./main"]
