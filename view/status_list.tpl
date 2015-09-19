{{define "content"}}
<!-- <meta http-equiv="refresh" content="30"> -->
<div class="p-staticList mdl-grid">
  <div class="mdl-cell mdl-cell--1-col mdl-cell--hide-phone mdl-cell--hide-tablet"></div>
  <div class="page mdl-cell mdl-cell--10-col mdl-cell--4-col-phone mdl-shadow--2dp">

    <form accept-charset="UTF-8" id="search_form">
      <div class="mdl-grid">

        <div class="mdl-cell mdl-cell--2-col mdl-cell--2-col-phone">
          <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
            <input class="mdl-textfield__input" name="search_uid" value="{{.SearchUid}}" type="text" id="sample1" />
            <label class="mdl-textfield__label" for="sample1">User</label>
          </div>
        </div>

        <div class="mdl-cell mdl-cell--2-col mdl-cell--2-col-phone">
          <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
            <input class="mdl-textfield__input" name="search_pid" value="{{.SearchPid}}" type="text" pattern="-?[0-9]*(\.[0-9]+)?" id="sample2" />
            <label class="mdl-textfield__label" for="sample2">Problem</label>
            <span class="mdl-textfield__error">请输入正确的题目ID</span>
          </div>
        </div>

        <div class="mdl-cell mdl-cell--3-col mdl-cell--2-col-phone">
          <div class="go-select-title">Result</div>
          <select name="search_judge" class="go-select">
            <option value="0">All</option>
            <option value="1" {{if .SearchJudge0}}selected{{end}}>Pending</option>
            <option value="2" {{if .SearchJudge1}}selected{{end}}>Running &amp;Judging</option>
            <option value="3" {{if .SearchJudge2}}selected{{end}}>Compile Error</option>
            <option value="4" {{if .SearchJudge3}}selected{{end}}>Accepted</option>
            <option value="5" {{if .SearchJudge4}}selected{{end}}>Runtime Error</option>
            <option value="6" {{if .SearchJudge5}}selected{{end}}>Wrong Answer</option>
            <option value="7" {{if .SearchJudge6}}selected{{end}}>Time Limit Exceeded</option>
            <option value="8" {{if .SearchJudge7}}selected{{end}}>Memory Limit Exceeded</option>
            <option value="9" {{if .SearchJudge8}}selected{{end}}>Output Limit Exceeded</option>
            <option value="10" {{if .SearchJudge9}}selected{{end}}>Presentation Error</option>
            <option value="11" {{if .SearchJudge10}}selected{{end}}>System Error</option>
          </select>
        </div>

        <div class="mdl-cell mdl-cell--2-col mdl-cell--2-col-phone">
          <div class="go-select-title">Language</div>
          <select name="search_language" class="go-select">
            <option value="0" {{if .SearchLanguage0}}selected{{end}}>All</option>
            <option value="1" {{if .SearchLanguage1}}selected{{end}}>C</option>
            <option value="2" {{if .SearchLanguage2}}selected{{end}}>C++</option>
            <option value="3" {{if .SearchLanguage3}}selected{{end}}>Java</option>
          </select>
        </div>
        <div class="search-btn mdl-cell mdl-cell--1-col mdl-cell--4-col-phone">
          <button class="mdl-button mdl-js-button mdl-button--icon" type="submit">
            <i class="material-icons">search</i>
          </button>
        </div>
      </div>
    </form>


    {{template "pagination" .}}

    <div class="table-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone mdl-shadow--2dp">
      <table class="go-table text-center">
        <thead>
          <tr>
            <th class="mdl-layout--large-screen-only">Run ID</th>
            <th>User</th>
            <th>Problem</th>
            <th>Result</th>
            <th>Time</th>
            <th>Memory</th>
            <th>Language</th>
            <th class="mdl-layout--large-screen-only">Code Length</th>
            <th>Submit Time</th>
          </tr>
        </thead>
        <tbody>
          {{$privilege := .Privilege}}
          {{$uid := .Uid}}
          {{with .Solution}}  
            {{range .}} 
              {{if ShowStatus .Status}}
                <tr>
                  <td class="mdl-layout--large-screen-only">{{.Sid}}</td>
                  <td><a href="/users/{{.Uid}}">{{.Uid}}</a></td>
                  <td><a href="/problems/{{.Pid}}">{{.Pid}}</a></td>
                  <td class="static-{{.Judge}}">{{ShowJudge .Judge}}</td>
                  <td>{{.Time}} MS</td>
                  <td>{{.Memory}} KB</td>
                  <td>
                    {{ShowLanguage .Language}}
                    {{if or (or (eq $uid .Uid) (LargePU $privilege)) .Share}}
                      <a href="/status/code?sid={{.Sid}}">[view]</a>
                    {{end}}
                  </td>
                  <td class="mdl-layout--large-screen-only">{{.Length}}B</td>
                  <td>{{ShowTime .Create}}</td>
                </tr>
              {{end}}
            {{end}}
          {{else}}
            <td class="mdl-layout--large-screen-only"></td>
            <td></td>
            <td></td>
            <td></td>
            <td>无</td>
            <td></td>
            <td></td>
            <td></td>
            <td class="mdl-layout--large-screen-only"></td>
          {{end}}
        </tbody>
      </table>
    </div>

  </div>
</div>
<script type="text/javascript">
  $('#search_form').submit( function(e) {
    e.preventDefault();
    var uid = $('[name=search_uid]').val();
    var pid = $('[name=search_pid]').val();
    var judge = $('[name=search_judge]').val();
    var language = $('[name=search_language]').val();
    var url = '/status?';
    if (uid != '')
      url += 'uid=' + uid + "&";
    if (pid != '')
      url += 'pid=' + pid + "&";
    if (judge > 0){
      judge = judge-1;
      url += 'judge=' + judge + "&";
    }
    if (language > 0)
      url += 'language=' + language + "&";
    location.href = url;
  });
  </script>
{{end}}
