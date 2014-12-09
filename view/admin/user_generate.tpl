{{define "content"}}
<h1>Team Account Generate</h1>

<form accept-charset="UTF-8" id="search_form" action="/admin/users/generation" method="post">
<div style="margin:0;padding:0;display:inline">
	<input name="utf8" type="hidden" value="✓">
</div>
<div class="field">
	<label for="user_handle">User Prefix Handle</label><font color="red">*</font>
	<br/>
	<input id="user_handle" name="prefix" size="30" type="text" required>
</div>	
<div class="field">
	<label for="user_amount">Amount</label><font color="red">*</font>
	<br/>
	<input id="user_amount" name="amount" size="30" type="text" pattern="^[0-9]+" required>
</div>
<div class="field">
<label for="module">Account Type</label><br/>
<select id="module" name="module">
<option value="0">Normal</option>
<option value="1">Team</option>
</select>
</div>
<div class="actions">
	<input name="user_password" type="submit" value="Submit" required="required" />
</div>
</form>
{{end}}