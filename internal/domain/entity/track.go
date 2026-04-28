package entity

import "time"

type Track struct {
	ID       string
	Title    string
	URL      string
	Duration time.Duration
}
