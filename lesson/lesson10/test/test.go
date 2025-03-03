package test

import ans "lesson10/test/answer"

func (ans.AnswerB) Show() {}

func MakeTest(question string, answer string) bool {
	if question == "a" {
		if answer == ans.AnswerA {
			return true
		}
	}
	return false
}
