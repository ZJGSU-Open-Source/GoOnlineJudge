{{define "content"}}
<h1>Problem Import</h1>
<form name="uploadfiles" enctype="multipart/form-data" method="post" action="/admin/problems/importor">
<div class="actions">
<label><input type="file" multiple="" size="80" name="fps.xml" style="background-color:white;color:black" />
<input name="commit" type="submit" value="upload" /> </label>
</div>
</form>

{{end}}