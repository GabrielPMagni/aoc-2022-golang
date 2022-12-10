package shared

import (
	"aoc/go2022/shared"
	"os"
)

type MethodHandler interface {
	Execute(inputFile *os.File)
	GetAnswer() interface{}
}

func GetFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	shared.CheckError(err)
	return file
}

func HandleFile(inputFile *os.File, handlingMethod MethodHandler) interface{} {
	handlingMethod.Execute(inputFile)
	return handlingMethod.GetAnswer()
}
