{{define "content"}}
{{with .Solution}}
<script src="/static/js/prettify.js"></script>
<div class="p-code-detail mdl-grid"> 
  <div class="mdl-cell mdl-cell--1-col mdl-cell--hide-phone mdl-cell--hide-tablet"></div>
  <div class="page mdl-cell mdl-cell--10-col mdl-cell--4-col-phone mdl-shadow--2dp mdl-grid">
    <div class="go-title-area border text-center mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
      <div class="title">
        View Code of Problem 
        <a href="/problems/{{.Pid}}">{{.Pid}}</a>
      </div>
    </div>
    <div class="small mdl-cell mdl-cell--12-col mdl-cell--4-col-phone mdl-shadow--2dp">{{.Code}}</div>
    <div class="large mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
      Double click to view unformatted code
      <textarea class="source J_source mdl-shadow--2dp" readonly>{{.Code}}</textarea>
      <div class="J_code">
        <pre class="prettyprint linenums code mdl-shadow--2dp">{{.Code}}</pre>
      </div>
    </div>
  </div>
</div>
{{end}}
{{end}}
