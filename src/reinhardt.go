package reinhardt

import (
	"fmt"
	"github.com/r0fls/reinhardt/src/config"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func new_project(name string) {
	base := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "r0fls", "reinhardt")
	os.MkdirAll(filepath.Join(name, "app", "models"), 0700)
	text, err := ioutil.ReadFile(filepath.Join(base, "src", "app_files", "models.go"))
	check(err)
	err = ioutil.WriteFile(filepath.Join(name, "app", "models", "models.go"), text, 0644)
	check(err)

	os.Mkdir(filepath.Join(name, "app", "views"), 0700)

	text, err = ioutil.ReadFile(filepath.Join(base, "src", "app_files", "views.go"))
	check(err)
	err = ioutil.WriteFile(filepath.Join(name, "app", "views", "views.go"), text, 0644)
	check(err)

	os.Mkdir(filepath.Join(name, "app", "temps"), 0700)

	text, err = ioutil.ReadFile(filepath.Join(base, "src", "app_files", "home.html"))
	check(err)
	err = ioutil.WriteFile(filepath.Join(name, "app", "temps", "home.html"), text, 0644)
	check(err)

	text, err = ioutil.ReadFile(filepath.Join(base, "src", "app_files", "settings.json"))
	check(err)
	dir, _ := os.Getwd()
	gopath := os.Getenv("GOPATH")
	local, _ := filepath.Rel(gopath, dir)
	local, _ = filepath.Rel("src", local)
	if string([]rune(local)[0]) == "/" || string([]rune(local)[0]) == "\\" {
		local = local[1:]
	}
	local = filepath.Join(local, name)
	err = ioutil.WriteFile(filepath.Join(name, "settings.json"), []byte(fmt.Sprintf(string(text), dir, name, gopath, local)), 0644)
	check(err)

	c := config.Load_config(filepath.Join(name, "settings.json"))
	text, err = ioutil.ReadFile(filepath.Join(base, "src", "app_files", "manager.go"))
	check(err)
	tmpl, _ := template.New("manager").Parse(string(text))
	f, err := os.Create(filepath.Join(name, "manager.go"))
	check(err)
	err = tmpl.Execute(f, c)
	check(err)

	text, err = ioutil.ReadFile(filepath.Join(base, "src", "app_files", "urls.go"))
	check(err)
	tmpl, _ = template.New("urls").Parse(string(text))
	f, err = os.Create(filepath.Join(name, "app", "urls.go"))
	check(err)
	err = tmpl.Execute(f, c)
	check(err)
}

func Reinhardt() {
	if len(os.Args) < 3 {
		s := []string{"Usage: ", os.Args[0], " new <projectname>"}
		fmt.Println(strings.Join(s, ""))
	} else if os.Args[1] == "new" {
		s := []string{"Created project: ", os.Args[2]}
		fmt.Println(strings.Join(s, ""))
		new_project(os.Args[2])
	}
}
