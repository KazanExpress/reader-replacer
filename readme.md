# Reader Replacer

Replace bytes token from `io.Reader`.

## Usage

```golang
package main

import (
	"bytes"
	"io"
	"os"

	replacer "github.com/KazanExpress/reader-replacer/pkg"
)

func main() {
	var s = "hello world ozon hola some ozon top ozon"
	var reader = replacer.NewReaderReplacer(bytes.NewBufferString(s), []byte("ozon"), []byte("kazanexpress"))
	io.Copy(os.Stdout, reader)

}

```

Prints:

`hello world kazanexpress hola some kazanexpress top kazanexpress`