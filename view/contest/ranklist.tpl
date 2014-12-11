{{define "content"}}
{{$cid := .Cid}}
<h1 style="text-align:center">Contest RankList -- {{.Contest}}</h1>
<h5><a href="/contests/{{.Cid}}/rankfile">Export ranklist</a></h5>
<table id="contest_list" class="table table-bordered table-striped table-hover">
  <thead>
    <tr>
      <th class="header">Rank</th>
      <th class="header">Team</th>
      <th class="header">Solved</th>
      <th class="header">Penalty</th>
      {{with .ProblemList}}
      {{range $idx,$pid:= .}}
      <th class="header"><a href="/contests/{{$cid}}/problems/{{$idx}}">{{$idx}}</a></th>
      {{end}}
      {{end}}
    </tr>
  </thead>
  <tbody>
    {{with .UserList}}
      {{range $idx,$v := .}} 
          <tr>
            <td>{{NumAdd $idx 1}}</td>
            <td><a href="/users/{{$v.Uid}}">{{$v.Uid}}</a></td>
            <td><a href="/contests/{{$cid}}/status?uid={{$v.Uid}}&judge=3">{{$v.Solved}}</a></td>
            <td>{{ShowGapTime $v.Time}}</td>
            {{with $v.ProblemList}}
            {{range .}}
            {{if .}}
              {{if eq .Judge 3}}<td id="ac"> {{ShowGapTime .Time}}/({{.Count}})</td>
              {{else}}<td>0/({{.Count}})</td>
              {{end}}
            {{else}}<td>0/(0){{end}}</td>
            {{end}}
            {{end}}
          </tr>
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}