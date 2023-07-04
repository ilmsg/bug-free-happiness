package main

import "fmt"

type MediaPlayer interface {
	Play(audioType string, fileName string) error
}

type AdvancedMediaPlayer interface {
	PlayVideo(fileName string) error
	PlayAudio(fileName string) error
}

type VLCPlayer struct{}

func (v *VLCPlayer) PlayVideo(fileName string) error {
	fmt.Printf("Playing video file. Name: %s\n", fileName)
	return nil
}

func (v *VLCPlayer) PlayAudio(fileName string) error {
	fmt.Printf("Playing audio file. Name: %s\n", fileName)
	return nil
}

type MediaAdapter struct {
	advanceMusicPlayer AdvancedMediaPlayer
}

func (m *MediaAdapter) play(audioType string, fileName string) error {
	if audioType == "vlc" {
		return m.advanceMusicPlayer.PlayVideo(fileName)
	} else if audioType == "mp4" {
		return m.advanceMusicPlayer.PlayAudio(fileName)
	}
	return fmt.Errorf("invalid media. %s format not supported", audioType)
}

type AudioPlayer struct {
	mediaAdapter MediaPlayer
}

func (a *AudioPlayer) Play(audioType string, fileName string) error {
	if audioType == "mp3" {
		fmt.Printf("Playing mp3 file. Name: %s\n", fileName)
		return nil
	} else if audioType == "vlc" || audioType == "mp4" {
		a.mediaAdapter = &MediaPlayer{&VLCPlayer{}}
		return a.mediaAdapter.Play(audioType, fileName)
	}
	return fmt.Errorf("invalid media. %s format not supported", audioType)
}

func main() {

}
