package nim2go

import(
  "path/filepath"
  "io/ioutil"
  "os"
  "strings"
  "os/exec"
)

const TempPrefix string = "go2nim"

const MacroHeader string = `
import macros

dumpTree:
`

// A basic data structure for the Nim parser.
type Parser struct {
  AppendMacros bool
  CompilerOutput []byte
}

// Initializes a new parser.
func NewParser() Parser {
  parser := Parser{
  }
  return parser
}

// Takes an input source file and preprends the macros required to output the AST.
func WrapWithMacros(filePath string) (tempFile *os.File, err error) {
  tempFile, err = ioutil.TempFile( "/tmp", TempPrefix )

  var inputSource []byte
  inputSource, err = ioutil.ReadFile(filePath)

  var inputSourceString string
  inputSourceString = string(inputSource)
  inputSourceString = strings.Replace(inputSourceString, "\n", "\n ", -1)
  inputSourceString = strings.Join([]string{" ", inputSourceString}, "")

  var newTempFilePath string
  newTempFilePath = strings.Join([]string{tempFile.Name(), ".nim"}, "")

  if err == nil {
    tempFile.Write([]byte(MacroHeader))
    tempFile.Write([]byte(inputSourceString))
    tempFile.Close()

    os.Rename(tempFile.Name(), newTempFilePath)

    tempFile, err = os.Open(newTempFilePath)
  }

  return tempFile, err
}

// Calls the Nim compiler and returns the stdout data.
func NimCompile(filePath string) (output []byte, err error) {
  output, err = exec.Command("nim", "compile", filePath).Output()
  return output, err
}


// Parses Nim code directly, eval-like. It makes use of a temporary file.
func(p *Parser) ParseString(source string) (err error) {
  return err
}

// Parses Nim code from a given file.
func(p *Parser) ParseFile(filePath string) (err error) {
  var sourceFile *os.File
  sourceFile, err = WrapWithMacros(filePath)

  if err == nil {
    p.CompilerOutput, err = NimCompile(sourceFile.Name())
    os.Remove(sourceFile.Name())
  }

  return err
}

// This will call ParseFile if you pass a string with ".nim".
func(p *Parser) Parse(input string) (err error) {
  ext := filepath.Ext(input)
  if ext == ".nim" {
    err = p.ParseFile(input)
  } else {
    err = p.ParseString(input)
  }
  return err
}
