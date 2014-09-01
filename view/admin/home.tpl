{{define "content"}}
	{{if .IsAdmin}}
	<div class="flash notice">Hi, Admin.</div>
	{{end}}

	{{if .IsTeacher}}
	<div class="flash notice">Hi, Teacher.</div>
	{{end}}
{{end}}