{{define "content"}}
<table class="definitions">
	<tbody>
		<tr><th>ID</th><th>Title</th><th>Status</th></tr>
		{{with .Problem}}
			{{range .}}
				<tr><td>{{.Pid}}</td><td><a href="/problem/detail?pid={{.Pid}}">{{.Title}}</a></td><td>{{.Solve}}/{{.Submit}}</td></tr>
			{{end}}
		{{end}}
	</tbody>
</table>
{{end}}