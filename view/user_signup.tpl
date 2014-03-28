{{define "content"}}
<h1 class="compact">Sign Up</h1>
<form accept-charset="UTF-8" class="new_user" id="new_user" method="post">
  <div style="margin:0;padding:0;display:inline">
    <input name="utf8" type="hidden" value="✓">
  </div>
	<div class="field">
  	<label for="user_handle">Handle</label><br>
 		<input id="user_handle" name="user[handle]" size="30" type="text">
 		<label>*</label>
 		<label id="user_warning_handle"></label>
	</div>
	<div class="field">
		<label for="user_nick">Nick</label><br>
		<input id="user_nick" name="user[nick]" size="30" type="text">
		<label>*</label>
		<label id="user_warning_nick"></label>
	</div>
	<div class="field">
		<label for="user_password">Password</label><br>
		<input id="user_password" name="user[password]" size="30" type="text">
		<label>*</label>
		<label id="user_warning_password">Password should contain at least six characters</label>
	</div>
	<div class="field">
		<label for="user_confirmPassword">Confirm Password</label><br>
		<input id="user_confirmPassword" name="user[confirmPassword]" size="30" type="text">
		<label>*</label>
		<label id="user_warning_confirmPassword"></label>
	</div>
	<div class="field">
		<label for="user_email">Email</label><br>
		<input id="user_email" name="user[email]" size="30" type="text">
		<label></label>
		<label id="user_warning_email"></label>
	</div>
	<div class="field">
		<label for="user_school">School</label><br>
		<input id="user_school" name="user[school]" size="30" type="text">
		<label></label>
		<label id="user_warning_school"></label>
	</div>
	<div class="field">
		<label for="user_motto">Motto</label><br>
		<input id="user_motto" name="user[motto]" size="30" type="text">
		<label></label>
		<label id="user_warning_motto"></label>
	</div>
	<div class="actions">
  	<input name="user_signup" type="submit" value="Sign Up">
	</div>
</form>
{{end}}