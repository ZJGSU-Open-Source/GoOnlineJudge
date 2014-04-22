var ConfirmDelete = function (url, msg) {
	ret = window.confirm(msg);
	if (ret == true) {
		//alert(url);
		//window.location.href = url;
		//indow.location.href = "127.0.0.1:8888/admin/news/delete/nid/9";
		//window.location.assign(url);
		//window.open(url);
		$.ajax({
			url: url,
			type: 'POST',
			data: {},
			error: function () {
				alert('error');
			},
			success: function () {
				alert('success!');
			}
		});
	}
}