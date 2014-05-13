{{define "content"}}
<h1>Admin - Contest Edit</h1>
{{with .Detail}}
<form accept-charset="UTF-8" class="" id="new_contest" method="post" action="/admin/contest/update/cid/{{.Cid}}">
	<div style="margin:0;padding:0;display:inline">
    	<input name="utf8" type="hidden" value="✓">
    </div>
    <div class="field">
    	<label for="content_title">Title</label><br>
    	<input id="content_title" name="title" size="60" type="text" value="{{.Title}}">
    </div>
    <div class="field">
    	<label for="content_type" value="Public">Contest Type</label><br>
    	<select id="content_type" name="type">
    		<option value="public" {{if .IsPublic}}selected{{end}}>Public</option>
    		<option value="private" {{if .IsPrivate}}selected{{end}}>Private</option>
    		<option value="password" {{if .IsPassword}}selected{{end}}>Password</option>
    	</select>
    </div>
    <div class="field">
      <label for="content_startTimeYear">Start Time[Year-Month-Day Hour:Minute:Second]</label><br>
      <input id="content_startTimeYear" name="startTimeYear" size="3" type="text" value="{{.StartTimeYear}}">-
      <input id="content_startTimeMonth" name="startTimeMonth" size="1" type="text" value="{{.StartTimeMonth}}">-
      <input id="content_startTimeDay" name="startTimeDay" size="1" type="text" value="{{.StartTimeDay}}">   
      <input id="content_startTimeHour" name="startTimeHour" size="1" type="text" value="{{.StartTimeHour}}">:
      <input id="content_startTimeMinute" name="startTimeMinute" size="1" type="text" value="{{.StartTimeMinute}}">:
      <input id="content_startTimeSecond" name="startTimeSecond" size="1" type="text" value="{{.StartTimeSecond}}">
    </div>
    <div class="field">
      <label for="content_endTimeYear">End Time[Year-Month-Day Hour:Minute:Second]</label><br>
      <input id="content_endTimeYear" name="endTimeYear" size="3" type="text" value="{{.EndTimeYear}}">-
      <input id="content_endTimeMonth" name="endTimeMonth" size="1" type="text" value="{{.EndTimeMonth}}">-
      <input id="content_endTimeDay" name="endTimeDay" size="1" type="text" value="{{.EndTimeDay}}">   
      <input id="content_endTimeHour" name="endTimeHour" size="1" type="text" value="{{.EndTimeHour}}">:
      <input id="content_endTimeMinute" name="endTimeMinute" size="1" type="text" value="{{.EndTimeMinute}}">:
      <input id="content_endTimeSecond" name="endTimeSecond" size="1" type="text" value="{{.EndTimeSecond}}">
    </div>
    <div class="field">
    	<label for="content_problemList">Problem List[Please using ";" split each problem.]</label><br>
    	<input id="content_problemList" name="problemList" size="100" type="text" value="{{.ProblemList}}">
    </div>
    <div class="actions">
      <input name="commit" type="submit" value="Submit">
    </div>
</form>
{{end}}
{{end}}