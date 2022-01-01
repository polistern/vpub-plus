// +build ignore

package main

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

const tpl = `// Code generated by go generate; DO NOT EDIT.

package {{ .Package }}

var {{ .Map }} = map[string]string{
{{ range $constant, $content := .Files }}` + "\t" + `"{{ $constant }}": ` + "`{{ $content }}`" + `,
{{ end }}}
`

var bundleTpl = template.Must(template.New("").Parse(tpl))

type Bundle struct {
	Package string
	Map     string
	Files   map[string]string
}

func (b *Bundle) Write(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bundleTpl.Execute(f, b)
}

func NewBundle(pkg, mapName string) *Bundle {
	return &Bundle{
		Package: pkg,
		Map:     mapName,
		Files:   make(map[string]string),
	}
}

func stripExtension(filename string) string {
	filename = strings.TrimSuffix(filename, path.Ext(filename))
	return strings.Replace(filename, " ", "_", -1)
}

func readFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

func glob(pattern string) []string {
	files, _ := filepath.Glob(pattern)
	for i := range files {
		if strings.Contains(files[i], "\\") {
			files[i] = filepath.ToSlash(files[i])
		}
	}
	return files
}

func generateMap(target string, pkg string, mapName string, srcFiles []string) {
	bundle := NewBundle(pkg, mapName)
	for _, srcFile := range srcFiles {
		data := readFile(srcFile)
		filename := stripExtension(path.Base(srcFile))
		bundle.Files[filename] = string(data)
	}
	bundle.Write(target)
}

func main() {
	generateMap(path.Join("storage", "sql.go"), "storage", "SqlMap", glob("storage/sql/*.sql"))
	generateMap(path.Join("web", "handler", "html.go"), "handler", "TplMap", glob("web/handler/html/*.html"))
	generateMap(path.Join("web", "handler", "common.go"), "handler", "TplCommonMap", glob("web/handler/html/common/*.html"))
	generateMap(path.Join("assets", "style.go"), "assets", "AssetsMap", glob("assets/*"))
}
