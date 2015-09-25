{{define "content"}}

{{with .Links}}
{{range .}}
<div class="row link-area">
    <div class="vote">
      <p class="news">
        <a href="#"><span class="hint icon-material-arrow-drop-up"></span></a>
        <br>
        40
        <br>
        <a href="#"> <span class="hint icon-material-arrow-drop-down"></span></a>
      </p>
    </div>
    <div class="link">
      <p class="news">
        <a href="{{.Link}}"><span class="href">{{.Title}}</span> </a>
        &nbsp;<a href="{{ShowHostUrl .Link}}"><span class="hint">({{ShowHost .Link}})</span></a>
        <br>
        &nbsp;<span class="hint">submitted at {{ShowTime .Create}} by</span><a href="/users/{{.Uid}}"><span class="hint"> {{.Uid}}</span></a>
        <br>
        &nbsp;<a href="/links/{{.Lid}}"><span class="link-comment">0 Comments</span></a>
      </p>
    </div>
</div>
<br>
{{end}}
{{end}}
<hr>

<form action="/links" method="POST">
  <fieldset>
    <legend>Share a new link</legend>
     <label for="title">Tiltle</label>
     <br>
     <textarea id="title" name="title" required style="width:50%;height:0%"></textarea>
    <br>
    <label for="link">Link</label>
    <br>
    <input type="url" id="link" name="link" required style="width:50%;"/>
  </fieldset>
  <div class="actions">
      <input name="submit" type="submit" value="Submit">
  </div>
</form>

<style type="text/css">
.link-area{
  /*border:1px solid #888;*/
}
.href{
  color:#428bca;
  font-size: 18px;
}
.hint{
  color: #888;
}
.vote{
  width:4%;
  padding-top: 0.5%;
  margin-left: 10px;
  float:left;
}
.link{
  float:left;
  width:80%;
}
.link-comment{
  color: #888;
  font-weight: bold;
}
</style>
{{end}}
