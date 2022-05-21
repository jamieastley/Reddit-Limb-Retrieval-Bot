package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const subreddit = "golang"

var mockDate = time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)

type fakeClock struct{}

func (f *fakeClock) NowUTC() time.Time {
	return mockDate
}

func TestRepository_Insert(t *testing.T) {

	db := setupDb(t)
	_ = db.AutoMigrate(&BannedSubreddit{})

	data := repository{
		db:    db,
		clock: &fakeClock{},
	}
	repo := Repository{
		BannedSubreddit: &bannedSubredditHandler{data},
	}

	bs, err := repo.BannedSubreddit.Insert(subreddit)
	assert.Equal(t, bs.Subreddit, subreddit)
	assert.Equal(t, bs.InsertedAt, mockDate)
	assert.NoError(t, err)
}

func TestRepository_Get(t *testing.T) {
	sub := BannedSubreddit{
		Subreddit:  subreddit,
		InsertedAt: mockDate,
	}

	db := setupDb(t)
	_ = db.AutoMigrate(&BannedSubreddit{})

	data := repository{
		db:    db,
		clock: &fakeClock{},
	}
	repo := Repository{
		BannedSubreddit: &bannedSubredditHandler{data},
	}

	//t.Run("No results", func(t *testing.T) {
	//	result, err := repo.BannedSubreddit.Get(subreddit)
	//	assert.Nil(t, result)
	//	assert.Error(t, err)
	//})

	t.Run("Returns result", func(t *testing.T) {
		db.Create(&sub)
		result, err := repo.BannedSubreddit.Get(subreddit)
		assert.Equal(t, sub.Subreddit, result.Subreddit)
		assert.NoError(t, err)
	})

}

func setupDb(t *testing.T) *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open("file::memory:?cache=shared"),
		&gorm.Config{
			Logger: logger.Default,
		})
	if err != nil {
		t.Error(err)
	}
	return db
}
