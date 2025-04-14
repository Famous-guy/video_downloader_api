# # # Use official Golang image
# # FROM golang:1.22.5

# # # Install yt-dlp, curl, ffmpeg, python3 and pip
# # RUN apt-get update && \
# #     apt-get install -y --no-install-recommends \
# #         curl \
# #         ffmpeg \
# #         python3 \
# #         python3-pip \
# #         ca-certificates && \
# #     pip3 install --break-system-packages --no-cache-dir yt-dlp && \
# #     apt-get clean && rm -rf /var/lib/apt/lists/*

# # # Set working directory
# # WORKDIR /app

# # # Copy Go module files and download dependencies
# # COPY go.mod ./
# # COPY go.sum ./
# # RUN go mod download

# # # Copy the rest of the code
# # COPY . .

# # # Build the Go application
# # RUN go build -o main .

# # # Expose application port
# # EXPOSE 3000

# # # Command to run the app
# # CMD ["./main"]


# # Use official Golang image
# FROM golang:1.22.5 AS builder

# # Install yt-dlp, curl, ffmpeg, python3 and pip
# RUN apt-get update && \
#     apt-get install -y --no-install-recommends \
#         curl \
#         ffmpeg \
#         python3 \
#         python3-pip \
#         ca-certificates && \
#     # Install yt-dlp using pip3 for the latest version
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

# # Final image to reduce size
# FROM golang:1.22.5 AS final

# # Install ffmpeg and yt-dlp as they are needed to run the application
# RUN apt-get update && \
#     apt-get install -y --no-install-recommends \
#         curl \
#         ffmpeg \
#         python3 \
#         python3-pip \
#         ca-certificates && \
#     pip3 install --break-system-packages --no-cache-dir yt-dlp && \
#     apt-get clean && rm -rf /var/lib/apt/lists/*

# # Set the working directory in the final image
# WORKDIR /app

# # Copy the built binary


# Use official Golang image
FROM golang:1.22.5

# Install yt-dlp, curl, ffmpeg, python3 and pip
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
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose application port
EXPOSE 3000

# Command to run the app
CMD ["./main"]
