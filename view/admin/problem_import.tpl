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
      <div class="link current">
        <a>Import</a>
      </div>
      {{if .RejudgePrivilege}}
      <div class="link">
        <a href="/admin/rejudger">Rejudge</a>
      </div>
      {{end}}
    </div>
  </div>

  <div class="page mdl-cell mdl-cell--8-col mdl-cell--6-col-tablet mdl-cell--4-col-phone mdl-shadow--2dp mdl-grid J_list">
    <form accept-charset="UTF-8" class="J_addForm mdl-cell mdl-cell--12-col mdl-cell--4-col-phone"  enctype="multipart/form-data" action="/admin/problems/importor" method="post">
      <div class="go-title-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <div class="title">Problem Import</div>
      </div>
      <input name="utf8" type="hidden" value="✓">
      <div class="mdl-cell mdl-cell--12-col mdl-cell--4-col-phone text-center">
        <input type="file" id="file" name="fps.xml"/>
        <!-- <label class="mdl-textfield__label" for="file">File</label> -->
      </div>
      <div class="btn-area mdl-cell--12-col mdl-cell--4-col-phone">
        <!-- Accent-colored raised button with ripple -->
        <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored J_submit" type="submit">submit</button>
      </div>
    </form>
  </div>

  <div class="mdl-cell mdl-cell--2-col mdl-cell--4-col-phone"></div>
</div>
{{end}}