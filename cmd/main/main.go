package main

import (
	"aoc/go2022/functions"
	sharedFunctions "aoc/go2022/functions/shared"
	"fmt"
)

func question1() {
	file := sharedFunctions.GetFile("inputs/input1.txt")
	result := sharedFunctions.HandleFile(file, &functions.GetBiggerCalloriesCount{})
	fmt.Println(result)
}

func question2() {
	file := sharedFunctions.GetFile("inputs/input1.txt")
	result := sharedFunctions.HandleFile(file, &functions.Get3BiggerColloriesCount{})
	fmt.Println(result)
}

func question3() {
	file := sharedFunctions.GetFile("inputs/input2.txt")
	result := sharedFunctions.HandleFile(file, &functions.GetScoreAccordingToStrategyGuide{})
	fmt.Println(result)
}

func main() {
	fmt.Println("Question 1")
	question1()
	fmt.Println("Question 2")
	question2()
	fmt.Println("Question 3")
	question3()
}
