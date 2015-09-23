{{define "content"}}
<div id="contestinfo">
<h1>{{.Contest}}</h1>
<p>Start Time : {{ShowTime .Start}} &nbsp;&nbsp;End Time : {{ShowTime .End}}</p>
<p>Current Time : {{ShowTime .Time}}</p>
</div>
<table id="contest_list" class="table table-bordered table-striped table-hover">
  <thead>
    <tr>
      <th class="header">Flag</th>
      <th class="header">ID</th>
      <th class="header">Title</th>
      <th class="header">Ratio(Accept/Submit)</th>
    </tr>
  </thead>
  <tbody>
    {{$cid := .Cid}}
    {{with .Problem}}  
      {{range .}}
      {{if .}} 
            <tr>
              <td>
              {{if ShowACFlag .Flag}} <span class="icon icon-material-check"></span>
              {{else if ShowErrFlag .Flag}}
              <span class="icon icon-material-clear"></span>
              {{end}}
              </td>
              <td>{{.Pid}}</td>
              <td><a href="/contests/{{$cid}}/problems/{{.Pid}}">{{.Title}}</a></td>
              <td>{{ShowRatio .Solve .Submit}} ({{.Solve}}/{{.Submit}})</td>
            </tr>
      {{end}}
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}