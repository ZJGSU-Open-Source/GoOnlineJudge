{{define "content"}}
	<h1 class="compact">Edit Password</h1>
	<form accept-charset="UTF-8" class="new_user" id="new_user">
		<div style="margin:0;padding:0;display:inline">
			<input name="utf8" type="hidden" value="✓">
		</div>
		<div class="field">
			<label for="user_Handle">User Handle</label><font color="red">*</font>
			<font id="user_warning_Handle" color="red"></font><br>
			<input id="user_Handle" name="user[Handle]" size="30" type="text">
		</div>	
		<div class="field">
			<label for="user_newPassword">New Password</label><font color="red">*</font>
			<font id="user_warning_newPassword" color="red"></font><br>
			<input id="user_newPassword" name="user[newPassword]" size="30" type="password" placeholder="at least six characters.">
		</div>	
		<div class="field">
			<label for="user_confirmPassword">Confirm Password</label><font color="red">*</font>
			<font id="user_warning_confirmPassword" color="red"></font><br>
			<input id="user_confirmPassword" name="user[confirmPassword]" size="30" type="password">	
		</div>		
		<div class="actions">
	  		<input name="user_password" type="submit" value="Edit">
		</div>
	</form>

	<script type="text/javascript">
	$('#new_user').submit( function(e) {
		e.preventDefault();
		$.ajax({
			type:'PUT',
			url:'/admin/users/password',
			data:$(this).serialize(),
			error: function(response) {
				var json = eval('('+response.responseText+')');
				if(json.uid != null) {
					$('#user_Handle').css({"border-color": "red"});
					$('#user_warning_Handle').text(json.uid);
				} else {
					$('#user_warning_Handle').text('');
				}
				if(json.newPassword != null) {
					$('#user_newPassword').css({"border-color": "red"});
					$('#user_warning_newPassword').text(json.newPassword);
				} else {
					$('#user_warning_newPassword').text('');
				}
				if(json.confirmPassword != null) {
					$('#user_confirmPassword').css({"border-color": "red"});
					$('#user_warning_confirmPassword').text(json.confirmPassword);
				} else {
					$('#user_warning_confirmPassword').text('');
				}	
			},
			success: function(response) {
				//var json = eval('('+response+')');
				alert("Success");
				window.location.href='/admin/users'
			}
		});
	});
	</script>
{{end}}