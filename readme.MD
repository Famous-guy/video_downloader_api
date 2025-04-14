
# Video Downloader & Uploader API

This Go application is a simple web API that downloads videos from various platforms (YouTube, TikTok, Telegram, etc.), uploads them to Cloudinary, and returns the uploaded video URLs.

## Features

- Accepts video download requests via a POST endpoint
- Supports popular platforms: YouTube, TikTok, Telegram, Facebook, X (Twitter)
- Automatically infers platform from URL
- Uploads videos to Cloudinary
- Progress bar for command line feedback

## Prerequisites

- [Go](https://golang.org/dl/)
- [yt-dlp](https://github.com/yt-dlp/yt-dlp) (for downloading videos from YouTube, TikTok, etc.)
- `curl` (for downloading from Telegram)
- A [Cloudinary](https://cloudinary.com/) account

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Famous-guy/video_downloader_api.git
cd video_downloader_api
```

2. Create a .env file in the root directory

   ```env
   CLOUDINARY_URL=cloudinary://<API_KEY>:<API_SECRET>@<CLOUD_NAME>
   ```

3. Install Go dependencies:
   ```bash
   go mod tidy
   ```

### Running the App

```bash
go run main.go
```

The server will start at http://localhost:3000.

### API Usage

#### Endpoint

```bash
POST /download
```

#### Request Body

```json
{
  "tasks": [
    {
      "url": "https://www.youtube.com/watch?v=example"
    },
    {
      "url": "https://t.me/your_telegram_video"
    }
  ],
  "folder": "./downloads"
}
```

#### Response

```json
{
  "message": "Download process completed",
  "results": [
    "https://res.cloudinary.com/.../video1.mp4",
    "https://res.cloudinary.com/.../video2.mp4"
  ]
}
```



