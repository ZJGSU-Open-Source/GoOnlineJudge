{{define "content"}}
<h1>{{.Contest}}</h1>
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
    {{$cid := .Cid}}
    {{with .Solution}}  
      {{range .}} 
        {{if ShowStatus .Status}} 
          <tr>
            <td>{{.Sid}}</td>
            <td>{{.Uid}}</td>
            <td><a href="/contest/problem/detail/cid/{{$cid}}/pid/{{.Pid}}">{{.Pid}}</a></td>
            <td>{{ShowJudge .Judge}}</td>
            <td>{{.Time}}ms</td>
            <td>{{.Memory}}kB</td>
            <td>{{ShowLanguage .Language}}<a href="/contest/status/code/cid/{{$cid}}/sid/{{.Sid}}">[view]</a></td>
            <td>{{.Length}}B</td>
            <td>{{.Create}}</td>
          </tr>
        {{end}}
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}
