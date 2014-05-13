{{define "content"}}
  <h1>User Detail</h1> 
  <table>
    <tbody>
      {{with .Detail}}
        <tr>
          <th>Handle</th>
          <td>{{.Uid}}</td>
        </tr>
        <tr>
          <th>Nick</th>
          <td>{{.Nick}}</td>
        </tr>
        <tr>
          <th>Email</th>
          <td>{{.Mail}}</td>
        </tr>
        <tr>
          <th>Motto</th>
          <td>{{.Motto}}</td>
        </tr>
        <tr>
          <th>Problems Submitted</th>
          <td>{{.Submit}}</td>
        </tr>
        <tr>
          <th>Problems Solved</th>
          <td>{{.Solve}}</td>
        </tr>
      {{end}}
        <tr>
          <th>Achieve</th>
          <td>
            {{with .List}}
              {{range .}}
                <a href="/problem/detail/pid/{{.}}">{{.}}</a> 
              {{end}}
            {{end}}
          </td>
        </tr>
    </tbody>
  </table>
  
{{end}}