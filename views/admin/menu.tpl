<!DOCTYPE html>
<html>
	<head>
			<meta charset="utf-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<link rel="stylesheet" href="/static/css/admin.css">
			<title>Admin</title>
	</head>
	<body>
		<div class="container">
			<div class="menu">
				<ul>
					<li>修改公告</li>
					<li>编辑新闻</li>
					<li>编辑题目</li>
					<li>编辑练习</li>
					<li>编辑竞赛</li>
					<li>编辑权限</li>
					<li>修改密码</li>
					<li>题目重判</li>
					<li>账号生成</li>
				</ul>
			</div>
			<div class="content">
				{{template "content" .}}
			</div>
		</div>
	</body>
</html>