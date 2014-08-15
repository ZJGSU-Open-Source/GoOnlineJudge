{{define "content"}}
<h1>{{.Title}}</h1>
	<table id="contest_list">
  <thead>
    <tr>
      <th class="header">Data File</th>
      <th class="header">Delete</th>
      <th class="header">Download</th>
    </tr>
  </thead>
  <tbody>
    {{$Pid := .Pid}}
    {{with .Files}}
    {{range .}}
    <tr>
      <td><a>{{.}}</a></td>
      <td><a class="testdata_delete" href="#" data-type="{{.}}">[Delete]</a></td>
      <td><a href="/admin/testdata?download/pid?{{$Pid}}/type?{{.}}">[Download] </a></td>
    </tr>
    {{end}}
    {{end}}     
  </tbody>
</table>
<form name="uploadfiles" enctype="multipart/form-data" method="post" action="/admin/testdata?upload/pid?{{.Pid}}">
<div class="actions">
<label><input type="file" multiple="" size="80" name="testfiles" style="background-color:white;color:black" />
<input name="commit"type="submit" value="upload" /> </label>
</div>
</form>
	<div class="flash notice">You can just add test.in and test.out</div>

<script type="text/javascript">
$('.testdata_delete').on('click', function() {
  var type = $(this).data('type');
  var ret = confirm('Delete the '+ type +'?');
   if (ret == true) {
               var pid = {{.Pid}}
               $.ajax({
                type: 'POST',
                url: '/admin/testdata?delete/pid?' + pid + '/type?' + type,
                data:$(this).serialize(),
                error: function() {
                    alert('failed!');
                },
                success: function() {
                    window.location.reload();
                }
            });
   }
});
</script>
{{end}}
