// +build ignore

package main

import (
	"log"
	"github.com/oddlid/alcolator/srv/assets"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(assets.Assets, vfsgen.Options{
		PackageName: "assets",
		BuildTags: "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
