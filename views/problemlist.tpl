{{template "head" .}}
	<div class="container">
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
	</div>
{{template "foot"}}