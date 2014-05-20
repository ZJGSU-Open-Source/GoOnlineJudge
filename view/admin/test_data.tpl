{{define "content"}}
<h1>{{.Title}}</h1>
	<table id="contest_list">
  <thead>
    <tr>
      <th class="header">Data File</th>
      <th class="header">Delete</th>
      <th class="header">Get</th>
    </tr>
  </thead>
  <tbody>
  {{range .Files}}
        <tr>
          <td>{{.}}</td>
          <td><a>[Delete]</a></td>
          <td><a>[Get]</a></td>
        </tr>
  {{end}}
  </tbody>
</table>
<form name="uplodafiles" enctype="multipart/form-data" method="post" action="/admin/testdata/upload/pid/{{.Pid}}">
<label><input type="file" multiple="" size="80" name="testfiles"/> <input type="submit" value="upload" /> </label>
</form>
	<div class="flash notice">You can just add test.in and test.out</div>
{{end}}