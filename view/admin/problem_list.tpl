{{define "content"}}
<h1>Admin - Problem List</h1>
{{$isAdmin := .IsAdmin}}

<div class="pagination">
  {{$current := .CurrentPage}}
  {{if .IsPreviousPage}}
  <a href="?page={{NumSub .CurrentPage 1}}">Prev</a>
  {{else}}
  <span>Prev</span>
  {{end}}
  &nbsp;
  {{if .IsPageHead}}
    {{with .PageHeadList}}
      {{range .}}
        {{if eq . $current}}
          <span>{{.}}</span>
        {{else}}
          <a href="?page={{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}

  {{if .IsPageMid}}
  ...
    {{with .PageMidList}}
      {{range .}}
        {{if eq . $current}}
          <span>{{.}}</span>
        {{else}}
          <a href="?page={{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}

  {{if .IsPageTail}}
  ...
    {{with .PageTailList}}
      {{range .}}
        {{if eq . $current}}
          <span>{{.}}</span>
        {{else}}
          <a href="?page={{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}
  &nbsp;
  {{if .IsNextPage}}
  <a href="?page={{NumAdd .CurrentPage 1}}">Next</a>
  {{else}}
  <span>Next</span>
  {{end}}
</div>

<table id="problem_list">
  <thead>
    <tr>
      <th class="header">ID</th>
      <th class="header">Title</th>
      {{if $isAdmin}}
      <th class="header">Status</th>
      <th class="header">Delete</th>
      <th class="header">Edit</th>
      {{end}}
      <th class="header">Data</th>
    </tr>
  </thead>
  <tbody>
    {{with .Problem}}  
      {{range .}} 
        <tr>
          <td>{{.Pid}}</td>
          <td><a href="/problems/{{.Pid}}">{{.Title}}</a></td>
          {{if $isAdmin}}
          <td><a class="problem_status" href="#" data-id="{{.Pid}}">[{{if ShowStatus .Status}}Available{{else}}Reserved{{end}}]</a></td>
          <td><a class="problem_delete" href="#" data-id="{{.Pid}}">[Delete]</a></td>
          <td><a class="problem_edit" href="#" data-id="{{.Pid}}">[Edit]</a></td>
          {{end}}
          <td><a class="test_data" href="/admin/testdata/{{.Pid}}">[Test Data]</a></td>
        </tr>
      {{end}}  
    {{end}}
  </tbody>
</table>
<script type="text/javascript">
$('.problem_status').on('click', function() {
  var pid = $(this).data('id');
  $.ajax({
    type:'POST',
    url:'/admin/problems/'+pid+'/status',
    data:$(this).serialize(),
    error: function(){
      alert('failed!');
    },
    success: function(){
      window.location.reload();
    }
  });
});
$('.problem_delete').on('click', function() {
  var ret = confirm('Delete the Problem?');
  if (ret == true) {
    var pid = $(this).data('id');
    $.ajax({
      type:'DELETE',
      url:'/admin/problems/'+pid,
      data:$(this).serialize(),
      error: function() {
        alert('failed!');
      },
      success: function() {
        window.location.reload();
      }
    });
  }
});
$('.problem_edit').on('click', function() {
  var pid = $(this).data('id');
  window.location.href = '/admin/problems/'+pid;
});
</script>
{{end}}
