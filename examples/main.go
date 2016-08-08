package main

import(
  "fmt"

  "github.com/matiasinsaurralde/nim2go"
)

func main() {
  fmt.Println("Let's compile hello_world.nim!")

  var err error

  parser := nim2go.Parser{
    AppendMacros: true,
  }

  err = parser.Parse("hello_world.nim")

  if err == nil {
    fmt.Println("Compiler output:")
    fmt.Println( string(parser.CompilerOutput))
  } else {
    panic(err)
  }
}
