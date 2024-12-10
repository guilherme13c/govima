package config

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var ffmpegArgs = ffmpeg.KwArgs{
	"c:v":     "libx264",
	"pix_fmt": "yuv420p",
	"preset":  "slow",
	"crf":     18,
}

type Config_t struct {
	Width      float64
	Height     float64
	FrameDir   string
	OutputDir  string
	FFmpegArgs ffmpeg.KwArgs
}

var Config = Config_t{
	Width:      1920,
	Height:     1080,
	FrameDir:   "frames",
	OutputDir:  "output",
	FFmpegArgs: ffmpegArgs,
}