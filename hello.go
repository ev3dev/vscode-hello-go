package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func resetConsole() {
	// escape code that resets the console
	fmt.Print("\x1Bc")
}

func setCursor(on bool) {
	// use escape code to turn cursor on or off
	if on {
		fmt.Print("\x1B[?25h")
	} else {
		fmt.Print("\x1B[?25l")
	}
}

func setFont(name string) {
	// use setfont command to change the font
	// run `ls /usr/share/consolefonts` in a terminal to get a list of available fonts
	exec.Command("setfont", name).Run()
}

func main() {
	// redirect log to stderr so that it prints messages in VS Code
	log.SetOutput(os.Stderr)

	// configure the console
	resetConsole()
	setFont("Lat15-Terminus16")
	setCursor(false)

	// print some messages
	fmt.Println("Hello EV3!")
	log.Println("Hello VS Code!")

	// give some time too look at the screen before the program exits
	time.Sleep(time.Second * 5)
}
