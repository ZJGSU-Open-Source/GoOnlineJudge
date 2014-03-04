{{define "content"}}
<h1>Problem List</h1>
<table id="contest_list">
  <thead>
    <tr>
      <th class="header">ID</th>
      <th class="header">TITLE</th>
      <th class="header">RATIO</th>
    </tr>
  </thead>
  <tbody>
    {{with .Problem}}  
      {{range .}}  
        <tr>
          <td>{{.Pid}}</td>
          <td>{{.Title}}</td>
          <td>{{.Solve}}/{{.Submit}}</td>
        </tr>
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}
