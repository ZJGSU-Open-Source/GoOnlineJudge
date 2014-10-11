<html>
  <head>
    <meta charset="utf-8">
      <meta http-equiv="X-UA-Compatible" content="IE=edge">
      <meta name="viewport" content="width=device-width, initial-scale=1">

      <link rel="ico" href="/static/favicon.ico" mce_href="/static/favicon.ico" type="image/x-icon">
      <link rel="shortcut icon" href="/static/favicon.ico" mce_href="/static/favicon.ico" type="image/x-icon">
      <title>{{.Title}}</title>
      <link href="/static/css/style.css" rel="stylesheet" type="text/css">
      <link href="/static/css/bootstrap.css" rel="stylesheet">
    
      <link href="/static/material/css/material.css" rel="stylesheet">


      <script src="/static/js/jquery.min.js" type="text/javascript"></script>
      <script src="/static/js/action.js" type="text/javascript"></script>

      {{if .IsCode}}
      <link href="/static/prettify/prettify.css" rel="stylesheet" type="text/css" />
      <script src="/static/prettify/prettify.js" type="text/javascript"></script>
      {{end}}
  </head>

  <body {{if .IsCode}}onload="prettyPrint()"{{end}}>

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
                 <li><a href="/contestlist?type=exercise">Exercise</a></li>
                
                <li><a href="/osc">OSC</a></li>
                <li><a href="/faq">FAQ</a></li>
                </ul>
                <ul class="nav navbar-nav navbar-right">
                {{if .IsCurrentUser}}
                  {{if .IsShowAdmin}}<li><a href="/admin/" class="icon icon-material-star">[Admin]</a></li>{{end}}
                  {{if .IsShowTeacher}}<li><a href="/admin/" class="icon icon-material-star-half">[Teacher]</a></li>{{end}}
                  <li><a href="/user/settings" class="icon icon-material-settings">[{{.CurrentUser}}]</a></li>
                  <li><a class="user_signout icon icon-material-chevron-right" href="#">[Sign Out]</a></li>
                {{else}}
                  {{if .IsUserSignIn}}{{else}}<li><a href="/user/signin">[Sign In]</a></li>{{end}}
                  {{if .IsUserSignUp}}{{else}}<li><a href="/user/signup">[Sign Up]</a></li>{{end}}
                {{end}}
                </ul>
            </div><!-- /.navbar-collapse -->
          </div><!-- /.container-fluid -->
      </nav>

      <div id="body"> 
      
                {{if .IsContestDetail}}
                <div class="pinned note" >
                    {{if .IsContestProblem}}<span>Problem</sapn>{{else}}<a href="/contest/problem/list?cid={{.Cid}}">Problem</a>{{end}}
                    {{if .IsContestStatus}}<span>Status</sapn>{{else}}<a href="/contest/status/list?cid={{.Cid}}">Status</a>{{end}}
                    {{if .IsContestRanklist}}<span>Ranklist</sapn>{{else}}<a href="/contest/ranklist?cid={{.Cid}}">Ranklist</a>{{end}}
                </div>
                  {{end}}
                {{if .IsExerciseDetail}}
                <div class="pinned note" >
                    {{if .IsExerciseProblem}}<span>Problem</sapn>{{else}}<a href="/contest/problem/list?cid={{.Cid}}">Problem</a>{{end}}
                    {{if .IsExerciseStatus}}<span>Status</sapn>{{else}}<a href="/contest/status/list?cid={{.Cid}}">Status</a>{{end}}
                    {{if .IsExerciseRanklist}}<span>Ranklist</sapn>{{else}}<a href="/contest/ranklist?cid={{.Cid}}">Ranklist</a>{{end}}
                </div>
                {{end}}
                {{if .IsCurrentUser}}
                  {{if .IsSettings}}
                  <div class="pinned note" >
                    {{if .IsSettingsDetail}}<span>Detail</span>{{else}}<a href="/user/detail?uid={{.CurrentUser}}">Detail</a>{{end}}
                    {{if .IsSettingsEdit}}<span>Edit Info</span>{{else}}<a href="/user/edit">Edit Info</a>{{end}}
                    {{if .IsSettingsPassword}}<span>Password</span>{{else}}<a href="/user/pagepassword">Password</a>{{end}}
                  </div>
                  {{end}}
                {{end}}
</div> 
<script src="/static/js/jquery.pin.js" type="text/javascript"></script>
<script type="text/javascript">
  $(".pinned").pin();
</script>

      <div class="jumbotron">
        {{template "content" .}}
        </div>
      </div>
      <div id="pageFooter" class="center">
        <hr class="nomarginbottom">
        <div id="footerContainer">
          <center><div class="center">ZJGSU Online Judge Beta 4.0 @ <a href="https://github.com/ZJGSU-Open-Source/GoOnlineJudge" target="_blank">Github</a></div></center>
            <center><div class="center">Copyright © 2013-2014 ZJGSU ACM Club</div></center>
            <center> <div class="center">Developer: <a href="https://github.com/memelee" target="_blank">@memelee</a> <a href="https://github.com/sakeven" target="_blank">@sakeven</a> <a href="https://github.com/JinweiClarkChao" target="_blank">@JinweiClarkChao</a> <a href="https://github.com/rex-zsd" target="_blank">@rex-zsd</a></div></center>
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
