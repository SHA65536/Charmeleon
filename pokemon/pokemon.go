package pokemon

import (
	"encoding/json"
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
	Index string `json:"idx"`
	Name  Name   `json:"name"`
	Slug  Name   `json:"slug"`
	Forms Forms  `json:"gen-8"`
}

type Name struct {
	Str string `json:"eng"`
}

type Forms struct {
	Entries []Form
}

type Form struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

func (f *Forms) UnmarshalJSON(data []byte) error {
	value := gjson.GetBytes(data, "forms")
	f.Entries = make([]Form, 0)
	value.ForEach(func(key, val gjson.Result) bool {
		if strings.Contains(val.String(), "is_alias_of") {
			return true
		}
		if key.String() == "$" {
			f.Entries = append(f.Entries, Form{"data/regular/{name}", "Regular"})
			f.Entries = append(f.Entries, Form{"data/shiny/{name}", "Shiny"})
			if strings.Contains(val.String(), `"has_female": true`) {
				f.Entries = append(f.Entries, Form{"data/regular/female/{name}", "Female"})
				f.Entries = append(f.Entries, Form{"data/shiny/female/{name}", "Shiny Female"})
			}
		} else {
			f.Entries = append(f.Entries, Form{"data/regular/{name}-" + key.String(), Pascal(key.String())})
			f.Entries = append(f.Entries, Form{"data/shiny/{name}-" + key.String(), Pascal("Shiny " + key.String())})
			if strings.Contains(val.String(), `"has_female": true`) {
				f.Entries = append(f.Entries, Form{"data/regular/female/{name}-" + key.String(), Pascal("Female " + key.String())})
				f.Entries = append(f.Entries, Form{"data/shiny/female/{name}-" + key.String(), Pascal("Shiny Female " + key.String())})
			}
		}
		return true
	})
	return nil
}

func (p *Pokemon) MarshalJSON() ([]byte, error) {
	for i := range p.Forms.Entries {
		p.Forms.Entries[i].Path = strings.Replace(p.Forms.Entries[i].Path, "{name}", p.Slug.Str, 1) + ".png"
	}
	return json.Marshal(&struct {
		Index string `json:"idx"`
		Name  string `json:"name"`
		Forms []Form `json:"forms"`
	}{
		Index: p.Index,
		Name:  p.Name.Str,
		Forms: p.Forms.Entries,
	})
}
