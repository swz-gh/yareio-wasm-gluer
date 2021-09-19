package main

import (
	"bytes"
	"crypto/rand"
	_ "embed"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/evanw/esbuild/pkg/api"
)

//go:embed glue.js
var glueCode string

type GlueVars struct {
	WasmContent string
	Unique      string
}

func RandomId() string {
	b := make([]byte, 5)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}

func main() {
	if len(os.Args) < 2 {
		panic("You need to provide a file")
	}
	wasmfile := os.Args[1]

	file, err := ioutil.ReadFile(wasmfile)
	if err != nil {
		panic(err)
	}
	content := base64.URLEncoding.EncodeToString(file)
	content = strings.ReplaceAll(content, "-", "+")
	content = strings.ReplaceAll(content, "_", "/")

	u := GlueVars{
		WasmContent: content,
		Unique:      RandomId(),
	}

	ut, err := template.New("Glue").Parse(glueCode)

	if err != nil {
		panic(err)
	}

	var temp bytes.Buffer
	err = ut.Execute(&temp, u)
	if err != nil {
		panic(err)
	}
	result := api.Transform(temp.String(), api.TransformOptions{MinifyWhitespace: true, MinifyIdentifiers: true, MinifySyntax: true})
	os.Stdout.Write(result.Code)
}
