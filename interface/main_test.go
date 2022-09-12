package _interface

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
3
4
5`

var testOkRes = `1
2
3
4
5
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err != nil {
		t.Errorf("test for OK failed")
	}
	result := out.String()
	if result != testOkRes {
		t.Errorf("test for OK failed - results don't match\n %v %v", result, testOkRes)
	}
}

var testFail = `1
2
1
`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err == nil {
		t.Errorf("test for Error failed: %v", err)
	}
}
