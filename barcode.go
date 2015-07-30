package barcode

import (
	"gobarcode/reader"
	"log"
	"os/exec"
	"strings"
)

var OUTCHAN chan string

func init() {
	OUTCHAN = make(chan string, 1)
}

func Run(doneChan chan string) {
	options := []string{
		"--nodisplay",
		"/dev/video0",
	}
	cmd := exec.Command("zbarcam", options...)

	// get the stdout
	processor := reader.New()
	processor.SetOutChannel(OUTCHAN)
	cmd.Stdout = processor

	// run the listener
	go listenAndPipe(OUTCHAN, doneChan)

	// run
	err := cmd.Run()
	if err != nil {
		log.Fatal("zbarcam process failed: ", err)
	}
}

func listenAndPipe(outCh <-chan string, doneCh chan<- string) {
	for {
		data := <-outCh
		splitted := strings.Split(strings.TrimSpace(data), ":")
		doneCh <- strings.Join(splitted[1:], ":")
	}
}
