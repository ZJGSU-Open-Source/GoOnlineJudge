{{define "content"}}
<div class="p-signin mdl-grid">

	<div class="mdl-cell mdl-cell--2-col mdl-cell--1-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link">
        <a href="/admin/problems">List</a>
      </div>
      <div class="link">
        <a href="/admin/problems/new">Add</a>
      </div>
      <div class="link">
        <a href="/admin/problems/importor">Import</a>
      </div>
      {{if .RejudgePrivilege}}
      <div class="link current">
        <a>Rejudge</a>
      </div>
      {{end}}
    </div>
  </div>

  <div class="page mdl-cell mdl-cell--8-col mdl-cell--6-col-tablet mdl-cell--4-col-phone mdl-shadow--2dp mdl-grid">
    <form accept-charset="UTF-8" class="J_addForm mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
      <div class="go-title-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <div class="title">Rejudge</div>
      </div>
      <input name="utf8" type="hidden" value="✓">

      <div class="contain-center mdl-cell mdl-cell--3-col mdl-cell--4-col-phone">
        <div class="go-select-title">search type</div>
        <select name="type" class="go-select J_type">
          <option value="Sid">Solution ID</option>
					<option value="Pid">Problem ID</option>
        </select>
      </div>

      <div class="contain-center mdl-cell mdl-cell--3-col mdl-cell--4-col-phone mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="text" id="J_id" name="id"/>
        <label class="mdl-textfield__label" for="J_id">Id</label>
        <span class="mdl-textfield__error">请输入账号</span>
      </div>
      
      <div class="btn-area mdl-cell--12-col mdl-cell--4-col-phone">
        <!-- Accent-colored raised button with ripple -->
        <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored J_submit" type="submit">rejudge</button>
      </div>
    </form>
  </div>
  <div class="mdl-cell mdl-cell--2-col mdl-cell--4-col-phone"></div>
</div>
{{end}}
