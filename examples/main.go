package main

import(
  "fmt"
  "os"

  nim "github.com/matiasinsaurralde/nim2go"
)

func main() {
  fmt.Println("Let's compile hello_world.nim!")

  var err error

  parser := nim.Parser{
    AppendMacros: true,
  }

  err = parser.Parse("hello_world.nim")

  // err = parser.Parse("greet_module.nim")

  os.Exit(0)

  stmtList := (nim.Node)(&nim.StmtList{
    List: []nim.Statement{
      nim.ImportStmt{
        ModuleName: "greet_module",
      },
      nim.Command{
        Ident: nim.Ident{"echo"},
        Arguments: nim.StrLit{"hello world"},
      },
      nim.Command{
        Ident: nim.Ident{"echo"},
        Arguments: nim.IntLit{1},
      },
      nim.CallStmt{
        DotExpr: nim.DotExpr{
          Idents: []nim.Ident{
            nim.Ident{"greet"},
            nim.Ident{"greet"},
          },
        },
      },
    },
  })

  ast := nim.NimSource{
    Root: &stmtList,
  }

  fmt.Println(*ast.Root)



  if err == nil {
    fmt.Println("Compiler output:")
    fmt.Println( string(parser.CompilerOutput))
  } else {
    panic(err)
  }
}
