{{define "content"}}
<!-- <h1>Ranklist</h1> -->
<div class="pagination">
  {{$current := .CurrentPage}}
  {{$url := .URL}}
  {{if .IsPreviousPage}}
  <a href="{{$url}}page={{NumSub .CurrentPage 1}}" class="icon-material-arrow-back"></a>
  {{else}}
  <span class="icon-material-arrow-back"></span>
  {{end}}
  &nbsp;&nbsp;&nbsp;
  {{if .IsNextPage}}
  <a href="{{$url}}page={{NumAdd .CurrentPage 1}}" class="icon-material-arrow-forward"></a>
  {{else}}
  <span class="icon-material-arrow-forward"></span>
  {{end}}
</div>
<div class="table-responsive">
  <table id="ranklist" class="table table-bordered table-striped table-hover">
    <thead>
      <tr>
        <th class="header">Rank</th>
        <th class="header">User</th>
        <th class="header">Nick</th>
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
              <td><a href="/users/{{.Uid}}">{{.Uid}}</a></td>
              <td>{{.Nick}}</td>
              <td id="motto" >{{.Motto}}</td>
              <td>{{ShowRatio .Solve .Submit}} (<a href="/status?uid={{.Uid}}&judge=3">{{.Solve}}</a>/<a href="/status?uid={{.Uid}}">{{.Submit}}</a>)</td>
            </tr>
          {{end}}
        {{end}}  
      {{end}}
    </tbody>
  </table>
 </div>
{{end}}