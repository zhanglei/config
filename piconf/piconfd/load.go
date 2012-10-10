// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

var startVar = []byte("package main; var (\n")

const endVar = ')'

func Load(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	content = append(startVar, content...)
	content = append(content, endVar)

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("configuration in %q is not valid: %s", filename, err)
	}

	genDecl := node.Decls[0].(*ast.GenDecl)
	for _, spec := range genDecl.Specs {
		valueSpec := spec.(*ast.ValueSpec)

		for i, name := range valueSpec.Names {
			key := name.Name
fmt.Print(key, ": ")

			if valueSpec.Doc == nil {
				return fmt.Errorf("variable %q has not documentation in %q", key, filename)
			}

			switch expr := valueSpec.Values[i].(type) {
			case *ast.BasicLit:
				fmt.Println(expr.Value)
			case *ast.Ident:
				fmt.Println(expr.Name)
			default:
				return fmt.Errorf("expression no valid for variable %q", key)
			}

			if typ, ok := valueSpec.Type.(*ast.Ident); ok {
				fmt.Println(typ.Name)
			}
		}
	}


//	ast.Print(fset, node)

	return nil
}
