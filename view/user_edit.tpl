{{define "content"}}
	<h1 class="compact">Sign Up</h1>
	{{with .Detail}}
		<form accept-charset="UTF-8" class="new_user" id="new_user">
		<div style="margin:0;padding:0;display:inline">
			<input name="utf8" type="hidden" value="✓">
		</div>
			<div class="field">
			  	<label for="user_handle">Handle</label><br>
			 	<input id="user_handle" name="user[handle]" size="30" type="text" value="{{.Uid}}" readonly>
			</div>
			<div class="field">
				<label for="user_nick">Nick</label><font color="red">*</font>
				<font id="user_warning_nick" color="red"></font><br>
				<input id="user_nick" name="user[nick]" size="30" type="text" value="{{.Nick}}">
			</div>
			<div class="field">
				<label for="user_mail">Email</label>
				<font id="user_warning_mail" color="red"></font><br>
				<input id="user_mail" name="user[mail]" size="30" type="text" value="{{.Mail}}">	
			</div>
			<div class="field">
				<label for="user_school">School</label>
				<font id="user_warning_school" color="red"></font><br>
				<input id="user_school" name="user[school]" size="30" type="text" value="{{.School}}">	
			</div>
			<div class="field">
				<label for="user_motto">Motto</label>
				<font id="user_warning_motto" color="red"></font><br>
				<input id="user_motto" name="user[motto]" size="30" type="text" value="{{.Motto}}">
			</div>
		
			<div class="actions">
		  		<input name="user_edit" type="submit" value="Edit">
			</div>
		</form>
	{{end}}
	<script type="text/javascript">
	$('#new_user').submit( function(e) {
		e.preventDefault();
		$.ajax({
			type:'POST',
			url:'/user/update',
			data:$(this).serialize(),
			error: function(response) {
				var json = eval('('+response.responseText+')');	
				if(json.nick != null) {
					$('#user_nick').css({"border-color": "red"});
					$('#user_warning_nick').text(json.nick);
				} else {
					$('#user_warning_nick').text('');
				}			
			},
			success: function(result) {
				var json = eval('('+result+')');
				window.location.href = '/user/detail/uid/'+json.uid;
			}
		});
	});
	</script>
{{end}}