package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository interface {
	BannedSubredditHandler
}

type BannedSubredditHandler interface {
	Insert(subreddit string) (BannedSubreddit, error)
	Get(subreddit string) (BannedSubreddit, error)
}

type repository struct {
	db    *gorm.DB
	clock Clock
}

func NewRepository(dsn string) (*repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &repository{
		db:    db,
		clock: NewClock(),
	}, nil
}

func (r *repository) Insert(subreddit string) (*BannedSubreddit, error) {
	sub := BannedSubreddit{
		Subreddit:  subreddit,
		InsertedAt: r.clock.NowUTC(),
	}
	if err := r.db.Create(&sub).Error; err != nil {
		return nil, err
	}

	return &sub, nil
}

func (r *repository) Get(subreddit string) (*BannedSubreddit, error) {
	var sub BannedSubreddit
	if err := r.db.Where("subreddit = ?", subreddit).First(&sub).Error; err != nil {
		return nil, err
	}

	return &sub, nil
}
