package functions

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type MethodHandler interface {
	Execute(inputFile *os.File)
	GetAnswer() interface{}
}

func GetFileLines(filePath string) []string {
	checkErr := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	fp, err := os.Open(filePath)
	checkErr(err)

	raw, err := ioutil.ReadAll(fp)
	checkErr(err)

	return strings.Split(string(raw), "\n")
}

func HandleFile(inputFile *os.File, handlingMethod MethodHandler) interface{} {
	handlingMethod.Execute(inputFile)
	return handlingMethod.GetAnswer()
}
