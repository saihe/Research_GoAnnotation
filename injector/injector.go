/*
使用できるアノテーションは `service` のみとする
*/
package injector

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

	"golang.org/x/tools/go/packages"
)

type InjectParameters struct {
	targetPackages []string
}

var Services map[string]interface{}

func Initialize(params InjectParameters) {
	target := params.targetPackages
	if target == nil {
		target = []string{
			"goannotation/service",
		}
	}
	for _, packagename := range target {
		cfg := &packages.Config{
			Mode: packages.NeedFiles,
			// Mode:  packages.NeedTypes,
			Tests: false,
		}
		// パッケージ内を取得
		pkgs, err := packages.Load(cfg, packagename)
		if err != nil {
			log.Fatalf("failed to load package: %v", err)
			return
		}
		for _, p := range pkgs {
			// ファイルを取得
			for _, f := range p.GoFiles {
				parsedFile, parseerr := parser.ParseFile(token.NewFileSet(), f, nil, parser.ParseComments)
				if parseerr != nil {
					log.Fatal(parseerr)
				}
				// アノテーションが付与された構造体を取得する
				v1 := visitorFunc(func(node ast.Node) (w ast.Visitor) {
					fmt.Printf("%T\n", node)
					return nil
				})
				for _, d := range parsedFile.Decls {
					ast.Walk(v1, d)
				}

				// 構造体に service アノテーションがついている場合、インジェクトする
				for _, c := range parsedFile.Comments {
					for _, c2 := range strings.Split(c.Text(), "\n") {
						if strings.Contains(c2, "`") && strings.Contains(c2, "service") {
							annotations := strings.Split(strings.ReplaceAll(c2, "`", ""), ":")
							fmt.Printf("%v\n", annotations)
						}
					}
				}
			}
		}
	}
}

type visitorFunc func(node ast.Node) (w ast.Visitor)

func (f visitorFunc) Visit(node ast.Node) (w ast.Visitor) {
	return f(node)
}
