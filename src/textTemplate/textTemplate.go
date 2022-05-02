package main

import (
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func upperCase(s string) string {
	return strings.ToUpper(s)
}

func main() {

	filicity := Person{"filicity", 23}
	oliver := Person{"oliver", 23}

	people := []Person{filicity, oliver}
	const greetPerson = `Hi {{.Name | upperCase}}. {{.Age}} years old{{"\n"}}`
	const greetPreple = `{{range .}}Hi {{.Name}}. {{.Age}} {{"\n"}}{{end}}`

	// var maps template.FuncMap != make(template.FuncMap)
	maps := make(template.FuncMap)
	maps["upperCase"] = upperCase
	greetTemplate := template.Must(template.New("greetFromPerson").Funcs(maps).Parse(greetPerson))
	greetPpTemplate := template.Must(template.New("greetFromPeople").Parse(greetPreple))

	greetTemplate.Execute(os.Stdout, filicity)
	greetTemplate.Execute(os.Stdout, oliver)

	greetPpTemplate.Execute(os.Stdout, people)
}
