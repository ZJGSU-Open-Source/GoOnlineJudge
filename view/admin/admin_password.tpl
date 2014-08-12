{{define "content"}}
	<h1 class="compact">Edit Password</h1>
	<form accept-charset="UTF-8" class="new_user" id="new_user">
		<div style="margin:0;padding:0;display:inline">
			<input name="utf8" type="hidden" value="✓">
		</div>
		<div class="field">
			<label for="user_Handler">User Handler</label><font color="red">*</font>
			<font id="user_warning_Handler" color="red"></font><br>
			<input id="user_Handler" name="user[Handler]" size="30" type="text">
		</div>	
		<div class="field">
			<label for="user_newPassword">New Password</label><font color="red">*</font>
			<font id="user_warning_newPassword" color="red">Password should contain at least six characters.</font><br>
			<input id="user_newPassword" name="user[newPassword]" size="30" type="password">
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
			type:'POST',
			url:'/admin/user/password',
			data:$(this).serialize(),
			error: function(response) {
				var json = eval('('+response.responseText+')');
				if(json.Handler != null) {
					$('#user_Handler').css({"border-color": "red"});
					$('#user_warning_Handler').text(json.Handler);
				} else {
					$('#user_warning_Handler').text('');
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
				var json = eval('('+response+')');
				window.location.href = '/user/detail/uid/'+json.uid;
				//window.location.href='/user/list'
			}
		});
	});
	</script>
{{end}}