{{define "content"}}
<h1>Problem List</h1>
<form accept-charset="UTF-8" id="search_form">
Search: <input id="search" name="search" size="30" type="text" value="{{.SearchValue}}">
<select id="option" name="option">
  <option value="pid" {{if .SearchPid}}selected{{end}}>ID</option>
  <option value="title" {{if .SearchTitle}}selected{{end}}>Title</option>
  <option value="source" {{if .SearchSource}}selected{{end}}>Source</option>
</select>
<input name="commit" type="submit" value="Go">
</form>
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
  <script type="text/javascript">
  $('#search_form').submit( function(e) {
    e.preventDefault();
    var value = $('#search').val();
    var key = $('#option').val();
    window.location.href = '/problem/list/'+key+'/'+value;
  });
  </script>
{{end}}
