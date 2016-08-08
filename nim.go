package nim2go

/* Basic structs */

type NimSource struct {
  Root *Node
}

type Node interface {
}

/* Tokens */

type Token int

const (
  _ = iota
  STMTLIST
  IMPORTSTMT
  INFIX
  IDENT
  COMMAND
  INTLIT
  STRLIT
  EMPTY
  LETSECTION
  IDENTDEFS
  DOTEXPR
  CALL
)

var tokens = [...]string{
  STMTLIST: "StmtList",
  IMPORTSTMT: "ImportStmt",
  INFIX: "Infix",
  IDENT: "Ident",
  COMMAND: "Command",
  INTLIT: "IntLit",
  STRLIT: "StrLit",
  EMPTY: "Empty",
  LETSECTION: "LetSection",
  IDENTDEFS: "IdentDefs",
  DOTEXPR: "DotExpr",
  CALL: "Call",
}

func TokenLookup(input string) Token {
  for i, v := range tokens {
    if v == input {
      return Token(i)
      break
    }
  }
  return -1
}

/* Statements */

type Statement interface {
  Node
}

type StmtList struct {
  List []Statement
}


type ImportStmt struct {
  ModuleName string
}

type CallStmt struct {
  DotExpr DotExpr
}

type Ident struct {
  Name string
}

type DotExpr struct {
  Idents []Ident
}

type Command struct {
  Ident Ident
  Arguments Node
}

/* Literals */

type StrLit struct {
  Value string
}

type IntLit struct {
  Value int
}
