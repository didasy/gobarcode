# gobarcode
A barcode and qr-code reader using libzbar.

You need to install `libzbar-dev` and `zbar-tools` before using this package.

### Example
-----------------

```
package main

import (
	"fmt"
	"github.com/JesusIslam/gobarcode"
)

func main() {
	strChan := make(chan string)
	gobarcode.Run(strChan)

	for {
		str := <-strChan
		fmt.Println(str)
	}
}
```