package problems

import (
	"testing"
)

func TestOnlineElections(t *testing.T) {
	persons := []int{0, 1, 1, 0, 0, 1, 0}
	times := []int{0, 5, 10, 15, 20, 25, 30}

	testCases := []struct {
		input    int
		expected int
	}{
		//{3, 0},
		//{12, 1},
		{25, 1},
		//{15, 0},
		//{24, 0},
		//{8, 1},
	}

	topVotedCandidate := Constructor(persons, times)

	for _, tc := range testCases {
		res := topVotedCandidate.Q(tc.input)
		if res != tc.expected {
			t.Errorf("input: %v. expected %v, got %v", tc.input, tc.expected, res)
		}
	}
}

type TopVotedCandidate struct {
	leaders map[int]int
	times   []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {

	leaderScores := make(map[int]int)
	for _, p := range persons {
		leaderScores[p] = 0
	}

	leaders := make(map[int]int, len(times))

	var leader int

	for p, t := range times {
		votedPerson := persons[p]
		leaderScores[votedPerson]++
		if leaderScores[votedPerson] >= leaderScores[leader] {
			leader = votedPerson
		}

		leaders[t] = leader
	}

	return TopVotedCandidate{leaders: leaders, times: times}
}

func (tvc *TopVotedCandidate) Q(time int) int {
	var lastVotedTime int
	for _, t := range tvc.times {
		if t <= time {
			lastVotedTime = t
		} else {
			break
		}
	}

	return tvc.leaders[lastVotedTime]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
