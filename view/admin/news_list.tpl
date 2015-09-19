{{define "content"}}
<div class="p-adminNewsList mdl-grid">

  <div class="mdl-cell mdl-cell--2-col mdl-cell--1-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
    	<div class="link current">
		    <a href="/admin/news">List</a>
		  </div>
		  <div class="link">
		    <a href="/admin/news/new">Add</a>
		  </div>
    </div>
  </div>

  <div class="page mdl-cell mdl-cell--8-col mdl-cell--6-col-tablet mdl-cell--4-col-phone mdl-shadow--2dp J_list">
    <div class="go-title-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
      <div class="title">News List</div>
    </div>
    
    <div class="table-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone mdl-shadow--2dp">
      <table class="go-table text-center">
        <thead>
          <tr>
            <th>Title</th>
				    <th>Date</th>
				    <th>Status</th>
				    <th>Delete</th>
				    <th>Edit</th>
          </tr>
        </thead>
        <tbody>
          {{with .News}}
					{{range .}}
            <tr>
              <td><a href="/news/{{.Nid}}">{{.Title}}</a></td>
							<td>{{.Create}}</td>
							<td>
								<a class="J_status" href="#" data-id="{{.Nid}}">
									{{if ShowStatus .Status}}
										Available
									{{else}}
										Reserved
									{{end}}
								</a>
							</td>
							<td><a class="J_delete" href="#" data-id="{{.Nid}}">Delete</a></td>
							<td><a class="J_edit" href="/admin/news/{{.Nid}}">Edit</a></td>
            </tr>
          {{end}}
          {{end}}
        </tbody>
      </table>
    </div>
    
  </div>

  <div class="mdl-cell mdl-cell--2-col mdl-cell--4-col-phone"></div>

</div>
{{end}}