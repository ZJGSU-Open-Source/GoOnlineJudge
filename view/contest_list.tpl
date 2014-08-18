{{define "content"}}
<h1>Contest List</h1>
<table id="contest_list">
  <thead>
    <tr>
      <th class="header">ID</th>
      <th class="header">Title</th>
      <th class="header">Status</th>
      <th class="header">Type</th>
    </tr>
  </thead>
  <tbody>
    {{$time := .Time}}
    {{$privilege := .Privilege}}
    {{with .Contest}}
      {{range .}} 
        {{if or (ShowStatus .Status) (LargePU $privilege)}}
          <tr>
            <td>{{.Cid}}</td>
            <td><a href="/contest/problem?list/cid?{{.Cid}}">{{.Title}}</a></td>
            <td>{{if ge $time .End }}<font color="green">Ended@{{ShowTime .End}}</font>{{else}}{{if ge .Start $time}}<font color="blue">Start@{{ShowTime .Start}}</font>{{else}}<font color="red">Running</font>{{end}}{{end}}</td>
            <td>{{ShowEncrypt .Encrypt}}</td>
          </tr>
        {{end}}
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}
