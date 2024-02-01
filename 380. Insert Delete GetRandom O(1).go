package main

import (
	"math/rand"
)

type RandomizedSet struct {
	m map[int]int
	l []int
}

func Constructor() RandomizedSet {
	m := make(map[int]int)
	l := make([]int, 0)
	return RandomizedSet{m, l}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.m[val]; ok {
		return false
	}
	s.m[val] = len(s.l)
	s.l = append(s.l, val)
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	if _, ok := s.m[val]; !ok {
		return false
	}
	n, i := len(s.l), s.m[val]
	s.l[i], s.l[n-1] = s.l[n-1], s.l[i]
	s.m[s.l[i]] = i
	s.l = s.l[:n-1]
	delete(s.m, val)
	return true
}

func (s *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(s.l))
	return s.l[i]
}
