{{define "content"}}
{{with .Detail}}
	<div class="row">
	    <div class="col-lg-6">
			<form accept-charset="UTF-8" class="new_user form-horizontal" id="new_user">
				<legend>Edit Info</legend>
				<div style="margin:0;padding:0;display:inline">
					<input name="utf8" type="hidden" value="✓">
				</div>
				<div class="form-group">
		            <label for="user_handle" class="col-lg-2 control-label">Handle</label>
		            <div class="col-lg-10">
		              <input type="text" class="form-control" id="user_handle" name="user[handle]" value="{{.Uid}}" readonly autofocus>
		            </div>
		         </div>
		         <div class="form-group">
		            <label for="user_nick" class="col-lg-2 control-label">Nick<font color="red">*</font></label>
		            <div class="col-lg-10">
		              <input type="text" class="form-control" id="user_nick" name="user[nick]" value="{{.Nick}}" required>
		            </div>
		          </div>
				<div class="form-group">
		            <label for="user_mail" class="col-lg-2 control-label">Email</label>
		            <div class="col-lg-10">
		              <input type="email" class="form-control" id="user_mail" name="user[mail]" value="{{.Mail}}" >
		            </div>
		         </div>

				<div class="form-group">
		            <label for="user_school" class="col-lg-2 control-label">School</label>
		            <div class="col-lg-10">
		              <input type="text" class="form-control" id="user_school" name="user[school]" value="{{.School}}">
		            </div>
		        </div>

				<div class="form-group">
		            <label for="user_motto" class="col-lg-2 control-label">Motto</label>
		            <div class="col-lg-10">
		              <input type="text" class="form-control" id="user_motto" name="user[motto]" value="{{.Motto}}">
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
				} 		
			},
			success: function(result) {
				window.location.href = '/user/detail?uid='+{{.CurrentUser}};
			}
		});
	});
	</script>
{{end}}