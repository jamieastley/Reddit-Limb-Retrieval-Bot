package repository

type BannedSubreddit struct {
	Subreddit  string `gorm:"primaryKey"`
	InsertedAt int64  `gorm:"autoCreateTime"`
}
