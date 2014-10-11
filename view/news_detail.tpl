{{define "content"}}
	{{with .Detail}}
		<h2>{{.Title}}</h2>
		<p><b>Date: </b>{{.Create}}</p>
		{{.Content}}
	{{end}}
{{end}}