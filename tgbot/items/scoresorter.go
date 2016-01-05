package items

import (
	"sort"
)

type ScoreSorter struct {
	userScores map[int]int64
	results    []int
}

// NewScoreSorter gives a new ScoreSorter object.
func NewScoreSorter(scores map[int]int64) *ScoreSorter {
	s := new(ScoreSorter)
	s.userScores = scores
	s.results = make([]int, len(s.userScores))

	i := 0
	for key := range s.userScores {
		s.results[i] = key
		i++
	}
	return s
}

// Len returns the length of users in the sorter.
func (s *ScoreSorter) Len() int {
	return len(s.userScores)
}

// Less checks if the user i has less score than user j.
func (s *ScoreSorter) Less(i, j int) bool {
	return s.userScores[s.results[i]] < s.userScores[s.results[j]]
}

// Swap will swap user i with user j in the scoreing list.
func (s *ScoreSorter) Swap(i, j int) {
	s.results[i], s.results[j] = s.results[j], s.results[i]
}

// Sorted will return the keys of the sorted score list.
func (s *ScoreSorter) Sorted() []int {
	sort.Sort(s) // might need stable instead of sort?

	return s.results
}
