package _interface

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func uniq(reader io.Reader, writer io.Writer) error {
	in := bufio.NewScanner(reader)
	var prev string
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("file not sorted")
		}
		prev = txt
		fmt.Fprintln(writer, txt)
	}
	return nil
}

func main() {
	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}
