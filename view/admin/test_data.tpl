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
        <tr>
        {{if .Files.testin}}
          <td><a>{{.Files.testin}}</a></td>
          <td><a class="testdata_in_delete" href="#">[Delete]</a></td>
          <td><input type="submit" value="Download" class="testdata_in_download" /></td>
        </tr>
        {{end}}

        {{if .Files.testout}}
          <td><a>{{.Files.testout}}</a></td>
          <td><a class="testdata_out_delete" href="#">[Delete]</a></td>
          <td><input type="submit" value="Download" class="testdata_out_download" /></td>
        </tr>
        {{end}}
  </tbody>
</table>
<form name="uploadfiles" enctype="multipart/form-data" method="post" action="/admin/testdata/upload/pid/{{.Pid}}">
<label><input type="file" multiple="" size="80" name="testfiles"/> <input type="submit" value="上传" /> </label>
</form>
	<div class="flash notice">You can just add test.in and test.out</div>

<script type="text/javascript">
$('.testdata_in_download').on('click', function() {
	var pid = {{.Pid}}
	$.ajax({
                type: 'POST',
                url: '/admin/testdata/Downloadin/pid/' + pid,
                data:$(this).serialize(),
                error: function() {
                	alert('failed!');
                },
                success: function() {
                	alert('Downloaded at $HOME directory');
                	window.location.reload();
                }
            });
});
</script>

<script type="text/javascript">
$('.testdata_out_download').on('click', function() {
	var pid = {{.Pid}}
	$.ajax({
                type: 'POST',
                url: '/admin/testdata/Downloadout/pid/' + pid,
                data:$(this).serialize(),
                error: function() {
                	alert('failed!');
                },
                success: function() {
                	alert('Downloaded at $HOME directory');
                	window.location.reload();
                }
            });
});
</script>

<script type="text/javascript">
$('.testdata_in_delete').on('click', function() {
	var ret = confirm('Delete the Testdata?');
	 if (ret == true) {
               var pid = {{.Pid}}
               $.ajax({
                type: 'POST',
                url: '/admin/testdata/Deletein/pid/' + pid,
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

<script type="text/javascript">
$('.testdata_out_delete').on('click', function() {
  var ret = confirm('Delete the Testdata?');
   if (ret == true) {
               var pid = {{.Pid}}
               $.ajax({
                type: 'POST',
                url: '/admin/testdata/Deleteout/pid/' + pid,
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
