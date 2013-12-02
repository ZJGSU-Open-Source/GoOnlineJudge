$(document).ready(function() {
    setMenu();  
});

// Cookie

var setMenu = function() {
    var uid = $.cookie('uid');
    if (uid == '' || uid == null) {
        prepareToLogin();
    } else {
        prepareToLogout(uid);
    }
};

var prepareToLogin = function() {
    $('#login').html('<div class="button">Sign In</div>\
        <div class="contents menu"><form id="loginForm">\
            <input type="text" name="uid" placeholder="User ID">\
            <input type="password" name="pwd" placeholder="Password">\
            <button class="minibutton ok" type="submit">Sign In</button>\
            <button class="minibutton" type="button">Sign Up</button>\
        </form></div>');

    $('#loginForm').submit(function(e) {
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

var prepareToLogout = function(uid) {
    $('#login').html('<div class="button">Hi, ' + uid + '</div>\
        <div class="contents menu"><form id="logoutForm">\
            <button class="minibutton ok" type="submit">Sign Out</button>\
        </form></div>');

    $('#logoutForm').submit(function(e) {
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

