{{define "content"}}
<h1>Admin - Contest Add</h1>
<form accept-charset="UTF-8" class="" id="new_contest" method="post" action="/admin/contest?insert/type?{{.Type}}">
	<div style="margin:0;padding:0;display:inline">
    	<input name="utf8" type="hidden" value="✓">
    </div>
    <div class="field">
    	<label for="content_title">Title</label><br>
    	<input id="content_title" name="title" size="60" type="text" required="">
    </div>
    <div class="field">
    	<label for="content_type">Contest Type</label><br>
    	<select id="content_type" name="type">
    		<option value="public">Public</option>
    		<option value="private">Private</option>
    		<option value="password">Password</option>
    	</select>
    </div>
    <div class="field">
      <label for="content_startTimeYear">Start Time[Year-Month-Day Hour:Minute:Second]</label><br>
      <input id="content_startTimeYear" name="startTimeYear" size="4" type="text">-
      <input id="content_startTimeMonth" name="startTimeMonth" size="4" type="text">-
      <input id="content_startTimeDay" name="startTimeDay" size="4" type="text">   
      <input id="content_startTimeHour" name="startTimeHour" size="4" type="text">:
      <input id="content_startTimeMinute" name="startTimeMinute" size="4" type="text">:
      <input id="content_startTimeSecond" name="startTimeSecond" size="4" type="text">
    </div>
    <div class="field">
      <label for="content_endTimeYear">End Time[Year-Month-Day Hour:Minute:Second]</label><br>
      <input id="content_endTimeYear" name="endTimeYear" size="4" type="text">-
      <input id="content_endTimeMonth" name="endTimeMonth" size="4" type="text">-
      <input id="content_endTimeDay" name="endTimeDay" size="4" type="text">   
      <input id="content_endTimeHour" name="endTimeHour" size="4" type="text">:
      <input id="content_endTimeMinute" name="endTimeMinute" size="4" type="text">:
      <input id="content_endTimeSecond" name="endTimeSecond" size="4" type="text">
    </div>
    <div class="field">
    	<label for="content_problemList">Problem List[Please using ";" split each problem.]</label><br>
    	<input id="content_problemList" name="problemList" size="100" type="text">
    </div>
    <div class="actions">
      <input name="commit" type="submit" value="Submit">
    </div>
</form>
{{end}}
