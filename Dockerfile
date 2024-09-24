# Start from the latest golang base image
FROM golang:latest

# Add maintainer info
LABEL maintainer="Quique <hello@pragmaticreviews.com>"

# Set the current working directory inside the container
WORKDIR /app

# Copy go modules dependency requirements file
COPY go.mod .

# Copy go modules expected hashes file
COPY go.sum .

# Download dependencies
RUN go mod Download

# Copy all the app sources (recursively copies files and directories from the host into)
COPY . .

# Set http port
ENV PORT 5000

# Build the app
RUN go Build

# Remove source files
RUN find . -name "*.go" -type f -delete

# Make port 5000 available to the world outside this container
EXPOSE $PORT

# Run the app
CMD ["./gin-go"]