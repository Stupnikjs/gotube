package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func main() {
	for {
		fmt.Print(":>")
		input := bufio.NewReader(os.Stdin)
		str, err := input.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if str == "exit" {
			os.Exit(1)
		}
		Wrapper(str[:len(str)-2])
	}

}

func ExampleClient(arg string) string {

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
	filename := ""

	filename = video.Title
	maxLength := 20
	if len(video.Title) > maxLength {
		filename = video.Title[0:maxLength]
	}

	file, err := os.Create(filename + ".mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
	return file.Name()

}

func Wrapper(arg string) {
	filename := ExampleClient(arg)
	defer os.Remove(filename)
	woutMp4 := strings.Split(filename, ".")[0]
	cmd := exec.Command("ffmpeg", "-i", filename, "-q:a", "0", "-map", "a", woutMp4+".mp3")
	cmd.Run()
}
