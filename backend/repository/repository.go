package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IRepository interface {
	Repository
}

type Repository struct {
	BannedSubreddit IBannedSubreddit
}

type repository struct {
	db    *gorm.DB
	clock Clock
}

type IBannedSubreddit interface {
	Insert(subreddit string) (*BannedSubreddit, error)
	Get(subreddit string) (*BannedSubreddit, error)
	GetAll() (*[]BannedSubreddit, error)
}

type bannedSubredditHandler struct {
	repository
}

func (b *bannedSubredditHandler) GetAll() (*[]BannedSubreddit, error) {
	var bannedSubreddits []BannedSubreddit
	results := b.db.Find(&bannedSubreddits)

	if results.Error != nil {
		return &[]BannedSubreddit{}, results.Error
	}

	return &bannedSubreddits, nil
}

func NewRepository(dsn string) (Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		return Repository{}, err
	}

	initErr := initTables(db)
	if err != nil {
		return Repository{}, initErr
	}

	repo := repository{
		db:    db,
		clock: NewClock(),
	}
	return Repository{
		BannedSubreddit: &bannedSubredditHandler{repo},
	}, nil
}

func (b *bannedSubredditHandler) Insert(subreddit string) (*BannedSubreddit, error) {
	sub := BannedSubreddit{
		Subreddit:  subreddit,
		InsertedAt: b.clock.NowUTC().Unix(),
	}
	if err := b.db.FirstOrCreate(&sub).Error; err != nil {
		return nil, err
	}

	return &sub, nil
}

func (b *bannedSubredditHandler) Get(subreddit string) (*BannedSubreddit, error) {
	var sub BannedSubreddit
	if err := b.db.Where("subreddit = ?", subreddit).First(&sub).Error; err != nil {
		fmt.Println(fmt.Sprintf("failed to query for subreddit: %s", subreddit))
		return &sub, err
	}

	return &sub, nil
}

func initTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&BannedSubreddit{},
		&IgnoredUser{},
	)
}
