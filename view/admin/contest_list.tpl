{{define "content"}}
<div class="p-adminConList mdl-grid">

  <div class="mdl-cell mdl-cell--2-col mdl-cell--1-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link current">
        <a>List</a>
      </div>
      <div class="link">
        <a href="/admin/contests/new">Add</a>
      </div>
    </div>
  </div>

  <div class="page mdl-cell mdl-cell--8-col mdl-cell--6-col-tablet mdl-cell--4-col-phone mdl-shadow--2dp J_list">
    <div class="go-title-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <div class="title">Contest List</div>
      </div>
    <div class="table-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone mdl-shadow--2dp">
      <table class="go-table text-center">
        <thead>
          <tr>
            <th>ID</th>
            <th>Title</th>
            <th>Creator</th>
            <th>Status</th>
            <th>Delete</th>
            <th>Edit</th>
          </tr>
        </thead>
        <tbody>
        {{with .Contest}}  
          {{range .}} 
          <tr>
            <td>{{.Cid}}</td>
            <td><a href="/contests/{{.Cid}}">{{.Title}}</a></td>
            <td><a href="/users/{{.Creator}}">{{.Creator}}</a></td>
            <td>
              <a class="J_status" href="#" data-id="{{.Cid}}">
                {{if ShowStatus .Status}}
                  Available
                {{else}}
                  Reserved
                {{end}}
              </a>
            </td>
            <td><a class="J_delete" href="#" data-id="{{.Cid}}">Delete</a></td>
            <td><a href="/admin/contests/{{.Cid}}">Edit</a></td>
          </tr>
          {{end}}
        {{else}}
          <td></td>
          <td></td>
          <td>无</td>
          <td></td>
          <td></td>
          <td></td>
        {{end}}
        </tbody>
      </table>
    </div>

  </div>
  <div class="mdl-cell mdl-cell--2-col mdl-cell--4-col-phone"></div>
</div>
{{end}}


