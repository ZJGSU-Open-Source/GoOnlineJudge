{{define "content"}}
<div class="p-conList mdl-grid">
  <div class="mdl-cell mdl-cell--1-col mdl-cell--hide-phone mdl-cell--hide-tablet"></div>
  <div class="page mdl-cell mdl-cell--10-col mdl-cell--4-col-phone mdl-shadow--2dp">
  
    <div class="table-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone mdl-shadow--2dp">
      <table class="go-table text-center">
        <thead>
          <tr>
            <th>ID</th>
            <th>Title</th>
            <th>Status</th>
            <th>Type</th>
          </tr>
        </thead>
        <tbody>
          {{$time := .Time}}
          {{$privilege := .Privilege}}
          {{with .Contest}}  
            {{range .}} 
              {{if or (ShowStatus .Status) (LargePU $privilege)}}
                <tr>
                  <td>{{.Cid}}</td>
                  <td><a href="/contests/{{.Cid}}">{{.Title}}</a></td>
                  <td>
                    {{if ge $time .End }}
                      <span class="static-4">Ended@{{ShowTime .End}}</span>
                    {{else}}
                      {{if ge .Start $time}}
                        <span class="static-1">Start@{{ShowTime .Start}}</span>
                      {{else}}
                        <span class="static-3">Running</span>
                      {{end}}
                    {{end}}
                  </td>
                  <td>{{ShowEncrypt .Encrypt}}</td>
                </tr>
              {{end}}
            {{end}}
          {{else}}
            <td></td>
            <td></td>
            <td>无</td>
            <td></td>
          {{end}}
        </tbody>
      </table>
    </div>

  </div>
</div>
{{end}}