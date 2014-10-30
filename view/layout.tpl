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
                {{if .IsNews}}<li class="active"><a href="/">Home</a></li>{{else}}<li><a href="/">Home</a></li>{{end}}
                {{if .IsProblem}}<li class="active"><a href="/problem/list">Problem</a>{{else}}<li><a href="/problem/list">Problem</a></li>{{end}}
                {{if .IsStatus}}<li class="active"><a href="/status/list">Status</a></li>{{else}}<li><a href="/status/list">Status</a></li>{{end}}
                {{if .IsRanklist}}<li class="active"><a href="/ranklist">Ranklist</a></li>{{else}}<li><a href="/ranklist">Ranklist</a></li>{{end}}
                {{if .IsContest}}<li class="active"><a href="/contestlist">Contest</a></li>{{else}}<li><a href="/contestlist">Contest</a></li>{{end}}
                {{if.IsOSC}}<li class="active"><a href="/osc">OSC</a></li>{{else}}<li><a href="/osc">OSC</a></li>{{end}}
                {{if.IsFAQ}}<li class="active"><a href="/faq">FAQ</a></li>{{else}}<li><a href="/faq">FAQ</a></li>{{end}}
                </ul>
                <ul class="nav navbar-nav navbar-right">
                {{if .IsCurrentUser}}
                  {{if .IsShowAdmin}}<li><a href="/admin/" class="icon icon-material-star">[Admin]</a></li>{{end}}
                  {{if .IsShowTeacher}}<li><a href="/admin/" class="icon icon-material-star-half">[Teacher]</a></li>{{end}}
                  <li><a href="/user/settings" class="icon icon-material-settings">[{{.CurrentUser}}]</a></li>
                  <li><a class="user_signout icon icon-material-chevron-right" href="#">[Sign Out]</a></li>
                {{else}}
                  {{if .IsUserSignIn}}{{else}}<li><a class="icon icon-material-account-circle" href="/user/signin">[Sign In]</a></li>{{end}}
                  {{if .IsUserSignUp}}{{else}}<li><a class="icon icon-material-person-add" href="/user/signup">[Sign Up]</a></li>{{end}}
                {{end}}
                </ul>
            </div><!-- /.navbar-collapse -->
          </div><!-- /.container-fluid -->
      </nav>
      {{if .Notice}}
     <center>
      <marquee style='width:60%;height:60px' scrollamount=2 direction=left scrolldelay=30 onMouseOver='this.stop()' onMouseOut='this.start()'>123</marquee>
     </center>
     {{end}}
      <div id="body"> 
      {{if .IsContestDetail}}
      <div class="pinned note" >
          <div class="icon icon-material-add-circle" style="float:right"></div>
          {{if .IsContestProblem}}<span>Problem</sapn>{{else}}<a href="/contest/problem/list?cid={{.Cid}}">Problem</a>
          {{end}}
          <br/>
          {{if .IsContestStatus}}<span>Status</sapn>{{else}}<a href="/contest/status/list?cid={{.Cid}}">Status</a>{{end}}
          {{if .IsContestRanklist}}<span>Ranklist</sapn>{{else}}<a href="/contest/ranklist?cid={{.Cid}}">Ranklist</a>{{end}}
      </div>
      {{end}}
      {{if .IsExerciseDetail}}
      <div class="pinned note" >
          <div class="icon icon-material-add-circle" style="float:right"></div>
          {{if .IsExerciseProblem}}<span>Problem</sapn>{{else}}<a href="/contest/problem/list?cid={{.Cid}}">Problem</a>{{end}}
          {{if .IsExerciseStatus}}<span>Status</sapn>{{else}}<a href="/contest/status/list?cid={{.Cid}}">Status</a>{{end}}
          {{if .IsExerciseRanklist}}<span>Ranklist</sapn>{{else}}<a href="/contest/ranklist?cid={{.Cid}}">Ranklist</a>{{end}}
      </div>
      {{end}}
      {{if .IsCurrentUser}}
        {{if .IsSettings}}
        <div class="pinned note" >
        <div class="icon icon-material-add-circle" style="float:right"></div>
          {{if .IsSettingsDetail}}<span>Detail</span>{{else}}<a href="/user/detail?uid={{.CurrentUser}}">Detail</a>{{end}}

          <br/>
          {{if .IsSettingsEdit}}<span>Edit Info</span>{{else}}<a href="/user/edit">Edit Info</a>{{end}}
          <br/>
          {{if .IsSettingsPassword}}<span>Password</span>{{else}}<a href="/user/pagepassword">Password</a>{{end}}
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
          <center><div class="center">ZJGSU Online Judge Version 14.10.12 @ <a href="https://github.com/ZJGSU-Open-Source/GoOnlineJudge" target="_blank">Github</a></div></center>
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
      <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="/static/js/jquery.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>
  </body>
