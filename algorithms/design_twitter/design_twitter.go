package designtwitter

import (
	"slices"
)

type feed []tweet

type tweet struct {
	order, uid, tid int
}

type Twitter struct {
	autoIncrement     int
	userFeeds         map[int]*feed
	tweets            []tweet
	userFollowers     map[int]map[int]struct{}
	userSubscriptions map[int]map[int]struct{}
}

func Constructor() Twitter {
	return Twitter{
		userFeeds:         make(map[int]*feed),
		userFollowers:     make(map[int]map[int]struct{}),
		userSubscriptions: make(map[int]map[int]struct{}),
	}
}

// PostTweet posts a tweet for a user.
// Time O(M*logN), where M is number of followers and N is tweets in feed(heap insertion is logN)
// Space O(M*N), M is users and N is total tweets.
func (tw *Twitter) PostTweet(userId int, tweetId int) {
	twt := tweet{tw.autoIncrement, userId, tweetId}
	tw.autoIncrement++

	tw.tweets = append(tw.tweets, twt)

	if _, ok := tw.userFeeds[userId]; !ok {
		tw.userFeeds[userId] = &feed{}
		Heapify(*tw.userFeeds[userId])
	}

	Push(tw.userFeeds[userId], twt)

	for followerID := range tw.userFollowers[userId] {
		Push(tw.userFeeds[followerID], twt)
	}
}

// GetNewsFeed retrieves the 10 most recent tweet ids for a user.
// Time O(1), since the heap popping
// Space O(1), since the fixed feed length.
func (tw *Twitter) GetNewsFeed(userId int) []int {
	var res []int

	if tw.userFeeds[userId] == nil {
		return res
	}

	if f := tw.userFeeds[userId]; f != nil {
		tmp := slices.Clone(*f)

		for len(tmp) > 0 && len(res) < 10 {
			t := Pop(&tmp)
			res = append(res, t.tid)
		}
	}

	return res
}

// Follow makes followerId follow followeeId.
// Time O(M*logN), where M is number of followees and N is tweets in feed(heap insertion is logN)
// Space O(M*N), M is followees and N is followee tweets.
func (tw *Twitter) Follow(followerId int, followeeId int) {
	if tw.userFollowers[followeeId] == nil {
		tw.userFollowers[followeeId] = make(map[int]struct{})
	}

	tw.userFollowers[followeeId][followerId] = struct{}{}

	if tw.userSubscriptions[followerId] == nil {
		tw.userSubscriptions[followerId] = make(map[int]struct{})
	}

	tw.userSubscriptions[followerId][followeeId] = struct{}{}

	tw.refreshFeed(followerId)
}

// Unfollow makes followerId unfollow followeeId.
// Time O(M*logN), where M is number of followees and N is tweets in feed(heap insertion is logN)
// Space O(M*N), M is followees and N is followee tweets.
func (tw *Twitter) Unfollow(followerId int, followeeId int) {
	delete(tw.userFollowers[followeeId], followerId)
	delete(tw.userSubscriptions[followerId], followeeId)
	tw.refreshFeed(followerId)
}

func (tw *Twitter) refreshFeed(userId int) {
	subscription := tw.userSubscriptions[userId]

	feed := new(feed)

	for i := range tw.tweets {
		if _, ok := subscription[tw.tweets[i].uid]; ok || tw.tweets[i].uid == userId {
			Push(feed, tw.tweets[i])
		}
	}

	tw.userFeeds[userId] = feed
}

func Heapify(arr []tweet) {
	for i := range arr {
		PercolateUp(arr[:i+1])
	}
}

func Push(heap *feed, node tweet) {
	*heap = append(*heap, node)
	PercolateUp(*heap)
}

func Pop(heap *feed) tweet {
	popped := (*heap)[0]
	(*heap)[0], *heap = (*heap)[len(*heap)-1], (*heap)[:len(*heap)-1]

	PercolateDown(*heap)

	return popped
}

func PercolateUp(arr []tweet) {
	const minHeapSize = 2

	if len(arr) < minHeapSize {
		return
	}

	parentIndex := (len(arr) - 1 - 1) / minHeapSize

	if arr[parentIndex].order > arr[len(arr)-1].order {
		return
	}

	arr[parentIndex], arr[len(arr)-1] = arr[len(arr)-1], arr[parentIndex]

	PercolateUp(arr[:parentIndex+1])
}

func PercolateDown(arr []tweet) {
	currentNodeIndex := 0

	var bfs func()

	const (
		leftChildOffset  = 1
		rightChildOffset = 2
		multiplier       = 2
	)

	bfs = func() {
		leftChildIndex := currentNodeIndex*multiplier + leftChildOffset
		rightChildIndex := currentNodeIndex*multiplier + rightChildOffset

		if leftChildIndex < len(arr) && arr[leftChildIndex].order > arr[currentNodeIndex].order {
			if rightChildIndex < len(arr) && arr[rightChildIndex].order > arr[leftChildIndex].order {
				arr[rightChildIndex], arr[currentNodeIndex] = arr[currentNodeIndex], arr[rightChildIndex]

				currentNodeIndex = rightChildIndex

				bfs()
			} else {
				arr[leftChildIndex], arr[currentNodeIndex] = arr[currentNodeIndex], arr[leftChildIndex]

				currentNodeIndex = leftChildIndex

				bfs()
			}

			return
		}

		if rightChildIndex < len(arr) && arr[rightChildIndex].order > arr[currentNodeIndex].order {
			arr[rightChildIndex], arr[currentNodeIndex] = arr[currentNodeIndex], arr[rightChildIndex]

			currentNodeIndex = rightChildIndex

			bfs()
		}
	}

	bfs()
}
