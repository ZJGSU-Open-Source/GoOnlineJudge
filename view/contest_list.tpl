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
    {{with .Contest}}  
      {{range .}} 
        {{if ShowStatus .Status}}
          <tr>
            <td>{{.Cid}}</td>
            <td><a href="/contest/problem/list/cid/{{.Cid}}">{{.Title}}</a></td>
            <td>{{if ShowExpire .End $time}}<font color="green">Ended@{{.End}}</font>{{else}}{{if ShowExpire $time .Start}}<font color="blue">Start@{{.Start}}</font>{{else}}<font color="red">Running</font>{{end}}{{end}}</td>
            <td>{{ShowEncrypt .Encrypt}}</td>
          </tr>
        {{end}}
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}
