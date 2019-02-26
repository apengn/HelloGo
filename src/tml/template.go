package main

import (
	"html/template"
	"os"
	"fmt"
	"strings"
	"log"
	"strconv"
)

type Friends struct {
	Fname string
}

type MyEmail struct {
	Email string
}

type Person struct {
	UserName string
	Emails   []MyEmail
	Friends  []*Friends
	Addr     []string
}

func main1() {
	t := template.New("x")
	t, err := t.Parse("hello {{.UserName}}!")
	if os.IsExist(err) {
		fmt.Println(err.Error())
	}
	p := Person{UserName: "xxxxxxxxxxx"}
	t.Execute(os.Stdout, p)
}
func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	var i int
	o:=false
	if len(args) == 2 {
		s, ok = args[1].(string)
		i, o= args[0].(int)
		if o {
			fmt.Println(i)
		}
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	// replace the @ by " at "
	return (substrs[0] + " at " + substrs[1]+strconv.Itoa(i))
}

func main() {
	f1 := Friends{Fname: "f1"}
	f2 := Friends{"f2"}
	t := template.New("filedname example")

	t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t.Parse(` {{ define 'name2' }} 
                   {{ $name := default .UserName  .UserName -}}
                   {{ printf "%s-%s" $name $name}} 
                   {{ end }}

`)
	t, err := t.Parse(`hello {{.UserName}}!
                   oooooooooo  {{ template "name2" .}}
                  {{range $index,$email:=.Emails}}
                       an email  {{ $xxx := emailDeal $index $email}}

                  xxx.txt  {{$xxx}}
                  {{end}}
                  {{range .Addr}}
                     {{.}}
                  {{else}}
                     empty
                  {{end}}
                  {{with .Friends}}
                     {{range $index,$value := .}}
                       {{$index}}my friend name is {{$value}}                   
                     {{end}}
                  {{end}}`)
	if err!=nil {
		fmt.Println(err)
		return
	}
	p := Person{Addr: []string{}, UserName: "wp", Emails: []MyEmail{{"xxxxxxxxx@qq.com"},{"tttttttt@qq.com"},{"fffffffffff@163.com"}}, Friends: []*Friends{&f1, &f2}}
	f, _ := os.OpenFile("template.conf", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	fmt.Sprint(t.Execute(f, p))
}

func main33() {
	tEmpty := template.New("template test")
	template.ParseFiles()
	tEmpty = template.Must(tEmpty.Parse(`空 pipeline if demo :
                                           {{if .IsNull}}
                                               {{.Name}}
                                           {{else}}
                                                else 部分
                                           {{end}}`))

	f, _ := os.OpenFile("nginx.conf", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	tEmpty.Execute(f, struct {
		Name   string
		IsNull bool
	}{Name: "xxx.txt", IsNull: false})

}
func main4() {
	//tmpl:=` $variable := pipeline`

	t := template.New("template test")
	tmpl1 := `{{with $x := "output" | printf "%q"}}{{$x}}{{end}}`
	//tmpl2 := `{{with $x := "output"}}{{printf "%q" $x}}{{end}}`
	//temp3 := `{{with $x := "output"}}{{$x |printf "%q"}}{{end}}`

	t, _ = t.Parse(tmpl1)

	//f, _ := ioutil.TempFile("./", "xx.tmpl")

	f, _ := os.OpenFile("nginx.conf", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)

	t.Execute(f, nil)

	//main2()
}

type Ex1 struct {
	Params string
}

func main9() {
	const (
		master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	)
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
}

func main6() {
	t, _ := template.New("xxx.txt").Parse(`{{.Name}}{{define "T1"}}{{.Name}}{{end}}
{{define "T2"}}TWO{{end}}
{{define "T3"}}{{template "T1"}}{{template "T2"}}{{end}}
{{template "T3"}}`)

	type Addr struct {
		Addr string
	}
	f, _ := os.OpenFile("nginx23.conf", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	t.Execute(f, struct {
		Name string
		Addr []*Addr
	}{Name: "xxxxxxx", Addr: []*Addr{&Addr{Addr: "addr"}}})
}
