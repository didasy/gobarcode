# gobarcode
### A barcode and qr-code reader using libzbar.
-----------------------------------------------

[![GoDoc](https://godoc.org/github.com/JesusIslam/gobarcode?status.svg)](https://godoc.org/github.com/JesusIslam/gobarcode)

You need to install `go get github.com/galaktor/gostwriter`, `libzbar-dev` and `zbar-tools` before using this package.

### Example
-----------------

```
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

```