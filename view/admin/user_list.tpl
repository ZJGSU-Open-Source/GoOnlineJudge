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
						<td>[Delete]</td>
					</tr>
				{{end}}
				{{end}}
			{{end}}
		</tbody>
</table>
{{end}}