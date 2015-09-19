{{define "content"}}
{{$isAdmin := .IsAdmin}}
<div class="p-adminTest mdl-grid">
  
  <div class="mdl-cell mdl-cell--2-col mdl-cell--1-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link">
        <a href="/admin/problems">List</a>
      </div>
      <div class="link">
        <a href="/admin/problems/new">Add</a>
      </div>
      <div class="link">
        <a href="/admin/problems/importor">Import</a>
      </div>
      {{if .RejudgePrivilege}}
      <div class="link">
        <a href="/admin/rejudger">Rejudge</a>
      </div>
      {{end}}
    </div>
  </div>

  <div class="page mdl-cell mdl-cell--8-col mdl-cell--4-col-phone mdl-shadow--2dp J_list">
    {{if $isAdmin}}
    <form accept-charset="UTF-8" enctype="multipart/form-data" method="post" action="/admin/testdata/{{.Pid}}">
      <div class="mdl-grid">
        <div class="mdl-cell mdl-cell--12-col mdl-cell--4-col-phone text-center">
          <input type="file" name="testfiles"/>
          <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored J_submit" type="submit">upload</button>
        </div>        
      </div>
    </form>
    {{end}}

    <div class="table-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone mdl-shadow--2dp">
      <table class="go-table text-center">
        <thead>
          <tr>
            <th>Data File</th>
            {{if $isAdmin}}
            <th>Delete</th>
            {{end}}
            <th>Download</th>
          </tr>
        </thead>
        <tbody>
          {{$Pid := .Pid}}
          {{with .Files}}
          {{range .}}
          <tr>
            <td><a>{{.}}</a></td>
            {{if $isAdmin}}
            <td><a class="J_delete" href="#" data-type="{{.}}">Delete</a></td>
            {{end}}
            <td><a href="/admin/testdata/{{$Pid}}/file?type={{.}}">Download</a></td>
          </tr>
          {{end}}
          {{end}}
        </tbody>
      </table>
    </div>

  </div>
</div>
<script type="text/javascript">
var deleteNode = $('.J_delete');
deleteNode.on('click', function() {
  var self = $(this);
  var type = $(this).data('type');
  if( confirm('Delete the '+ type +'?') ){
    var pid = {{.Pid}}
    $.ajax({
      type: 'DELETE',
      url: '/admin/testdata/' + pid + '?type=' + type,
      data:$(this).serialize(),
      error: function() {
          alert('failed!');
      },
      success: function() {
        self.parent().parent().remove();
        location.reload();
      }
    });
  }
});
</script>
{{end}}
