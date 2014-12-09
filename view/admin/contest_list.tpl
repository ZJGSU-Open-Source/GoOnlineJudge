{{define "content"}}
<h1>Admin - Contest List</h1>
<table id="contest_list">
  <thead>
    <tr>
      <th class="header">ID</th>
      <th class="header">Title</th>
      <th class="header">Creator</th>
      <th class="header">Status</th>
      <th class="header">Delete</th>
      <th class="header">Edit</th>
    </tr>
  </thead>
  <tbody>
    {{with .Contest}}  
      {{range .}} 
        <tr>
          <td>{{.Cid}}</td>
          <td><a href="/contests/{{.Cid}}">{{.Title}}</a></td>
          <td><a href="/users/{{.Creator}}">{{.Creator}}</a></td>
          <td><a class="contest_status" href="#" data-id="{{.Cid}}">[{{if ShowStatus .Status}}Available{{else}}Reserved{{end}}]</a></td>
          <td><a class="contest_delete" href="#" data-id="{{.Cid}}">[Delete]</a></td>
          <td><a class="contest_edit" href="#" data-id="{{.Cid}}">[Edit]</a></td>
        </tr>
      {{end}}  
    {{end}}
  </tbody>
</table>
<script type="text/javascript">
$('.contest_status').on('click', function() {
  var cid = $(this).data('id');
  $.ajax({
    type:'POST',
    url:'/admin/contests/'+cid+'/status',
    data:$(this).serialize(),
    error: function(){
      alert('failed!');
    },
    success: function(){
      window.location.reload();
    }
  });
});
$('.contest_delete').on('click', function() {
  var ret = confirm('Delete the contest?');
  if (ret == true) {
    var cid = $(this).data('id');
    $.ajax({
      type:'DELETE',
      url:'/admin/contests/'+cid,
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
$('.contest_edit').on('click', function() {
  var cid = $(this).data('id');
  window.location.href = '/admin/contests/'+cid;
});
</script>
{{end}}
