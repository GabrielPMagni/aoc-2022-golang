package shared

import "regexp"

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckRegex(compareThis string, withThis string) bool {
	regex, err := regexp.Compile(withThis)
	CheckError(err)
	return regex.MatchString(compareThis)
}
