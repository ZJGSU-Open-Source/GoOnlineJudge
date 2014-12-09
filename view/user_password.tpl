{{define "content"}}
	<div class="row">
	    <div class="col-lg-6">
			<form accept-charset="UTF-8" class="new_user form-horizontal" id="new_user">
				<legend>Edit Password</legend>
				<div style="margin:0;padding:0;display:inline">
					<input name="utf8" type="hidden" value="✓">
				</div>
				<div class="form-group">
			        <label for="user_oldPassword" class="col-lg-2 control-label">Old Password<font color="red">*</font>
			         <font id="user_warning_oldPassword" color="red"></font><br>
					     </label>
			        <div class="col-lg-10">
			          <input type="password" class="form-control" id="user_oldPassword" name="user[oldPassword]" required placeholder="Old Password" autofocus>
			        </div>
			    </div>
				<div class="form-group">
			        <label for="user_newPassword" class="col-lg-2 control-label">New Password<font color="red">*</font>
			         <font id="user_warning_newPassword" color="red"></font><br>
					     </label>
			        <div class="col-lg-10">
			          <input type="password" class="form-control" id="user_newPassword" name="user[newPassword]" required placeholder="At least six characters.">
			        </div>
			     </div>
			     <div class="form-group">
			        <label for="user_confirmPassword" class="col-lg-2 control-label">Confirm Password<font color="red">*</font></label>
			        <font id="user_warning_confirmPassword" color="red"></font><br>
			        <div class="col-lg-10">
			          <input type="password" class="form-control" id="user_confirmPassword" name="user[confirmPassword]" required placeholder="Confirm Password">
			        </div>
			    </div>
				<div class="form-group">
			        <div class="col-lg-10 col-lg-offset-5">
			          <div class="actions">
			            <input class="btn btn-info" name="user_signup" type="submit" value="Edit">
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
{{end}}
