<!DOCTYPE HTML>
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
      <script type="text/x-mathjax-config">MathJax.Hub.Config({tex2jax: {inlineMath: [['$','$'], ['\\(','\\)']]}});</script>
      <script type="text/javascript" src="http://cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML"></script>
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
                {{if .IsNews}}<li class="active"><a href="/">Home</a></li>{{else}}<li><a href="/">Home</a></li>{{end}}
                {{if .IsProblem}}<li class="active"><a href="/problems">Problems</a>{{else}}<li><a href="/problems">Problems</a></li>{{end}}
                {{if .IsStatus}}<li class="active"><a href="/status">Status</a></li>{{else}}<li><a href="/status">Status</a></li>{{end}}
                {{if .IsRanklist}}<li class="active"><a href="/ranklist">Ranklist</a></li>{{else}}<li><a href="/ranklist">Ranklist</a></li>{{end}}
                {{if .IsContest}}<li class="active"><a href="/contests">Contests</a></li>{{else}}<li><a href="/contests">Contests</a></li>{{end}}
                {{if.IsOSC}}<li class="active"><a href="/osc">OSC</a></li>{{else}}<li><a href="/osc">OSC</a></li>{{end}}
                {{if.IsFAQ}}<li class="active"><a href="/faq">FAQ</a></li>{{else}}<li><a href="/faq">FAQ</a></li>{{end}}
                </ul>
                <ul class="nav navbar-nav navbar-right">
                {{if .IsCurrentUser}}
                  {{if .IsShowAdmin}}<li><a href="/admin/" class="icon icon-material-star">[Admin]</a></li>{{end}}
                  {{if .IsShowTeacher}}<li><a href="/admin/" class="icon icon-material-star-half">[Teacher]</a></li>{{end}}
                  <li><a href="/settings" class="icon icon-material-settings">[{{.CurrentUser}}]</a></li>
                  <li><a class="user_signout icon icon-material-chevron-right" href="#">[Sign Out]</a></li>
                {{else}}
                  {{if .IsUserSignIn}}{{else}}<li><a class="icon icon-material-account-circle" href="/sess">[Sign In]</a></li>{{end}}
                  {{if .IsUserSignUp}}{{else}}<li><a class="icon icon-material-person-add" href="/users/new">[Sign Up]</a></li>{{end}}
                {{end}}
                </ul>
            </div><!-- /.navbar-collapse -->
          </div><!-- /.container-fluid -->
      </nav>
      {{if .Msg}}
     <center>
      <marquee style='width:60%;height:30px' scrollamount=2 direction=left scrolldelay=30 onMouseOver='this.stop()' onMouseOut='this.start()'>{{.Msg}}</marquee>
     </center>
     {{end}}
      <div id="body"> 
      {{if .IsContestDetail}}
      <div class="pinned note" >
          <div class="icon icon-material-add-circle" style="float:right"></div>
          {{if .IsContestProblem}}<span>Problem</sapn>{{else}}<a href="/contests/{{.Cid}}">Problem</a>
          {{end}}
          <br/>
          {{if .IsContestStatus}}<span>Status</sapn>{{else}}<a href="/contests/{{.Cid}}/status">Status</a>{{end}}
          {{if .IsContestRanklist}}<span>Ranklist</sapn>{{else}}<a href="/contests/{{.Cid}}/ranklist">Ranklist</a>{{end}}
      </div>
      {{end}}
      {{if .IsCurrentUser}}
        {{if .IsSettings}}
        <div class="pinned note" >
        <div class="icon icon-material-add-circle" style="float:right"></div>
          {{if .IsSettingsDetail}}<span>Detail</span>{{else}}<a href="/users/{{.CurrentUser}}">Detail</a>{{end}}

          <br/>
          {{if .IsSettingsEdit}}<span>Edit Info</span>{{else}}<a href="/profile">Edit Info</a>{{end}}
          <br/>
          {{if .IsSettingsPassword}}<span>Password</span>{{else}}<a href="/account">Password</a>{{end}}
        </div>
        {{end}}
      {{end}}
<script src="/static/js/jquery.pin.js" type="text/javascript"></script>
<script type="text/javascript">
  $(".pinned").pin();
</script>

      <div class="jumbotron">
        {{template "content" .}}
      </div>
      <div id="pageFooter" class="center">
        <hr class="nomarginbottom">
        <div id="footerContainer">
          <center><div class="center">ZJGSU Online Judge Version 15.01.31 @ <a href="https://github.com/ZJGSU-Open-Source/GoOnlineJudge" target="_blank">Github</a></div></center>
            <center><div class="center">Copyright © 2013-2014 ZJGSU ACM Club</div></center>
            <center> <div class="center">Developer: <a href="https://github.com/memelee" target="_blank">@memelee</a> <a href="https://github.com/sakeven" target="_blank">@sakeven</a> <a href="https://github.com/JinweiClarkChao" target="_blank">@JinweiClarkChao</a> <a href="https://github.com/rex-zsd" target="_blank">@rex-zsd</a></div></center>
        </div>
      </div>
      </div>
      <script type="text/javascript">
        $('.user_signout').on('click', function(e) {
            e.preventDefault();
            $.ajax({
              type:'DELETE',
              url:'/sess',
              data:$(this).serialize(),
              error: function() {
                  alert('Sign Out Failed.');
              },
              success: function() {
                  window.location.href = '/sess';
              }
            });
        });
      </script>
      <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="/static/js/jquery.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>
  </body>
</html>

