package main
import (
	"luna/bot"
)

func main() {
	bot.Start()
	<-make(chan struct{})
	return
}