{{define "head"}}
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<link type="text/css" rel="stylesheet" href="css/docs.css">
		<title>{{.Title}}</title>
	</head>
	<body>
		<div id="fadeout"></div>
		<div id="flybar">
			<div class="navigation no">
      			<div class="button">Home</div>
    		</div>
			<div class="navigation">
      			<div class="button">Online Practice</div>
      			<div class="contents menu">
			        <a href="/problem/list">Problem List</a>
			        <a href="#installation">Realtime Status</a>
			        <a href="#usage">User Ranklist</a>
			    </div>
      		</div>
			<div class="navigation">
      			<div class="button">Online Competition</div>
      			<div class="contents menu">
			        <a href="#overview">Standard Contest</a>
			        <a href="#installation">Teaching Exam</a>
			    </div>
      		</div>
      		<div class="navigation">
      			<div class="button">Online Help</div>
      			<div class="contents menu">
			        <a href="#overview">Judge Information</a>
			        <a href="#overview">Frequently Asked Questions</a>
			    </div>
      		</div>
			<div class="navigation">
      			<div class="button">{{.User}}</div>
      			<div class="contents menu">
      				<form id="login">
	      				<input type="text" name="uid" placeholder="User ID">
	      				<input type="password" name="pwd" placeholder="Password">
	      				<button class="minibutton ok" type="submit">Sign In</button>
	      				<button class="minibutton" type="button">Sign Up</button>
      				</form>
      			</div>
      		</div>
      		<div id="error" style="display: none;"></div>
		</div>
{{end}}