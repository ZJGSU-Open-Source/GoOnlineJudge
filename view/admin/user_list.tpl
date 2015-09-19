{{define "content"}}
<div class="p-adminUserList mdl-grid">

  <div class="mdl-cell mdl-cell--2-col mdl-cell--1-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link current">
        <a>Privilege</a>
      </div>
      <div class="link">
        <a href="/admin/users/pagepassword">Password</a>
      </div>
      <div class="link">
        <a href="/admin/users/generation">Generate</a>
      </div>
    </div>
  </div>

  <div class="page mdl-cell mdl-cell--8-col mdl-cell--6-col-tablet mdl-cell--4-col-phone mdl-shadow--2dp J_list">

    <form accept-charset="UTF-8" class="J_addForm">
      <div class="mdl-grid">

        <div class="mdl-cell mdl-cell--2-col mdl-cell--2-col-phone">
          <div class="go-select-title">user type</div>
          <select name="type" class="go-select">
            <option value="TC">Teacher</option>
						<option value="Admin">Admin</option>
          </select>
        </div>

        <div class="mdl-cell mdl-cell--2-col mdl-cell--2-col-phone">
          <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
		        <input class="mdl-textfield__input" type="text" id="user_handle" name="uid"/>
		        <label class="mdl-textfield__label" for="user_handle">Handle</label>
		      </div>
        </div>

        <div class="btn-area">
        	<button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored J_submit" type="submit">Add User</button>
        </div>
				
      </div>
    </form>
    
    <div class="table-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone mdl-shadow--2dp">
      <table class="go-table text-center">
        <thead>
          <tr>
            <th>Uid</th>
				    <th>Privilege</th>
				    <th>Delete</th>
          </tr>
        </thead>
        <tbody>
        	{{with .User}}
					{{range .}}
					{{if LargePU .Privilege}}
						<tr>
							<td><a href="/users/{{.Uid}}">{{.Uid}}</a></td>
							<td>{{ShowPrivilege .Privilege}}</td>
							<td><a class="J_delete" href="#" data-id="{{.Uid}}">Delete</a></td>
						</tr>
					{{end}}
					{{end}}
					{{end}}
        </tbody>
      </table>
    </div>

  </div>
</div>
{{end}}