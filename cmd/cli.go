package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/opensaucerer/bifrost"
	"github.com/spf13/cobra"
)

var (
	defaultBucket string
	provider      string
	accessKey     string
	secretKey     string
	enableDebug   bool
	publicRead    bool
	region        string
	zone          string
	filePath      string
	fileName      string
	version       string = "0.0.1"
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:     "upload",
	Version: version,
	Short:   "CLI tool for Bifrost",
	Long:    `This is a simple CLI tool for Bifrost - A Rainbow bridge used for shipping your file(s) to any cloud storage service.`,
	Run: func(cmd *cobra.Command, args []string) {
		if provider == "" {
			fmt.Print("Enter the cloud provider name (e.g. s3, gcs): ")
			fmt.Scanln(&provider)
		}
		if defaultBucket == "" {
			fmt.Print("Enter the default bucket name: ")
			fmt.Scanln(&defaultBucket)
		}
		if accessKey == "" {
			fmt.Print("Enter the access key: ")
			fmt.Scanln(&accessKey)
		}
		if secretKey == "" {
			fmt.Print("Enter the secret key: ")
			fmt.Scanln(&secretKey)
		}
		if region == "" {
			fmt.Print("Enter the region name: ")
			fmt.Scanln(&region)
		}
		if zone == "" {
			fmt.Print("Enter the zone name: ")
			fmt.Scanln(&zone)
		}

		bridge, err := bifrost.NewRainbowBridge(&bifrost.BridgeConfig{
			DefaultBucket: defaultBucket,
			Provider:      provider,
			AccessKey:     accessKey,
			SecretKey:     secretKey,
			EnableDebug:   enableDebug,
			PublicRead:    publicRead,
			Region:        region,
			Zone:          zone,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		defer bridge.Disconnect()

		fmt.Printf("Connected to %s\n", bridge.Config().Provider)

		for {
			fmt.Print("Enter the file path (or 'quit' to exit): ")
			fmt.Scanln(&filePath)
			if strings.ToLower(filePath) == "quit" {
				break
			}

			fmt.Print("Enter the filename (or 'quit' to exit): ")
			fmt.Scanln(&fileName)
			if strings.ToLower(fileName) == "quit" {
				break
			}

			fmt.Println("Uploading...")

			uploadedFile, err := bridge.UploadFile(bifrost.File{
				Path:     filePath,
				Filename: fileName,
			})
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("Uploaded file: %s to %s\n", uploadedFile.Name, uploadedFile.Preview)
		}

	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&enableDebug, "enable-debug", "d", false, "enable debug mode")
	rootCmd.PersistentFlags().BoolVarP(&publicRead, "public-read", "r", false, "make uploaded files public")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
