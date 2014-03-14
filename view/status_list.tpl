{{define "content"}}
<h1>Status List</h1>
<table id="contest_list">
  <thead>
    <tr>
      <th class="header">ID</th>
      <th class="header">User</th>
      <th class="header">Problem</th>
      <th class="header">Result</th>
      <th class="header">Time</th>
      <th class="header">Memory</th>
      <th class="header">Language</th>
      <th class="header">Code Length</th>
      <th class="header">Submit Time</th>
    </tr>
  </thead>
  <tbody>
    {{with .Solution}}  
      {{range .}} 
        {{if ShowStatus .Status}} 
          <tr>
            <td>{{.Sid}}</td>
            <td>{{.Uid}}</td>
            <td><a href="/problem/detail/pid/{{.Pid}}">{{.Pid}}</a></td>
            <td>{{ShowJudge .Judge}}</td>
            <td>{{.Time}}ms</td>
            <td>{{.Memory}}kB</td>
            <td>{{ShowLanguage .Language}}<a href="/status/code/sid/{{.Sid}}"></td>
            <td><a href="/status/code/sid/{{.Sid}}">{{.Length}}B</a></td>
            <td>{{.Create}}</td>
          </tr>
        {{end}}
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}
