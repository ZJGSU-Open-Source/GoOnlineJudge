{{define "content"}}
<table class="definitions">
	<tbody>
		<tr><th>ID</th><th>Title</th><th>Status</th></tr>
		{{with .Problem}}
			{{range .}}
				<tr><td>{{.Pid}}</td><td>{{.Title}}</td></tr>
			{{end}}
		{{end}}
	</tbody>
</table>
{{end}}