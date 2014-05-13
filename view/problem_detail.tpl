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
  <form accept-charset="UTF-8" method="post" id="problem_submit">
    <div style="margin:0;padding:0;display:inline">
      <input name="utf8" type="hidden" value="✓">
    </div>
    <div class="field">
      <label for="compiler_id">Compiler</label><br>
      <select id="compiler_id" name="compiler_id">
        <option value="1" selected="selected">C (gcc)</option>
        <option value="2">C++ (g++)</option>
        <option value="3">Java (gcj)</option>
    </div>
    <div class="field">
      <label for="code">Code</label><br>
      <textarea id="code" name="code" style="" required=""></textarea>
    </div>
    <div class="actions">
      <input name="submit" type="submit" value="Submit">
    </div>
  </form></div>

  <script type="text/javascript">
  function show_submission() {
    $('#submission').show();
    $('#submission_link').hide();
  };
  $('#problem_submit').submit(function(e) {
    e.preventDefault();
    $.ajax({
      type:'POST',
      url:'/problem/submit/pid/{{.Pid}}',
      data:$(this).serialize(),
      error: function(XMLHttpRequest) {
        if(XMLHttpRequest.status == 401){
          alert('Please Sign In.');
          window.location.href = '/user/signin';
        }
      },
      success: function(result) {
        window.location.href = '/status/list';
      }
    });
  });
  </script>
{{end}}
{{end}}