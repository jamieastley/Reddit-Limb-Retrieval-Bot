package repository

import "time"

type BannedSubreddit struct {
	ID         uint `gorm:"primaryKey"`
	Subreddit  string
	InsertedAt time.Time
}
