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
          <td><a class="testdata_delete" href="#">[Delete]</a></td>
          <td><a>[Download]</a></td>
        </tr>
        {{end}}

        {{if .Files.testout}}
          <td><a>{{.Files.testout}}</a></td>
          <td><a class="testdata_delete" href="#">[Delete]</a></td>
          <td><a>[Download]</a></td>
        </tr>
        {{end}}
  </tbody>
</table>
<form name="uplodafiles" enctype="multipart/form-data" method="post" action="/admin/testdata/upload/pid/{{.Pid}}">
<label><input type="file" multiple="" size="80" name="testfiles"/> <input type="submit" value="上传" /> </label>
</form>
	<div class="flash notice">You can just add test.in and test.out</div>

<script type="text/javascript">
$('.testdata_delete').on('click', function() {
	var ret = confirm('Delete the Testdata?');
	 if (ret == true) {
	 	//这里考虑直接通过js代码来删除指定目录下的文件，而不用通过数据库的操作
	 	//alert('Yes!');
	 	//var inpath = $(this).data('testin_path');
		//var outpath = $(this).data('testout_path');
		/*var fso, file;
		fso = new ActiveXObject("Scripting.FileSystemObject"); 
		if (inpath){
			file = fso.GetFile (inpath);
			file.Delete();
			
		} else if (outpath) {
			file = fso.GetFile(outpath);
			file.Delete();
		}
		*/
	 }
});
</script>
{{end}}