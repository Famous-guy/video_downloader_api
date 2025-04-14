package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

type DownloadRequest struct {
	Tasks  []DownloadTask `json:"tasks"`
	Folder string         `json:"folder"` // optional
}

type DownloadTask struct {
	URL      string `json:"url"`
	Platform string `json:"platform"` // optional
}

var platforms = []string{"Telegram", "TikTok", "YouTube", "Facebook", "X"}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file")
	}

	// Log the CLOUDINARY_URL to validate it
	cm := os.Getenv("CLOUDINARY_URL")
	if cm == "" {
		log.Fatal("CLOUDINARY_URL not set in the environment variables")
	}
	fmt.Println("CLOUDINARY_URL:", cm)

	// Set up Fiber app
	app := fiber.New()

	// Define POST route to handle download
	app.Post("/download", handleDownload)

	// Start the app on port 3000
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}

func handleDownload(c *fiber.Ctx) error {
	var req DownloadRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request format")
	}

	if req.Folder == "" {
		cwd, _ := os.Getwd()
		req.Folder = cwd
	}
	os.MkdirAll(req.Folder, os.ModePerm)

	// Initialize progress bar
	p := mpb.New()
	results := make([]string, len(req.Tasks))

	done := make(chan struct{})

	// Loop over each task and handle it asynchronously
	for i, task := range req.Tasks {
		if task.Platform == "" {
			inferred, err := inferPlatform(task.URL)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Unable to determine platform for URL %s", task.URL))
			}
			task.Platform = inferred
		}

		// Set up the progress bar
		bar := p.New(100,
			mpb.BarStyle().Rbound("▕").Filler("█").Tip("█").Padding("░").Lbound("▏"),
			mpb.PrependDecorators(
				decor.Name(fmt.Sprintf("[%d] %s ", i+1, task.Platform)),
				decor.Percentage(),
			),
		)

		// Start downloading asynchronously
		go func(index int, t DownloadTask, b *mpb.Bar) {
			url, err := downloadWithProgress(t.URL, t.Platform, req.Folder, b)
			if err != nil {
				results[index] = fmt.Sprintf("Failed to download %s: %v", t.URL, err)
			} else {
				results[index] = url
			}
			b.SetCurrent(100)
			done <- struct{}{}
		}(i, task, bar)
	}

	// Wait for all download tasks to complete
	for range req.Tasks {
		<-done
	}

	p.Wait()

	// Return the results as JSON
	return c.JSON(fiber.Map{
		"message": "Download process completed",
		"results": results,
	})
}

// 	// Set up the download command based on platform
// 	var cmd *exec.Cmd
// 	if platform == "Telegram" {
// 		cmd = exec.Command("curl", "-L", "-o", finalPath, videoURL)
// 	} else {
// 		// Pass the cookies file to yt-dlp
// 		cookiesFilePath := "./youtube.json" // Path to your cookies file (update if necessary)
// 		cmd = exec.Command("yt-dlp", "--cookies", cookiesFilePath, "-o", finalPath, videoURL)
// 	}

// 	// Simulate progress
// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			time.Sleep(30 * time.Millisecond)
// 			bar.IncrBy(1)
// 		}
// 	}()

// 	// Execute the download command
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	if err := cmd.Run(); err != nil {
// 		return "", err
// 	}

// 	// Upload to Cloudinary
// 	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create cloudinary instance: %w", err)
// 	}

// 	resp, err := cld.Upload.Upload(context.Background(), finalPath, uploader.UploadParams{})
// 	if err != nil {
// 		return "", fmt.Errorf("cloudinary upload failed: %w", err)
// 	}

// 	// Optional: remove local file after upload
// 	os.Remove(finalPath)

//		return resp.SecureURL, nil
//	}
func downloadWithProgress(videoURL, platform, folder string, bar *mpb.Bar) (string, error) {
	tempName := fmt.Sprintf("tempfile_%d.mp4", time.Now().UnixNano())
	finalPath := filepath.Join(folder, tempName)

	// Set up the download command based on platform
	var cmd *exec.Cmd
	if platform == "Telegram" {
		cmd = exec.Command("curl", "-L", "-o", finalPath, videoURL)
	} else {
		// Assuming you have the cookies file in Netscape format at `cookies.txt`
		cmd = exec.Command("yt-dlp", "--cookies", "cookies.txt", "-o", finalPath, videoURL)
	}

	// Simulate progress
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(30 * time.Millisecond)
			bar.IncrBy(1)
		}
	}()

	// Execute the download command
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}

	// Upload to Cloudinary
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return "", fmt.Errorf("failed to create cloudinary instance: %w", err)
	}

	resp, err := cld.Upload.Upload(context.Background(), finalPath, uploader.UploadParams{})
	if err != nil {
		return "", fmt.Errorf("cloudinary upload failed: %w", err)
	}

	// Optional: remove local file after upload
	os.Remove(finalPath)

	return resp.SecureURL, nil
}

func inferPlatform(link string) (string, error) {
	parsed, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	host := parsed.Hostname()
	domainMap := map[string]string{
		"t.me":     "Telegram",
		"telegram": "Telegram",
		"youtube":  "YouTube",
		"youtu.be": "YouTube",
		"facebook": "Facebook",
		"fb.watch": "Facebook",
		"tiktok":   "TikTok",
		"twitter":  "X",
		"x.com":    "X",
	}

	// Infer the platform based on the URL's host
	for domain, platform := range domainMap {
		if strings.Contains(host, domain) {
			return platform, nil
		}
	}

	return "", fmt.Errorf("unknown domain")
}
