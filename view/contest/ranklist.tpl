{{define "content"}}
{{$cid := .Cid}}
<h1>{{.Contest}}</h1>
<h5><a class="ranklist_export" href="#">Export ranklist to csv</a></h5>
<table id="contest_list">
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

<script type="text/javascript">
$('.ranklist_export').on('click', function() {
  var ret = confirm('Export this ranklist to csv?');
  if (ret == true) {
    var table = document.getElementById("contest_list");


    $.ajax({
      type: 'POST',
      url: '/contest/exportranklist/cid?='+{{.Cid}},
      data:$(this).serialize(),
      error: function(response) {
        var json = eval('('+response.responseText+')');
        if (json.hint != null) {
          alert(json.hint);
        } else {
          alert('failed!');
        }
      },
      success: function(response) {
        window.location.reload();
      }
    });

  }
});
</script>

{{end}}