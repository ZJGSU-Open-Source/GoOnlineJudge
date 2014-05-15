{{define "content"}}
<h1>Ranklist</h1>
  
<div class="pagination">
  {{$current := .CurrentPage}}
  {{$url := .URL}}
  {{if .IsPreviousPage}}
  <a href="{{$url}}/page/{{NumSub .CurrentPage 1}}">Prev</a>
  {{else}}
  <span>Prev</span>
  {{end}}

  {{if .IsPageHead}}
    {{with .PageHeadList}}
      {{range .}}
        {{if NumEqual . $current}}
          <span>{{.}}</span>
        {{else}}
          <a href="{{$url}}/page/{{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}

  {{if .IsPageMid}}
  ...
    {{with .PageMidList}}
      {{range .}}
        {{if NumEqual . $current}}
          <span>{{.}}</span>
        {{else}}
          <a href="{{$url}}/page/{{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}

  {{if .IsPageTail}}
  ...
    {{with .PageTailList}}
      {{range .}}
        {{if NumEqual . $current}}
          <span>{{.}}</span>
        {{else}}
          <a href="{{$url}}/page/{{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}

  {{if .IsNextPage}}
  <a href="{{$url}}/page/{{NumAdd .CurrentPage 1}}">Next</a>
  {{else}}
  <span>Next</span>
  {{end}}
</div>

  <table id="ranklist">
    <thead>
      <tr>
        <th class="header">Rank</th>
        <th class="header">User</th>
        <th class="header">Motto</th>
        <th class="header">Ratio</th>
      </tr>
    </thead>
    <tbody>
      {{with .User}}  
        {{range .}}
          {{if ShowStatus .Status}}
            <tr>
              <td>{{.Index}}</td>
              <td><a href="/user/detail/uid/{{.Uid}}">{{.Uid}}</a></td>
              <td>{{.Motto}}</td>
              <td>{{ShowRatio .Solve .Submit}} (<a href="/status/list/uid/{{.Uid}}/judge/3">{{.Solve}}</a>/<a href="/status/list/uid/{{.Uid}}">{{.Submit}}</a>)</td>
            </tr>
          {{end}}
        {{end}}  
      {{end}}
    </tbody>
  </table>
{{end}}