# sasa

String utilities for Go

## Usage

### Requirements

Go 1.20+

### TrimMargin

[trimMargin](https://kotlinlang.org/api/latest/jvm/stdlib/kotlin.text/trim-margin.html) in Kotlin, or [stripMargin](https://www.scala-lang.org/api/current/scala/collection/StringOps.html#stripMargin:String) in Scala

```go
func main() {
	s := sasa.TrimMargin(`
	|foo
	| bar
	|  baz
	`)
	fmt.Println(s)
    // =>
    //foo
    // bar
    //  baz

	t := sasa.TrimMargin(`foo
	# bar
	#  baz`, sasa.MarginPrefix("#"))
	fmt.Println(t)
    // =>
    //foo
    // bar
    //  baz
}
```

### TrimIndent

[trimIndent](https://kotlinlang.org/api/latest/jvm/stdlib/kotlin.text/trim-indent.html) in Kotlin

```go
func main() {
	s := sasa.TrimIndent(`
	foo
	 bar
	  baz
	`)
	fmt.Println(s)
	// =>
	//foo
	// bar
	//  baz
}
```

## License

MIT
