package pokemon

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

func Pascal(input string) string {
	var output, last string
	input = strings.ReplaceAll(input, "-", " ")
	for k, v := range input {
		if k == 0 || last == " " {
			output += strings.ToUpper(string(input[k]))
		} else {
			output += string(v)
		}
		last = string(v)
	}
	return output
}

type Pokemon struct {
	Index string
	Name  string
	Slug  string
	Forms []Form
}

type Form struct {
	Name string
	Png  string
	Cow  string
}

func (p *Pokemon) UnmarshalJSON(data []byte) error {
	var parsed map[string]map[string]string
	var path string
	p.Index = gjson.GetBytes(data, "idx").String()
	p.Name = gjson.GetBytes(data, "name.eng").String()
	p.Slug = gjson.GetBytes(data, "slug.eng").String()
	p.Forms = make([]Form, 0)
	forms := gjson.GetBytes(data, "gen-8.forms")
	json.Unmarshal([]byte(forms.String()), &parsed)
	for name, val := range parsed {
		if _, ok := val["is_alias_of"]; ok {
			continue
		}
		if name == "$" {
			path = fmt.Sprintf("data/regular/%s.png", p.Slug)
			p.Forms = append(p.Forms, Form{"Regular", path, path + ".cow"})
			path = fmt.Sprintf("data/shiny/%s.png", p.Slug)
			p.Forms = append(p.Forms, Form{"Shiny", path, path + ".cow"})
			if _, ok := val["has_female"]; ok {
				path = fmt.Sprintf("data/regular/female/%s.png", p.Slug)
				p.Forms = append(p.Forms, Form{"Female", path, path + ".cow"})
				path = fmt.Sprintf("data/shiny/female/%s.png", p.Slug)
				p.Forms = append(p.Forms, Form{"Shiny Female", path, path + ".cow"})
			}
		} else {
			pname := Pascal(name)
			path = fmt.Sprintf("data/regular/%s-%s.png", p.Slug, name)
			p.Forms = append(p.Forms, Form{"Regular " + pname, path, path + ".cow"})
			path = fmt.Sprintf("data/shiny/%s-%s.png", p.Slug, name)
			p.Forms = append(p.Forms, Form{"Shiny " + pname, path, path + ".cow"})
			if _, ok := val["has_female"]; ok {
				path = fmt.Sprintf("data/regular/female/%s-%s.png", p.Slug, name)
				p.Forms = append(p.Forms, Form{"Female " + pname, path, path + ".cow"})
				path = fmt.Sprintf("data/shiny/female/%s-%s.png", p.Slug, name)
				p.Forms = append(p.Forms, Form{"Shiny Female " + pname, path, path + ".cow"})
			}
		}
	}
	return nil
}
