package main

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
	"strings"
	"unsafe"
)

func main() {

	RunAndPrint("pidof bash")
	RunAndPrint("pidof java")
	return

	// check to see if program is running
	out, code := Run("cat fakeproc")
	fmt.Printf("exited with code: %d\n", code)
	if strings.Contains(out, "12345") {
		// program is running, okay to call stop
		out, code = Run("bash stopme.sh")
		fmt.Printf("exited with code: %d\n", code)
	}
}

func Run(s string) (string, int) {
	name, args := sanitizeInput(s)
	cmd := exec.Command(name, args...)
	out, err := cmd.Output()
	var code int = 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			log.Printf("recieved bad exit status: %v\n", exitErr.Error())
			code = exitErr.ExitCode()
		}
		log.Printf("encountered error: %v\n", err)
	}
	return BytesToString(out), code
}

func RunAndPrint(s string) {
	out, code := Run(s)
	fmt.Printf("> %s\n%s\nexited %d\n", s, out, code)
}

func sanitizeInput(s string) (string, []string) {
	if s != "" {
		ss := strings.Split(s, " ")
		if len(ss) == 1 {
			return ss[0], nil
		}
		if len(ss) > 1 {
			return ss[0], ss[1:]
		}
	}
	return "echo", []string{"'encountered error, not enough arguments!'"}
}

func BytesToString(bytes []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}))
}

func StringToBytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}))
}
