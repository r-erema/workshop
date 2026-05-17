package designtwitter_test

import (
	"testing"

	designtwitter "github.com/r-erema/workshop/algorithms/design_twitter"
	"github.com/stretchr/testify/assert"
)

func TestDesignTwitter(t *testing.T) {
	t.Parallel()

	twitter := designtwitter.Constructor()
	twitter.PostTweet(1, 5)
	feed := twitter.GetNewsFeed(1)
	assert.Equal(t, []int{5}, feed)
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	feed = twitter.GetNewsFeed(1)
	assert.Equal(t, []int{6, 5}, feed)
	twitter.Unfollow(1, 2)
	feed = twitter.GetNewsFeed(1)
	assert.Equal(t, []int{5}, feed)
}
