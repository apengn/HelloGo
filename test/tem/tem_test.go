package tem

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"testing"
)

var (
	xx *string
)

type M map[string]interface{}

type DM struct {
	M M
}

func TestTem(t *testing.T) {
	i := "fffffff"
	xx = &i

	D := map[string]interface{}{
		"xxx": xx,
	}
	d := DM{M: D}

	file, err := ioutil.ReadFile("./1.tpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	te := template.Must(template.New("xxx").Parse(string(file)))

	var buf bytes.Buffer

	err = te.Execute(&buf, d)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(buf.String())
}
