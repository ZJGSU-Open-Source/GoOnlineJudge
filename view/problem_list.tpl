{{define "content"}}
<!-- <h1>Problem List</h1> -->

<form accept-charset="UTF-8" id="search_form" >
Search: <input id="search" name="search" size="30" type="text" value="{{.SearchValue}}">
<select id="option" name="option">
  <option value="pid" {{if .SearchPid}}selected{{end}}>ID</option>
  <option value="title" {{if .SearchTitle}}selected{{end}}>Title</option>
  <option value="source" {{if .SearchSource}}selected{{end}}>Source</option>
</select>
<input name="commit" type="submit" value="Go">
</form>

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

<table id="contest_list" class="table table-bordered table-striped table-hover">
  <thead>
    <tr>
      <th class="header">Flag</th>
      <th class="header">ID</th>
      <th class="header">Title</th>
      <th class="header">Ratio(Solve/Submit)</th>
    </tr>
  </thead>
  <tbody>
    {{$time := .Time}}
    {{$privilege := .Privilege}}
    {{with .Problem}}  
      {{range .}} 
        {{if or (ShowStatus .Status) (LargePU $privilege)}}
          {{/*if ShowExpire .Expire $time*/}}
            <tr>
              <td><span class="icon icon-material-check"></span><span class="icon icon-material-clear" ></span></td>
              <td>{{.Pid}}</td>
              <td><a href="/problem/detail?pid={{.Pid}}">{{.Title}}</a></td>
              <td>{{ShowRatio .Solve .Submit}} (<a href="/status/list?pid={{.Pid}}&judge=3">{{.Solve}}</a>/<a href="/status/list?pid={{.Pid}}">{{.Submit}}</a>)</td>
            </tr>
          {{/*end*/}}
        {{end}}
      {{end}}  
    {{end}}
  </tbody>
</table>

<script type="text/javascript">
  $('#search_form').submit( function(e) {
    e.preventDefault();
    var value = $('#search').val();
    var key = $('#option').val();
    value = encodeURIComponent(value);
    window.location.href = '/problem/list?'+key+'='+value;
  });
  </script>
{{end}}
