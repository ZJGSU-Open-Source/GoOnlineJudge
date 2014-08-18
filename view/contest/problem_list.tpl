{{define "content"}}
<h1>{{.Contest}}</h1>
<div class="flash notice">Start@{{ShowTime .Start}} End@{{ShowTime .End}}</div>
<table id="contest_list">
  <thead>
    <tr>
      <th class="header">ID</th>
      <th class="header">Title</th>
      <th class="header">Ratio(Accept/Submit)</th>
    </tr>
  </thead>
  <tbody>
    {{$cid := .Cid}}
    {{with .Problem}}  
      {{range .}} 
            <tr>
              <td>{{.Pid}}</td>
              <td><a href="/contest/problem?detail/cid?{{$cid}}/pid?{{.Pid}}">{{.Title}}</a></td>
              <td>{{ShowRatio .Solve .Submit}} ({{.Solve}}/{{.Submit}})</td>
            </tr>
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}