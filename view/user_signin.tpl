{{define "content"}}
  <h1>Sign In</h1>
  <form accept-charset="UTF-8" class="new_user" id="new_user" method="post">
    <div style="margin:0;padding:0;display:inline">
      <input name="utf8" type="hidden" value="✓">
    </div>
    <div class="field">
      <label for="user_handle">User</label><br>
      <input id="user_handle" name="user[handle]" size="30" type="text">
    </div>
    <div class="field">
      <label for="user_password">Password</label><br>
      <input id="user_password" name="user[password]" size="30" type="password">
    </div>
    <div class="field">
      <label for="user_remember_me">Remember me</label><br>
      <input id="user_remember_me" name="user[remember_me]" type="checkbox" value="1">
    </div>
    <div class="actions">
      <input name="commit" type="submit" value="Sign In">
    </div>
  </form>
  <a href="/user/signup">Register a new account.</a>
{{end}}
