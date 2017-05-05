
-- Created by genevent.

--- {{.Comment}}
local {{.Name}} = {
{{range .Defs}} 
	{{.Name}} = {{.Value}}, --{{.Comment}}
{{end}}
}

return {{.Name}}
