// // // package main

// // // import (
// // // 	"context"
// // // 	"fmt"
// // // 	"log"
// // // 	"net/url"
// // // 	"os"
// // // 	"os/exec"
// // // 	"path/filepath"
// // // 	"strings"
// // // 	"time"
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // //
// // // type DownloadRequest struct {
// // // 	Tasks  []DownloadTask `json:"tasks"`
// // // 	Folder string         `json:"folder"` // optional
// // // }
// // // type DownloadTask struct {
// // // 	URL      string `json:"url"`
// // // 	Platform string `json:"platform"` // optional
// // // }
// // // var platforms = []string{"Telegram", "TikTok", "YouTube", "Facebook", "X"}
// // // func main(
// // // func main(
// // // func main(
// // // func main(

// // // 	"github.com/cloudinary/cloudinary-go/v2"
// // // 	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
// // // 	"github.com/gofiber/fiber/v2"
// // // 	"github.com/joho/godotenv"
// // // 	"github.com/vbauerster/mpb/v8"
// // // 	"github.com/vbauerster/mpb/v8/decor"
// // // func main() {
// // // 	// Load .env file
// // // 	if err := godotenv.Load(); err != nil {
// // // 		log.Printf("Error loading .env file")
// // // 	}

// // // 	// Log the CLOUDINARY_URL to validate it
// // // 	cm := os.Getenv("CLOUDINARY_URL")
// // // 	if cm == "" {
// // // 		log.Fatal("CLOUDINARY_URL not set in the environment variables")
// // // 	}
// // // 	fmt.Println("CLOUDINARY_URL:", cm)

// // // 	// Set up Fiber app
// // // 	app := fiber.New()

// // // 	// Define POST route to handle download
// // // 	app.Post("/download", handleDownload)

// // // 	// Start the app on port 3000
// // // 	if err := app.Listen(":3000"); err != nil {
// // // 		panic(err)
// // // 	}
// // // }

// // // func handleDownload(c *fiber.Ctx) error {
// // // 	var req DownloadRequest
// // // 	if err := c.BodyParser(&req); err != nil {
// // // 		return fiber.NewError(fiber.StatusBadRequest, "Invalid request format")
// // // 	}

// // // 	if req.Folder == "" {
// // // 		cwd, _ := os.Getwd()
// // // 		req.Folder = cwd
// // // 	}
// // // 	os.MkdirAll(req.Folder, os.ModePerm)

// // // 	// Initialize progress bar
// // // 	p := mpb.New()
// // // 	results := make([]string, len(req.Tasks))

// // // 	done := make(chan struct{})

// // // 	// Loop over each task and handle it asynchronously
// // // 	for i, task := range req.Tasks {
// // // 		if task.Platform == "" {
// // // 			inferred, err := inferPlatform(task.URL)
// // // 			if err != nil {
// // // 				return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Unable to determine platform for URL %s", task.URL))
// // // 			}
// // // 			task.Platform = inferred
// // // 		}

// // // 		// Set up the progress bar
// // // 		bar := p.New(100,
// // // 			mpb.BarStyle().Rbound("▕").Filler("█").Tip("█").Padding("░").Lbound("▏"),
// // // 			mpb.PrependDecorators(
// // // 				decor.Name(fmt.Sprintf("[%d] %s ", i+1, task.Platform)),
// // // 				decor.Percentage(),
// // // 			),
// // // 		)

// // // 		// Start downloading asynchronously
// // // 		go func(index int, t DownloadTask, b *mpb.Bar) {
// // // 			url, err := downloadWithProgress(t.URL, t.Platform, req.Folder, b)
// // // 			if err != nil {
// // // 				results[index] = fmt.Sprintf("Failed to download %s: %v", t.URL, err)
// // // 			} else {
// // // 				results[index] = url
// // // 			}
// // // 			b.SetCurrent(100)
// // // 			done <- struct{}{}
// // // 		}(i, task, bar)
// // // 	}

// // // 	// Wait for all download tasks to complete
// // // 	for range req.Tasks {
// // // 		<-done
// // // 	}

// // // 	p.Wait()

// // // 	// Return the results as JSON
// // // 	return c.JSON(fiber.Map{
// // // 		"message": "Download process completed",
// // // 		"results": results,
// // // 	})
// // // }

