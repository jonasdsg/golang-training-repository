package util

import "log"

func ErrorHandler() {
	//Handle an errer when it occurs
	err := recover()
	if err != nil {
		log.Fatal("Something went wrong..", err)
	}
}
