package functions

import (
	"aoc/go2022/shared"
	"bufio"
	"os"
	"strconv"
)

type GetBiggerCalloriesCount struct {
	value int
}

func (gbcc *GetBiggerCalloriesCount) Execute(inputFile *os.File) {
	biggerCalories := 0
	calloriesCounter := 0
	index := 0
	file := bufio.NewScanner(inputFile)
	file.Split(bufio.ScanLines)
	for file.Scan() {
		content := file.Text()
		if shared.CheckRegex(content, `^\d+$`) {
			callories, err := strconv.Atoi(content)
			shared.CheckError(err)
			calloriesCounter += callories
		} else {
			if calloriesCounter > biggerCalories || index == 0 {
				biggerCalories = calloriesCounter
			}
			calloriesCounter = 0
			index++
		}
	}
	inputFile.Close()
	gbcc.value = biggerCalories
}

func (gbcc *GetBiggerCalloriesCount) GetAnswer() interface{} {
	return gbcc.value
}
