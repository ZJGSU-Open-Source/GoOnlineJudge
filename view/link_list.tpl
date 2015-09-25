{{define "content"}}

<div class="row link-area">
    <div class="vote">
      <p class="news">
              <a href="/links/1/up"><span class="icon-material-arrow-drop-up"></span></a>
              <br>
              99
              <br>
              <a href="/links/1/down"> <span class=" icon-material-arrow-drop-down"></span></a>
      </p>
    </div>
    <div class="link">
      <p class="news">
        <a href="/status/code?sid=7365" class="href">Share code of Problem 1</a>
        &nbsp;<span class="hint">(<a href="/">acm.zjgsu.edu.cn)</a></span>
        <br>
        &nbsp;<span class="hint">submitted at 2015-09-25 12:01:40 by</span> <a href="/users/sakeven">sakeven</a>
        <br>
        &nbsp;<a href="/links/1">0 Comments</a>
      </p>
    </div>
</div>
<br>

{{with .Links}}
{{range .}}
<div class="row link-area">
    <div class="vote">
      <p class="news">
              <a href="/links/{{.Lid}}/up"><span class="icon-material-arrow-drop-up"></span></a>
              <br>
              40
              <br>
              <a href="/links/{{.Lid}}/down"> <span class=" icon-material-arrow-drop-down"></span></a>
      </p>
    </div>
    <div class="link">
      <p class="news">
        <a href="{{.Link}}" class="href">{{.Title}}</a>
        &nbsp;<span class="hint">(<a href="{{ShowHostUrl .Link}}">{{ShowHost .Link}})</a></span>
        <br>
        &nbsp;<span class="hint">submitted at {{ShowTime .Create}} by</span> <a href="/users/{{.Uid}}">{{.Uid}}</a>
        <br>
        &nbsp;<a href="/links/{{.Lid}}">0 Comments</a>
      </p>
    </div>
</div>
<br>
{{end}}
{{end}}
<style type="text/css">
.link-area{
  /*border:1px solid #888;*/
}
.href{
  color: blue;
  font-size: 18px;
}
.hint{
  color: #888;
}
.vote{
  width:4%;
  padding-top: 1%;
  margin-left: 10px;
  float:left;
}
.link{
  float:left;
  width:80%;
}
</style>
{{end}}
