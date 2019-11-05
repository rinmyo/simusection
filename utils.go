package main

import (
	"sort"
)

const (
	L = iota
	LU
	U
	HU
	H
)

type Line []*Train

type Train struct {
	id               string //車次
	occupyingSection int
} //佔用一個分區的火車

func NewLine() Line {
	return make(Line, 0)
}

func (l *Line) NewTrain(id string, sect int) {
	*l = append(*l, &Train{id: id, occupyingSection: sect})
	sort.Sort(*l)
}

func (l Line) Len() int {
	return len(l)
}

func (l Line) Less(i, j int) bool {
	return l[i].occupyingSection < l[j].occupyingSection
}

func (l Line) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l Line) GetSignal() (result []int) {
	result = make([]int, l[len(l)-1].occupyingSection)
	for i := len(l)-1; i >= 0; i-- {
		for j := 0; j<4 ;j++ {
			result[l[i].occupyingSection - 1- j] = j
		}
	}
	return
}
