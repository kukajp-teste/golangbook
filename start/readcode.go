package main

import (
	"bufio"
	"os"
)

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/Users/samchen/Music/text.txt")
	defer f.Close()
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		n, _ := r.Read(buf)
		if n == 0 {
			break
		}
		w.Write(buf[0:n])
	}
}
