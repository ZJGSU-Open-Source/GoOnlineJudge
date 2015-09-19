{{define "content"}}
	{{with .Detail}}
	<div class="p-new-detail mdl-grid">	
		<div class="page mdl-cell mdl-cell--8-col mdl-cell--4-col-phone mdl-shadow--2dp mdl-grid mdl-grid--no-spacing">
			<div class="mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
				<div class="go-title-area border text-center mdl-cell mdl-cell--12-col mdl-cell--4-col-phone">
					<div class="title">{{.Title}}</div>
					<div><b>Date: </b>{{.Create}}</div>
				</div>
				<section class="text">{{.Content}}</section>
			</div>
		</div>
	</div>
	{{end}}
{{end}}