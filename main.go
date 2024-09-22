package main

import (
	"fmt"
	"time"
)

func cliReset() {
	fmt.Print("\033[0m")
}

func cliPrintInGreen() {
	fmt.Print("\033[32m")
}

func cliClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func cliPrintInRed() {
	fmt.Print("\033[31m")
}

func printProgressLine(start int, end int) {
	for j := start; j < end; j++ {
		fmt.Print("-")
	}
}

type Loader struct {
	progress      int
	total         int
	loaderMessage string
}

func (l *Loader) PrintTotal() {
	cliReset()
	cliPrintInRed()
	printProgressLine(l.progress, l.total)
}

func (l *Loader) PrintLoaderMessage() {
	cliReset()
	cliPrintInGreen()
	percentage := float64(l.progress) / float64(l.total)
	fmt.Print(fmt.Sprintf("%s %.2f", l.loaderMessage, percentage*100), "%: ")
}

func (l *Loader) PrintProgress() {
	cliReset()
	cliPrintInGreen()
	printProgressLine(0, l.progress)
}

func (l *Loader) MakeProgress() {
	l.progress = l.progress + 1
}

func (l *Loader) StopLoading() bool {
	return l.progress > l.total
}

func load(l *Loader, channel chan bool, cb func()) {

	for !l.StopLoading() {
		cliClearScreen()
		l.PrintLoaderMessage()
		l.PrintProgress()
		l.PrintTotal()
		cb()

	}
	channel <- true
}

func main() {
	loader := Loader{
		progress:      0,
		total:         150,
		loaderMessage: "Loading",
	}
	channel := make(chan bool)
	fmt.Println("Started loading")
	go load(&loader, channel, func() {
		loader.MakeProgress()
		time.Sleep(50 * time.Millisecond)
	})
	<-channel
	fmt.Println("Finished Loading")
}
