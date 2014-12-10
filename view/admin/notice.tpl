{{define "content"}}
<h1>Admin - Notice</h1>
<form accept-charset="UTF-8" class="new_news" id="msg_form" method="post" action="/admin/notice">
 	<div class="field">
    	<label for="title">Title</label><br>
    	<input id="msg" name="msg" size="60" type="text" value="{{.Msg}}">
    </div>
    <div class="actions">
      <input name="commit" type="submit" value="Submit">
    </div>
</form>
{{end}}