// // // // 	// Set up the download command based on platform
// // // // 	var cmd *exec.Cmd
// // // // 	if platform == "Telegram" {
// // // // 		cmd = exec.Command("curl", "-L", "-o", finalPath, videoURL)
// // // // 	} else {
// // // // 		// Pass the cookies file to yt-dlp
// // // // 		cookiesFilePath := "./youtube.json" // Path to your cookies file (update if necessary)
// // // // 		cmd = exec.Command("yt-dlp", "--cookies", cookiesFilePath, "-o", finalPath, videoURL)
// // // // 	}

// // // // 	// Simulate progress
// // // // 	go func() {
// // // // 		for i := 0; i < 100; i++ {
// // // // 			time.Sleep(30 * time.Millisecond)
// // // // 			bar.IncrBy(1)
// // // // 		}
// // // // 	}()

// // // // 	// Execute the download command
// // // // 	cmd.Stdout = os.Stdout
// // // // 	cmd.Stderr = os.Stderr
// // // // 	if err := cmd.Run(); err != nil {
// // // // 		return "", err
// // // // 	}

// // // // 	// Upload to Cloudinary
// // // // 	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
// // // // 	if err != nil {
// // // // 		return "", fmt.Errorf("failed to create cloudinary instance: %w", err)
// // // // 	}

// // // // 	resp, err := cld.Upload.Upload(context.Background(), finalPath, uploader.UploadParams{})
// // // // 	if err != nil {
// // // // 		return "", fmt.Errorf("cloudinary upload failed: %w", err)
// // // // 	}

// // // // 	// Optional: remove local file after upload
// // // // 	os.Remove(finalPath)

// // // //		return resp.SecureURL, nil
// // // //	}
// // // func downloadWithProgress(videoURL, platform, folder string, bar *mpb.Bar) (string, error) {
// // // 	tempName := fmt.Sprintf("tempfile_%d.mp4", time.Now().UnixNano())
// // // 	finalPath := filepath.Join(folder, tempName)

// // // 	// Set up the download command based on platform
// // // 	var cmd *exec.Cmd
// // // 	if platform == "Telegram" {
// // // 		cmd = exec.Command("curl", "-L", "-o", finalPath, videoURL)
// // // 	} else {
// // // 		// Assuming you have the cookies file in Netscape format at `cookies.txt`
// // // 		cmd = exec.Command("yt-dlp", "--cookies", "cookies.txt", "-o", finalPath, videoURL)
// // // 	}

// // // 	// Simulate progress
// // // 	go func() {
// // // 		for i := 0; i < 100; i++ {
// // // 			time.Sleep(30 * time.Millisecond)
// // // 			bar.IncrBy(1)
// // // 		}
// // // 	}()

// // // 	// Execute the download command
// // // 	cmd.Stdout = os.Stdout
// // // 	cmd.Stderr = os.Stderr
// // // 	if err := cmd.Run(); err != nil {
// // // 		return "", err
// // // 	}

// // // 	// Upload to Cloudinary
// // // 	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
// // // 	if err != nil {
// // // 		return "", fmt.Errorf("failed to create cloudinary instance: %w", err)
// // // 	}

// // // 	resp, err := cld.Upload.Upload(context.Background(), finalPath, uploader.UploadParams{})
// // // 	if err != nil {
// // // 		return "", fmt.Errorf("cloudinary upload failed: %w", err)
// // // 	}

// // // 	// Optional: remove local file after upload
// // // 	os.Remove(finalPath)

// // // 	return resp.SecureURL, nil
// // // }

// // // func inferPlatform(link string) (string, error) {
// // // 	parsed, err := url.Parse(link)
// // // 	if err != nil {
// // // 		return "", err
// // // 	}

// // // 	host := parsed.Hostname()
// // // 	domainMap := map[string]string{
// // // 		"t.me":     "Telegram",
// // // 		"telegram": "Telegram",
// // // 		"youtube":  "YouTube",
// // // 		"youtu.be": "YouTube",
// // // 		"facebook": "Facebook",
// // // 		"fb.watch": "Facebook",
// // // 		"tiktok":   "TikTok",
// // // 		"twitter":  "X",
// // // 		"x.com":    "X",
// // // 	}

// // // 	// Infer the platform based on the URL's host
// // // 	for domain, platform := range domainMap {
// // // 		if strings.Contains(host, domain) {
// // // 			return platform, nil
// // // 		}
// // // 	}

// // // 	return "", fmt.Errorf("unknown domain")
// // // }

// // package main

