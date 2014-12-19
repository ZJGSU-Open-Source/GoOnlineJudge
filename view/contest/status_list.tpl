{{define "content"}}
<h1 style="text-align:center">Contest Status -- {{.Contest}}</h1>

<form accept-charset="UTF-8" id="search_form" class="form-inline">
  <span> User: </span><input id="search_uid" name="search_uid" size="20" type="text" value="{{.SearchUid}}">
  <span style="margin-left:10px">Problem: </span><input id="search_pid" name="search_pid" size="10" type="text" value="{{.SearchPid}}">
  <span style="margin-left:10px">Result: </span>
    <select id="search_judge" name="search_judge">
      <option value="0">All</option>
      <option value="1" {{if .SearchJudge0}}selected{{end}}>Pending</option>
      <option value="2" {{if .SearchJudge1}}selected{{end}}>Running &amp;Judging</option>
      <option value="3" {{if .SearchJudge2}}selected{{end}}>Compile Error</option>
      <option value="4" {{if .SearchJudge3}}selected{{end}}>Accepted</option>
      <option value="5" {{if .SearchJudge4}}selected{{end}}>Runtime Error</option>
      <option value="6" {{if .SearchJudge5}}selected{{end}}>Wrong Answer</option>
      <option value="7" {{if .SearchJudge6}}selected{{end}}>Time Limit Exceeded</option>
      <option value="8" {{if .SearchJudge7}}selected{{end}}>Memory Limit Exceeded</option>
      <option value="9" {{if .SearchJudge8}}selected{{end}}>Output Limit Exceeded</option>
      <option value="10" {{if .SearchJudge9}}selected{{end}}>Presentation Error</option>
      <option value="11" {{if .SearchJudge10}}selected{{end}}>System Error</option>
    </select>
    <span style="margin-left:10px">Language: </span>
    <select id="search_language" name="search_language">
      <option value="0" {{if .SearchLanguage0}}selected{{end}}>All</option>
      <option value="1" {{if .SearchLanguage1}}selected{{end}}>C</option>
      <option value="2" {{if .SearchLanguage2}}selected{{end}}>C++</option>
      <option value="3" {{if .SearchLanguage3}}selected{{end}}>Java</option>
    </select>
  <input name="commit" type="submit" value="Go" style="margin-left:10px">
</form>
<script type="text/javascript">
  $('#search_form').submit( function(e) {
    e.preventDefault();
    var uid = $('#search_uid').val();
    var pid = $('#search_pid').val();
    var judge = $('#search_judge').val();
    var language = $('#search_language').val();
    var url = 'status?';
    if (uid != '')
      url += 'uid=' + uid + "&";
    if (pid != '')
      url += 'pid=' + pid + "&";
    if (judge > 0){
      judge = judge-1;
      url += 'judge=' + judge + "&";
    }
    if (language > 0)
      url += 'language=' + language;
    window.location.href = url;
  });
  </script>

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

<table id="contest_list" class="table table-striped table-hover">
  <thead>
    <tr>
      <th class="header">ID</th>
      <th class="header">User</th>
      <th class="header">Problem</th>
      <th class="header">Result</th>
      <th class="header">Time</th>
      <th class="header">Memory</th>
      <th class="header">Language</th>
      <th class="header">Code Length</th>
      <th class="header">Submit Time</th>
    </tr>
  </thead>
  <tbody>
    {{$cid := .Cid}}
    {{$privilege := .Privilege}}
    {{$uid := .CurrentUser}}
    
    {{with .Solution}}  
      {{range .}} 
        {{if ShowStatus .Status}} 
          <tr>

            <td>{{.Sid}}</td>
            <td><a href="/users/{{.Uid}}">{{.Uid}}</a></td>
            <td><a href="/contests/{{$cid}}/problems/{{.Pid}}">{{.Pid}}</a></td>
            <td><span class="submitRes-{{.Judge}}">{{if (LargePU $privilege)}}{{if .Sim}}*{{end}}{{end}}{{ShowJudge .Judge}}</span>
            {{if (LargePU $privilege)}} {{if .Sim}} ID:{{.Sim_s_id}} [{{.Sim}}％] {{end}} {{end}}</td>
            <td>{{.Time}}ms</td>
            <td>{{.Memory}}kB</td>
            <td>{{ShowLanguage .Language}}{{if or (eq .Uid $uid) (LargePU $privilege)}}<a href="/contests/{{$cid}}/status/{{.Sid}}/code">[view]</a>{{end}}</td>
            <td>{{.Length}}B</td>
            <td>{{ShowTime .Create}}</td>
          </tr>
        {{end}}
      {{end}}  
    {{end}}
  </tbody>
</table>
{{end}}
