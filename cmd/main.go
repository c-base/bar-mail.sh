package main

import (
	_ "time/tzdata"

	"github.com/c-base/bar-mail/pkg/barmail"
)

func main() {
	err := barmail.GetBarMail()
	if err != nil {
		panic(err)
	}
}
