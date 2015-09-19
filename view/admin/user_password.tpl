{{define "content"}}
<div class="p-signin mdl-grid">
  <div class="mdl-cell mdl-cell--2-col mdl-cell--1-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link">
        <a href="/admin/users">Privilege</a>
      </div>
      <div class="link current">
        <a>Password</a>
      </div>
      <div class="link">
        <a href="/admin/users/generation">Generate</a>
      </div>
    </div>
  </div>
  <div class="page mdl-cell mdl-cell--8-col mdl-cell--6-col-tablet mdl-cell--4-col-phone mdl-shadow--2dp mdl-grid">
    <form accept-charset="UTF-8" class="J_addForm mdl-cell mdl-cell--12-col mdl-cell--4-col-phone" method="post">
      <div class="go-title-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <div class="title">Edit Password</div>
      </div>
      <input name="utf8" type="hidden" value="✓">
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="text" id="user_handle" name="user[Handle]"/>
        <label class="mdl-textfield__label" for="user_handle">Handle</label>
      </div>
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="password" id="user_password" name="user[newPassword]"/>
        <label class="mdl-textfield__label" for="user_password">New Password</label>
      </div>
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="password" id="user_password" name="user[confirmPassword]"/>
        <label class="mdl-textfield__label" for="user_password">New Password</label>
      </div>
      <div class="btn-area mdl-cell--12-col mdl-cell--4-col-phone">
        <!-- Accent-colored raised button with ripple -->
        <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored J_submit" type="submit">submit</button>
      </div>
    </form>
  </div>
</div>
{{end}}