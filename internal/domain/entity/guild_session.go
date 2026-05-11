package entity

import "errors"

type PlaybackState string

const (
	PlaybackStateIdle    PlaybackState = "idle"
	PlaybackStatePaused  PlaybackState = "paused"
	PlaybackStatePlaying PlaybackState = "playing"
)

type GuildID string

type GuildSession struct {
	guildID       GuildID
	queue         Queue
	playbackState PlaybackState
}

func NewGuildSession(guildID GuildID) (*GuildSession, error) {

	if guildID == "" {
		return nil, errors.New("guildID is required")
	}

	return &GuildSession{guildID: guildID, queue: NewQueue(), playbackState: PlaybackStateIdle}, nil
}
