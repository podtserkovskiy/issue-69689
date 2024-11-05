package main

import (
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	rawCGoFile, err := parser.ParseFile(fset, "./cgo_objdir/test1.cgo1.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Raw CGo file parsed")

	filename := fset.Position(rawCGoFile.Pos()).Filename
	if filename != "testdata/test1.go" {
		log.Fatalf("Wrong file name %s, expected: testdata/test1.go\n", filename)
	}
	log.Println("//line resolved correctly")
}
