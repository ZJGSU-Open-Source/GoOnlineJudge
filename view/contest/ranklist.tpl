{{define "content"}}
{{$cid := .Cid}}
<h1 style="text-align:center">Contest RankList -- {{.Contest}}</h1>
<h5><a href="/contest/ranklist/download?cid={{.Cid}}">Export ranklist</a></h5>
<table id="contest_list" class="table table-bordered table-striped table-hover">
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
            <td>{{ShowGapTime $v.Time}}</td>
            {{with $v.ProblemList}}
            {{range .}}
            <td>{{if .}}{{if eq .Judge 3}}{{ShowGapTime .Time}}{{else}}0{{end}}/({{.Count}}){{else}}0/(0){{end}}</td>
            {{end}}
            {{end}}
          </tr>
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}