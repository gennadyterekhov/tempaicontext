FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY . .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["/app/linux_amd64"]