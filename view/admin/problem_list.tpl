{{define "content"}}
<h1>Admin - Problem List</h1>
<table id="contest_list">
  <thead>
    <tr>
      <th class="header">ID</th>
      <th class="header">Title</th>
      <th class="header">Status</th>
      <th class="header">Delete</th>
      <th class="header">Edit</th>
      <th class="header">Data</th>
    </tr>
  </thead>
  <tbody>
    {{with .Problem}}  
      {{range .}} 
        <tr>
          <td>{{.Pid}}</td>
          <td><a href="/admin/problem/detail/pid/{{.Pid}}">{{.Title}}</a></td>
          <td><a href="/admin/problem/status/pid/{{.Pid}}">[{{if ShowStatus .Status}}Available{{else}}Reserved{{end}}]</a></td>
          <td><a href="#" onclick="ConfirmDelete('/admin/problem/delete/pid/{{.Pid}}', 'Detele The Problem ?')">[Delete]</a></td>
          <td><a href="/admin/problem/edit/pid/{{.Pid}}">[Edit]</a></td>
          <td>[Test Data]</td>
        </tr>
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}
