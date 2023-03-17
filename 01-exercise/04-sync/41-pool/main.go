package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

//TODO: create pool of bytes.Buffers which can be reused.

// https://pkg.go.dev/sync#Pool
var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, val string) {
	// var b bytes.Buffer
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset() // Must call this, otherwise debug-string1 is printed again ( which shows the reuse of the buffer )

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	w.Write(b.Bytes())
	bufPool.Put(b)
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}
