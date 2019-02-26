package main

import (
	"html/template"
	"fmt"
	"os"
)

type Location struct {
	Error    bool
	Api      string
	Location string //匹配的location
}

type Upstream struct {
	Name     string
	Server   []string //在负载均衡的地址
	Location []*Location
}

type Server struct {
	Listen     string
	ServerName string
	Upstream   []*Upstream
}
type NginxConfig struct {
	Server   []*Server
	Upstream []*Upstream
}

const tmpl = `
              {{range .Upstream}}
                     upstream  {{.Name}} {
                       {{range .Server}}
                          server {{.}};
                        {{end}}
                     }
              {{end}}

                   {{range .Server}}
                        server {
                         listen       {{.Listen}};
                         server_name  {{.ServerName}};

                      {{range $upstream := .Upstream}}
                         {{range $upstream.Location }}
                              {{if .Error }}

                                error_page   500 502 503 504  {{.Location}};
                                location = {{.Location}} {
                                   root   /usr/share/nginx/html;
                                }
                                   
                              {{else}}

                                  location {{.Location}} {
                                     proxy_pass http://{{$upstream.Name}}/{{.Api}};
                                 }

                              {{end}}
                               
                         {{end}} 
                      {{end}}
                     }

                   {{end}}
                 `

const tmpl1 = `{{range .Upstream}}
	  upstream  {{.Name}} { {{range .Server}}
	    server {{.}};{{end}}
	}{{end}}{{range $s := .Server}}
	  server {
		listen       {{$s.Listen}};
		server_name  {{$s.ServerName}};  {{range $upstream := $s.Upstream}}	{{range $upstream.Location }}  {{if .Error }}

		error_page   500 502 503 504  {{.Location}};
		location = {{.Location}} {
		   root   /usr/share/nginx/html;
		}{{else}}
		  
		location {{.Location}} {
		  proxy_pass http://{{$upstream.Name}}/{{.Api}};
		} {{end}}{{end}} {{end}}
	   }
   {{end}}`

func main() {

	u := &Upstream{Server: []string{"192.168.0.1:8080", "192.168.0.1:8080",},
		Name: "xxxxxx",
		Location: []*Location{
			&Location{Error: true, Location: "/50x.html", Api: "oo"},
			&Location{Error: false, Location: "/test1", Api: "test1"},
			&Location{Error: false, Location: "/test2", Api: "test2"},
			&Location{Error: false, Location: "/test3", Api: "test3"},
			&Location{Error: false, Location: "/test4", Api: "test4"},
		}}

	u2 := &Upstream{Server: []string{"192.168.0.1:9999", "192.168.0.1:7777",},
		Name: "xxxxx2",
		Location: []*Location{
			&Location{Error: false, Location: "/50x.html", Api: "oo"},
			&Location{Error: false, Location: "/test1", Api: "test1"},
			&Location{Error: false, Location: "/test2", Api: "test2"},
			&Location{Error: false, Location: "/test3", Api: "test3"},
			&Location{Error: false, Location: "/test4", Api: "test4"},
		}}

	up := []*Upstream{u, u2}
	nginx := NginxConfig{Upstream: up,
		Server: []*Server{
			&Server{Listen: "8080", ServerName: "localhost",
				Upstream: up},
			&Server{Listen: "8032", ServerName: "localhost3232",
				Upstream: up},}}


	t, err := template.ParseFiles("nginx.tmpl")
	if err != nil {
		fmt.Println("error", err.Error())
	}

	f, _ := os.OpenFile("nginx.conf", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	fmt.Println("execute error:", t.Execute(f, nginx))
}
