{{define "content"}}
{{$compiler_id := .Compiler_id}}
<div class="p-proDetail mdl-grid">
  <div class="mdl-cell mdl-cell--2-col mdl-cell--4-col-tablet mdl-cell--2-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link">
        <a href="/contests/{{.Cid}}">Problem</a>
      </div>
      <div class="link">
        <a href="/contests/{{.Cid}}/status">Status</a>
      </div>
      <div class="link">
        <a href="/contests/{{.Cid}}/ranklist">Ranklist</a>
      </div>
    </div>
  </div>
  
  <div class="mdl-cell mdl-cell--4-col-tablet mdl-cell--2-col-phone mdl-cell--hide-desktop">
    {{with .Detail}}
    <div class="info mdl-shadow--2dp">
      <div>Time Limit</div>
      <div>{{.Time}}s</div>
      <div>Memory Limit</div>
      <div>{{.Memory}}KB</div>
      <div>Judge Program</div>
      <div>{{ShowSpecial .Special}}</div>
      <div>Ratio(Solve/Submit)</div>
      <div>{{ShowRatio .Solve .Submit}}( {{.Solve}} / {{.Submit}} )</div>
    </div>
    {{end}}
  </div>

  <div class="page mdl-cell mdl-cell--8-col mdl-cell--8-col-tablet mdl-cell--4-col-phone mdl-shadow--2dp J_list">
    {{with .Detail}}
    <div class="go-title-area border mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
      <div class="title">{{.Title}}</div>
    </div>
    
    <div class="item mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
      <div class="tip">Description:</div>
      <section class="text">{{.Description}}</section>

      <div class="tip">Input:</div>
      <section class="text">{{.Input}}</section>

      <div class="tip">Output:</div>
      <section class="text">{{.Output}}</section>

      <div class="tip">Sample Input:</div>
      <section class="code mdl-shadow--2dp">{{.In}}</section>

      <div class="tip">Sample Output:</div>
      <section class="code mdl-shadow--2dp">{{.Out}}</section>
      {{if .Hint}}
        <div class="tip">Hint:</div>
        <section class="text">{{.Hint}}</section>
      {{end}}
    </div>
    {{end}}
    <div class="btn-area mdl-cell mdl-cell--1-col mdl-cell--2-col-phone">
      <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored J_extend" >Submit</button>
    </div>
    <div class="J_submission" style="display: none;">
      <form accept-charset="UTF-8" id="problem_submit" data-pid="{{.Pid}}" data-cid="{{.Cid}}">
        <input name="utf8" type="hidden" value="✓">

        <div class="mdl-grid">
          <div class="mdl-cell mdl-cell--2-col mdl-cell--1-col-phone">
            <div class="go-select-title">Compiler</div>
            <select id="compiler_id" name="compiler_id" class="go-select">
              <option value="1" {{if eq $compiler_id "1"}}selected="selected"{{end}}>C</option>
              <option value="2" {{if eq $compiler_id "2"}}selected="selected"{{end}}>C++</option>
              <option value="3" {{if eq $compiler_id "3"}}selected="selected"{{end}}>Java</option>
            </select>
          </div>
          <label class="check-area mdl-cell mdl-cell--4-col mdl-cell--3-col-phone mdl-grid">
            <i class="material-icons J_label">check_box</i>
            <div class="text">Use advanced editor</div>
            <input checked id="advanced_editor" name="advanced_editor" type="checkbox" value="1" hidden />
          </label>
          <div class="mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
            <textarea id="code" name="code" class="textarea"></textarea>
          </div>
        </div>
        <div class="btn-area mdl-cell mdl-cell--1-col mdl-cell--2-col-phone">
          <button class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored" type="submit">Submit</button>
        </div>     
      </form>
    </div>

  </div>
  
  <div class="mdl-cell mdl-cell--2-col mdl-cell--2-col-tablet mdl-cell--4-col-phone mdl-cell--hide-phone mdl-cell--hide-tablet">
    {{with .Detail}}
    <div class="info J_static mdl-shadow--2dp">
      <div>Time Limit</div>
      <div>{{.Time}}s</div>
      <div>Memory Limit</div>
      <div>{{.Memory}}KB</div>
      <div>Judge Program</div>
      <div>{{ShowSpecial .Special}}</div>
      <div>Ratio(Solve/Submit)</div>
      <div>{{ShowRatio .Solve .Submit}}( {{.Solve}} / {{.Submit}} )</div>
    </div>
    {{end}}
  </div>
  
</div>

<link rel="stylesheet" href="/static/css/codemirror.css">
<script src="/static/js/codemirror.js"></script>

{{end}}