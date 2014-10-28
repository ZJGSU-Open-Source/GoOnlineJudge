{{define "content"}}
  <h1>User Detail</h1> 
  <table class="table table-bordered table-striped table-hover">
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
          <td><a href="mailto:{{.Mail}}">{{.Mail}}</a></td>
        </tr>
        <tr>
          <th>Motto</th>
          <td>{{.Motto}}</td>
        </tr>
        <tr>
          <th>Problems Submitted</th>
          <td><a href="/status/list?uid={{.Uid}}">{{.Submit}}</a></td>
        </tr>
        <tr>
          <th>Problems Solved</th>
          <td>{{.Solve}}</td>
        </tr>
      {{end}}
      <table class="table table-bordered table-striped table-hover">
        <tr>
          <th><center>Login IP</center></th>
          <th><center>Login Time</center></th>
        </tr>
          <td>
            {{with .IpList}}
              {{range .}}
              {{if .}}
                <li>{{.}}</li>
              {{end}}
              {{end}}
            {{end}}
          </td>
          <td>
          <!--login time模块-->
          </td>
       </table>
      <table class="table table-bordered table-striped table-hover">
        <tr>
          <th><center>Achieve</center></th>
        </tr>
        <tr>
          <td>
            {{with .List}}
              {{range .}}
                <a href="/problem/detail?pid={{.}}">{{.}}</a> 
              {{end}}
            {{end}}
          </td>
        </tr>
      </table>
        
    </tbody>
  </table>
  
{{end}}
