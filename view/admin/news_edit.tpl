{{define "content"}}
<h1>Admin - News Edit</h1>
{{with .Detail}}
<form accept-charset="UTF-8" class="new_news" id="new_news" method="post" action="/admin/news/{{.Nid}}">
    <div style="margin:0;padding:0;display:inline">
      <input name="utf8" type="hidden" value="✓">
    </div>   
    <div class="field">
    	<label for="news_title">Title</label><br>
    	<input id="news_title" name="title" size="60" type="text" value="{{.Title}}">
    </div>
    <div class="field">
    	<label for="news_content">Content</label><br>
    	<textarea id="news_content" name="content" style="width:640px;height:200px;">{{.Content}}</textarea> 
    <div class="actions">
      <input name="commit" type="submit" value="Submit">
    </div>
</form>
{{end}}
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
    window.editor = K.create('#news_content', options);
});
</script>
{{end}}
