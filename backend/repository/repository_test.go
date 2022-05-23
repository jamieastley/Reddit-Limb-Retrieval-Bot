package repository

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var subredditList = []string{
	"golang",
	"programming",
	"programminghorror",
}

var ignoredUsers = []string{
	"BobbyTables",
	"MichaelScott",
	"AnitaHuginkiss",
}

var mockDate = time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)

type fakeClock struct{}

func (f *fakeClock) NowUTC() time.Time {
	return mockDate
}

func TestBannedSubredditHandler_Success(t *testing.T) {
	db := setupDb(t)
	sqlDb, _ := db.DB()
	defer func(sqlDb *sql.DB) {
		_ = sqlDb.Close()
	}(sqlDb)

	data := repository{
		db:    db,
		clock: &fakeClock{},
	}
	repo := Repository{
		BannedSubreddit: &bannedSubredditHandler{data},
	}

	t.Run(`Insert banned subreddits`, func(t *testing.T) {
		for _, sub := range subredditList {
			res, err := repo.BannedSubreddit.Insert(sub)
			assert.Equal(t, res.Subreddit, sub)
			assert.Equal(t, res.InsertedAt, mockDate.Unix())
			assert.NoError(t, err)
		}
	})

	t.Run(`Get banned subreddit`, func(t *testing.T) {
		res, err := repo.BannedSubreddit.Get(subredditList[0])
		assert.Equal(t, res.Subreddit, subredditList[0])
		assert.Equal(t, res.InsertedAt, mockDate.Unix())
		assert.NoError(t, err)
	})

	t.Run(`Get all banned subreddits`, func(t *testing.T) {
		res, err := repo.BannedSubreddit.GetAll()
		for i, sub := range subredditList {
			assert.Equal(t, res[i].Subreddit, sub)
		}
		assert.NoError(t, err)
	})
}

func TestBannedSubredditHandler_NoResults(t *testing.T) {
	db := setupDb(t)
	sqlDb, _ := db.DB()
	defer func(sqlDb *sql.DB) {
		_ = sqlDb.Close()
	}(sqlDb)

	data := repository{
		db:    db,
		clock: &fakeClock{},
	}
	repo := Repository{
		BannedSubreddit: &bannedSubredditHandler{data},
	}

	t.Run(`Get banned subreddit`, func(t *testing.T) {
		res, err := repo.BannedSubreddit.Get(subredditList[0])
		assert.Equal(t, res, &BannedSubreddit{})
		assert.NoError(t, err)
	})

	t.Run(`Get all banned subreddits`, func(t *testing.T) {
		res, err := repo.BannedSubreddit.GetAll()
		assert.Equal(t, res, []BannedSubreddit{})
		assert.NoError(t, err)
	})
}

func TestIgnoredUserHandler_Success(t *testing.T) {
	db := setupDb(t)
	sqlDb, _ := db.DB()
	defer func(sqlDb *sql.DB) {
		_ = sqlDb.Close()
	}(sqlDb)

	data := repository{
		db:    db,
		clock: &fakeClock{},
	}
	repo := Repository{
		IgnoredUser: &ignoredUserHandler{data},
	}

	t.Run(`Insert banned users`, func(t *testing.T) {
		for _, user := range ignoredUsers {
			res, err := repo.IgnoredUser.Insert(user)
			assert.Equal(t, res.Username, user)
			assert.Equal(t, res.IgnoredAt, mockDate.Unix())
			assert.NoError(t, err)
		}
	})

	t.Run(`Get ignored user`, func(t *testing.T) {
		res, err := repo.IgnoredUser.Get(ignoredUsers[0])
		assert.Equal(t, res.Username, ignoredUsers[0])
		assert.Equal(t, res.IgnoredAt, mockDate.Unix())
		assert.NoError(t, err)
	})

	t.Run(`Get all ignored users`, func(t *testing.T) {
		res, err := repo.IgnoredUser.GetAll()
		for i, sub := range ignoredUsers {
			assert.Equal(t, res[i].Username, sub)
		}
		assert.NoError(t, err)
	})
}

func TestIgnoredUserHandler_NoResults(t *testing.T) {
	db := setupDb(t)
	sqlDb, _ := db.DB()
	defer func(sqlDb *sql.DB) {
		_ = sqlDb.Close()
	}(sqlDb)

	data := repository{
		db:    db,
		clock: &fakeClock{},
	}
	repo := Repository{
		IgnoredUser: &ignoredUserHandler{data},
	}

	t.Run(`Get ignored user`, func(t *testing.T) {
		res, err := repo.IgnoredUser.Get(ignoredUsers[0])
		assert.Equal(t, res, &IgnoredUser{})
		assert.NoError(t, err)
	})

	t.Run(`Get all ignored users`, func(t *testing.T) {
		res, err := repo.IgnoredUser.GetAll()
		assert.Equal(t, res, []IgnoredUser{})
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
	_ = db.AutoMigrate(&BannedSubreddit{}, &IgnoredUser{})
	return db
}
