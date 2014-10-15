{{define "content"}}
  {{with .Detail}}
    <h1 style="text-align: center">{{.Title}}</h1>
    <div id="problemInfo" class="rfloat" title="Problem Information">
      <div class="limit">
        <div class="key">Time Limit</div>
        <div class="value">{{.Time}}ms<br></div>
      </div>
      <div class="limit">
        <div class="key">Memory Limit</div>
        <div class="value">{{.Memory}}KB<br></div>
      </div>
      <div class="checker">
        <div class="key">Judge Program</div>
        <div class="value">
          <span>{{ShowSpecial .Special}}</span>
        </div>
      </div>
      <div class="checker">
        <div class="key">Ratio(Solve/Submit)</div>
        <div class="value">
          <span>{{ShowRatio .Solve .Submit}}({{.Solve}}/{{.Submit}})</span>
        </div>
      </div>
    </div>
    <div id="problemContent">
      <div class="problemIteam">Description:</div>
    <p>{{.Description}}</p>
    <div class="problemIteam">Input:</div>
    <p>{{.Input}}</p>
    <div class="problemIteam">Output:</div>
    <p>{{.Output}}</p>
    <div class="problemIteam">Sample Input:</div>
    <pre class="sample">{{.In}}</pre>
    <div class="problemIteam">Sample Output:</div>
    <pre class="sample">{{.Out}}</pre>
    {{if .Hint}}
      <div class="problemIteam">Hint:</div>
      <p>{{.Hint}}</p>
    {{end}}
    </div>
    {{end}}
    <hr>
  <a href="#" id="submission_link" onclick="show_submission(); return false;">Submit</a>
  <script src="/static/js/codemirror.js" type="text/javascript"></script>
  <div id="submission" style="display: none;">
  <form accept-charset="UTF-8" method="post" id="problem_submit">
    <div style="margin:0;padding:0;display:inline">
      <input name="utf8" type="hidden" value="✓">
    </div>
    <div class="field">
      <label for="compiler_id">Compiler</label><br>
      <select id="compiler_id" name="compiler_id">
        <option value="1" selected="selected">C</option>
        <option value="2">C++</option>
        <option value="3">Java</option>
      </select>
      <font  id="warning" color="red"></font>
    </div>
    <div class="field">
      <div class="rfloat">
       <input checked="checked" id="advanced_editor" name="advanced_editor" onchange="toggle_editor()" onclick="toggle_editor()" type="checkbox" value="1" />
        use advanced editor
      </div>
      <label for="code">Code</label><br>
      <textarea id="code" name="code"  autofocus=""></textarea>
    </div>
    <div class="actions">
      <input name="submit" type="submit" value="Submit">
    </div>
  </form></div>
  <script type="text/javascript">
  var editor;
  function show_submission() {
    $('#submission').show();
    $('#submission_link').hide();
    editor = CodeMirror.fromTextArea(document.getElementById("code"), {
      lineNumbers: true,
    }); 
    $('#code').blur(function(){editor.setValue($('#code').val());});
    $('#compiler_id').change(set_mode);
    set_mode();
    toggle_editor();
  };
  $('#problem_submit').submit(function(e) {
    $('#code').val(editor.getValue());
    e.preventDefault();
    $.ajax({
      type:'POST',
      url:'/contest/problem/submit?pid={{.Pid}}&cid={{.Cid}}',
      data:$(this).serialize(),
      error: function(XMLHttpRequest) {
        if(XMLHttpRequest.status == 401){
          alert('Please Sign In.');
          window.location.href = '/user/signin';
        }else {
          var json = eval('('+XMLHttpRequest.responseText+')');
          if(json.info != null) {
            $('#warning').text(json.info);
          } else {
            $('#warning').text('');
          }
        }
      },
      success: function(result) {
        $('textarea').val('')
        window.location.href = '/contest/status/list?cid={{.Cid}}';
      }
    });
  });
  function toggle_editor() {
    var cm=$('.CodeMirror'), c=$('#code');
    if($('#advanced_editor').prop('checked')) {
      cm.show();
      editor.setValue(c.val());
      c.hide();
    } else {
      c.val(editor.getValue()).show();
      cm.hide();
    };
    return true;
  }
  function set_mode() {
    var compiler=$('#compiler_id option:selected').text();
    var modes=[ 
    'Javascript', 'Haskell', 'Lua', 'Pascal', 'Python', 'Ruby', 'Scheme', 'Smalltalk', 'Clojure',
    ['C', 'text/x-csrc'],
    ['C++', 'text/x-c++src'],
    ['Java', 'text/x-java'],
    ['', 'text/plain'] ];
    for(var i=0;i!=modes.length;++i){
      var n=modes[i], m=modes[i];
      if($.isArray(n)) { m=n[1]; n=n[0]; }
      if(compiler.indexOf(n)>=0){editor.setOption('mode',m.toLowerCase());break;}
    }
  };
  </script>
{{end}}