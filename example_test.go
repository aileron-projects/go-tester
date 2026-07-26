package tester_test

import (
	"fmt"
	"strings"

	"github.com/aileron-projects/go-tester"
)

func ExampleNewMaxReader() {
	s := strings.NewReader("1234567890")
	r := tester.NewMaxReader(s, 5, tester.ErrMaxRead)

	buf := make([]byte, 3)
	fmt.Println(r.Read(buf))
	fmt.Println(r.Read(buf))
	fmt.Println(r.Read(buf))
	// Output:
	// 3 <nil>
	// 2 go-tester/tester: max readable byte reached
	// 0 go-tester/tester: max readable byte reached
}

func ExampleNewMaxWriter() {
	w := tester.NewMaxWriter(5, tester.ErrMaxWritten)

	fmt.Println(w.Write([]byte("123")))
	fmt.Println(w.Write([]byte("456")))
	fmt.Println(w.Write([]byte("789")))
	// Output:
	// 3 <nil>
	// 2 go-tester/tester: max writable byte reached
	// 0 go-tester/tester: max writable byte reached
}
