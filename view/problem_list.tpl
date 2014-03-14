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
    {{$time := .Time}}
    {{with .Problem}}  
      {{range .}} 
        {{if ShowStatus .Status}}
          {{if ShowExpire .Expire $time}}
            <tr>
              <td>{{.Pid}}</td>
              <td><a href="/problem/detail/pid/{{.Pid}}">{{.Title}}</a></td>
              <td>{{ShowRatio .Solve .Submit}}({{.Solve}}/{{.Submit}})</td>
            </tr>
          {{end}}
        {{end}}
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}
