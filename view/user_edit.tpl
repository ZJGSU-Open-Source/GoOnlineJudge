{{define "content"}}
<div class="p-user-edit mdl-grid">
  <div class="mdl-cell mdl-cell--2-col mdl-cell--4-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link">
        <a href="/settings">Detail</a>
      </div>
      <div class="link current">
        <a>Edit Info</a>
      </div>
      <div class="link">
        <a href="/account">Password</a>
      </div>
    </div>
  </div>
  <div class="page mdl-cell mdl-cell--8-col mdl-cell--4-col-phone mdl-shadow--2dp mdl-grid">
    <form accept-charset="UTF-8" class="J_addForm mdl-cell mdl-cell--12-col mdl-cell--4-col-phone" action="/profile" method="post">
    {{with .Detail}}
      <div class="go-title-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <div class="title">Edit Info</div>
      </div>
      <input name="utf8" type="hidden" value="✓">
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="text" id="user_handle" name="user[handle]"  value="{{.Uid}}" readonly="true" />
        <label class="mdl-textfield__label" for="user_handle">Handle</label>
      </div>
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="text" id="user_handle" name="user[nick]" value="{{.Nick}}"/>
        <label class="mdl-textfield__label" for="user_handle">Nick</label>
      </div>
      <div class="check-area">
      	<label class="mdl-checkbox mdl-js-checkbox mdl-js-ripple-effect" for="checkbox-2">
			  <input type="checkbox" id="checkbox-2" class="mdl-checkbox__input" name="user[share_code]" value="true" {{if .ShareCode}} checked="checked" {{end}}/>
			  <span class="mdl-checkbox__label">Share Code</span>
			</label>
      </div>
      
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="text" id="user_password" name="user[mail]" value="{{.Mail}}" />
        <label class="mdl-textfield__label" for="user_password">Email</label>
      </div>
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="text" id="user_handle" name="user[school]" value="{{.School}}"/>
        <label class="mdl-textfield__label" for="user_handle">School</label>
      </div>
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="text" id="user_handle" name="user[motto]" value="{{.Motto}}"/>
        <label class="mdl-textfield__label" for="user_handle">Motto</label>
      </div>
      <div class="btn-area mdl-cell--12-col mdl-cell--4-col-phone">
        <!-- Accent-colored raised button with ripple -->
        <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored J_submit" type="submit">edit</button>
      </div>
      {{end}}
    </form>
  </div>
</div>
{{end}}
