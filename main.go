package main

import "fmt"

func main() {
	testLine := NewLine()
	testLine.NewTrain("打法", 6)
	testLine.NewTrain("ddd", 10)
	testLine.NewTrain("aaa", 20)


	for _,x := range testLine.GetSignal() {
		fmt.Println(x)
	}
}
