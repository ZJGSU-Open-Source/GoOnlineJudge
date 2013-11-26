// Menu
var closeMenus = function() {
    $('.navigation.active').removeClass('active');
};

$('.navigation').click(function(e) {
	if (e.target.tagName.toLowerCase() == 'input')
		return false;
    if ($(this).hasClass('active')) {
      	closeMenus();
    } else {
      	closeMenus();
      	$(this).addClass('active');
    }
});

// Login

$("#login").submit(function(e) {
    e.preventDefault();
    $.post('/user/login', $(e.target).serialize(), function(json) {
        if (json.Ok) {
        } else {
            alert('Failed!');
        }
    });
});