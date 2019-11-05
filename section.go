package main

import (
	"sort"
	"strconv"
)

const (
	L = iota
	LU
	U
	HU
)

type Line []Train

type Train struct {
	id               string //車次
	occupyingSection int
} //佔用一個分區的火車

func NewLine() Line {
	return make(Line, 0)
}

func (l *Line) NewTrain(id string, sect int) {
	*l = append(*l, Train{id: id, occupyingSection: sect})
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
	for i := len(l) - 1; i >= 0; i-- {
		for j := 1; j < 4 && l[i].occupyingSection >= j+1; j++ {
			result[l[i].occupyingSection-1-j] = 4 - j
		}
	}
	return
}

func (l Line) toString() (result []string) {
	sig := l.GetSignal()
	result = make([]string, len(sig))
	for i := len(result) - 1; i >= 0; i-- {

		switch sig[i] {
		case LU:
			result[i] = "LU"
			break
		case U:
			result[i] = "U"
			break
		case HU:
			result[i] = "HU"
			break
		case L:
			//首個車的區間
			if i == len(result)-1 {
				result[i] = "L5"
				break
			}

			//前一區間爲綠黃碼
			if sig[i+1] == LU {
				result[i] = "L"
				break
			}

			//前一區間爲綠碼
			if sig[i+1] == L {
				f := 2
				if sig[i+2] == L {
					f = 1
				}
				if result[i+1] == "L5" {
					f = 0
				}
				preL, _ := strconv.Atoi(result[i+1][1:])
				result[i] = "L" + strconv.Itoa(preL+f)
				break
			}
		}
	}
	return
}
