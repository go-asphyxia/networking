package query

import (
	"log"
	"testing"

	"github.com/go-asphyxia/core/bytes"
	"github.com/valyala/fasthttp"
)

var (
	simple       = "one=two&two=three"
	simpleBuffer = bytes.Buffer(simple)
)

func TestQuery(t *testing.T) {
	query := Decode(simpleBuffer)

	log.Println(simple)
	log.Println(Encode(query).String())
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Decode(simpleBuffer)
	}
}

func BenchmarkDecodeFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := fasthttp.Args{}
		a.Parse(simple)
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := Query{
			Argument{bytes.Buffer("one"), bytes.Buffer("two")},
			Argument{bytes.Buffer("two"), bytes.Buffer("three")},
		}

		_ = Encode(a)
	}
}

func BenchmarkEncodeFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := fasthttp.Args{}
		a.Set("one", "two")
		a.Set("two", "three")

		_ = a.String()
	}
}
