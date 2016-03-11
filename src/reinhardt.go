package reinhardt

import (
	"fmt"
	"github.com/r0fls/reinhardt/src/config"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func new_project(name string) {
	m := []string{name, "app", "models"}
	os.MkdirAll(strings.Join(m, "/"), 0700)

	s := []string{name, "app", "models", "models.go"}
	text, err := ioutil.ReadFile("src/app_files/models.go")
	check(err)
	err = ioutil.WriteFile(strings.Join(s, "/"), text, 0644)
	check(err)

	v := []string{name, "app", "views"}
	os.Mkdir(strings.Join(v, "/"), 0700)

	s = []string{name, "app", "views", "views.go"}
	text, err = ioutil.ReadFile("src/app_files/views.go")
	check(err)
	err = ioutil.WriteFile(strings.Join(s, "/"), text, 0644)
	check(err)

	t := []string{name, "app", "temps"}
	os.Mkdir(strings.Join(t, "/"), 0700)

	s = []string{name, "app", "temps", "home.html"}
	text, err = ioutil.ReadFile("src/app_files/home.html")
	check(err)
	err = ioutil.WriteFile(strings.Join(s, "/"), text, 0644)
	check(err)

	s = []string{name, "settings.json"}
	text, err = ioutil.ReadFile("src/app_files/settings.json")
	check(err)
	dir, _ := os.Getwd()
	gopath := os.Getenv("GOPATH")
	local := strings.Replace(dir, gopath, "", 1)
	local = strings.Replace(local, "src", "", 1)
	local = strings.Join([]string{strings.Replace(local, "//", "", 1), name}, "/")
	err = ioutil.WriteFile(strings.Join(s, "/"), []byte(fmt.Sprintf(string(text), dir, name, gopath, local)), 0644)
	check(err)

	c := config.Load_config(strings.Join(s, "/"))
	s = []string{name, "manager.go"}
	text, err = ioutil.ReadFile("src/app_files/manager.go")
	check(err)
	tmpl, _ := template.New("manager").Parse(string(text))
	f, err := os.Create(strings.Join(s, "/"))
	check(err)
	err = tmpl.Execute(f, c)
	check(err)

	s = []string{name, "app", "urls.go"}
	text, err = ioutil.ReadFile("src/app_files/urls.go")
	check(err)
	tmpl, _ = template.New("urls").Parse(string(text))
	f, err = os.Create(strings.Join(s, "/"))
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
