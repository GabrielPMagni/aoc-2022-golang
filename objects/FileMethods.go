package objects

import (
	"aoc/go2022/shared"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type MethodHandler interface {
	Execute(inputFile *os.File)
	GetAnswer() interface{}
}

type GetBiggerCalloriesCount struct {
	value int
}

type Get3BiggerColloriesCount struct {
}

func (gbcc *GetBiggerCalloriesCount) Execute(inputFile *os.File) {
	// elfArray := []int{}
	counter := 0
	file := bufio.NewScanner(inputFile)
	file.Split(bufio.ScanLines)
	for file.Scan() {
		content := file.Text()
		if shared.CheckRegex(content, `^[\r?\n]+$`) {
			fmt.Println("A")

		} else if shared.CheckRegex(content, `^\d+$`) {
			val, err := strconv.Atoi(content)
			shared.CheckError(err)
			counter += val
			fmt.Println("B")
		} else {
			fmt.Println("Pattern not found")
		}

	}
	gbcc.value = 100
}

func (gbcc *GetBiggerCalloriesCount) GetAnswer() interface{} {
	return gbcc.value
}

func (Get3BiggerColloriesCount) Execute(inputFile *os.File) {
}
