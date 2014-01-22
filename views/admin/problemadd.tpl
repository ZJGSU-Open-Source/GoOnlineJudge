{{define "content"}}
<p><b>Add Problem:</b></p>
<form id="problemForm">
    <div class="section">
        <span class="label">Title:</span>
        <input type="text" name="title" style="width: 640px;">
    </div>
    <div class="section">
        <span class="label">Source:</span>
        <input type="text" name="source" style="width: 640px;">
    </div>
    <div class="section">
        <span class="label">Time:</span>
        <input type="text" name="time" style="width: 300px;">
        <span class="label">MS</span>
    </div>
    <div class="section">
        <span class="label">Memory:</span>
        <input type="text" name="memory" style="width: 300px;">
        <span class="label">KB</span>
    </div>
    <div class="section">
        <span>Description:</span>
        <textarea id="description" name="description"></textarea>
    </div>
    <div class="section">
        <span>Input:</span>
        <textarea id="input" name="input"></textarea>
    </div>
    <div class="section">
        <span class="label">Output:</span>
        <textarea id="output" name="output"></textarea>
    </div>
    <div class="section">
        <span>Sample Input:</span>
        <textarea id="sampleInput" name="sampleInput" style="width: 744px; height: 300px;"></textarea>
    </div>
    <div class="section">
        <span class="label">Sample Output:</span>
        <textarea id="sampleOutput" name="sampleOutput" style="width: 744px; height: 300px;"></textarea>
    </div>
    <div class="section">
        <span>Hint:</span>
        <textarea id="hint" name="hint"></textarea>
    </div>
    <div class="section">
        <button class="minibutton ok" type="submit">Submit & Add Test Cases</button>
    </div>
</form>

<script>
    var options = {
        width: '748px',
        height: '300px',
    	resizeType: 1,
    	items: [
            'source', '|', 'undo', 'redo', '|', 'preview', 'print', 'code', 'cut', 'copy', 'paste',
            'plainpaste', '|', 'justifyleft', 'justifycenter', 'justifyright',
            'justifyfull', 'insertorderedlist', 'insertunorderedlist', 'subscript',
            'superscript', 'clearhtml', '|', 'fullscreen', '/',
            'formatblock', 'fontname', 'fontsize', '|', 'forecolor', 'hilitecolor', 'bold',
            'italic', 'underline', 'strikethrough', 'lineheight', '|', 'image', 'multiimage',
            'table', 'hr', 'emoticons', 'baidumap', 'link', 'unlink', '|', 'about'
    	]
    };

    KindEditor.ready(function(k) {
        var editor = k.create('#description',options);
        var editor = k.create('#input',options);
        var editor = k.create('#output',options);
        var editor = k.create('#hint',options);

        $('#problemForm').submit(function(e) {
            e.preventDefault();
            var target = e.target;
            var action = '/problemAjax/insert';
            $.post(action, $(target).serialize(), function(json) {
                if (json.ok) {
                    alert('Successful!');
                } else {
                    alert('Failed!');
                }
            });
        });
    });

    
</script>
{{end}}