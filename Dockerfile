# FROM golang:1.22.5

# # Install yt-dlp, curl, ffmpeg, python3 and pip
# RUN apt-get update && \
#     apt-get install -y --no-install-recommends \
#         curl \
#         ffmpeg \
#         python3 \
#         python3-pip \
#         ca-certificates && \
#     pip3 install --break-system-packages --no-cache-dir yt-dlp && \
#     apt-get clean && rm -rf /var/lib/apt/lists/*

# # Set working directory
# WORKDIR /app

# # Copy Go module files and download dependencies
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# # Copy the rest of the code
# COPY . .

# # Build the Go application
# RUN go build -o main .

# # Expose application port
# EXPOSE 3000

# # Command to run the app
# CMD ["./main"]


FROM golang:1.23.0

# Install yt-dlp, curl, ffmpeg, python3, and pip
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        curl \
        ffmpeg \
        python3 \
        python3-pip \
        ca-certificates && \
    pip3 install --break-system-packages --no-cache-dir yt-dlp && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod ./
COPY go.sum ./
# COPY .env ./
COPY proxies.txt ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose application port
EXPOSE 3000

# Command to run the app
CMD ["./main"]