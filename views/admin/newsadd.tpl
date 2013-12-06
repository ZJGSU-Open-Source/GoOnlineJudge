{{define "content"}}
<p>Add News:</p>
<form id="newsForm">
    <textarea id="news" name="notice"></textarea>
    <div class="section">
        <button class="minibutton ok" type="submit">Submit</button>
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
        var editor = k.create('#news',options);

        $('#newsForm').submit(function(e) {
            e.preventDefault();
            var target = e.target;
            alert($(target).serialize());
        });
    });
</script>
{{end}}