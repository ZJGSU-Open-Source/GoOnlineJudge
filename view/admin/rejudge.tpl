{{define "content"}}
<h1>Rejudge</h1>

<form accept-charset="UTF-8" id="search_form">
<select id="type" name="type">
<option value="Sid">Solution ID</option>
<option value="Pid">Problem ID</option>
</select>
<br>
<input id="id" name="id" size="20" type="text" pattern="^[0-9]+" required>
<div class="actions">
	<input name="rejudgebutton" type="submit" value="Rejudge" required="required" />
</div>
</form>

<script type="text/javascript">
$('#search_form').submit( function(e) {
	e.preventDefault();
	var id = $('#id').val();
	var type = $('#type').val();
	$.ajax({
		type:'POST',
		url:'/admin/rejudger?type='+type+'&id='+id,
		data:$(this).serialize(),
		error:function(response){
			var json = eval('('+response.responseText+')');
			if(json.info != null) {
				alert(json.info);
			}
		},
		success:function(response){
			alert("Rejudge Complete")
			//window.location.reload();
			window.location.href = '/status'
		}
	});
});
</script>
{{end}}