FROM golang:1.22.5

# Install yt-dlp and curl
RUN apt-get update && \
    apt-get install -y curl ffmpeg python3 python3-pip && \
    pip3 install yt-dlp

# Set working directory
WORKDIR /app

# Copy Go module files and download deps
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy rest of the code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port
EXPOSE 3000

# Run the app
CMD ["./main"]
