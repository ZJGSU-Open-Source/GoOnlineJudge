{{define "content"}}
<h1>Admin - Privilege User List</h1>
<table id="contest_list">
	<thead>
		<tr>
		    <th class="header">Uid</th>
		    <th class="header">Privilege</th>
		    <th class="header">Delete</th>
		</tr>
	</thead>
		<tbody>
			{{with .User}}
				{{range .}}
				{{if LargePU .Privilege}}
					<tr>
						<td><a href="/user?detail/uid?{{.Uid}}" target="_new">{{.Uid}}</a></td>
						<td>{{PriToString .Privilege}}</td>
						<td><a class="admin_user_delete" href="#" data-id="{{.Uid}}">[Delete]</a></td>
					</tr>
				{{end}}
				{{end}}
			{{end}}
		</tbody>
</table>

<form accept-charset="UTF-8" id="search_form">
Add Admin: <input id="user" name="user" size="20" type="text">
<select id="type" name="type">
<option value="Admin">Admin</option>
<option value="Source broswer">Source broswer</option>
</select>
<input name="commit" type="submit" value="Add">
</form>

<script type="text/javascript">
$('#search_form').submit( function(e) {
	e.preventDefault();
	var user = $('#user').val();
	var type = $('#type').val();
	$.ajax({
		type:'POST',
		url:'/admin/user/privilege/'+type+'/uid?'+user,
		data:$(this).serialize(),
		error:function(){
			if (user == ""){
				alert("Handle must not be empty!")
			}
		},
		success:function(){
			window.location.reload();
		}
	});
});
</script>

<script type="text/javascript">
$('.admin_user_delete').on('click', function() {
	var uid = $(this).data("id");
	var ret = confirm('Delete the user '+uid+'?');
	if (ret == true) {
		$.ajax({
			type: 'POST',
			url: '/admin/user?privilege/pu/uid?' + uid,
			data:$(this).serialize(),
			error: function() {
				alert('failed!');
			},
			success: function() {
				window.location.reload();
			}
		});
	}
});
</script>
{{end}}