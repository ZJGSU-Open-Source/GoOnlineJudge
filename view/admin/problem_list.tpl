{{define "content"}}
{{$isAdmin := .IsAdmin}}
<div class="p-adminProList mdl-grid">

  <div class="mdl-cell mdl-cell--2-col mdl-cell--1-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link current">
        <a>List</a>
      </div>
      <div class="link">
        <a href="/admin/problems/new">Add</a>
      </div>
      <div class="link">
        <a href="/admin/problems/importor">Import</a>
      </div>
      {{if .RejudgePrivilege}}
      <div class="link">
        <a href="/admin/rejudger">Rejudge</a>
      </div>
      {{end}}
    </div>
  </div>

  <div class="page mdl-cell mdl-cell--8-col mdl-cell--6-col-tablet mdl-cell--4-col-phone mdl-shadow--2dp J_list">
    {{template "pagination" .}}
    
    <div class="table-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone mdl-shadow--2dp">
      <table class="go-table text-center">
        <thead>
          <tr>
            <th>ID</th>
            <th>Title</th>
            {{if $isAdmin}}
            <th>Status</th>
            <th>Delete</th>
            <th>Edit</th>
            {{end}}
            <th>Data</th>
          </tr>
        </thead>
        <tbody>
        {{with .Problem}}  
          {{range .}} 
            <tr>
              <td>{{.Pid}}</td>
              <td>
                <a href="/problems/{{.Pid}}">{{.Title}}</a>
              </td>
              {{if $isAdmin}}
              <td>
                <a class="J_status" href="#" data-id="{{.Pid}}">
                {{if ShowStatus .Status}}
                  Available
                {{else}}
                  Reserved
                {{end}}
                </a>
              </td>
              <td>
                <a class="J_delete" href="#" data-id="{{.Pid}}">Delete</a>
              </td>
              <td>
                <a href="/admin/problems/{{.Pid}}">Edit</a>
              </td>
              {{end}}
              <td>
                <a href="/admin/testdata/{{.Pid}}">Test Data</a>
              </td>
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

