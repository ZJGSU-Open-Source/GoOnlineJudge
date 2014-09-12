<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>{{.Title}}</title>
    <link href="/static/css/style.css" rel="stylesheet" type="text/css" />
    <!--[if lt IE 8]><link href="/static/css/style-lt8.css" media="screen" rel="stylesheet" type="text/css" /><![endif]-->
    <!--[if lt IE 7]><link href="/static/css/style-lt7.css" media="screen" rel="stylesheet" type="text/css" /><![endif]-->
    <script src="/static/js/jquery.min.js" type="text/javascript"></script>
    <script src="/static/js/action.js" type="text/javascript"></script>
    <!-- link href='http://fonts.googleapis.com/css?family=Droid+Sans:400,700' rel='stylesheet' type='text/css' -->
    {{if .IsEdit}}
      <script src="/static/kindeditor/kindeditor-min.js" type="text/javascript"></script>
      <script src="/static/kindeditor/lang/en.js" type="text/javascript"></script>
    {{end}}
  </head>
  <body>
    <div class="container">
      <div id="pageHeader">
        <div id="logo" class="lfloat">
          <a href="/"><img alt="Logo" src="/static/img/logo.png" /></a>
        </div>
        <div id="headerInfo" class="rfloat">
          {{if .IsCurrentUser}}
            <a href="/user/settings">[{{.CurrentUser}}]</a>
            <a class="user_signout" href="#">[Sign Out]</a>
          {{else}}
            {{if .IsUserSignIn}}[Sign In]{{else}}<a href="/user/signin">[Sign In]</a>{{end}}
            {{if .IsUserSignUp}}[Sign Up]{{else}}<a href="/user/signup">[Sign Up]</a>{{end}}
          {{end}}
        </div>
        <hr> 
        </div>
        <div id="navibar" class="span-3">
        <ul>
          <li>{{if .IsHome}}<span>Home</span>{{else}}<a href="/">Home</a>{{end}}</li>
          {{if .IsAdmin}}
          <li><a href="/admin/news/list">News</a></li>
          {{if .IsNews}}
            <div id="psnavi">
              <ul>
                <li>{{if .IsList}}<span>List</sapn>{{else}}<a href="/admin/news/list">List</a>{{end}}</li>
                <li>{{if .IsAdd}}<span>Add</sapn>{{else}}<a href="/admin/news/add">Add</a>{{end}}</li>
              </ul>
            </div>
          {{end}}
          {{end}}
          <li><a href="/admin/problem/list">Problem</a></li>
          {{if .IsProblem}}
            <div id="psnavi">
              <ul>
                <li>{{if .IsList}}<span>List</sapn>{{else}}<a href="/admin/problem/list">List</a>{{end}}</li>
                {{if .IsAdmin}}
                <li>{{if .IsAdd}}<span>Add</sapn>{{else}}<a href="/admin/problem/add">Add</a>{{end}}</li>
                {{end}}
                {{if .RejudgePrivilege}}
                <li>{{if .IsRejudge}}<span>Rejudge</span>{{else}}<a href="/admin/problem/rejudgepage">Rejudge</a>{{end}}</li>
                {{end}}
              </ul>
            </div>
          {{end}}
          <li><a href="/admin/contest/list?type=contest">Contest</a></li>
          {{if .IsContest}}
            <div id="psnavi">
              <ul>
                <li>{{if .IsList}}<span>List</sapn>{{else}}<a href="/admin/contest/list?type=contest">List</a>{{end}}</li>
                <li>{{if .IsAdd}}<span>Add</sapn>{{else}}<a href="/admin/contest/add?type=contest">Add</a>{{end}}</li>
              </ul>
            </div>
          {{end}}
           <li><a href="/admin/contest/list?type=exercise">Exercise</a></li>
          {{if .IsExercise}}
            <div id="psnavi">
              <ul>
                <li>{{if .IsList}}<span>List</sapn>{{else}}<a href="/admin/contest/list?type=exercise">List</a>{{end}}</li>
                <li>{{if .IsAdd}}<span>Add</sapn>{{else}}<a href="/admin/contest/add?type=exercise">Add</a>{{end}}</li>
              </ul>
            </div>
          {{end}}
          {{if .IsAdmin }}
          <li><a href="/admin/user/list">User</a></li>
          {{if .IsUser}}
            <div id="psnavi">
              <ul>
                <li>{{if .IsList}}<span>Privilege</sapn>{{else}}<a href="/admin/user/list">Privilege</a>{{end}}</li>
                <li>{{if .IsPwd}}<span>Password</sapn>{{else}}<a href="/admin/user/pagepassword">Password</a>{{end}}</li>
                <li>{{if .IsGenerate}}<span>Generate</sapn>{{else}}<a href="/admin/user/generate">Generate</a>{{end}}</li>
              </ul>
            </div>
            {{end}}
          {{end}}
        </ul>
      </div>
      <div id="body" class="span-22 last">
        {{template "content" .}}
      </div>
      <div id="pageFooter" class="center">
        <hr class="nomarginbottom">
        <div id="footerContainer">
         <div class="center">ZJGSU Online Judge Beta 2.0 @ <a href="https://github.com/ZJGSU-Open-Source/GoOnlineJudge" target="_new">Github</a></div>
          <div class="center">Copyright &copy; 2013-2014 ZJGSU ACM Club</div>
        </div>
      </div>
    </div>
    <script type="text/javascript">
    $('.user_signout').on('click', function(e) {
      e.preventDefault();
      $.ajax({
        type:'POST',
        url:'/user/logout',
        data:$(this).serialize(),
        error: function() {
          alert('Sign Out Failed.');
        },
        success: function() {
          window.location.href = '/user/signin';
        }
      });
    });
    </script>
  </body>
</html>

