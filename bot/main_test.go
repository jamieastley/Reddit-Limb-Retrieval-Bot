package main

//
//import (
//	"github.com/turnage/graw"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func TestCheckContainsShrug(t *testing.T) {
//
//	t.Run("unescaped left arm", func(t *testing.T) {
//		assert.Equal(t, MissingLeftArmPattern, checkContainsShrug(`¯\_(ツ)_/¯`, ""))
//		assert.Equal(t, MissingLeftArmPattern, checkContainsShrug(`asd ¯\_(ツ)_/¯ some text`, ""))
//		assert.Equal(t, MissingLeftArmPattern, checkContainsShrug(`¯\_(ツ)_/¯asdasd`, ""))
//		assert.Equal(t, MissingLeftArmPattern, checkContainsShrug(`asdasd¯\_(ツ)_/¯asdasd`, ""))
//	})
//
//	t.Run("missing shoulders", func(t *testing.T) {
//		assert.Equal(t, MissingShouldersPattern, checkContainsShrug(`¯\\_(ツ)_/¯`, ""))
//		assert.Equal(t, MissingShouldersPattern, checkContainsShrug(`asd ¯\\_(ツ)_/¯ some text`, ""))
//		assert.Equal(t, MissingShouldersPattern, checkContainsShrug(`¯\\_(ツ)_/¯asdasd`, ""))
//		assert.Equal(t, MissingShouldersPattern, checkContainsShrug(`asdasd¯\\_(ツ)_/¯asdasd`, ""))
//	})
//}
//
//type MockLimbBot struct {
//	bot MockBot
//}
//
//type MockBot struct {
//}
//
//func TestInitBot(t *testing.T) {
//	var mockBot = &MockLimbBot{MockBot{}}
//	cfg := graw.Config{SubredditComments: []string{"News"}}
//	initBot(mockBot, cfg)
//}
