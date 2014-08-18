{{define "content"}}
{{$cid := .Cid}}
<table id="contest_list">
  <thead>
    <tr>
      <th class="header">Rank</th>
      <th class="header">Team</th>
      <th class="header">Solved</th>
      {{with .ProblemList}}
      {{range .}}
      <th class="header"><a href="/contest/problem?detail/cid?{{$cid}}/pid?{{.}}">{{.}}</a></th>
      {{end}}
      {{end}}
    </tr>
  </thead>
  <tbody>
    {{with .UserList}}
      {{range $idx,$v := .}} 
          <tr>
            <td>{{NumAdd $idx 1}}</td>
            <td><a href="/user?detail/uid?{{$v.Uid}}">{{$v.Uid}}</a></td>
            <td><a href="/contest/status?list/cid?{{$cid}}/uid?{{$v.Uid}}/solved?3">{{$v.Solved}}</a></td>
            {{with $v.ProblemList}}
            {{range .}}
            <td>{{.Time}}({{.Count}})</td>
            {{end}}
            {{end}}
          </tr>
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}