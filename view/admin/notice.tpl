{{define "content"}}
<div class="p-adminNotice mdl-grid">
  <div class="mdl-cell mdl-cell--2-col mdl-cell--hide-phone mdl-cell--hide-tablet"></div>
  <div class="page mdl-cell mdl-cell--8-col mdl-cell--4-col-phone mdl-shadow--2dp mdl-grid">
    <form accept-charset="UTF-8" class="mdl-cell mdl-cell--12-col mdl-cell--4-col-phone" action="/admin/notice" method="post">
      <div class="go-title-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <div class="title">Notice</div>
      </div>
      <input name="utf8" type="hidden" value="✓">
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="text" name="msg" value="{{.Msg}}"/>
        <label class="mdl-textfield__label" for="user_handle">message</label>
      </div>
      <div class="btn-area" style="text-align: center">
        <!-- Accent-colored raised button with ripple -->
        <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored" type="submit">submit</button>
      </div>
    </form>
  </div>
</div>
{{end}}
