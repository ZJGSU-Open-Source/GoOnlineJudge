{{define "content"}}
<h1>Admin - Contest Add</h1>
<form accept-charset="UTF-8" class="" id="new_contest" method="post" action="/admin/contests">
	<div style="margin:0;padding:0;display:inline">
    	<input name="utf8" type="hidden" value="✓">
    </div>
    <div class="field">
    	<label for="content_title">Title</label><br>
    	<input id="content_title" name="title" size="60" style="width:640px" type="text" required="">
    </div>
    <div class="field">
    	<label for="content_type">Contest Type</label><br>
    	<select id="content_type" name="type" onchange="show(); return false;">
    		<option value="public">Public</option>
    		<option value="private">Private</option>
    		<option value="password">Password</option>
    	</select>
    </div>
    <div class="field" id="passwd_field" style="display: none;">
      <label for="password">Password</label><br>
      <input id="password" name="password" size="20" type="text">
    </div>
     
    <div class="field">
      <label for="content_startTimeYear">Start Time[Year-Month-Day Hour:Minute]</label><br>
      <input id="content_startTimeYear" name="startTimeYear" size="4" type="number" max="2100" min="2014" value="{{.StartYear}}">-
      <input id="content_startTimeMonth" name="startTimeMonth" size="4" type="number" max="12" min="1" value="{{.StartMonth}}">-
      <input id="content_startTimeDay" name="startTimeDay" type="number" max="31" min="1" value="{{.StartDay}}">
      <input id="content_startTimeHour" name="startTimeHour" size="4" type="number" max="23" min="0" value="{{.StartHour}}">:
      <input id="content_startTimeMinute" name="startTimeMinute" size="4" type="number" max="59" min="0" value="0">:
    </div>
    <div class="field">
      <label for="content_endTimeYear">End Time[Year-Month-Day Hour:Minute]</label><br>
      <input id="content_endTimeYear" name="endTimeYear" size="4" type="number" max="2100" min="2014" value="{{.EndYear}}">-
      <input id="content_endTimeMonth" name="endTimeMonth" size="4" type="number" max="12" min="1" value="{{.EndMonth}}">-
      <input id="content_endTimeDay" name="endTimeDay" type="number" max="31" min="1" value="{{.EndDay}}">
      <input id="content_endTimeHour" name="endTimeHour" size="4" type="number" max="23" min="0" value="{{.EndHour}}">:
      <input id="content_endTimeMinute" name="endTimeMinute" size="4" size="4" type="number" max="59" min="0" value="0">:
    </div>
    <div class="field">
    	<label for="content_problemList">Problem List[Please using ";" split each problem.]</label><br>
    	<input id="content_problemList" name="problemList" size="100" style="width:640px" type="text">
    </div>
    <div class="field" id="userlist_field" style="display: none;">
      <label for="userlist">User List</label><br>
      <textarea id="userlist" name="userlist" style="width:630px;height:200px;"></textarea> 
    </div>
    <div class="actions">
      <input name="commit" type="submit" value="Submit">
    </div>
</form>
<script type="text/javascript">
function show() {
    if($('select').val() == "password"){
      $('#userlist_field').hide();
      $('#passwd_field').show();
    }
    else if($('select').val() == "private"){
      $('#userlist_field').show();
      $('#passwd_field').hide();
    }
    else if($('select').val() == "public"){
      $('#userlist_field').hide();
      $('#passwd_field').hide();
    }
  };
</script>
{{end}}
