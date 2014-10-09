<html>
  <head>
    <meta charset="utf-8">
      <meta http-equiv="X-UA-Compatible" content="IE=edge">
      <meta name="viewport" content="width=device-width, initial-scale=1">

      <link rel="ico" href="/static/favicon.ico" mce_href="/static/favicon.ico" type="image/x-icon">
      <link rel="shortcut icon" href="/static/favicon.ico" mce_href="/static/favicon.ico" type="image/x-icon">
      <title>{{.Title}}</title>
      <link href="/static/css/style.css" rel="stylesheet" type="text/css">
      <link href="/static/css/bootstrap.min.css" rel="stylesheet">  
    
      <script src="/static/js/jquery.min.js" type="text/javascript"></script>
      <script src="/static/js/action.js" type="text/javascript"></script>

      {{if .IsEdit}}
      <script src="/static/kindeditor/kindeditor-min.js" type="text/javascript"></script>
      <script src="/static/kindeditor/lang/en.js" type="text/javascript"></script>
    {{end}}
  </head>

  <body>
      <div class="container"> 
      <div id="logo" class="lfloat">
            <a href="/"><img alt="Logo" src="/static/img/logo.png"></a>
          </div>
          <hr>
      <nav class="navbar navbar-default" role="navigation">
          <div class="container-fluid">
            <!-- Brand and toggle get grouped for better mobile display -->
            <div class="navbar-header">
              <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                </button>
            </div>
            <!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                <ul class="nav navbar-nav">
                <li><a href="/">Home</a></li>
                <li><a href="/problem/list">Problem</a></li>
                <li><a href="/status/list">Status</a></li>
                <li><a href="/ranklist">Ranklist</a></li>
                <li><a href="/contestlist?type=contest">Contest</a></li>
                {{if .IsContestDetail}}
                  <div id="psnavi">
                  <ul>
                    <li>{{if .IsContestProblem}}<span>Problem</sapn>{{else}}<a href="/contest/problem/list?cid={{.Cid}}">Problem</a>{{end}}</li>
                    <li>{{if .IsContestStatus}}<span>Status</sapn>{{else}}<a href="/contest/status/list?cid={{.Cid}}">Status</a>{{end}}</li>
                    <li>{{if .IsContestRanklist}}<span>Ranklist</sapn>{{else}}<a href="/contest/ranklist?cid={{.Cid}}">Ranklist</a>{{end}}</li>
                  </ul>
                  </div>
                {{end}}
                 <li><a href="/contestlist?type=exercise">Exercise</a></li>
                {{if .IsExerciseDetail}}
                  <div id="psnavi">
                  <ul>
                    <li>{{if .IsExerciseProblem}}<span>Problem</sapn>{{else}}<a href="/contest/problem/list?cid={{.Cid}}">Problem</a>{{end}}</li>
                    <li>{{if .IsExerciseStatus}}<span>Status</sapn>{{else}}<a href="/contest/status/list?cid={{.Cid}}">Status</a>{{end}}</li>
                    <li>{{if .IsExerciseRanklist}}<span>Ranklist</sapn>{{else}}<a href="/contest/ranklist?cid={{.Cid}}">Ranklist</a>{{end}}</li>
                  </ul>
                  </div>
                {{end}}
                {{if .IsCurrentUser}}
                  {{if .IsSettings}}
                  <li><a href="/user/settings">Settings</a></li>
                  <div id="psnavi">
                  <ul>
                    <li>{{if .IsSettingsDetail}}<span>Detail</span>{{else}}<a href="/user/detail?uid={{.CurrentUser}}">Detail</a>{{end}}</li>
                    <li>{{if .IsSettingsEdit}}<span>Edit Info</span>{{else}}<a href="/user/edit">Edit Info</a>{{end}}</li>
                    <li>{{if .IsSettingsPassword}}<span>Password</span>{{else}}<a href="/user/pagepassword">Password</a>{{end}}</li>
                  </ul>
                  </div>
                  {{end}}
                {{end}}
                <li><a href="/osc">OSC</a></li>
                <li><a href="/faq">FAQ</a></li>
                </ul>
                <ul class="nav navbar-nav navbar-right">
                {{if .IsCurrentUser}}
                  {{if .IsShowAdmin}}<li><a href="/admin/">[Admin]</a></li>{{end}}
                  {{if .IsShowTeacher}}<li><a href="/admin/">[Teacher]</a></li>{{end}}
                  <li><a href="/user/settings">[{{.CurrentUser}}]</a></li>
                  <li><a class="user_signout" href="#">[Sign Out]</a></li>
                {{else}}
                  {{if .IsUserSignIn}}{{else}}<li><a href="/user/signin">[Sign In]</a></li>{{end}}
                  {{if .IsUserSignUp}}{{else}}<li><a href="/user/signup">[Sign Up]</a></li>{{end}}
                {{end}}
                </ul>
            </div><!-- /.navbar-collapse -->
          </div><!-- /.container-fluid -->
      </nav>
      <div id="body" class="span-22 last">
        {{template "content" .}}
      </div>
      <div id="pageFooter" class="center">
        <hr class="nomarginbottom">
        <div id="footerContainer">
          <center><div class="center">ZJGSU Online Judge Beta 4.0 @ <a href="https://github.com/ZJGSU-Open-Source/GoOnlineJudge" target="_blank">Github</a></div></center>
            <center><div class="center">Copyright © 2013-2014 ZJGSU ACM Club</div></center>
            <center> <div class="center">Developer: <a href="https://github.com/memelee" target="_blank">@memelee</a> <a href="https://github.com/sakeven" target="_blank">@sakeven</a> <a href="https://github.com/JinweiClarkChao" target="_blank">@JinweiClarkChao</a></div></center>
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
