{{define "content"}}
<div class="p-adminUserGener mdl-grid">
  <div class="mdl-cell mdl-cell--2-col mdl-cell--1-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link">
        <a href="/admin/users">Privilege</a>
      </div>
      <div class="link">
        <a href="/admin/users/pagepassword">Password</a>
      </div>
      <div class="link current">
        <a>Generate</a>
      </div>
    </div>
  </div>
  <div class="page mdl-cell mdl-cell--8-col mdl-cell--6-col-tablet mdl-cell--4-col-phone mdl-shadow--2dp mdl-grid">
    <form accept-charset="UTF-8" class="J_addForm mdl-cell mdl-cell--12-col mdl-cell--4-col-phone" method="post" action="/admin/users/generation">
      <div class="go-title-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <div class="title">Team Account Generate</div>
      </div>
      <input name="utf8" type="hidden" value="✓">
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="text" id="user_handle" name="prefix"/>
        <label class="mdl-textfield__label" for="user_handle">User Prefix Handle</label>
      </div>
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="password" id="user_password" name="amount"/>
        <label class="mdl-textfield__label" for="user_password">Amount</label>
      </div>

			<div class="select-area">
        <div class="go-select-title">user type</div>
        <select name="module" class="go-select">
          <option value="0">Normal</option>
					<option value="1">Team</option>
        </select>
      </div>

      <div class="btn-area mdl-cell--12-col mdl-cell--4-col-phone">
        <!-- Accent-colored raised button with ripple -->
        <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored J_submit" type="submit">submit</button>
      </div>
    </form>
  </div>
</div>
{{end}}