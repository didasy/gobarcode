# gobarcode
A barcode and qr-code reader using libzbar.

You need to install `libzbar-dev` and `zbar-tools` before using this package.

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