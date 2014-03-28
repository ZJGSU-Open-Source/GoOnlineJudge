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
						<td><a href="/news/detail/nid/{{.Nid}}">{{.Title}}</a></td>
						<td>{{.Create}}</td>
						<td>[{{if ShowStatus .Status}}Available{{else}}Reserved{{end}}]</td>
						<td>[Delete]</td>
						<td>[Edit]</td>
					</tr>
				{{end}}
			{{end}}
		</tbody>
</table>

{{end}}