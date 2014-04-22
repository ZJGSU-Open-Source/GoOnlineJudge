{{define "content"}}
	<h1 class="compact">Edit Password</h1>
	<form accept-charset="UTF-8" class="new_user" id="new_user">
		<div style="margin:0;padding:0;display:inline">
			<input name="utf8" type="hidden" value="✓">
		</div>
		<div class="field">
			<label for="user_oldPassword">Old Password</label><font color="red">*</font>
			<font id="user_warning_oldPassword" color="red"></font><br>
			<input id="user_oldPassword" name="user[oldPassword]" size="30" type="password">
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
			url:'/user/password',
			data:$(this).serialize(),
			error: function(response) {
				var json = eval('('+response.responseText+')');
				if(json.oldPassword != null) {
					$('#user_oldPassword').css({"border-color": "red"});
					$('#user_warning_oldPassword').text(json.oldPassword);
				} else {
					$('#user_warning_oldPassword').text('');
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
			}
		});
	});
	</script>
{{end}}