{{define "content"}}
<div class="row">

<h2 class="page-header">News List</h2>
<div class="col-md-8">
{{with .News}}

	{{range .}}
		{{if ShowStatus .Status}}

			<p class="news">
				<span class="flag"></span>
				<span class="date">{{.Create}}</span>		
				<br><a href="/news/{{.Nid}}">{{.Title}}</a>
			</p>
		{{end}}
	{{end}}
{{end}}
</div>
<div class="col-md-4">
<span>Check every 10 minutes</span>
	<table id="problem_list" class="table table-bordered table-striped table-hover">
  <thead>
    <tr>
      <th class="header">OJ</th>
      <th class="header">STATUS</th>
    </tr>
  </thead>
  <tbody>
  {{with .OJStatus}}
  	{{range .}}
    <tr>
      <td>{{.Name}}</td>
      <td>{{if eq .Status 0}} 
      		<span class="submitRes-3"><strong>Ok</strong></span>
      		{{else}}
      		<span class="submitRes-4"> <strong>Unavailable</strong></span>
      		{{end}}
      </td>
    </tr>
    {{end}}
   {{end}}
  </tbody>
</table>
</div>
</div>
{{end}}
