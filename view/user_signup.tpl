{{define "content"}}
    <div class="row">
      <div class="col-lg-6">

        <form accept-charset="UTF-8" class="new_user form-horizontal" id="new_user">
          <legend>Sign Up</legend>
          <div style="margin:0;padding:0;display:inline">
            <input name="utf8" type="hidden" value="✓">
          </div>

          <div class="form-group">
            <label for="user_handle" class="col-lg-2 control-label">Handle<font color="red">*</font>
		      	</label>
            <div class="col-lg-10">
              <input type="text" class="form-control" id="user_handle" name="user[handle]" placeholder="Handle" pattern="\w+" autofocus required>
            </div>
            <div class="col-lg-10">
                  <font  id="user_warning_handle" color="red"></font>
            </div>
          </div>

		    	<div class="form-group">
            <label for="user_nick" class="col-lg-2 control-label">Nick<font color="red">*</font>
            </label>
            <div class="col-lg-10">
              <input type="text" class="form-control" id="user_nick" name="user[nick]" placeholder="Nick" required>
            </div>
            <div class="col-lg-10">
            <font  id="user_warning_nick" color="red"></font>              
            </div>
          </div>

          <div class="form-group">
            <label for="user_password" class="col-lg-2 control-label">Password<font color="red">*</font>
			     </label>
            <div class="col-lg-10">
              <input type="password" class="form-control" id="user_password" name="user[password]" required placeholder="At least six characters.">
            </div>
            <div class="col-lg-10">
              <font  id="user_warning_password" color="red"></font>              
            </div>
          </div>

          <div class="form-group">
            <label for="user_confirmPassword" class="col-lg-2 control-label">Confirm Password<font color="red">*</font></label>
            <div class="col-lg-10">
              <input type="password" class="form-control" id="user_confirmPassword" name="user[confirmPassword]" required placeholder="Confirm Password">
            </div>
            <div class="col-lg-10">
              <font  id="user_warning_confirmPassword" color="red"></font>              
            </div>
          </div>

		      <div class="form-group">
            <label for="user_mail" class="col-lg-2 control-label">Email</label>
            <div class="col-lg-10">
              <input type="email" class="form-control" id="user_mail" name="user[mail]" placeholder="Mail" >
            </div>
            <div class="col-lg-10">
              <font  id="user_warning_mail" color="red"></font>              
            </div>
          </div>

			     <div class="form-group">
            <label for="user_school" class="col-lg-2 control-label">School</label>
            <div class="col-lg-10">
              <input type="text" class="form-control" id="user_school" name="user[school]" placeholder="School">
            </div>
            <div class="col-lg-10">
              <font  id="user_warning_school" color="red"></font>              
            </div>
          </div>

				<div class="form-group">
            <label for="user_motto" class="col-lg-2 control-label">Motto</label>
            <div class="col-lg-10">
              <input type="text" class="form-control" id="user_motto" name="user[motto]" placeholder="Motto">
            </div>
            <div class="col-lg-10">
              <font  id="user_motto" color="red"></font>              
            </div>
          </div>
          <div class="form-group">
            <div class="col-lg-10 col-lg-offset-5">
              <div class="actions">
                <input class="btn btn-info" name="user_signup" type="submit" value="Sign Up">
               </div>
            </div>
          </div>
        </form>

      </div>
  </div>
  <script src="/static/js/bootstrap.min.js"></script>
  <script src="/static/material/js/material.min.js"></script>

	<script type="text/javascript">
	$('#new_user').submit( function(e) {
		e.preventDefault();
		$.ajax({
			type:'POST',
			url:'/users',
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
        if(json.mail != null) {
          $('#user_mail').css({"border-color": "red"});
          $('#user_warning_mail').text(json.mail);
        } else {
          $('#user_warning_mail').text('');
        }				
			},
			success: function() {
				window.location.href = '/sess';
			}
		});
	});
	</script>
{{end}}