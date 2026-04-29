package entity

type PlaybackState string

const (
	PlaybackStateIdle    PlaybackState = "idle"
	PlaybackStatePaused  PlaybackState = "paused"
	PlaybackStatePlaying PlaybackState = "playing"
)

type GuildSession struct {
	guildID       string
	queue         Queue
	playbackState PlaybackState
}
