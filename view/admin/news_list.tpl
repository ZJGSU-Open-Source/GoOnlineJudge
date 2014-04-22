{{define "content"}}
<h1>Admin - News List</h1>
<table id="news_list">
	<thead>
		<tr>
		    <th class="header">Title</th>
		    <th class="header">Date</th>
		    <th class="header">Status</th>
		    <th class="header">Delete</th>
		    <th class="header">Edit</th>
		</tr>
	</thead>
		<tbody>
			{{with .News}}
				{{range .}}
					<tr>
						<td><a href="/admin/news/detail/nid/{{.Nid}}">{{.Title}}</a></td>
						<td>{{.Create}}</td>
						<td><a class="news_status" href="#">[{{if ShowStatus .Status}}Available{{else}}Reserved{{end}}]</a></td>
						<td><a class="delete" href="" onclick="ConfirmDelete('/admin/news/delete/nid/{{.Nid}}', 'Delete The News ?')">[Delete]</a></td>
						<td><a href="/admin/news/edit/nid/{{.Nid}}">[Edit]</a></td>
					</tr>
				{{end}}
			{{end}}
		</tbody>
</table>
<script type="text/javascript">
$('.news_status').onclick( function(e) {
	e.preventDefault();
	$.ajax({
		type:'POST',
		url:'/admin/news/status/nid/{{.Nid}}',
		data:$(this).serialize(),
		error: function
	});
});
</script>

{{end}}