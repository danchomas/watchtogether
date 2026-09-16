package video

import "time"

type Video struct {
	Id       string
	Title    string
	Source   SourceType
	Url      string
	Duration time.Duration
}
