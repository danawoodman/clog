package clog

import (
	"fmt"
	"time"

	"github.com/agtorre/gocolorize"
)

func Log(label string, msg string, color string) {
	plain := gocolorize.NewColor(color).Paint
	bold := gocolorize.NewColor(color + "+b").Paint
	t := time.Now().Format("Jan 2, 2006 at 3:04pm")
	l := "[" + bold(label) + "]"
	fmt.Println(l, "\t", plain(t), "\t", plain(msg))
}

func Info(msg string) {
	Log("info", msg, "cyan")
}

func Success(msg string) {
	Log("ok", msg, "green")
}

func Warn(msg string) {
	Log("warn", msg, "yellow")
}

func Error(msg string) {
	Log("err", msg, "red")
}
