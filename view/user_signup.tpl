{{define "content"}}
	<h1 class="compact">Sign Up</h1>
	<form accept-charset="UTF-8" class="new_user" id="new_user">
		<div style="margin:0;padding:0;display:inline">
			<input name="utf8" type="hidden" value="✓">
		</div>
		<div class="field">
		  	<label for="user_handle">Handle</label><font color="red">*</font>
		  	<font  id="user_warning_handle" color="red"></font><br>
		 	<input id="user_handle" name="user[handle]" size="30" type="text" autofocus required="required" />
		</div>
		<div class="field">
			<label for="user_nick">Nick</label><font color="red">*</font>
			<font id="user_warning_nick" color="red"></font><br>
			<input id="user_nick" name="user[nick]" size="30" type="text">
		</div>
		<div class="field">
			<label for="user_password">Password</label><font color="red">*</font>
			<font id="user_warning_password" color="red"></font><br>
			<input id="user_password" name="user[password]" size="30" type="password" required="required" placeholder="at least six characters.">
		</div>	
		<div class="field">
			<label for="user_confirmPassword">Confirm Password</label><font color="red">*</font>
			<font id="user_warning_confirmPassword" color="red"></font><br>
			<input id="user_confirmPassword" name="user[confirmPassword]" size="30" type="password" required="required">	
		</div>
		<div class="field">
			<label for="user_mail">Email</label>
			<font id="user_warning_mail" color="red"></font><br>
			<input id="user_mail" name="user[mail]" size="30" type="email">	
		</div>
		<div class="field">
			<label for="user_school">School</label>
			<font id="user_warning_school" color="red"></font><br>
			<input id="user_school" name="user[school]" size="30" type="text">	
		</div>
		<div class="field">
			<label for="user_motto">Motto</label>
			<font id="user_warning_motto" color="red"></font><br>
			<input id="user_motto" name="user[motto]" size="30" type="text">
		</div>
		<div class="actions">
	  	<input name="user_signup" type="submit" value="Sign Up">
		</div>
	</form>

	<script type="text/javascript">
	$('#new_user').submit( function(e) {
		e.preventDefault();
		$.ajax({
			type:'POST',
			url:'/user/register',
			data:$(this).serialize(),
			error: function(response) {
				var json = eval('('+response.responseText+')');
				if(json.uid != null) {
					$('#user_handle').css({"border-color": "red"});
					$('#user_warning_handle').text(json.uid);
				} else {
					$('#user_warning_handle').text('');
				}				
				if(json.nick != null) {
					$('#user_nick').css({"border-color": "red"});
					$('#user_warning_nick').text(json.nick);
				} else {
					$('#user_warning_nick').text('');
				}
				if(json.pwd != null) {
					$('#user_password').css({"border-color": "red"});
					$('#user_warning_password').text(json.pwd);
				} else {
					$('#user_warning_password').text('');
				}
				if(json.pwdConfirm != null) {
					$('#user_confirmPassword').css({"border-color": "red"});
					$('#user_warning_confirmPassword').text(json.pwdConfirm);
				} else {
					$('#user_warning_confirmPassword').text('');
				}				
			},
			success: function() {
				window.location.href = '/';
			}
		});
	});
	</script>
{{end}}