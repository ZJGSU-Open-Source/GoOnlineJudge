{{define content}}
{{$cid := .Cid}}
<table id="contest_list">
  <thead>
    <tr>
      <th class="header">Rank</th>
      <th class="header">Team</th>
      <th class="header">Solved</th>
      {{with .ProblemList}}
      {{range .}}
      <th class="header"><a href="/contest/problem?detail/cid?{{$cid}}/pid?{{.Pid}}">{{.}}</a></th>
      {{end}}
      {{end}}
    </tr>
  </thead>
  <tbody>
    {{with .UserList}}
    {{$rank := 1}}  
      {{range .}} 
          <tr>
            <td>{{$rank}}</td>
            <td><a href="/user?detail/uid?{{.Uid}}">{{.Uid}}</a></td>
            <td><a href="/contest/status?list/uid?{{.Uid}}/solved?3">{{.Solved}}</a></td>
            {{with .ProblemList}}
            {{range .}}
            <td>{{.Time}}({{.Count}})</td>
            {{$rank = NumAdd $rank 1}}
            {{end}}
            {{end}}
          </tr>
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}