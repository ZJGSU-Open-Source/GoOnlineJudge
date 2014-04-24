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
          <th>{{.Submit}}</th>
        </tr>
        <tr>
          <th>Problems Solved</th>
          <th>{{.Solve}}</th>
        </tr>
      {{end}}
        <tr>
          <th>Achieve</th>
          <th>
            {{with .List}}
              {{range .}}
                <a href="/problem/detail/pid/{{.}}">{{.}} </a>
              {{end}}
            {{end}}
          </th>
        </tr>
    </tbody>
  </table>
  
{{end}}