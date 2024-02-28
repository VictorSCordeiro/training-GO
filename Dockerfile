# Use the official Alpine image as the base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable to the container
COPY www /app/www

# Expose the port on which the Go service will run
EXPOSE 3333

# Command to run the Go service
CMD ["./www"]