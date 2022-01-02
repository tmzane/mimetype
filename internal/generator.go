package main

//go:generate go run generator.go

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"unicode"

	"github.com/ettle/strcase"
)

const mtTemplate = `package mimetype

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

const (
{{- range $i, $line := .In}}
	// {{ index $line 0 | constName $.Document }} {{ index $line 0 }} mime type.
	{{ index $line 0 | constName $.Document }} = "{{if index $line 1 }}{{ index $line 1 }}{{else}}{{ $.Document }}/{{index $line 0}}{{end}}"
{{- end}}
)

// Is{{ toGoPascal .Document }} checks if the mime types is {{ .Document }}.
func Is{{ toGoPascal .Document }}(mt string) bool {
	switch mt {
{{- range $index, $value := .MimeTemplates }}
	case "{{ $value }}":
		return true
{{- end}}
	default:
		return false
	}
}
`

func main() {
	documents := []string{
		"application",
		"audio",
		"font",
		"image",
		"message",
		"model",
		"multipart",
		"text",
		"video",
	}

	for _, document := range documents {
		log.Println(document)
		err := generate(document)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// document

func generate(document string) error {
	resp, err := http.Get(fmt.Sprintf("https://www.iana.org/assignments/media-types/%s.csv", document))
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		all, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("problem during download (%d): %s", resp.StatusCode, string(all))
	}

	reader := csv.NewReader(resp.Body)

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("read CSV entries: %w", err)
	}

	uniq := make(map[string]struct{})

	for _, record := range records[1:] {
		r := record[1]
		if r == "" {
			r = document + "/" + record[0]
		}
		uniq[r] = struct{}{}
	}

	var mts []string

	for k := range uniq {
		mts = append(mts, k)
	}

	sort.Strings(mts)

	model := map[string]interface{}{
		"In":            records[1:],
		"Document":      document,
		"MimeTemplates": mts,
	}

	content, err := generateFileContent(document, model)
	if err != nil {
		return fmt.Errorf("generate file content: %w", err)
	}

	file, err := os.Create(filepath.Join("..", document+".go"))
	if err != nil {
		return err
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func generateFileContent(document string, model map[string]interface{}) (string, error) {
	funcMap := template.FuncMap{
		"constName":  constName,
		"toGoPascal": strcase.ToGoPascal,
	}

	base := template.New(document + ".go")
	tmpl, err := base.Funcs(funcMap).Parse(mtTemplate)
	if err != nil {
		return "", fmt.Errorf("parse template: %w", err)
	}

	b := &bytes.Buffer{}
	err = tmpl.Execute(b, model)
	if err != nil {
		return "", fmt.Errorf("execute template: %w", err)
	}

	// gofmt
	source, err := format.Source(b.Bytes())
	if err != nil {
		log.Println(b.String())
		return "", fmt.Errorf("format source: %w", err)
	}

	return string(source), nil
}

func constName(document, text string) string {
	if text[len(text)-1] == '+' {
		text += "plus"
	}

	fargs := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	return strcase.ToGoPascal(document + "_" + strings.Join(strings.FieldsFunc(text, fargs), "_"))
}
