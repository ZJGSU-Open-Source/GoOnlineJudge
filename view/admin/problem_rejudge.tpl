{{define "content"}}
<h1>Rejudge</h1>

<form accept-charset="UTF-8" id="search_form">
<select id="type" name="type">
<option value="Pid">Problem ID</option>
<option value="Sid">Solution ID</option>
</select>
<br>
<input id="id" name="id" size="20" type="text">
<div class="actions">
	<input name="user_password" type="submit" value="Rejudge">
</div>
</form>

<script type="text/javascript">
$('#search_form').submit( function(e) {
	e.preventDefault();
	var id = $('#id').val();
	var type = $('#type').val();

	if (id == "") {
		alert("id should not be empty!")
	}
	
	$.ajax({
		type:'POST',
		url:'/problem?rejudge/type?'+type+'/id?'+id,
		data:$(this).serialize(),
		error:function(response){
			var json = eval('('+response.responseText+')');
			if(json.uid != null) {
				alert(json.uid);
			}
		},
		success:function(response){
			alert("Rejudge Complete")
			//window.location.reload();
			window.location.href = '/status?list'
		}
	});
});
</script>
{{end}}