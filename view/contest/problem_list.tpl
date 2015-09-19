{{define "content"}}
<div class="p-conProList mdl-grid">
  
  <div class="mdl-cell mdl-cell--1-col mdl-cel--1-col-tablet mdl-cell--4-col-phone">
    <div class="m-link J_static mdl-shadow--2dp">
      <div class="link current">
        <a>Problem</a>
      </div>
      <div class="link">
        <a href="/contests/{{.Cid}}/status">Status</a>
      </div>
      <div class="link">
        <a href="/contests/{{.Cid}}/ranklist">Ranklist</a>
      </div>
    </div>
  </div>


  <div class="page mdl-cell mdl-cell--10-col mdl-cell--6-col-tablet mdl-cell--4-col-phone mdl-shadow--2dp J_list">
    <div class="go-title-area mdl-cell--12-col mdl-cell--4-col-phone mdl-grid">
    
      <div class="title mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">{{.Contest}}</div>
      <div class="time mdl-cell mdl-cell--12-col">
        Start Time : <span class="J_start static-1">{{ShowTime .Start}}</span>
        <span class="blank"></span>
        End Time : <span class="J_end static-4"> {{ShowTime .End}}</span>
      </div>
      <div class="text-center mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        Current Time : <span class="J_time"></span>
      </div>
      <div class="text-center mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
        <input class="J_process mdl-slider mdl-js-slider" type="range" min="0" max="100" readonly/>      
      </div>
      
    </div>
    
    <div class="table-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone mdl-shadow--2dp">
      <table class="go-table text-center">
        <thead>
          <tr>
            <th>ID</th>
            <th>Title</th>
            <th>Ratio(Accept/Submit)</th>
          </tr>
        </thead>
        <tbody>
          {{$cid := .Cid}}
          {{with .Problem}}  
            {{range .}}
            {{if .}} 
              <tr>
                <td>{{.Pid}}</td>
                <td><a href="/contests/{{$cid}}/problems/{{.Pid}}">{{.Title}}</a></td>
                <td>{{ShowRatio .Solve .Submit}} ({{.Solve}}/{{.Submit}})</td>
              </tr>
            {{end}}
            {{end}}  
          {{end}}
        </tbody>
      </table>
    </div>

  </div>
  <div class="mdl-cell mdl-cell--1-col mdl-cel--1-col-tablet mdl-cell--4-col-phone"></div>
</div>
{{end}}