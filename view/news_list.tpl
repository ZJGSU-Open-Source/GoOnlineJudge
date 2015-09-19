{{define "content"}}

<div class="p-index">

  <div class="mdl-grid">
    <div class="img-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
    </div>
  </div>

  <div class="mdl-grid">

    <div class="mdl-cell mdl-cell--3-col mdl-cell--2-col-tablet mdl-cell--4-col-phone">
      <table class="J_static static-table mdl-data-table mdl-js-data-table mdl-shadow--2dp">
        <thead>
          <tr>
            <th>OJ</th>
            <th>STATUS</th>
          </tr>
        </thead>
        <tbody>
          {{with .OJStatus}}
            {{range .}}
              <tr>
                <td>{{.Name}}</td>
                <td>
                  {{if eq .Status 0}} 
                    <span>Ok</span>
                  {{else}}
                    <span>Unavailable</span>
                  {{end}}
                </td>
              </tr>
            {{end}}
          {{end}}
        </tbody>
      </table>
    </div>

    <div class="new-list mdl-cell mdl-grid--no-spacing mdl-cell--6-col mdl-cell--4-col-tablet mdl-cell--4-col-phone J_list">
      {{with .News}}
        {{range .}}
          {{if ShowStatus .Status}}
            <section class="new-card mdl-card mdl-cell--12-col mdl-shadow--2dp">
              <div class="mdl-card__title mdl-card--expand">
                <h2 class="mdl-card__title-text">{{.Title}}</h2>
              </div>
              <div class="mdl-card__supporting-text">{{.Create}}</div>
              <div class="mdl-card__actions mdl-card--border">
                <a class="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect" href="/news/{{.Nid}}">detail</a>
              </div>
              <!-- <div class="mdl-card__menu">
                <button class="mdl-button mdl-button--icon mdl-js-button mdl-js-ripple-effect">
                  <i class="material-icons">share</i>
                </button>
              </div> -->
            </section>
          {{end}}
        {{end}}
      {{end}}
    </div>
  </div>

</div>
{{end}}
