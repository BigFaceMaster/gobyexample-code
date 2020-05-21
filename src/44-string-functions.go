package main

import (
	"fmt"
	sf "strings"
)



var p = fmt.Println

func main() {
	p("Contains: ", sf.Contains("TEST", "ES") )
	p("Count: ", sf.Count("test", "t"))
	p("HasPrefix: ", sf.HasPrefix("test", "te"))
	p("hasSuffix: ", sf.HasSuffix("test", "st"))
	p("Index: ", sf.Index("test", "e"))
	p("Join: ", sf.Join([]string{"a", "b"}, "-"))
	p("Replace:   ", sf.Replace("foo", "o", "0", -1))
	p("Replace:   ", sf.Replace("foo", "o", "0", 1))
	p("Split:     ", sf.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", sf.ToLower("TEST"))
	p("ToUpper:   ", sf.ToUpper("test"))
	p()
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
