{{define "content"}}
<h1>Admin - Problem List</h1>
<table id="problem_list">
  <thead>
    <tr>
      <th class="header">ID</th>
      <th class="header">Title</th>
      <th class="header">Status</th>
      <th class="header">Delete</th>
      <th class="header">Edit</th>
      <th class="header">Data</th>
    </tr>
  </thead>
  <tbody>
    {{with .Problem}}  
      {{range .}} 
        <tr>
          <td>{{.Pid}}</td>
          <td><a href="/admin/problem/detail/pid/{{.Pid}}">{{.Title}}</a></td>
          <td><a class="problem_status" href="#" data-id="{{.Pid}}">[{{if ShowStatus .Status}}Available{{else}}Reserved{{end}}]</a></td>
          <td><a class="problem_delete" href="#" data-id="{{.Pid}}">[Delete]</a></td>
          <td><a class="problem_edit" href="#" data-id="{{.Pid}}">[Edit]</a></td>
          <td><a class="test_data" href="/admin/testdata/list/pid/{{.Pid}}">[Test Data]</a></td>
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
    url:'/admin/problem/status/pid/'+pid,
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
      type:'POST',
      url:'/admin/problem/delete/pid/'+pid,
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
  window.location.href = '/admin/problem/edit/pid/'+pid;
});
</script>
{{end}}
