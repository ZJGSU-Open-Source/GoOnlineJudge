{{define "content"}}
<div class="p-rankList mdl-grid">
  <div class="mdl-cell mdl-cell--1-col mdl-cell--hide-phone mdl-cell--hide-tablet"></div>
  <div class="page mdl-cell mdl-cell--10-col mdl-cell--4-col-phone mdl-shadow--2dp">
    {{template "pagination" .}}

    <div class="table-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone mdl-shadow--2dp">
      <table class="go-table text-center">
        <thead>
          <tr>
            <th>Rank</th>
            <th>User</th>
            <th class="mdl-layout--large-screen-only">Nick</th>
            <th class="mdl-layout--large-screen-only">Motto</th>
            <th>
              <span class="mdl-layout--large-screen-only">Ratio</span>
              (Solve / Submit)
            </th>
          </tr>
        </thead>
        <tbody>
          {{$time := .Time}}
          {{$privilege := .Privilege}}
          {{with .User}}  
            {{range .}} 
              {{if ShowStatus .Status}}
                <tr>
                  <td>{{.Index}}</td>
                  <td><a href="/users/{{.Uid}}">{{.Uid}}</a></td>
                  <td class="mdl-layout--large-screen-only">{{.Nick}}</td>
                  <td class="mdl-layout--large-screen-only">{{.Motto}}</td>
                  <td>
                    <span class="mdl-layout--large-screen-only">{{ShowRatio .Solve .Submit}}</span>
                    ( <a href="/status?uid={{.Uid}}&judge=3">{{.Solve}}</a> /
                    <a href="/status?uid={{.Uid}}">{{.Submit}}</a> )
                  </td>
                </tr>
              {{end}}
            {{end}}
          {{else}}
            <td></td>
            <td class="mdl-layout--large-screen-only"></td>
            <td>无</td>
            <td class="mdl-layout--large-screen-only"></td>
            <td></td>
          {{end}}
        </tbody>
      </table>
    </div>
    
  </div>
</div>
{{end}}