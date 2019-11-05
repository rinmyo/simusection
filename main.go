package main

import "fmt"

func main() {
	testLine := NewLine()
	testLine.NewTrain("列車1", 3)
	testLine.NewTrain("列車ee", 2)
	testLine.NewTrain("列車e1", 10)
	testLine.NewTrain("列車341", 19)
	testLine.NewTrain("列車133", 50)
	testLine.NewTrain("列車09", 60)

	fmt.Println(testLine.toString())
	fmt.Println(testLine)
}
