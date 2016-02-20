package reinhardt

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func new_project(name string) {
	m := []string{name, "/app", "/models"}
	os.MkdirAll(strings.Join(m, ""), 0700)
	mf := []string{name, "/app", "/models", "/models.go"}
	os.Create(strings.Join(mf, ""))

	v := []string{name, "/app", "/views"}
	os.Mkdir(strings.Join(v, ""), 0700)
	vf := []string{name, "/app", "/views", "/views.go"}
	os.Create(strings.Join(vf, ""))

	t := []string{name, "/app", "/temps"}
	os.Mkdir(strings.Join(t, ""), 0700)

	s := []string{name, "/settings.toml"}
	text, err := ioutil.ReadFile("app_files/settings.toml")
	check(err)
	err = ioutil.WriteFile(strings.Join(s, ""), text, 0644)
	check(err)

	s = []string{name, "/manager.go"}
	text, err = ioutil.ReadFile("app_files/manager.go")
	check(err)
	err = ioutil.WriteFile(strings.Join(s, ""), text, 0644)
	check(err)

	u := []string{name, "/urls.go"}
	os.Create(strings.Join(u, ""))
}

func Reinhardt() {
	if len(os.Args) < 2 {
		s := []string{"Usage: ", os.Args[0], " new <projectname>"}
		fmt.Println(strings.Join(s, ""))
	} else if os.Args[1] == "new" {
		s := []string{"Making project: ", os.Args[2]}
		fmt.Println(strings.Join(s, ""))
		new_project(os.Args[2])
	}
}
