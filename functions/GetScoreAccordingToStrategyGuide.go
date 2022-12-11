package functions

import (
	"bufio"
	"os"
	"strings"
)

type GetScoreAccordingToStrategyGuide struct {
	value uint32
}

type scorePoint uint32

const (
	gameLostPoints scorePoint = 0
	gameDrawPoints scorePoint = 3
	gameWonPoints  scorePoint = 6
)

type gameConstants struct {
	QuestionTag string
	AnswerTag   string
	points      scorePoint
}

type gameRuleConstants struct {
	item   gameConstants
	winOf  gameConstants
	loseOf gameConstants
}

func (gbcc *GetScoreAccordingToStrategyGuide) Execute(inputFile *os.File) {
	file := bufio.NewScanner(inputFile)
	file.Split(bufio.ScanLines)
	var totalScore scorePoint = 0
	for file.Scan() {
		content := strings.Split(file.Text(), " ")
		question, answer := getPlayedConstants(content[0], content[1])
		totalScore += verifyVictory(question, answer)
	}
	gbcc.value = uint32(totalScore)
}

func verifyVictory(question, answer gameRuleConstants) scorePoint {
	if question.winOf == answer.loseOf {
		return gameWonPoints + answer.item.points
	} else if question.loseOf == answer.winOf {
		return gameLostPoints + answer.item.points
	} else {
		return gameDrawPoints + answer.item.points
	}
}

func getPlayedConstants(question, answer string) (gameRuleConstants, gameRuleConstants) {
	Rock := gameConstants{"A", "X", 1}
	Paper := gameConstants{"B", "Y", 2}
	Scissor := gameConstants{"C", "Z", 3}

	RuleRock := gameRuleConstants{Rock, Scissor, Paper}
	RulePaper := gameRuleConstants{Paper, Rock, Scissor}
	RuleScissor := gameRuleConstants{Scissor, Paper, Rock}

	Rules := []gameRuleConstants{RuleRock, RulePaper, RuleScissor}

	var returnedQuestion gameRuleConstants
	var returnedAnswer gameRuleConstants
	for _, item := range Rules {
		if item.item.QuestionTag == question {
			returnedQuestion = item
		}
	}

	for _, item := range Rules {
		if item.item.AnswerTag == answer {
			returnedAnswer = item
		}
	}
	return returnedQuestion, returnedAnswer
}

func (gbcc *GetScoreAccordingToStrategyGuide) GetAnswer() interface{} {
	return gbcc.value
}
