{{define "content"}}
<div class="p-user-detail mdl-grid">
  <div class="mdl-cell mdl-cell--2-col mdl-cell--4-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link current">
        <a>Detail</a>
      </div>
      <div class="link">
        <a href="/profile">Edit Info</a>
      </div>
      <div class="link">
        <a href="/account">Password</a>
      </div>
    </div>
  </div>
  <div class="page mdl-cell mdl-cell--8-col mdl-shadow--2dp mdl-grid J_list">

    <div class="go-title-area border mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
      <div class="title">User Detail</div>
    </div>

    <div class="mdl-cell mdl-cell--6-col mdl-cell--4-col-tablet mdl-cell--2-col-phone right">
      <div class="tip">Handle</div>
      <div class="tip">Nick</div>
      <div class="tip">Email</div>
      <div class="tip">Motto</div>
      <div class="tip">School</div>
      <div class="tip">Problems Submitted</div>
      <div class="tip">Problems Solved</div>

      <div class="tip margin-top">Login Time</div>
      {{with .IpList}}
        {{range .}}
        {{if .}}
          <div class="text">{{ShowTime .Time}}</div>
        {{end}}
        {{end}}
      {{end}}
    </div>

    
    <div class="mdl-cell mdl-cell--6-col mdl-cell--4-col-tablet mdl-cell--2-col-phone">
      {{with .Detail}}
      <div class="text">{{.Uid}}</div>
      <div class="text">{{.Nick}}</div>
      <div class="text">{{.Mail}}</div>
      <div class="text">{{.Motto}}</div>
      <div class="text">{{.School}}</div>
      <div class="text">{{.Submit}}</div>
      <div class="text">{{.Solve}}</div>
      {{end}}

      <div class="tip margin-top">Login IP</div>
      {{with .IpList}}
        {{range .}}
        {{if .}}
          <div class="text">{{.IP}}</div>
        {{end}}
        {{end}}
      {{end}}
    </div>
    <div class="mdl-cell mdl-cell--12-col center">
      <div class="tip">Achieve</div>
      {{with .List}}
      {{range .}}
        <a class="btn mdl-button mdl-js-button mdl-js-ripple-effect button" href="/problems/{{.}}">
        {{.}}
        </a>
      {{end}}
      {{end}}
    </div>
    
  </div>
</div>
{{end}}
