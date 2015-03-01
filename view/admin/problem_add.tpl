{{define "content"}}
<h1>Admin - Problem Add</h1>
<form accept-charset="UTF-8" class="new_problem" id="new_problem" method="post" action="/admin/problems/">
    <div style="margin:0;padding:0;display:inline">
      <input name="utf8" type="hidden" value="✓">
    </div>
    <div class="field">
      <label for="problem_title">Title</label><br>
      <input id="problem_title" name="title" size="60" type="text">
    </div>
    <div class="field">
      <label for="problem_time">Time Limit</label><br>
      <input id="problem_time" name="time" size="30" type="number" value="1"> S
    </div>
    <div class="field">
      <label for="problem_memory">Memory Limit</label><br>
      <input id="problem_memory" name="memory" size="30" type="number" value="32768"> kB
    </div>
    <div class="field">
      <label for="problem_source">Source</label><br>
      <input id="problem_source" name="source" size="60" type="text">
    </div>
    <div class="field">
      <label for="problem_special">Special Judge</label><br>
      <input id="problem_special" name="special" type="checkbox" value="1">
    </div>
    <div class="field">
      <label for="problem_description">Description</label><br>
      <input id="problem_description" name="description" type="text">
    </div>
    <div class="field">
      <label for="problem_input">Input</label><br>
      <input id="problem_input" name="input" type="text">
    </div>
     <div class="field">
      <label for="problem_output">Output</label><br>
      <input id="problem_output" name="output" type="text">
    </div>
    <div class="field">
      <label for="problem_hint">Hint</label><br>
      <input id="problem_hint" name="hint" type="text">
    </div>
    <div class="field">
      <label for="problem_in">Sample Input</label><br>
      <textarea id="problem_in" name="in" style="width:640px;height:200px;"></textarea>
    </div>
    <div class="field">
      <label for="problem_out">Sample Output</label><br>
      <textarea id="problem_out" name="out" style="width:640px;height:200px;"></textarea>
    </div>
    <div class="actions">
      <input name="commit" type="submit" value="Submit">
    </div>
</form>
<script>
var options = {
	height: '250px',
	langType : 'en',
	items: [
        'source', '|', 'undo', 'redo', '|', 
        'preview', 'code', 'cut', 'copy', 'paste', 'plainpaste', 'wordpaste', '|', 
        'justifyleft', 'justifycenter', 'justifyright', 'justifyfull', 
        'insertorderedlist', 'insertunorderedlist', 'subscript', 'superscript', 
        'clearhtml', '|', 'fullscreen', '/', 'formatblock', 'fontname', 'fontsize', '|', 
        'forecolor', 'hilitecolor', 'bold', 'italic', 'underline', 'strikethrough', 
        'removeformat', '|', 'image', 'table', 'hr', 
        'emoticons', 'baidumap', 'link', 'unlink', '|', 'about'
	]
}

KindEditor.ready(function(K) {
    window.editor = K.create('#problem_description', options);
    window.editor = K.create('#problem_input', options);
    window.editor = K.create('#problem_output', options);
    window.editor = K.create('#problem_hint', options);
});
</script>
{{end}}
