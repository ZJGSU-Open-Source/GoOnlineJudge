{{define "content"}}
<div class="p-proList mdl-grid">
  <div class="mdl-cell mdl-cell--1-col mdl-cell--hide-phone mdl-cell--hide-tablet"></div>
  <div class="page mdl-cell mdl-cell--10-col mdl-cell--4-col-phone mdl-shadow--2dp">

    <form accept-charset="UTF-8" class="J_searchForm">
      <div class="mdl-grid">
        <div class="mdl-cell mdl-cell--2-col mdl-cell--2-col-phone">
          <div class="go-select-title">search type</div>
          <select name="option" class="go-select">
            <option value="title" {{if .SearchTitle}}selected{{end}}>Title</option>
            <option value="pid" {{if .SearchPid}}selected{{end}}>ID</option>
            <option value="source" {{if .SearchSource}}selected{{end}}>Source</option>
            <option value="page">Page</option>
          </select>
        </div>
        <!-- <div class="search-text">search:</div> -->
        <div class="mdl-cell mdl-cell--2-col mdl-cell--2-col-phone">
          <div class="mdl-textfield mdl-js-textfield mdl-textfield--expandable" >
            <label class="mdl-button mdl-js-button mdl-button--icon" for="sample6">
              <i class="material-icons">search</i>
            </label>
            <div class="mdl-textfield__expandable-holder">
              <input class="mdl-textfield__input" type="text" id="sample6" name="search" value="{{.SearchValue}}"/>
              <label class="mdl-textfield__label" for="sample6">Expandable Input</label>
            </div>
          </div>
        </div>
      </div>
    </form>

    {{template "pagination" .}}
    
    <div class="table-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone mdl-shadow--2dp">
      <table class="go-table text-center">
        <thead>
          <tr>
            <!-- <th>flag</th> -->
            <th>ID</th>
            <th>Title</th>
            <th class="mdl-layout--large-screen-only">Ratio(Solve/Submit)</th>
            <th class="mdl-layout--large-screen-only">OJ</th>
            <th class="mdl-layout--large-screen-only">VPID</th>
          </tr>
        </thead>
        <tbody>
          {{$time := .Time}}
          {{$privilege := .Privilege}}
          {{with .Problem}}  
            {{range .}} 
              {{if or (ShowStatus .Status) (LargePU $privilege)}}
                <tr>
                  <!-- <td></td> -->
                  <td>{{.Pid}}</td>
                  <td><a href="/problems/{{.Pid}}">{{.Title}}</a></td>
                  <td class="mdl-layout--large-screen-only">
                    {{ShowRatio .Solve .Submit}} (
                    <a href="/status?pid={{.Pid}}&judge=3">{{.Solve}}</a> /
                    <a href="/status?pid={{.Pid}}">{{.Submit}}</a> )
                  </td>
                  <td class="mdl-layout--large-screen-only">{{.ROJ}}</td>
                  <td class="mdl-layout--large-screen-only">{{.RPid}}</td>
                </tr>
              {{end}}
            {{end}}
          {{else}}  
            <!-- <td></td> -->
            <td></td>
            <td class="mdl-layout--large-screen-only"></td>
            <td>无</td>
            <td class="mdl-layout--large-screen-only"></td>
            <td class="mdl-layout--large-screen-only"></td>
          {{end}}
        </tbody>
      </table>
    </div>

  </div>
</div>
{{end}}
