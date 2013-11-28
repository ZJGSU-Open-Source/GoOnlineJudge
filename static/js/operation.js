$(document).ready(function() {
    setMenu();  
});

// Cookie

var setMenu = function() {
    var uid = $.cookie('uid');
    if (uid == '' || uid == null) {
        setAlreadyLogout();
    } else {
        setAlreadyLogin(uid);
    }
};

var setAlreadyLogout = function() {
    $('#loginLabel').html('Sign In');
    $('#loginContent').html('<form id="login">\
        <input type="text" name="uid" placeholder="User ID">\
        <input type="password" name="pwd" placeholder="Password">\
        <button class="minibutton ok" type="submit">Sign In</button>\
        <button class="minibutton" type="button">Sign Up</button>\
        </form>');

    $('#login').submit(function(e) {
        e.preventDefault();
        var target = e.target;
        var action = '/user/login';
        $.post(action, $(target).serialize(), function(json) {
            if (json.Ok) {
                setMenu();
            } else {
                alert('Failed!');
            }
        });
    });
};

var setAlreadyLogin = function(uid) {
    $('#loginLabel').html('Hi, ' + uid);
    $('#loginContent').html('<form id="logout">\
        <button class="minibutton ok" type="submit">Sign Out</button>\
        </form>');

    $('#logout').submit(function(e) {
        e.preventDefault();
        var target = e.target;
        var action = '/user/logout';
        $.post(action, $(target).serialize(), function(json) {
            if (json.Ok) {
                setMenu();
            } else {
                alert('Failed!');
            }
        });
    });
};

// Menu

var closeMenu = function() {
    $('.navigation.active').removeClass('active');
};

$('.navigation').click(function(e) {
    if ($(this).hasClass('home')) {
        window.location.href = "/home";
        return false;
    }
    if (e.target.tagName.toLowerCase() == 'input') {
		return false;
    }
    if ($(this).hasClass('active')) {
      	closeMenu();
    } else {
      	closeMenu();
      	$(this).addClass('active');
    }
});

