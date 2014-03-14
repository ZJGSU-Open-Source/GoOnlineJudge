{{define "content"}}
  <h1>Ranklist</h1>
  <div class="pagination">
    <span class="previous_page disabled">上一页</span> <em class="current">1</em> <a class="next_page" rel="next">下一页</a>
  </div>
  <table id="ranklist">
    <thead>
      <tr>
        <th class="header">Rank</th>
        <th class="header">User</th>
        <th class="header">Motto</th>
        <th class="header">Ratio</th>
      </tr>
    </thead>
    <tbody>
      {{$rank := 1}}
      {{with .User}}  
        {{range .}}
          {{if ShowStatus .Status}}
            <tr>
              <td>{{$rank}}</td>
              <td>{{.Uid}}</td>
              <td>{{.Motto}}</td>
              <td>{{ShowRatio .Solve .Submit}} ({{.Solve}}/{{.Submit}})</td>
            </tr>
          {{end}}
        {{end}}  
      {{end}}
    </tbody>
  </table>
{{end}}