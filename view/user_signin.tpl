{{define "content"}}
<div class="p-signin mdl-grid">
  <div class="mdl-cell mdl-cell--2-col mdl-cell--hide-phone mdl-cell--hide-tablet"></div>
  <div class="page mdl-cell mdl-cell--8-col mdl-cell--4-col-phone mdl-shadow--2dp mdl-grid">
    <form accept-charset="UTF-8" class="J_addForm mdl-cell mdl-cell--12-col mdl-cell--4-col-phone" action="/sess" method="post">
      <div class="go-title-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <div class="title">Sign In</div>
      </div>
      <input name="utf8" type="hidden" value="✓">
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="text" id="user_handle" name="user[handle]"/>
        <label class="mdl-textfield__label" for="user_handle">Handle</label>
        <span class="mdl-textfield__error">请输入账号</span>
      </div>
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="password" id="user_password" name="user[password]"/>
        <label class="mdl-textfield__label" for="user_password">Password</label>
        <span class="mdl-textfield__error">请输入密码</span>
      </div>
      <div class="btn-area mdl-cell--12-col mdl-cell--4-col-phone">
        <!-- Accent-colored raised button with ripple -->
        <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored J_submit" type="submit">signin</button>
        <!-- Accent-colored raised button with ripple -->
        <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent J_register" type="button">register</button>
      </div>
    </form>
  </div>
</div>
{{end}}
