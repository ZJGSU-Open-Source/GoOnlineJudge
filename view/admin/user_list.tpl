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
						<td><a href="/user/detail/uid/{{.Uid}}" target="_new">{{.Uid}}</a></td>
						<td>{{PriToString .Privilege}}</td>
						<!--><td><a class="admin_user_delete" href="#" data-type="{{.}}">[Delete]</a></td><-->
						<td><a href="/admin/user/deleteuser/uid/{{.Uid}}">[Delete]</a></td>
					</tr>
				{{end}}
				{{end}}
			{{end}}
		</tbody>
</table>

<script type="text/javascript">
$('.admin_user_delete').on('click', function() {
  var ret = confirm('Delete the user ?');
   if (ret == true) {
               //var pid = {{.User}}
               var uid = $(this).data("Uid");
               alert(uid);
               $.ajax({
                type: 'POST',
                url: '/admin/user/deleteuser/uid/' + uid,
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