<!DOCTYPE html>
<html>
	<head>
			<meta charset="utf-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<link rel="stylesheet" href="/static/css/docs.css">
			<link rel="stylesheet" href="/static/css/admin.css">
			<title>Admin</title>
	</head>
	<body>
		<div id="fadeout"></div>
		<div id="flybar">
			<div class="navigation home">
      			<div class="button">Home</div>
    		</div>
			<div class="navigation">
      			<div class="button">Online Practice</div>
      			<div class="contents menu">
			        <a href="/problem">Problem List</a>
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
			<div id="login" class="navigation"></div>
      		<div id="error" style="display: none;"></div>
		</div>
		<div class="container">
			<div class="side">
				<ul>
					<li><a href="/admin/notice">Edit Notice</a></li>
					<li><a href="/admin/news/add">Add News</a></li>
					<li><a href="/admin/news/edit">Edit News</a></li>
					<li><a href="/admin/problem/add">Add  Problem</a></li>
					<li>Edit Problem</li>
					<li>Add  Exam</li>
					<li>Edit Exam</li>
					<li>Add  Contest</li>
					<li>Edit Contest</li>
					<li>Add  Privilege</li>
					<li>Edit Privilege</li>
					<li><hr></li>
					<li>Generate User ID</li>
					<li>Modify   Password</li>
					<li>Rejudge  Problem</li>
				</ul>
			</div>
			<script src="/static/js/jquery.min.js" type="text/javascript"></script>
			<script charset="utf-8" src="/static/kindeditor/kindeditor.js" type="text/javascript"></script>
			<script charset="utf-8" src="/static/kindeditor/lang/zh_CN.js" type="text/javascript"></script>
			<div class="content">
				{{template "content" .}}
			</div>
			<div class="clearboth"></div>
		</div>
		<script src="/static/js/jquery.cookie.js" type="text/javascript"></script>
		<script src="/static/js/operation.js" type="text/javascript"></script>
	</body>
</html>