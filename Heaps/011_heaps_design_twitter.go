package heaps

/*
Create a Twitter tweets
Optmizations idea:
- When a tweet is posted we have the time, we can pick the next highest and solve our problem
- This next highest meaning the next latest in time as we increment for every post
*/

type Tweet struct {
	userId  int
	tweetId int
	time    int
}

type Relation struct {
	followerId int
	followeeId int
}

type Twitter struct {
	tweets    []Tweet
	relations []Relation
	time      int
}

type TwitterMaxHeap struct {
	nodes []Tweet
}

func (mh *TwitterMaxHeap) Insert(key Tweet) {
	mh.nodes = append(mh.nodes, key)

	idx := len(mh.nodes) - 1

	for idx > 0 {
		parentIdx := (idx - 1) / 2

		if parentIdx >= 0 && mh.nodes[parentIdx].time < mh.nodes[idx].time {
			mh.nodes[parentIdx], mh.nodes[idx] = mh.nodes[idx], mh.nodes[parentIdx]
		}
		idx = parentIdx
	}
}

func (mh *TwitterMaxHeap) Pop() Tweet {
	peekElement := mh.nodes[0]
	mh.nodes[0] = mh.nodes[len(mh.nodes)-1]
	mh.nodes = mh.nodes[:len(mh.nodes)-1]

	idx := 0

	for true {
		leftChild := 2*idx + 1
		rightChild := 2*idx + 2
		largest := idx

		if leftChild < len(mh.nodes) && mh.nodes[leftChild].time > mh.nodes[largest].time {
			largest = leftChild
		}

		if rightChild < len(mh.nodes) && mh.nodes[rightChild].time > mh.nodes[largest].time {
			largest = rightChild
		}

		if largest != idx {
			mh.nodes[largest], mh.nodes[idx] = mh.nodes[idx], mh.nodes[largest]
			idx = largest
		} else {
			break
		}

	}

	return peekElement
}

func (mh *TwitterMaxHeap) Peek() Tweet {
	return mh.nodes[0]
}

func (mh *TwitterMaxHeap) Length() int {
	return len(mh.nodes)
}

func Constructor() Twitter {
	return Twitter{
		tweets:    []Tweet{},
		relations: []Relation{},
		time:      0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets = append(this.tweets, Tweet{userId: userId, tweetId: tweetId, time: this.time})
	this.time += 1
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	results := []int{}
	maxHeap := TwitterMaxHeap{}

	userFollowees := make(map[int]bool)
	for _, val := range this.relations {
		if val.followerId == userId {
			userFollowees[val.followeeId] = true
		}
	}

	for _, val := range this.tweets {
		if _, ok := userFollowees[val.userId]; ok || val.userId == userId {
			maxHeap.Insert(val)
		}
	}

	count := 10

	for count > 0 && maxHeap.Length() > 0 {
		tweet := maxHeap.Pop()
		if _, ok := userFollowees[tweet.userId]; ok {
			results = append(results, tweet.tweetId)
			count -= 1
		}
		if tweet.userId == userId {
			results = append(results, tweet.tweetId)
			count -= 1
		}
	}

	return results
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	this.relations = append(this.relations, Relation{followerId: followerId, followeeId: followeeId})
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	resIdx := -1
	for idx, val := range this.relations {
		if val.followerId == followerId && val.followeeId == followeeId {
			resIdx = idx
			break
		}
	}

	if resIdx != -1 {
		this.relations = append(this.relations[:resIdx], this.relations[resIdx+1:]...)
	}

}
