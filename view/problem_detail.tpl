{{define "content"}}
  {{with .Detail}}
    <h1>{{.Title}}</h1>
    <div id="problemInfo" class="rfloat">
      <div class="limit">
        <div class="key">Time Limit</div>
        <div class="value">{{.Time}}ms<br></div>
      </div>
      <div class="limit">
        <div class="key">Memory Limit</div>
        <div class="value">{{.Memory}}kB<br></div>
      </div>
      <div class="checker">
        <div class="key">Judge Program</div>
        <div class="value">
          <span title="纯文本对比">{{ShowSpecial .Special}}</span>
        </div>
      </div>
      <div class="checker">
        <div class="key">Ratio</div>
        <div class="value">
          <span title="纯文本对比">{{ShowRatio .Solve .Submit}}({{.Solve}}/{{.Submit}})</span>
        </div>
      </div>
    </div>
    <div id="problemContent">
      <p><b>Description:</b></p>
      <p>{{.Description}}</p>
      <p><b>Input:</b></p>
      <p>{{.Input}}</p>
      <p><b>Output:</b></p>
      <p>{{.Output}}</p>
      <b>Sample Input:</b>
      <pre>{{.In}}</pre>
      <b>Sample Output:</b>
      <pre>{{.Out}}</pre>
      {{if .Hint}}
        <p><b>Hint:</b></p>
        <p>{{.Hint}}</p>
      {{end}}
      {{if .Source}}
        <p><b>Source:</b></p>
        <p>{{.Source}}</p>
      {{end}}
    </div>
    <hr>

     <a href="#" id="submission_link" onclick="show_submission(); return false;">Submit</a>

    <div id="submission" style="display: none;">
    <form accept-charset="UTF-8" action="/contests/32/1001/submit" method="post"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="✓"><input name="authenticity_token" type="hidden" value="QFsO399GT5G51F2jUHdTmZLGsFPbzB6bgIgxz6KpSJg="></div>
      <div class="field">
        <label for="compiler_id">Compiler</label><br>
        <select id="compiler_id" name="compiler_id">
          <option value="3" selected="selected">C (gcc)</option>
          <option value="12">C# (mcs)</option>
          <option value="2">C++ (g++)</option>
          <option value="14">Go (gccgo)</option>
          <option value="9">Haskell (ghc)</option>
          <option value="24">Java (gcj)</option>
          <option value="21">Javascript (nodejs)</option>
          <option value="16">Lisp (clisp)</option>
          <option value="7">Lua (lua)</option>
          <option value="11">Pascal (fpc)</option>
          <option value="20">PHP (php)</option>
          <option value="4">Python (python3)</option>
          <option value="5">Python (python2)</option>
          <option value="6">Ruby (ruby)</option>
          <option value="17">Scheme (guile)</option>
          <option value="8">Shell (bash)</option>
          <option value="15">Vala (valac)</option>
          <option value="13">VisualBasic (vbnc)</option></select>
      </div>
      <div class="field">
        <label for="code">代码</label><br>
        <textarea id="code" name="code" style=""></textarea><div class="CodeMirror" style="display: none;"><div style="overflow: hidden; position: relative; width: 1px; height: 0px; top: 0px; left: 0px;"><textarea style="position: absolute; width: 2px;" wrap="off" autocorrect="off" autocapitalize="off"></textarea></div><div class="CodeMirror-scroll cm-s-default"><div style="position: relative; height: 1px;"><div style="position: absolute; height: 0; width: 0; overflow: hidden;"><pre><span></span></pre></div><div style="position: relative; top: 0px;"><div class="CodeMirror-gutter" style="height: 250px;"><div class="CodeMirror-gutter-text"><pre>1</pre></div></div><div class="CodeMirror-lines"><div style="position: relative; margin-left: 25px;"><pre class="CodeMirror-cursor" style="top: 0px; left: 0px;">&nbsp;</pre><div style=""><pre> </pre></div></div></div></div></div></div></div>
      </div>
      <div class="actions">
        <input name="commit" type="submit" value="提交代码">
      </div>
    </form></div>

    <script type="text/javascript">
      function show_submission() {
        $('#submission').show();
        $('#submission_link').hide();
      };
      var editor;
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
        ['PHP', 'text/x-php'],
        ['C ', 'text/x-csrc'],
        ['C++ ', 'text/x-c++src'],
        ['Java ', 'text/x-java'],
        ['C#', 'text/x-csharp'],
        ['', 'text/plain'] ];
        for(var i=0;i!=modes.length;++i){
          var n=modes[i], m=modes[i];
          if($.isArray(n)) { m=n[1]; n=n[0]; }
          if(compiler.indexOf(n)>=0){editor.setOption('mode',m.toLowerCase());break;}
        }
      };
      $(document).ready(function() {
        editor = CodeMirror.fromTextArea(document.getElementById("code"), {
          lineNumbers: true,
        }); 
        $('#code').blur(function(){editor.setValue($('#code').val());});
        $('#compiler_id').change(set_mode);
        set_mode();
        toggle_editor();
      $('#submission').hide();
      });
    </script>
  {{end}}
{{end}}