package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/modules"
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/tmpl"
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/utils"
)

type data struct {
	Module  modules.Module
	Modules modules.Modules
}

func main() {
	var (
		workingDir = utils.ProjectDir()
		mods       = modules.GetModules()
		funcMap    = template.FuncMap{
			"ToCamelCase":      utils.ToCamelCase,
			"ToLowerCamelCase": utils.ToLowerCamelCase,
			"ToSnakeCase":      utils.ToSnakeCase,
			"ToLowerCase":      utils.ToLowerCase,
			"ToPlural":         utils.ToPlural,
		}
	)

	temps := prepareTemplate(workingDir, templatePaths, mods)
	for _, temp := range temps {
		if !temp.IsOverWrite && utils.FileExists(temp.Destination) {
			continue
		}
		dir := filepath.Dir(temp.Destination)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.Mkdir(dir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		file, err := os.Create(temp.Destination)
		if err != nil {
			log.Fatal(err)
		}
		buffer := new(bytes.Buffer)
		content, err := ioutil.ReadFile(temp.Source)
		if err != nil {
			log.Fatal(err)
		}
		err = temp.ExecuteTemplate(buffer, content, funcMap, temp.Data)
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.Write(buffer.Bytes())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func prepareTemplate(workingDir string, paths Paths, mods modules.Modules) tmpl.Templates {
	var (
		items tmpl.Templates
	)
	for _, mod := range mods {
		for _, templatePath := range templatePaths {
			source := path.Join(workingDir, templatePath.Source)
			destination := path.Join(workingDir, templatePath.Destination)
			destination = strings.ReplaceAll(destination, "{{module}}", utils.ToLowerCase(mod.Name))
			items = append(items, tmpl.Template{
				Source:      source,
				Destination: destination,
				Data: data{
					Module:  mod,
					Modules: mods,
				},
				IsOverWrite: templatePath.IsOverWrite,
			})
		}
	}
	return items
}