// // import (
// // 	"bytes"
// // 	"context"
// // 	"fmt"
// // 	"log"
// // 	"math/rand"
// // 	"net/url"
// // 	"os"
// // 	"os/exec"
// // 	"path/filepath"
// // 	"strings"
// // 	"time"

// // 	"github.com/cloudinary/cloudinary-go/v2"
// // 	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
// // 	"github.com/gofiber/fiber/v2"
// // 	"github.com/joho/godotenv"
// // 	"github.com/vbauerster/mpb/v8"
// // 	"github.com/vbauerster/mpb/v8/decor"
// // 	"golang.org/x/sync/semaphore"
// // )

// // type DownloadRequest struct {
// // 	Tasks  []DownloadTask `json:"tasks"`
// // 	Folder string         `json:"folder"` // optional
// // }

// // type DownloadTask struct {
// // 	URL      string `json:"url"`
// // 	Platform string `json:"platform"` // optional
// // }

// // var platforms = []string{"Telegram", "TikTok", "YouTube", "Facebook", "X"}

// // // List of proxies (replace with your proxy URLs)
// // var proxies = []string{
// // 	"http://8.220.204.215:808",
// // 	"http://39.101.65.228:5001",
// // 	"http://39.101.65.228:50001",
// // 	"http://101.37.12.43:8888",
// // 	"http://47.74.46.81:102",
// // 	"http://39.101.65.228:7890",
// // 	"http://8.138.125.130:8443",
// // 	"http://47.74.46.81:45",
// // 	"http://39.102.208.23:3128",
// // 	"http://8.138.131.110:8081",
// // 	"http://8.213.134.213:3128",
// // 	"http://39.102.208.236:8081",
// // 	"http://39.101.65.228:9443",
// // 	"http://8.220.204.215:3128",
// // 	"http://47.122.65.254:8080",
// // 	"http://39.101.65.228:3127",
// // 	"http://39.101.65.228:4006",
// // 	"http://8.213.134.213:6666",
// // 	"http://47.74.46.81:11310",
// // 	"http://39.104.27.89:8080",
// // 	"http://8.213.128.9:18081",
// // 	"http://8.213.134.213:4145",
// // 	"http://47.122.61.139:8080",
// // 	"http://47.104.27.249:8008",
// // 	"http://39.101.65.228:8009",
// // 	"http://47.104.27.249:9098",
// // 	"http://8.213.134.213:8081",
// // 	"http://8.138.125.13:80",
// // 	"http://39.101.65.228:8899",
// // 	"http://47.252.11.233:8080",
// // 	"http://47.90.167.27:102",
// // 	"http://8.130.36.245:4006",
// // 	"http://72.10.164.178:5137",
// // 	"http://67.43.236.19:1117",
// // 	"http://8.130.36.245:9090",
// // }

// // func main() {
// // 	// Load .env file
// // 	if err := godotenv.Load(); err != nil {
// // 		log.Printf("Warning: Error loading .env file: %v", err)
// // 	}

// // 	// Validate CLOUDINARY_URL
// // 	cm := os.Getenv("CLOUDINARY_URL")
// // 	if cm == "" {
// // 		log.Fatal("CLOUDINARY_URL not set in environment variables")
// // 	}
// // 	log.Println("CLOUDINARY_URL:", cm)

// // 	// Set up Fiber app
// // 	app := fiber.New()

// // 	// Define POST route
// // 	app.Post("/download", handleDownload)

// // 	// Start server
// // 	if err := app.Listen(":3000"); err != nil {
// // 		log.Fatalf("Failed to start server: %v", err)
// // 	}
// // }

// // func handleDownload(c *fiber.Ctx) error {
// // 	var req DownloadRequest
// // 	if err := c.BodyParser(&req); err != nil {
// // 		return fiber.NewError(fiber.StatusBadRequest, "Invalid request format")
// // 	}

// // 	// Set default folder if not provided
// // 	if req.Folder == "" {
// // 		cwd, _ := os.Getwd()
// // 		req.Folder = cwd
// // 	}
// // 	if err := os.MkdirAll(req.Folder, os.ModePerm); err != nil {
// // 		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed to create folder: %v", err))
// // 	}

// // 	// Initialize progress bar
// // 	p := mpb.New()
// // 	results := make([]string, len(req.Tasks))
// // 	done := make(chan struct{})
// // 	sem := semaphore.NewWeighted(2) // Limit to 2 concurrent downloads

