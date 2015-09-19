{{define "content"}}
	<script src="/static/js/bootstrap.min.js"></script>
  	<script src="/static/material/js/material.min.js"></script>
	<script type="text/javascript">
	$('#new_user').submit( function(e) {
		e.preventDefault();
		$.ajax({
			type:'POST',
			url:'/account',
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
				alert("Success");
				window.location.href = '/users/'+json.uid;
			}
		});
	});
	</script>

<div class="p-signin mdl-grid">
  <div class="mdl-cell mdl-cell--2-col mdl-cell--hide-phone mdl-cell--hide-tablet"></div>
  <div class="page mdl-cell mdl-cell--8-col mdl-cell--4-col-phone mdl-shadow--2dp mdl-grid">
    <form accept-charset="UTF-8" class="J_addForm mdl-cell mdl-cell--12-col mdl-cell--4-col-phone" action="/account" method="post">
      <div class="go-title-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <div class="title">Edit Password</div>
      </div>
      <input name="utf8" type="hidden" value="✓">
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="password" id="user_handle" name="user[oldPassword]"/>
        <label class="mdl-textfield__label" for="user_handle">Old Password</label>
      </div>
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="password" id="user_password" name="user[newPassword]"/>
        <label class="mdl-textfield__label" for="user_password">New Password</label>
      </div>
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="password" id="user_password" name="user[confirmPassword]"/>
        <label class="mdl-textfield__label" for="user_password">Confirm Password</label>
      </div>
      <div class="btn-area mdl-cell--12-col mdl-cell--4-col-phone">
        <!-- Accent-colored raised button with ripple -->
        <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored J_submit" type="submit">submit</button>
      </div>
    </form>
  </div>
</div>
{{end}}

