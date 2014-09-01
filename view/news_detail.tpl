{{define "content"}}
	{{with .Detail}}
		<h1>{{.Title}}</h1>
		<p><b>Date: </b>{{.Create}}</p>
		{{.Content}}
	{{end}}
{{end}}