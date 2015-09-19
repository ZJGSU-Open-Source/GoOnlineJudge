{{define "content"}}

<!-- <form accept-charset="UTF-8" class="" id="new_contest" method="post" action="/admin/contests">
  <div style="margin:0;padding:0;display:inline">
      <input name="utf8" type="hidden" value="✓">
    </div>
    <div class="field">
      <label for="content_title">Title</label><br>
      <input id="content_title" name="title" size="60" style="width:640px" type="text" required="">
    </div>
    <div class="field">
      <label for="content_type">Contest Type</label><br>
      <select id="content_type" name="type" onchange="show(); return false;">
        <option value="public">Public</option>
        <option value="private">Private</option>
        <option value="password">Password</option>
      </select>
    </div>
    <div class="field" id="passwd_field" style="display: none;">
      <label for="password">Password</label><br>
      <input id="password" name="password" size="20" type="text">
    </div>
     
    <div class="field">
      <label for="content_startTimeYear">Start Time[Year-Month-Day Hour:Minute]</label><br>
      <input id="content_startTimeYear" name="startTimeYear" size="4" type="number" max="2100" min="2014" value="{{.StartYear}}">-
      <input id="content_startTimeMonth" name="startTimeMonth" size="4" type="number" max="12" min="1" value="{{.StartMonth}}">-
      <input id="content_startTimeDay" name="startTimeDay" type="number" max="31" min="1" value="{{.StartDay}}">
      <input id="content_startTimeHour" name="startTimeHour" size="4" type="number" max="23" min="0" value="{{.StartHour}}">:
      <input id="content_startTimeMinute" name="startTimeMinute" size="4" type="number" max="59" min="0" value="0">:
    </div>
    <div class="field">
      <label for="content_endTimeYear">End Time[Year-Month-Day Hour:Minute]</label><br>
      <input id="content_endTimeYear" name="endTimeYear" size="4" type="number" max="2100" min="2014" value="{{.EndYear}}">-
      <input id="content_endTimeMonth" name="endTimeMonth" size="4" type="number" max="12" min="1" value="{{.EndMonth}}">-
      <input id="content_endTimeDay" name="endTimeDay" type="number" max="31" min="1" value="{{.EndDay}}">
      <input id="content_endTimeHour" name="endTimeHour" size="4" type="number" max="23" min="0" value="{{.EndHour}}">:
      <input id="content_endTimeMinute" name="endTimeMinute" size="4" size="4" type="number" max="59" min="0" value="0">:
    </div>
    <div class="field">
      <label for="content_problemList">Problem List[Please using ";" split each problem.]</label><br>
      <input id="content_problemList" name="problemList" size="100" style="width:640px" type="text">
    </div>
    <div class="field" id="userlist_field" style="display: none;">
      <label for="userlist">User List</label><br>
      <textarea id="userlist" name="userlist" style="width:630px;height:200px;"></textarea> 
    </div>
    <div class="actions">
      <input name="commit" type="submit" value="Submit">
    </div>
</form> -->

<div class="p-adminConAdd mdl-grid">

  <div class="mdl-cell mdl-cell--2-col mdl-cell--1-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link">
        <a href="/admin/contests">List</a>
      </div>
      <div class="link current">
        <a>Add</a>
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
          <input class="mdl-textfield__input" type="text" id="title" name="title"/>
          <label class="mdl-textfield__label" for="title">Title</label>
        </div>
      </div>
      
      <div class="mdl-cell mdl-cell--3-col mdl-cell--2-col-phone">
        <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
          <input class="mdl-textfield__input" type="text" id="J_start" readonly />
          <label class="mdl-textfield__label" for="J_start">Start Date</label>
        </div>
      </div>
      
      <div class="mdl-cell mdl-cell--3-col mdl-cell--2-col-phone">
        <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
          <input class="mdl-textfield__input" type="text" id="J_end" readonly />
          <label class="mdl-textfield__label" for="J_end">End Date</label>
        </div>
      </div>

      <div class="mdl-cell mdl-cell--3-col mdl-cell--2-col-phone">
        <div class="go-select-title">Contest Type</div>
        <select id="J_type" name="type" class="go-select">
          <option value="public">Public</option>
          <option value="private">Private</option>
          <option value="password">Password</option>
        </select>
      </div>
      

      <div class="mdl-cell mdl-cell--3-col mdl-cell--2-col-phone">
        <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label J_passwdArea" style="display: none">
          <input class="mdl-textfield__input" id="password" name="password" type="text">
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
              <input class="mdl-textfield__input J_proInput" type="text" id="problem"/>
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
        <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label J_privateArea" style="display: none;">
          <textarea class="mdl-textfield__input textarea" name="userlist" type="text" id="sample5"></textarea>
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