// // 	// Process each task
// // 	for i, task := range req.Tasks {
// // 		if task.Platform == "" {
// // 			inferred, err := inferPlatform(task.URL)
// // 			if err != nil {
// // 				return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Unable to determine platform for URL %s", task.URL))
// // 			}
// // 			task.Platform = inferred
// // 		}

// // 		// Set up progress bar
// // 		bar := p.New(100,
// // 			mpb.BarStyle().Rbound("▕").Filler("█").Tip("█").Padding("░").Lbound("▏"),
// // 			mpb.PrependDecorators(
// // 				decor.Name(fmt.Sprintf("[%d] %s ", i+1, task.Platform)),
// // 				decor.Percentage(),
// // 			),
// // 		)

// // 		// Start download asynchronously
// // 		go func(index int, t DownloadTask, b *mpb.Bar) {
// // 			defer sem.Release(1)
// // 			if err := sem.Acquire(context.Background(), 1); err != nil {
// // 				results[index] = fmt.Sprintf("Failed to acquire semaphore: %v", err)
// // 				done <- struct{}{}
// // 				return
// // 			}

// // 			url, err := downloadWithProgress(t.URL, t.Platform, req.Folder, b)
// // 			if err != nil {
// // 				results[index] = fmt.Sprintf("Failed to download %s: %v", t.URL, err)
// // 			} else {
// // 				results[index] = url
// // 			}
// // 			b.SetCurrent(100)
// // 			done <- struct{}{}
// // 		}(i, task, bar)

// // 		// Random delay between 2-5 seconds to avoid rate limiting
// // 		rand.Seed(time.Now().UnixNano())
// // 		time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)
// // 	}

// // 	// Wait for all downloads to complete
// // 	for range req.Tasks {
// // 		<-done
// // 	}

// // 	p.Wait()

// // 	// Return results
// // 	return c.JSON(fiber.Map{
// // 		"message": "Download process completed",
// // 		"results": results,
// // 	})
// // }

// // func downloadWithProgress(videoURL, platform, folder string, bar *mpb.Bar) (string, error) {
// // 	tempName := fmt.Sprintf("tempfile_%d.mp4", time.Now().UnixNano())
// // 	finalPath := filepath.Join(folder, tempName)

// // 	// Set up download command
// // 	var cmd *exec.Cmd
// // 	if platform == "Telegram" {
// // 		cmd = exec.Command("curl", "-L", "-o", finalPath, videoURL)
// // 	} else {
// // 		// Use a random proxy for non-Telegram platforms
// // 		rand.Seed(time.Now().UnixNano())
// // 		var proxy string
// // 		if len(proxies) > 0 {
// // 			proxy = proxies[rand.Intn(len(proxies))]
// // 			cmd = exec.Command("yt-dlp", "--proxy", proxy, "-o", finalPath, videoURL)
// // 		} else {
// // 			// Fallback to no proxy if none provided
// // 			cmd = exec.Command("yt-dlp", "-o", finalPath, videoURL)
// // 		}
// // 	}

// // 	// Capture output for debugging
// // 	var stdout, stderr bytes.Buffer
// // 	cmd.Stdout = &stdout
// // 	cmd.Stderr = &stderr

// // 	// Simulate progress (placeholder; consider parsing yt-dlp output for real progress)
// // 	go func() {
// // 		for i := 0; i < 100; i++ {
// // 			time.Sleep(30 * time.Millisecond)
// // 			bar.IncrBy(1)
// // 		}
// // 	}()

// // 	// Execute command
// // 	if err := cmd.Run(); err != nil {
// // 		log.Printf("Error downloading %s (%s): %v\nStdout: %s\nStderr: %s", videoURL, platform, err, stdout.String(), stderr.String())
// // 		return "", fmt.Errorf("download failed: %w", err)
// // 	}

// // 	// Upload to Cloudinary
// // 	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
// // 	if err != nil {
// // 		return "", fmt.Errorf("failed to create Cloudinary instance: %w", err)
// // 	}

// // 	resp, err := cld.Upload.Upload(context.Background(), finalPath, uploader.UploadParams{})
// // 	if err != nil {
// // 		return "", fmt.Errorf("Cloudinary upload failed: %w", err)
// // 	}

// // 	// Clean up local file
// // 	if err := os.Remove(finalPath); err != nil {
// // 		log.Printf("Warning: Failed to remove local file %s: %v", finalPath, err)
// // 	}

// // 	return resp.SecureURL, nil
// // }

