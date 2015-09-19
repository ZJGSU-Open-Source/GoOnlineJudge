{{define "content"}}
{{with .Detail}}
<div class="p-adminConAdd mdl-grid">

  <div class="mdl-cell mdl-cell--2-col mdl-cell--1-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link">
        <a href="/admin/contests">List</a>
      </div>
      <div class="link">
        <a href="/admin/contests/new">Add</a>
      </div>
    </div>
  </div>

  <div class="page mdl-cell mdl-cell--8-col mdl-cell--6-col-tablet mdl-cell--4-col-phone mdl-shadow--2dp mdl-grid J_list">

      <div class="go-title-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <div class="title">Contest Add</div>
      </div>
      <input name="utf8" type="hidden" value="✓">

      <div class="mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <div class="contain-center mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
          <input class="mdl-textfield__input" type="text" id="title" name="title" value="{{.Title}}"/>
          <label class="mdl-textfield__label" for="title">Title</label>
        </div>
      </div>
      
      <div class="mdl-cell mdl-cell--3-col mdl-cell--2-col-phone">
        <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
          <input class="mdl-textfield__input" type="text" id="J_start" value="{{.StartTimeYear}}/{{.StartTimeMonth}}/{{.StartTimeDay}} {{.StartTimeHour}}:{{.StartTimeMinute}}" readonly />
          <label class="mdl-textfield__label" for="J_start">Start Date</label>
        </div>
      </div>

      <div class="mdl-cell mdl-cell--3-col mdl-cell--2-col-phone">
        <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
          <input class="mdl-textfield__input" type="text" id="J_end" value="{{.EndTimeYear}}/{{.EndTimeMonth}}/{{.EndTimeDay}} {{.EndTimeHour}}:{{.EndTimeMinute}}" readonly />
          <label class="mdl-textfield__label" for="J_end">End Date</label>
        </div>
      </div>

      <div class="mdl-cell mdl-cell--3-col mdl-cell--2-col-phone">
        <div class="go-select-title">Contest Type</div>
        <select id="J_type" name="type" class="go-select">
          <option value="public" {{if .IsPublic}}selected{{end}}>Public</option>
          <option value="private" {{if .IsPrivate}}selected{{end}}>Private</option>
          <option value="password" {{if .IsPassword}}selected{{end}}>Password</option>
        </select>
      </div>
      

      <div class="mdl-cell mdl-cell--3-col mdl-cell--2-col-phone">
        <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label J_passwdArea"  {{if not .IsPassword}}hidden{{end}}>
          <input class="mdl-textfield__input" id="password" name="password" type="text" value="{{.Argument}}">
          <label class="mdl-textfield__label" for="password">Password</label>
        </div>
      </div>
      
      <div class="mdl-cell mdl-cell--6-col mdl-cell--4-col-phone">
        <div class="group">
          <div class="icon-area">
            <i class="material-icons" id="tt1">error_outline</i>
            <div class="mdl-tooltip" for="tt1">
              可单题或多题输入<br>
              多题输入请用任意非数字字符分割题号<br>
              输入后按回车即可录入题目<br>
              点击标签中的×可删除题目<br>
              可拖动标签为题目重新排序
            </div>
          </div>
          
          <div class="input-area">
            <div class="mdl-textfield mdl-js-textfield">
              <input class="mdl-textfield__input J_proInput" type="text" id="problem"  value="{{.ProblemList}}"/>
              <label class="mdl-textfield__label" for="problem">Problem Id</label>
            </div>
          </div>
        </div>

        <div id="J_proList" style="display: none">
          <div class="item J_pro">
            problem
            <div class="item-del J_del">×</div>
          </div>
        </div>

      </div>

      <div class="mdl-cell mdl-cell--6-col mdl-cell--4-col-phone">
        <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label J_privateArea" {{if not .IsPrivate}}hidden{{end}}>
          <textarea class="mdl-textfield__input textarea" name="userlist" type="text" id="sample5">{{.Argument}}</textarea>
          <label class="mdl-textfield__label" for="sample5">User List</label>
        </div>
      </div>

      <div class="btn-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <!-- Accent-colored raised button with ripple -->
        <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored J_submit" type="submit">submit</button>
      </div>

  </div>
  <div class="mdl-cell mdl-cell--2-col mdl-cell--4-col-phone"></div>
</div>
<link rel="stylesheet" href="/static/css/mobiscroll.css">
<script src="/static/js/mobiscroll.js"></script>
<script src="/static/js/jquery.dragsort-0.5.2.js"></script>
{{end}}
{{end}}