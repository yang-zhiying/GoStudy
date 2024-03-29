package main

import s "strings"
import "fmt"

var p = fmt.Println

func main(){
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index: ", s.Index("test", "e"))
	p("Join: ", s.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat: ", s.Repeat("a", 5))
	p("Split: ", s.Split("a-b-c-d-e", "-"))
	p("Replace: ", s.Replace("foo", "o", "0", -1))
	p("ToLower: ", s.ToLower("TeSt"))
	p("ToLower: ", s.ToUpper("test"))
	p("Len: ", len("hello"))
	p("Char: ", "hello"[1])
}