package echo

import (
	"testing"
)

var ARGS = []string{"hello", "world"}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(ARGS)
	}
}

/*
    5341            244151 ns/op
PASS
ok      the-go-programming-language-notes/chapter1/exercise1.3  2.022s
*/

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(ARGS)
	}
}

/*
    5929            222979 ns/op
PASS
ok      the-go-programming-language-notes/chapter1/exercise1.3  2.298s
*/

/*
go env -w GO111MODULE=auto
go test -bench=Echo1
go test -bench=Echo3
*/
