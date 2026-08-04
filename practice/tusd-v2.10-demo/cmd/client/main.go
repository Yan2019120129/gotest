package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sort"
	"syscall"

	"example.com/tusd-demo/internal/tusclient"
)

const defaultCheckpointDir = ".tus-checkpoints"

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var err error
	switch os.Args[1] {
	case "upload", "resume":
		err = runUpload(ctx, os.Args[2:])
	case "info":
		err = runInfo(ctx, os.Args[2:])
	case "cancel":
		err = runCancel(ctx, os.Args[2:])
	case "download":
		err = runDownload(ctx, os.Args[2:])
	case "options":
		err = runOptions(ctx, os.Args[2:])
	default:
		usage()
		os.Exit(2)
	}
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func runUpload(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("upload", flag.ContinueOnError)
	filePath := fs.String("file", "", "file to upload")
	endpoint := fs.String("endpoint", "http://localhost:8080/files/", "tus creation endpoint")
	token := fs.String("token", "demo-token", "bearer token")
	chunkSize := fs.Int64("chunk-size", 8<<20, "PATCH chunk size in bytes")
	retries := fs.Int("retries", 5, "maximum retries per chunk")
	checkpointDir := fs.String("checkpoint-dir", defaultCheckpointDir, "local checkpoint directory")
	restart := fs.Bool("restart", false, "ignore the existing checkpoint and create a new upload")
	description := fs.String("description", "", "optional Upload-Metadata description")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *filePath == "" {
		return fmt.Errorf("-file is required")
	}

	client, err := tusclient.New(*endpoint, *token, *chunkSize, *retries, *checkpointDir)
	if err != nil {
		return err
	}
	uploadURL, err := client.UploadFile(ctx, *filePath, tusclient.UploadOptions{
		Restart:     *restart,
		Description: *description,
		Progress: func(uploaded, total int64) {
			percent := float64(100)
			if total > 0 {
				percent = float64(uploaded) / float64(total) * 100
			}
			fmt.Printf("\ruploaded %d/%d bytes (%.1f%%)", uploaded, total, percent)
		},
	})
	fmt.Println()
	if err != nil {
		if uploadURL != "" {
			fmt.Printf("upload URL retained for resume: %s\n", uploadURL)
		}
		return err
	}
	fmt.Printf("completed: %s\n", uploadURL)
	return nil
}

func runInfo(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("info", flag.ContinueOnError)
	uploadURL := fs.String("url", "", "upload resource URL")
	token := fs.String("token", "demo-token", "bearer token")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *uploadURL == "" {
		return fmt.Errorf("-url is required")
	}
	client, err := tusclient.New("http://localhost/", *token, 8<<20, 0, defaultCheckpointDir)
	if err != nil {
		return err
	}
	info, err := client.Head(ctx, *uploadURL)
	if err != nil {
		return err
	}
	fmt.Printf("URL: %s\noffset: %d\nlength: %d\n", info.URL, info.Offset, info.Length)
	keys := make([]string, 0, len(info.Metadata))
	for key := range info.Metadata {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("metadata.%s: %s\n", key, info.Metadata[key])
	}
	return nil
}

func runCancel(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("cancel", flag.ContinueOnError)
	uploadURL := fs.String("url", "", "upload resource URL")
	token := fs.String("token", "demo-token", "bearer token")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *uploadURL == "" {
		return fmt.Errorf("-url is required")
	}
	client, err := tusclient.New("http://localhost/", *token, 8<<20, 0, defaultCheckpointDir)
	if err != nil {
		return err
	}
	if err := client.Cancel(ctx, *uploadURL); err != nil {
		return err
	}
	fmt.Println("upload terminated")
	return nil
}

func runDownload(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("download", flag.ContinueOnError)
	uploadURL := fs.String("url", "", "upload resource URL")
	out := fs.String("out", "", "destination file")
	token := fs.String("token", "demo-token", "bearer token")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *uploadURL == "" || *out == "" {
		return fmt.Errorf("-url and -out are required")
	}
	client, err := tusclient.New("http://localhost/", *token, 8<<20, 0, defaultCheckpointDir)
	if err != nil {
		return err
	}
	if err := client.Download(ctx, *uploadURL, *out); err != nil {
		return err
	}
	fmt.Printf("downloaded to %s\n", *out)
	return nil
}

func runOptions(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("options", flag.ContinueOnError)
	endpoint := fs.String("endpoint", "http://localhost:8080/files/", "tus creation endpoint")
	if err := fs.Parse(args); err != nil {
		return err
	}
	client, err := tusclient.New(*endpoint, "", 8<<20, 0, defaultCheckpointDir)
	if err != nil {
		return err
	}
	options, err := client.Options(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("versions: %s\nextensions: %s\nmax-size: %s\n", options.Versions, options.Extensions, options.MaxSize)
	return nil
}

func usage() {
	fmt.Fprintln(os.Stderr, `Usage:
  client upload   -file ./archive.zip [-restart]
  client resume   -file ./archive.zip
  client info     -url http://localhost:8080/files/<id>
  client cancel   -url http://localhost:8080/files/<id>
  client download -url http://localhost:8080/files/<id> -out ./archive.zip
  client options  -endpoint http://localhost:8080/files/`)
}
