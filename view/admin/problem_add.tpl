{{define "content"}}
<div class="p-adminProblem mdl-grid">

  <div class="mdl-cell mdl-cell--2-col mdl-cell--1-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link">
        <a href="/admin/problems">List</a>
      </div>
      <div class="link current">
        <a href="/admin/problems/new">Add</a>
      </div>
      <div class="link">
        <a href="/admin/problems/importor">Import</a>
      </div>
      {{if .RejudgePrivilege}}
      <div class="link">
        <a href="/admin/rejudger">Rejudge</a>
      </div>
      {{end}}
    </div>
  </div>

  <div class="page mdl-cell mdl-cell--8-col mdl-cell--6-col-tablet mdl-cell--4-col-phone mdl-shadow--2dp mdl-grid J_list">
    <form accept-charset="UTF-8" class="mdl-cell mdl-cell--12-col mdl-cell--4-col-phone mdl-grid mdl-grid--no-spacing" action="/admin/problems/" method="post">
      <div class="go-title-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <div class="title">Problem Add</div>
      </div>
      <input name="utf8" type="hidden" value="✓">
      <div class="mdl-cell--6-col mdl-cell--4-col-phone">
        <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
          <input class="mdl-textfield__input" type="text" id="title" name="title"/>
          <label class="mdl-textfield__label" for="title">title</label>
          <span class="mdl-textfield__error">请输入标题</span>
        </div>
      </div>
      <div class="mdl-cell--6-col mdl-cell--4-col-phone">
        <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
          <input class="mdl-textfield__input" type="text" id="source" name="source"/>
          <label class="mdl-textfield__label" for="source">Source</label>
          <span class="mdl-textfield__error">请输入来源</span>
        </div>
      </div>
      <div class="mdl-cell--6-col mdl-cell--2-col-phone">
        <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
          <input class="mdl-textfield__input" type="number" id="time" name="time" value="1"/>
          <label class="mdl-textfield__label" for="time">Time Limit (S)</label>
          <span class="mdl-textfield__error">请输入时间限制</span>
        </div>
      </div>
      <div class="mdl-cell--6-col mdl-cell--2-col-phone">
        <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
          <input class="mdl-textfield__input" type="number" id="memory" name="memory"  value="32768"/>
          <label class="mdl-textfield__label" for="memory">Memory Limit (KB)</label>
          <span class="mdl-textfield__error">请输入时内存限制</span>
        </div>
      </div>
      <label class="check-area">
        <i class="material-icons J_label icon">check_box_outline_blank</i>
        <div class="text">Special Judge</div>
        <input id="J_special" name="special" type="checkbox" value="1" hidden />
      </label>

      <div class="block mdl-cell--12-col mdl-cell--4-col-phone">
        <label for="J_description title">Description</label>
        <input id="J_description" name="description" type="text" hidden>
        <div class="loading-area J_load">
          <div class="go-loading mdl-shadow--2dp contain-center">
            <i class="mdl-spinner mdl-js-spinner is-active"></i>
          </div>
        </div>
      </div>

      <div class="block mdl-cell--12-col mdl-cell--4-col-phone">
        <label for="J_input title">Input</label>
        <input id="J_input" name="input" type="text" hidden>
        <div class="loading-area J_load">
          <div class="go-loading mdl-shadow--2dp contain-center">
            <i class="mdl-spinner mdl-js-spinner is-active"></i>
          </div>
        </div>
      </div>

      <div class="block mdl-cell--12-col mdl-cell--4-col-phone">
        <label for="J_output title">Output</label>
        <input id="J_output" name="output" type="text" hidden>
        <div class="loading-area J_load">
          <div class="go-loading mdl-shadow--2dp contain-center">
            <i class="mdl-spinner mdl-js-spinner is-active"></i>
          </div>
        </div>
      </div>

      <div class="block mdl-cell--12-col mdl-cell--4-col-phone">
        <label for="J_hint title">Hint</label>
        <input id="J_hint" name="hint" type="text" hidden>
        <div class="loading-area J_load">
          <div class="go-loading mdl-shadow--2dp contain-center">
            <i class="mdl-spinner mdl-js-spinner is-active"></i>
          </div>
        </div>
      </div>

      <div class="block mdl-cell--12-col mdl-cell--4-col-phone">
        <label for="problem_in title">Sample Input</label><br>
        <textarea id="problem_in" class="textarea" name="in"></textarea>
      </div>

      <div class="block mdl-cell--12-col mdl-cell--4-col-phone">
        <label for="problem_out title">Sample Output</label><br>
        <textarea id="problem_out" class="textarea" name="out"></textarea>
      </div>

      <div class="btn-area mdl-cell--12-col mdl-cell--4-col-phone">
        <!-- Accent-colored raised button with ripple -->
        <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored" type="submit">submit</button>
      </div>
    </form>
  </div>
  <div class="mdl-cell mdl-cell--2-col mdl-cell--4-col-phone"></div>
</div>
<script src="/static/kindeditor/kindeditor.js" type="text/javascript"></script>
<script src="/static/kindeditor/lang/zh_CN.js" type="text/javascript"></script>
{{end}}

