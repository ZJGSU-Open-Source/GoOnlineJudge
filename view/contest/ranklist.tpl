{{define "content"}}
{{$cid := .Cid}}
<h1>{{.Contest}}</h1>
<table id="contest_list">
  <thead>
    <tr>
      <th class="header">Rank</th>
      <th class="header">Team</th>
      <th class="header">Solved</th>
      <th class="header">Penalty</th>
      {{with .ProblemList}}
      {{range .}}
      <th class="header"><a href="/contest/problem/detail?cid={{$cid}}&pid={{.}}">{{.}}</a></th>
      {{end}}
      {{end}}
    </tr>
  </thead>
  <tbody>
    {{with .UserList}}
      {{range $idx,$v := .}} 
          <tr>
            <td>{{NumAdd $idx 1}}</td>
            <td><a href="/user/detail?uid={{$v.Uid}}">{{$v.Uid}}</a></td>
            <td><a href="/contest/status/list?cid={{$cid}}&uid={{$v.Uid}}&solved=3">{{$v.Solved}}</a></td>
            <td>{{$v.Time}}</td>
            {{with $v.ProblemList}}
            {{range .}}
            <td>{{if .}}{{.Time}}/({{.Count}}){{else}}0/(0){{end}}</td>
            {{end}}
            {{end}}
          </tr>
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}