{{define "content"}}
<!-- <h1>Ranklist</h1> -->
  
<div class="pagination">
  {{$current := .CurrentPage}}
  {{$url := .URL}}
  {{if .IsPreviousPage}}
  <a href="{{$url}}page={{NumSub .CurrentPage 1}}" class="icon icon-material-arrow-back"></a>
  {{else}}
  <span class="icon icon-material-arrow-back"></span>
  {{end}}
  &nbsp;
  {{if .IsPageHead}}
    {{with .PageHeadList}}
      {{range .}}
        {{if eq . $current}}
          <span>{{.}}</span>
        {{else}}
          <a href="{{$url}}page={{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}

  {{if .IsPageMid}}
  ...
    {{with .PageMidList}}
      {{range .}}
        {{if eq . $current}}
          <span>{{.}}</span>
        {{else}}
          <a href="{{$url}}page={{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}

  {{if .IsPageTail}}
  ...
    {{with .PageTailList}}
      {{range .}}
        {{if eq . $current}}
          <span>{{.}}</span>
        {{else}}
          <a href="{{$url}}page={{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}
  &nbsp;
  {{if .IsNextPage}}
  <a href="{{$url}}page={{NumAdd .CurrentPage 1}}" class="icon icon-material-arrow-forward"></a>
  {{else}}
  <span class="icon icon-material-arrow-forward"></span>
  {{end}}
</div>
<div class="table-responsive">
  <table id="ranklist" class="table table-bordered table-striped table-hover">
    <thead>
      <tr>
        <th class="header">Rank</th>
        <th class="header">User</th>
        <th class="header">Motto</th>
        <th class="header">Ratio(Solve/Submit)</th>
      </tr>
    </thead>
    <tbody>
      {{with .User}}
        {{range .}}
          {{if ShowStatus .Status}}
            <tr>
              <td>{{.Index}}</td>
              <td><a href="/user/detail?uid={{.Uid}}" class="btn btn-flat btn-info btn-xs">{{.Uid}}</a></td>
              <td id="motto" >{{.Motto}}</td>
              <td>{{ShowRatio .Solve .Submit}} (<a href="/status/list?uid={{.Uid}}&judge=3">{{.Solve}}</a>/<a href="/status/list?uid={{.Uid}}">{{.Submit}}</a>)</td>
            </tr>
          {{end}}
        {{end}}  
      {{end}}
    </tbody>
  </table>
 </div>
{{end}}