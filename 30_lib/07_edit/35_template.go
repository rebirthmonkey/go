package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	var group = "app"
	ctx := map[string]interface{}{
		"group": group,
	}

	docGoTemplate := `
// +k8s:openapi-gen=true
// +groupName={{ .group }}.demo.wukong.com
// +kubebuilder:object:generate=true

package {{ .group }}
	
`
	tmpl, _ := template.New("doc").Parse(docGoTemplate)

	var buf bytes.Buffer
	writer := io.MultiWriter(&buf)
	err := tmpl.Execute(writer, ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("the code is: ", buf.String())

	// create the file
	apisPath := "apis"
	groupDir := filepath.Join(apisPath, group)
	if err = os.MkdirAll(groupDir, 0755); err != nil {
		fmt.Println(err)
	}

	_, err = os.Stat(groupDir)
	if err != nil {
		fmt.Println(err)
	}

	if err = ioutil.WriteFile(filepath.Join(groupDir, "doc.go"), buf.Bytes(), 0644); err != nil {
		return
	}

}