// // func inferPlatform(link string) (string, error) {
// // 	parsed, err := url.Parse(link)
// // 	if err != nil {
// // 		return "", fmt.Errorf("invalid URL: %w", err)
// // 	}

// // 	host := parsed.Hostname()
// // 	domainMap := map[string]string{
// // 		"t.me":     "Telegram",
// // 		"telegram": "Telegram",
// // 		"youtube":  "YouTube",
// // 		"youtu.be": "YouTube",
// // 		"facebook": "Facebook",
// // 		"fb.watch": "Facebook",
// // 		"tiktok":   "TikTok",
// // 		"twitter":  "X",
// // 		"x.com":    "X",
// // 	}

// // 	for domain, platform := range domainMap {
// // 		if strings.Contains(host, domain) {
// // 			return platform, nil
// // 		}
// // 	}

// // 	return "", fmt.Errorf("unknown domain: %s", host)
// // }

// package main

// import (
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"net/url"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// 	"strings"
// 	"time"

// 	"github.com/cloudinary/cloudinary-go/v2"
// 	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/joho/godotenv"
// 	"github.com/vbauerster/mpb/v8"
// 	"github.com/vbauerster/mpb/v8/decor"
// 	"golang.org/x/sync/semaphore"
// )

// type DownloadRequest struct {
// 	Tasks  []DownloadTask `json:"tasks"`
// 	Folder string         `json:"folder"` // optional
// }

// type DownloadTask struct {
// 	URL      string `json:"url"`
// 	Platform string `json:"platform"` // optional
// }

// var platforms = []string{"Telegram", "TikTok", "YouTube", "Facebook", "X"}

// // Default proxy list (fallback if proxies.txt is missing)
// var proxies = []string{
// 	"http://8.220.204.215:808",
// 	"http://39.101.65.228:5001",
// 	"http://39.101.65.228:50001",
// 	"http://101.37.12.43:8888",
// 	"http://47.74.46.81:102",
// 	"http://39.101.65.228:7890",
// 	"http://8.138.125.130:8443",
// 	"http://47.74.46.81:45",
// 	"http://39.102.208.23:3128",
// 	"http://8.138.131.110:8081",
// 	"http://8.213.134.213:3128",
// 	"http://39.102.208.236:8081",
// 	"http://39.101.65.228:9443",
// 	"http://8.220.204.215:3128",
// 	"http://47.122.65.254:8080",
// 	"http://39.101.65.228:3127",
// 	"http://39.101.65.228:4006",
// 	"http://8.213.134.213:6666",
// 	"http://47.74.46.81:11310",
// 	"http://39.104.27.89:8080",
// 	"http://8.213.128.9:18081",
// 	"http://8.213.134.213:4145",
// 	"http://47.122.61.139:8080",
// 	"http://47.104.27.249:8008",
// 	"http://39.101.65.228:8009",
// 	"http://47.104.27.249:9098",
// 	"http://8.213.134.213:8081",
// 	"http://8.138.125.13:80",
// 	"http://39.101.65.228:8899",
// 	"http://47.252.11.233:8080",
// 	"http://47.90.167.27:102",
// 	"http://8.130.36.245:4006",
// 	"http://72.10.164.178:5137",
// 	"http://67.43.236.19:1117",
// 	"http://8.130.36.245:9090",
// }

// func main() {
// 	// Load .env file
// 	if err := godotenv.Load(); err != nil {
// 		log.Printf("Warning: Error loading .env file: %v", err)
// 	}

// 	// Validate CLOUDINARY_URL
// 	cm := os.Getenv("CLOUDINARY_URL")
// 	if cm == "" {
// 		log.Fatal("CLOUDINARY_URL not set in environment variables")
// 	}
// 	log.Println("CLOUDINARY_URL:", cm)

// 	// Load proxies from proxies.txt (if exists)
// 	if newProxies, err := loadProxies("proxies.txt"); err == nil && len(newProxies) > 0 {
// 		proxies = newProxies
// 		log.Printf("Loaded %d proxies from proxies.txt", len(proxies))
// 	} else {
// 		log.Printf("Using default proxy list (%d proxies)", len(proxies))
// 	}

// 	// Set up Fiber app
// 	app := fiber.New()

// 	// Define POST route
// 	app.Post("/download", handleDownload)

// 	// Start server
// 	if err := app.Listen(":3000"); err != nil {
// 		log.Fatalf("Failed to start server: %v", err)
// 	}
// }

