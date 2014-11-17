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
						<td><a href="/news/{{.Nid}}">{{.Title}}</a></td>
						<td>{{.Create}}</td>
						<td><a class="news_status" href="#" data-id="{{.Nid}}">[{{if ShowStatus .Status}}Available{{else}}Reserved{{end}}]</a></td>
						<td><a class="news_delete" href="#" data-id="{{.Nid}}">[Delete]</a></td>
						<td><a class="news_edit" href="#" data-id="{{.Nid}}">[Edit]</a></td>
					</tr>
				{{end}}
			{{end}}
		</tbody>
</table>
<script type="text/javascript">
$('.news_status').on('click', function() {
	var nid = $(this).data('id');
	$.ajax({
		type:'POST',
		url:'/admin/news/'+nid+'/status',
		data:$(this).serialize(),
		error: function() {
			alert('failed!');
		},
		success: function() {
			window.location.reload();
		}
	});
});
$('.news_delete').on('click', function() {
	var ret = confirm('Delete the News?');
	if(ret == true) {
		var nid = $(this).data('id');
		$.ajax({
			type:'DELETE',
			url:'/admin/news/'+nid,
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
$('.news_edit').on('click', function() {
	var nid = $(this).data('id');
	window.location.href = '/admin/news/'+nid;
});
</script>

{{end}}