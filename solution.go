package main

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	return emoji.Sprint("Hello \U0001F5FA !")
}