// func loadProxies(filePath string) ([]string, error) {
// 	data, err := os.ReadFile(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	lines := strings.Split(string(data), "\n")
// 	var proxies []string
// 	for _, line := range lines {
// 		line = strings.TrimSpace(line)
// 		if line != "" && !strings.HasPrefix(line, "#") {
// 			if !strings.HasPrefix(line, "http://") {
// 				line = "http://" + line
// 			}
// 			proxies = append(proxies, line)
// 		}
// 	}
// 	return proxies, nil
// }

// func handleDownload(c *fiber.Ctx) error {
// 	var req DownloadRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, "Invalid request format")
// 	}

// 	// Set default folder if not provided
// 	if req.Folder == "" {
// 		cwd, _ := os.Getwd()
// 		req.Folder = cwd
// 	}
// 	if err := os.MkdirAll(req.Folder, os.ModePerm); err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed to create folder: %v", err))
// 	}

// 	// Initialize progress bar
// 	p := mpb.New()
// 	results := make([]string, len(req.Tasks))
// 	done := make(chan struct{})
// 	sem := semaphore.NewWeighted(2) // Limit to 2 concurrent downloads

// 	// Process each task
// 	for i, task := range req.Tasks {
// 		if task.Platform == "" {
// 			inferred, err := inferPlatform(task.URL)
// 			if err != nil {
// 				return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Unable to determine platform for URL %s", task.URL))
// 			}
// 			task.Platform = inferred
// 		}

// 		// Set up progress bar
// 		bar := p.New(100,
// 			mpb.BarStyle().Rbound("▕").Filler("█").Tip("█").Padding("░").Lbound("▏"),
// 			mpb.PrependDecorators(
// 				decor.Name(fmt.Sprintf("[%d] %s ", i+1, task.Platform)),
// 				decor.Percentage(),
// 			),
// 		)

// 		// Start download asynchronously
// 		go func(index int, t DownloadTask, b *mpb.Bar) {
// 			defer sem.Release(1)
// 			if err := sem.Acquire(context.Background(), 1); err != nil {
// 				results[index] = fmt.Sprintf("Failed to acquire semaphore: %v", err)
// 				done <- struct{}{}
// 				return
// 			}

// 			url, err := downloadWithProgress(t.URL, t.Platform, req.Folder, b)
// 			if err != nil {
// 				results[index] = fmt.Sprintf("Failed to download %s: %v", t.URL, err)
// 			} else {
// 				results[index] = url
// 			}
// 			b.SetCurrent(100)
// 			done <- struct{}{}
// 		}(i, task, bar)

// 		// Random delay between 2-5 seconds to avoid rate limiting
// 		rand.Seed(time.Now().UnixNano())
// 		time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)
// 	}

// 	// Wait for all downloads to complete
// 	for range req.Tasks {
// 		<-done
// 	}

// 	p.Wait()

// 	// Return results
// 	return c.JSON(fiber.Map{
// 		"message": "Download process completed",
// 		"results": results,
// 	})
// }

// func downloadWithProgress(videoURL, platform, folder string, bar *mpb.Bar) (string, error) {
// 	tempName := fmt.Sprintf("tempfile_%d.mp4", time.Now().UnixNano())
// 	finalPath := filepath.Join(folder, tempName)

// 	// Set up download command
// 	var cmd *exec.Cmd
// 	if platform == "Telegram" {
// 		cmd = exec.Command("curl", "-L", "-o", finalPath, videoURL)
// 	} else {
// 		// Try up to 3 proxies for non-Telegram platforms
// 		var err error
// 		for i := 0; i < 3; i++ {
// 			rand.Seed(time.Now().UnixNano())
// 			var proxy string
// 			if len(proxies) > 0 {
// 				proxy = proxies[rand.Intn(len(proxies))]
// 				// Add timeout and retries to yt-dlp
// 				cmd = exec.Command("yt-dlp", "--proxy", proxy, "--socket-timeout", "10", "--retries", "5", "-o", finalPath, videoURL)
// 			} else {
// 				// Fallback to no proxy
// 				cmd = exec.Command("yt-dlp", "--socket-timeout", "10", "--retries", "5", "-o", finalPath, videoURL)
// 			}

// 			// Capture output for debugging
// 			var stdout, stderr bytes.Buffer
// 			cmd.Stdout = &stdout
// 			cmd.Stderr = &stderr

