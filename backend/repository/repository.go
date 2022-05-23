package repository

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IRepository interface {
	Repository
}

type Repository struct {
	BannedSubreddit IBannedSubreddit
	IgnoredUser     IIgnoredUser
}

type repository struct {
	db    *gorm.DB
	clock Clock
}

type IBannedSubreddit interface {
	Insert(subreddit string) (*BannedSubreddit, error)
	Get(subreddit string) (*BannedSubreddit, error)
	GetAll() ([]BannedSubreddit, error)
}

type IIgnoredUser interface {
	Insert(username string) (*IgnoredUser, error)
	Get(username string) (*IgnoredUser, error)
	GetAll() ([]IgnoredUser, error)
	Remove(username string) (int64, error)
}

type bannedSubredditHandler struct {
	repository
}

type ignoredUserHandler struct {
	repository
}

func (b *bannedSubredditHandler) GetAll() ([]BannedSubreddit, error) {
	var bannedSubreddits []BannedSubreddit
	results := b.db.Find(&bannedSubreddits)

	if results.Error != nil {
		return []BannedSubreddit{}, results.Error
	}

	return bannedSubreddits, nil
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
		IgnoredUser:     &ignoredUserHandler{repo},
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
	err := b.db.Where("subreddit = ?", subreddit).First(&sub).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &sub, nil
	}

	return &sub, err
}

func (i *ignoredUserHandler) Insert(username string) (*IgnoredUser, error) {
	user := IgnoredUser{
		Username:  username,
		IgnoredAt: i.clock.NowUTC().Unix(),
	}

	if err := i.db.FirstOrCreate(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (i *ignoredUserHandler) Get(username string) (*IgnoredUser, error) {
	var user IgnoredUser
	err := i.db.Where("username = ?", username).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &user, nil
	}

	return &user, err
}

func (i *ignoredUserHandler) GetAll() ([]IgnoredUser, error) {
	var ignoredUsers []IgnoredUser

	result := i.db.Find(&ignoredUsers)

	return ignoredUsers, result.Error
}

func (i *ignoredUserHandler) Remove(username string) (int64, error) {
	rows := i.db.Where("username = ?", username).Delete(&IgnoredUser{})

	return rows.RowsAffected, rows.Error
}

func initTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&BannedSubreddit{},
		&IgnoredUser{},
	)
}
