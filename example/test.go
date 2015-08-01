package main

import (
	"../."
	"../keyboard"
)

func main() {
	strChan := make(chan string)
	go gobarcode.Run(strChan)

	kb, err := keyboard.New("scan", 10)
	if err != nil {
		panic(err)
	}
	defer kb.Destroy()

	for {
		str := <-strChan
		err = kb.Type(str)
		if err != nil {
			panic(err)
		}
	}
}
