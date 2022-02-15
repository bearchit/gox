//go:build ignore
// +build ignore

package main

import (
	"github.com/bearchit/gox/entx/available"
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := available.NewExtension()
	if err != nil {
		log.Fatalf("creating available extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ex))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
