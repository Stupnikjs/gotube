package main

import (
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func ExampleClient(arg string) {

	client := youtube.Client{}

	video, err := client.GetVideo(arg)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	file, err := os.Create(video.Title + ".mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}

func main() {
	arg := os.Args[1]
	ExampleClient(arg)
}
