<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>{{.Title}}</title>
    <link href="/static/css/style.css" rel="stylesheet" type="text/css" />
    <!--[if lt IE 8]><link href="/static/css/style-lt8.css" media="screen" rel="stylesheet" type="text/css" /><![endif]-->
    <!--[if lt IE 7]><link href="/static/css/style-lt7.css" media="screen" rel="stylesheet" type="text/css" /><![endif]-->
    <script src="/static/js/action.js" type="text/javascript"></script>
    <!-- link href='http://fonts.googleapis.com/css?family=Droid+Sans:400,700' rel='stylesheet' type='text/css' -->
  </head>
  <body>
    <div class="container">
      <div id="pageHeader">
        <div id="logo" class="lfloat">
          <a href="/"><img alt="Logo" src="/static/img/logo.png" /></a>
        </div>
        <div id="headerInfo" class="rfloat">
          {{if .IsUserLogin}}[Sign In]{{else}}<a href="/user/login">[Sign In]</a>{{end}}
          <a href="/user/register">[Sign Up]</a>
        </div>
        <hr> 
        </div>
        <div id="navibar" class="span-3">
        <ul>
          <li>{{if .IsHome}}<span>Home</span>{{else}}<a href="/">Home</a>{{end}}</li>
          <li>{{if .IsProblem}}<span>Problem</span>{{else}}<a href="/problem/list">Problem</a>{{end}}</li>
          {{if .IsProblem}}
            <div id="psnavi">
              <ul>
                <li><a href="/admin/problem/list">Problem List</a></li>
                <li><a href="/admin/problem/add">Add Problem</a></li>
              </ul>
            </div>
          {{end}}
        </ul>
      </div>
      <div id="body" class="span-22 last">
        {{template "content" .}}
      </div>
      <div id="pageFooter" class="center">
        <hr class="nomarginbottom">
        <div id="footerContainer">
          <div class="center">Copyright (C) 2013-2014 Zhejiang Gongshang University ACM Club</div>
        </div>
      </div>
    </div>
  </body>
</html>

