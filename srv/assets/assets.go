// +build dev

package assets

import (
	"net/http"
	"os"
	"strings"

	"github.com/shurcooL/httpfs/filter"
	"github.com/shurcooL/httpfs/union"
)

//var static http.FileSystem = filter.Keep(
//	http.Dir("assets/static"),
//	func(path string, fi os.FileInfo) bool {
//		return fi.IsDir() || strings.HasSuffix(path, ".css")
//	},
//)

var templates http.FileSystem = filter.Keep(
	http.Dir("templates"),
	func(path string, fi os.FileInfo) bool {
		return fi.IsDir() || strings.HasSuffix(path, ".html")
	},
)

var Assets http.FileSystem = union.New(map[string]http.FileSystem{
	"/templates": templates,
	//"/static":    static,
})

