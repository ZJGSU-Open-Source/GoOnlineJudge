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