{{define "content"}}
<div class="p-adminNews mdl-grid">

  <div class="mdl-cell mdl-cell--2-col mdl-cell--1-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link">
        <a href="/admin/news">List</a>
      </div>
      <div class="link">
        <a href="/admin/news/new">Add</a>
      </div>
    </div>
  </div>

  <div class="page mdl-cell mdl-cell--8-col mdl-cell--6-col-tablet mdl-cell--4-col-phone mdl-shadow--2dp mdl-grid">
    {{with .Detail}}
    <form accept-charset="UTF-8" class="mdl-cell mdl-cell--12-col mdl-cell--4-col-phone" action="/admin/news/{{.Nid}}" method="post">
      <div class="go-title-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <div class="title">News Edit</div>
      </div>
      <input name="utf8" type="hidden" value="✓">
      <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
        <input class="mdl-textfield__input" type="text" id="title" name="title" value="{{.Title}}" />
        <label class="mdl-textfield__label" for="title">Title</label>
        <span class="mdl-textfield__error">请输入标题</span>
      </div>
      <textarea id="J_content" name="content" hidden>{{.Content}}</textarea>
      <div class="loading J_load">编辑器加载中...</div>
      <div class="btn-area">
        <!-- Accent-colored raised button with ripple -->
        <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored" type="submit">submit</button>
      </div>
    </form>
    {{end}}
  </div>
  <div class="mdl-cell mdl-cell--2-col mdl-cell--4-col-phone"></div>
</div>
<script src="/static/kindeditor/kindeditor.js" type="text/javascript"></script>
<script src="/static/kindeditor/lang/zh_CN.js" type="text/javascript"></script>
{{end}}

