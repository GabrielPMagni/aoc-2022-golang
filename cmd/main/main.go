package main

import (
	"aoc/go2022/objects"
	"aoc/go2022/shared"
	"fmt"
	"os"
)

func getFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	shared.CheckError(err)
	return file
}

func handleFile(inputFile *os.File, handlingMethod objects.MethodHandler) interface{} {
	handlingMethod.Execute(inputFile)
	return handlingMethod.GetAnswer()
}

func question1() {
	file := getFile("inputs/input1.txt")
	result := handleFile(file, &objects.GetBiggerCalloriesCount{})
	fmt.Println(result)
}

func main() {
	question1()
}
