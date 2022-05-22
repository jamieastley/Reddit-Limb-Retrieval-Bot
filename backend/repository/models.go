package repository

type BannedSubreddit struct {
	Subreddit  string `gorm:"primaryKey"`
	InsertedAt int64  `gorm:"autoCreateTime"`
}

type IgnoredUser struct {
	Username  string `gorm:"primaryKey"`
	IgnoredAt int64  `gorm:"autoCreateTime"`
}
