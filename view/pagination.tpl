{{define "pagination"}}
<div class="m-pagination mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
  {{$current := .CurrentPage}}
  {{if .IsPreviousPage}}
    <a class="mdl-button mdl-js-button mdl-button--icon mdl-layout--large-screen-only" href="?page={{NumSub .CurrentPage 1}}">
      <i class="material-icons">arrow_back</i>
    </a>
  {{else}}
    <button class="mdl-button mdl-js-button mdl-button--icon mdl-layout--large-screen-only">
      <i class="material-icons">arrow_back</i>
    </button>
  {{end}}

  {{if .IsPageHead}}
    {{with .PageHeadList}}
      {{range .}}
        {{if eq . $current}}
          <button class="btn now mdl-button mdl-js-button mdl-js-ripple-effect">
            {{.}}
          </button>
        {{else}}
          <a class="btn mdl-button mdl-js-button mdl-js-ripple-effect" href="?page={{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}

  {{if .IsPageMid}}
    <button class="mdl-button mdl-js-button mdl-button--icon">
      <i class="material-icons">more_horiz</i>
    </button>
    {{with .PageMidList}}
      {{range .}}
        {{if eq . $current}}
          <div class="btn now mdl-button mdl-js-button mdl-js-ripple-effect">{{.}}</div>
        {{else}}
          <a class="btn mdl-button mdl-js-button mdl-js-ripple-effect" href="?page={{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}

  {{if .IsPageTail}}
    <button class="mdl-button mdl-js-button mdl-button--icon">
      <i class="material-icons">more_horiz</i>
    </button>
    {{with .PageTailList}}
      {{range .}}
        {{if eq . $current}}
          <button class="btn now mdl-button mdl-js-button mdl-js-ripple-effect">{{.}}</button>
        {{else}}
          <a class="btn mdl-button mdl-js-button mdl-js-ripple-effect" href="?page={{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}

  {{if .IsNextPage}}
    <a class="mdl-button mdl-js-button mdl-button--icon mdl-layout--large-screen-only" href="?page={{NumAdd .CurrentPage 1}}">
      <i class="material-icons">arrow_forward</i>
    </a>
  {{else}}
    <button class="mdl-button mdl-js-button mdl-button--icon mdl-layout--large-screen-only">
      <i class="material-icons">arrow_forward</i>
    </button>
  {{end}}
</div>
{{end}}