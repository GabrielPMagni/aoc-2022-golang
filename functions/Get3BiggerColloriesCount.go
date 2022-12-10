package functions

import (
	"aoc/go2022/shared"
	"bufio"
	"os"
	"sort"
	"strconv"
)

type Get3BiggerColloriesCount struct {
	value int
}

func (gbcc *Get3BiggerColloriesCount) Execute(inputFile *os.File) {
	elfArray := []int{}
	calloriesCounter := 0
	file := bufio.NewScanner(inputFile)
	file.Split(bufio.ScanLines)
	for file.Scan() {
		content := file.Text()
		if shared.CheckRegex(content, `^\d+$`) {
			callories, err := strconv.Atoi(content)
			shared.CheckError(err)
			calloriesCounter += callories
		} else {
			elfArray = append(elfArray, calloriesCounter)
			calloriesCounter = 0
		}
	}
	inputFile.Close()
	sort.Ints(elfArray)
	calloriesCounter = 0
	for i := range elfArray {
		calloriesCounter += elfArray[len(elfArray)-(i+1)]
		if i == 2 {
			break
		}
	}
	gbcc.value = calloriesCounter
}

func (gbcc *Get3BiggerColloriesCount) GetAnswer() interface{} {
	return gbcc.value
}