// 			// Simulate progress (placeholder)
// 			go func() {
// 				for i := 0; i < 100; i++ {
// 					time.Sleep(30 * time.Millisecond)
// 					bar.IncrBy(1)
// 				}
// 			}()

// 			// Execute command
// 			err = cmd.Run()
// 			if err == nil {
// 				// Success, proceed to upload
// 				break
// 			}
// 			log.Printf("Attempt %d failed for %s with proxy %s: %v\nStdout: %s\nStderr: %s", i+1, videoURL, proxy, err, stdout.String(), stderr.String())
// 			time.Sleep(1 * time.Second) // Wait before retrying
// 		}

// 		if err != nil {
// 			return "", fmt.Errorf("download failed after 3 attempts: %w", err)
// 		}
// 	}

// 	// Upload to Cloudinary
// 	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create Cloudinary instance: %w", err)
// 	}

// 	resp, err := cld.Upload.Upload(context.Background(), finalPath, uploader.UploadParams{})
// 	if err != nil {
// 		return "", fmt.Errorf("Cloudinary upload failed: %w", err)
// 	}

// 	// Clean up local file
// 	if err := os.Remove(finalPath); err != nil {
// 		log.Printf("Warning: Failed to remove local file %s: %v", finalPath, err)
// 	}

// 	return resp.SecureURL, nil
// }

// func inferPlatform(link string) (string, error) {
// 	parsed, err := url.Parse(link)
// 	if err != nil {
// 		return "", fmt.Errorf("invalid URL: %w", err)
// 	}

// 	host := parsed.Hostname()
// 	domainMap := map[string]string{
// 		"t.me":     "Telegram",
// 		"telegram": "Telegram",
// 		"youtube":  "YouTube",
// 		"youtu.be": "YouTube",
// 		"facebook": "Facebook",
// 		"fb.watch": "Facebook",
// 		"tiktok":   "TikTok",
// 		"twitter":  "X",
// 		"x.com":    "X",
// 	}

// 	for domain, platform := range domainMap {
// 		if strings.Contains(host, domain) {
// 			return platform, nil
// 		}
// 	}

// 	return "", fmt.Errorf("unknown domain: %s", host)
/
/ }

package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/rand"
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
	"golang.org/x/sync/semaphore"
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
var proxies []string // Global proxy list

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Validate CLOUDINARY_URL
	cm := os.Getenv("CLOUDINARY_URL")
	if cm == "" {
		log.Fatal("CLOUDINARY_URL not set in environment variables")
	}
	log.Println("CLOUDINARY_URL:", cm)

	// Load proxies from proxies.txt
	var err error
	proxies, err = loadProxies("proxies.txt")
	if err != nil || len(proxies) == 0 {
		log.Fatal("Failed to load proxies.txt or no valid proxies found")
	}
	log.Printf("Loaded %d proxies from proxies.txt", len(proxies))

	// Shuffle proxies
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(proxies), func(i, j int) {
		proxies[i], proxies[j] = proxies[j], proxies[i]
	})

	// Log yt-dlp version
	cmd := exec.Command("yt-dlp", "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Warning: Failed to get yt-dlp version: %v", err)
	} else {
		log.Printf("yt-dlp version: %s", strings.TrimSpace(string(output)))
	}

	// Set up Fiber app
	app := fiber.New()

	// Define POST route
	app.Post("/download", handleDownload)

	// Start server
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func loadProxies(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	var proxies []string
	// Preferred ports for HTTPS compatibility
	validPorts := map[string]bool{
		"8080": true,
		"3128": true,
		"8888": true,
		"8081": true,
		"6853": true,
		"7890": true,
		"8443": true,
	}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// Ensure proxy starts with http://
		if !strings.HasPrefix(line, "http://") {
			line = "http://" + line
		}
		// Extract port
		u, err := url.Parse(line)
		if err != nil {
			continue
		}
		parts := strings.Split(u.Host, ":")
		if len(parts) != 2 {
			continue
		}
		port := parts[1]
		// Only include proxies with valid ports
		if validPorts[port] {
			proxies = append(proxies, line)
		}
	}
	return proxies, nil
}