</html>
 <!-- <div class="bs-docs-section">
                
                <div class="row">
                    <div class="col-lg-6">
                        <div class="well bs-component">
                            <form class="form-horizontal">
                                <fieldset>
                                    <legend>Legend</legend>
                                    <div class="form-group">
                                        <label for="inputEmail" class="col-lg-2 control-label">Email</label>
                                        <div class="col-lg-10">
                                            <input type="email" class="form-control" id="inputEmail" placeholder="Email">
                                        </div>
                                    </div>
                                    <div class="form-group">
                                        <label for="inputPassword" class="col-lg-2 control-label">Password</label>
                                        <div class="col-lg-10">
                                            <input type="password" class="form-control" id="inputPassword" placeholder="Password">
                                            <div class="checkbox">
                                                <label>
                                                    <input type="checkbox"> Checkbox
                                                </label>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="form-group">
                                        <label for="inputFile" class="col-lg-2 control-label">File</label>
                                        <div class="col-lg-10">
                                            <input type="text" readonly class="form-control floating-label" placeholder="Browse...">
                                            <input type="file" id="inputFile" multiple>
                                        </div>
                                    </div>
                                    <div class="form-group">
                                        <label for="textArea" class="col-lg-2 control-label">Textarea</label>
                                        <div class="col-lg-10">
                                            <textarea class="form-control" rows="3" id="textArea"></textarea>
                                            <span class="help-block">A longer block of help text that breaks onto a new line and may extend beyond one line.</span>
                                        </div>
                                    </div>
                                    <div class="form-group">
                                        <label class="col-lg-2 control-label">Radios</label>
                                        <div class="col-lg-10">
                                            <div class="radio radio-primary">
                                                <label>
                                                    <input type="radio" name="optionsRadios" id="optionsRadios1" value="option1" checked="">
                                                    Option one is this
                                                </label>
                                            </div>
                                            <div class="radio radio-primary">
                                                <label>
                                                    <input type="radio" name="optionsRadios" id="optionsRadios2" value="option2">
                                                    Option two can be something else
                                                </label>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="form-group">
                                        <label for="select" class="col-lg-2 control-label">Selects</label>
                                        <div class="col-lg-10">
                                            <select class="form-control" id="select">
                                                <option>1</option>
                                                <option>2</option>
                                                <option>3</option>
                                                <option>4</option>
                                                <option>5</option>
                                            </select>
                                            <br>
                                            <select multiple="" class="form-control">
                                                <option>1</option>
                                                <option>2</option>
                                                <option>3</option>
                                                <option>4</option>
                                                <option>5</option>
                                            </select>
                                        </div>
                                    </div>
                                    <div class="form-group">
                                        <div class="col-lg-10 col-lg-offset-2">
                                            <button class="btn btn-default">Cancel</button>
                                            <button type="submit" class="btn btn-primary">Submit</button>
                                        </div>
                                    </div>
                                </fieldset>
                            </form>
                        </div>
                    </div>
                    
            </div>

           </div>


         <script src="/static/js/bootstrap.min.js"></script>

       
         <script src="/static/material/js/material.min.js"></script> -->
