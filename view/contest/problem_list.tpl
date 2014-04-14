{{define "content"}}
<h1>{{.Contest}}</h1>
<table id="contest_list">
  <thead>
    <tr>
      <th class="header">ID</th>
      <th class="header">Title</th>
      <th class="header">Ratio</th>
    </tr>
  </thead>
  <tbody>
    {{with .Problem}}  
      {{range .}} 
        {{if ShowStatus .Status}}
            <tr>
              <td>{{.Pid}}</td>
              <td><a href="/contest/problem/detail/pid/{{.Pid}}">{{.Title}}</a></td>
              <td>Here is ratio</td>
            </tr>
        {{end}}
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}
