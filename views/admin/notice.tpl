{{define "content"}}
<p><b>Edit Notice:</b></p>
<form id="noticeForm">
	<textarea id="notice" name="notice"></textarea>
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
        'source', '|', 'undo', 'redo', '|', 'cut', 'copy', 'paste',
        'plainpaste', '|', 'justifyleft', 'justifycenter', 'justifyright',
        'justifyfull', 'subscript', 'superscript', 'clearhtml', '|', 
        'formatblock', 'fontname', 'fontsize', '|', 'forecolor', 'hilitecolor', 'bold',
        'italic', 'underline', 'strikethrough', '|', 'emoticons', 'link', 'unlink', '|', 'about'
		]
	};

    KindEditor.ready(function(k) {
        var editor = k.create('#notice',options);

        $('#noticeForm').submit(function(e) {
        	e.preventDefault();
        	var target = e.target;
        	alert($(target).serialize());
    	});
    });
</script>
{{end}}