{{define "content"}}
<h1>Problem List</h1>
<div class="pagination">
  {{.CurrentPage}}
  {{if .IsPreviousPage}}
  <a href="/problem/list/page/{{NumSub .CurrentPage 1}}">Prev</a>
  {{else}}
  <span>Prev</span>
  {{end}}

  {{if .IsPageHead}}
    {{with .PageHeadList}}
      {{range .}}
          <a href="/problem/list/page/{{.}}">{{.}}</a>
      {{end}}
    {{end}}
  {{end}}

  {{if .IsPageMid}}
  ...
    {{with .PageMidList}}
      {{range .}}
        <a href="/problem/list/page/{{.}}">{{.}}</a>
      {{end}}
    {{end}}
  {{end}}

  {{if .IsPageTail}}
  ...
    {{with .PageTailList}}
      {{range .}}
        <a href="/problem/list/page/{{.}}">{{.}}</a>
      {{end}}
    {{end}}
  {{end}}

  {{if .IsNextPage}}
  <a href="/problem/list/page/{{NumAdd .CurrentPage 1}}">Next</a>
  {{else}}
  <span>Next</span>
  {{end}}
</div>
<table id="contest_list">
  <thead>
    <tr>
      <th class="header">ID</th>
      <th class="header">Title</th>
      <th class="header">Ratio</th>
    </tr>
  </thead>
  <tbody>
    {{$time := .Time}}
    {{with .Problem}}  
      {{range .}} 
        {{if ShowStatus .Status}}
          {{if ShowExpire .Expire $time}}
            <tr>
              <td>{{.Pid}}</td>
              <td><a href="/problem/detail/pid/{{.Pid}}">{{.Title}}</a></td>
              <td>{{ShowRatio .Solve .Submit}} ({{.Solve}}/{{.Submit}})</td>
            </tr>
          {{end}}
        {{end}}
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}