func handleDownload(c *fiber.Ctx) error {
	var req DownloadRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request format")
	}

	// Set default folder if not provided
	if req.Folder == "" {
		cwd, _ := os.Getwd()
		req.Folder = cwd
	}
	if err := os.MkdirAll(req.Folder, os.ModePerm); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed to create folder: %v", err))
	}

	// Initialize progress bar
	p := mpb.New()
	results := make([]string, len(req.Tasks))
	done := make(chan struct{})
	sem := semaphore.NewWeighted(2) // Limit to 2 concurrent downloads

	// Process each task
	for i, task := range req.Tasks {
		if task.Platform == "" {
			inferred, err := inferPlatform(task.URL)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Unable to determine platform for URL %s", task.URL))
			}
			task.Platform = inferred
		}

		// Set up progress bar
		bar := p.New(100,
			mpb.BarStyle().Rbound("▕").Filler("█").Tip("█").Padding("░").Lbound("▏"),
			mpb.PrependDecorators(
				decor.Name(fmt.Sprintf("[%d] %s ", i+1, task.Platform)),
				decor.Percentage(),
			),
		)

		// Start download asynchronously
		go func(index int, t DownloadTask, b *mpb.Bar) {
			defer sem.Release(1)
			if err := sem.Acquire(context.Background(), 1); err != nil {
				results[index] = fmt.Sprintf("Failed to acquire semaphore: %v", err)
				done <- struct{}{}
				return
			}

			url, err := downloadWithProgress(t.URL, t.Platform, req.Folder, b)
			if err != nil {
				results[index] = fmt.Sprintf("Failed to download %s: %v", t.URL, err)
			} else {
				results[index] = url
			}
			b.SetCurrent(100)
			done <- struct{}{}
		}(i, task, bar)

		// Random delay between 2-5 seconds to avoid rate limiting
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)
	}

	// Wait for all downloads to complete
	for range req.Tasks {
		<-done
	}

	p.Wait()

	// Return results
	return c.JSON(fiber.Map{
		"message": "Download process completed",
		"results": results,
	})
}

func downloadWithProgress(videoURL, platform, folder string, bar *mpb.Bar) (string, error) {
	tempName := fmt.Sprintf("tempfile_%d.mp4", time.Now().UnixNano())
	finalPath := filepath.Join(folder, tempName)

	// Set up download command
	var cmd *exec.Cmd
	if platform == "Telegram" {
		cmd = exec.Command("curl", "-L", "-o", finalPath, videoURL)
	} else {
		// Try up to 3 proxies, then fallback to no proxy
		var err error
		for i := 0; i < 4; i++ {
			var proxy string
			if i < 3 && len(proxies) > 0 {
				rand.Seed(time.Now().UnixNano())
				proxy = proxies[rand.Intn(len(proxies))]
				cmd = exec.Command("yt-dlp", "--proxy", proxy, "--socket-timeout", "10", "--retries", "5", "-o", finalPath, videoURL)
			} else {
				// Fallback to no proxy on 4th attempt
				cmd = exec.Command("yt-dlp", "--socket-timeout", "10", "--retries", "5", "-o", finalPath, videoURL)
				log.Printf("Attempt %d for %s: Falling back to no proxy", i+1, videoURL)
			}

			// Capture output for debugging
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			// Simulate progress (placeholder)
			go func() {
				for i := 0; i < 100; i++ {
					time.Sleep(30 * time.Millisecond)
					bar.IncrBy(1)
				}
			}()

			// Execute command
			err = cmd.Run()
			if err == nil {
				// Success, proceed to upload
				break
			}
			log.Printf("Attempt %d failed for %s with proxy %s: %v\nStdout: %s\nStderr: %s", i+1, videoURL, proxy, err, stdout.String(), stderr.String())
			time.Sleep(1 * time.Second) // Wait before retrying
		}

		if err != nil {
			return "", fmt.Errorf("download failed after 4 attempts: %w", err)
		}
	}

	// Upload to Cloudinary
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return "", fmt.Errorf("failed to create Cloudinary instance: %w", err)
	}

	resp, err := cld.Upload.Upload(context.Background(), finalPath, uploader.UploadParams{})
	if err != nil {
		return "", fmt.Errorf("Cloudinary upload failed: %w", err)
	}

	// Clean up local file
	if err := os.Remove(finalPath); err != nil {
		log.Printf("Warning: Failed to remove local file %s: %v", finalPath, err)
	}

	return resp.SecureURL, nil
}

func inferPlatform(link string) (string, error) {
	parsed, err := url.Parse(link)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
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

	for domain, platform := range domainMap {
		if strings.Contains(host, domain) {
			return platform, nil
		}
	}

	return "", fmt.Errorf("unknown domain: %s", host)
}
