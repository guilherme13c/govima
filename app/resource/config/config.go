package config

import (
	"log"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var ffmpegArgs = ffmpeg.KwArgs{
	"c:v":      "libx264",
	"pix_fmt":  "yuv420p",
	"preset":   "slow",
	"loglevel": "error",
	"crf":      18,
}

type Config_t struct {
	FrameDir   string
	OutputDir  string
	LatexDir   string
	FFmpegArgs ffmpeg.KwArgs
}

var Config Config_t

func Init() {
	Config = Config_t{
		FrameDir:   "tmp/frames",
		OutputDir:  "output",
		LatexDir:   "tmp/latex",
		FFmpegArgs: ffmpegArgs,
	}

	createBaseFolders()
}

func createBaseFolders() {
	if err := os.MkdirAll(Config.FrameDir, 0755); err != nil {
		log.Fatalf("Failed to create frame directory: %v", err)
	}

	if err := os.MkdirAll(Config.OutputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	if err := os.MkdirAll(Config.LatexDir, 0755); err != nil {
		log.Fatalf("Failed to create latex directory: %v", err)
	}
}
