//全局函数与变量
var PATH = location.pathname;
console.log(PATH);
//登出
var signoutBtn = $('.J_signout');
signoutBtn.on('click', function(e) {
  e.preventDefault();
  $.ajax({
    type: 'DELETE',
    url: '/sess',
    data: '',
    error: function(err) {
      console.log(err);
      alert('登出失败');
    },
    success: function(data) {
      console.log(data);
      location.href = '/sess';
    }
  });
});

//警告
var warnNode = $('.J_warn');
var warning = function(text, callback){
  warnNode.text(text).show();
  setTimeout(function(){
    warnNode.fadeOut();
    if( callback )
      callback();
  }, 1000);
}

var winNode = $(window);
var mainNode = $('.J_main');
var staticNode =$('.J_static');
var listNode = $('.J_list');
if( staticNode.length ){
  //固定VOJ框体
  var initTop = staticNode.offset().top;
  var keep = function(){
    var staticTop = staticNode.offset().top;
    var listTop = listNode.offset().top;
    if( listTop >= 70 ){
      staticNode.css('top', 0);
    }
    else{
      staticNode.css('top', mainNode.scrollTop()-initTop+80);
    }
  };
  //若设备宽度大于840挂载滚动事件
  $(function(){
    if( winNode.width() >= '840' ){
      mainNode.on('scroll', keep);
    }
  });
}

//重载Date对象Format方法
Date.prototype.Format = function(fmt)   
{
  var o = {   
    "M+" : this.getMonth()+1,                 //月份   
    "d+" : this.getDate(),                    //日   
    "h+" : this.getHours(),                   //小时   
    "m+" : this.getMinutes(),                 //分   
    "s+" : this.getSeconds(),                 //秒   
    "q+" : Math.floor((this.getMonth()+3)/3), //季度   
    "S"  : this.getMilliseconds()             //毫秒   
  };   
  if(/(y+)/.test(fmt))   
    fmt=fmt.replace(RegExp.$1, (this.getFullYear()+"").substr(4 - RegExp.$1.length));   
  for(var k in o)   
    if(new RegExp("("+ k +")").test(fmt))   
  fmt = fmt.replace(RegExp.$1, (RegExp.$1.length==1) ? (o[k]) : (("00"+ o[k]).substr((""+ o[k]).length)));   
  return fmt;   
}

//kindeditor配置
var options = {
  resizeType: 0,
  height: '250px',
  width: '100%',
  langType : 'zh_CN',
  items: [
    'source', '|', 'undo', 'redo', '|', 
    'preview', 'code', 'cut', 'copy', 'paste', 'plainpaste', 'wordpaste', '|', 
    'justifyleft', 'justifycenter', 'justifyright', 'justifyfull', 
    'insertorderedlist', 'insertunorderedlist', 'subscript', 'superscript', 
    'clearhtml', '|', 'fullscreen', '/', 'formatblock', 'fontname', 'fontsize', '|', 
    'forecolor', 'hilitecolor', 'bold', 'italic', 'underline', 'strikethrough', 
    'removeformat', '|', 'image', 'table', 'hr', 
    'emoticons', 'baidumap', 'link', 'unlink', '|', 'about'
  ]
}

if ( PATH=='/admin/contests/new' || /^\/admin\/contests\/(\d+)$/.test(PATH) ) {
  console.log(1);
  //时间插件
  var startNode = $('#J_start');
  var endNode = $('#J_end');
  $(function() {
    startNode.mobiscroll().datetime({
      theme: 'material',
      lang: 'zh',
      display: 'bubble',
      minDate: new Date(),
    });
    endNode.mobiscroll().datetime({
      theme: 'material',
      lang: 'zh',
      display: 'bubble',
      minDate: new Date(),
    });
  });
  //解决chrome下label的bug
  startNode.on('blur', function() {
    if( $(this).val() && !$(this).parent().hasClass('is-dirty') ){
      $(this).parent().addClass('is-dirty');
    }
  });
  endNode.on('blur', function() {
    if( $(this).val() && !$(this).parent().hasClass('is-dirty') ){
      $(this).parent().addClass('is-dirty');
    }
  });
  //比赛类型选择
  var typeNode = $('#J_type');
  var passwdNode = $('.J_passwdArea');
  var privateNode = $('.J_privateArea');
  typeNode.on('change', function() {
    type = $(this).val();
    if (type == 'password') {
      passwdNode.fadeIn();
      privateNode.fadeOut();
    } else if (type == 'private') {
      privateNode.fadeIn();
      passwdNode.fadeOut();
    } else {
      passwdNode.fadeOut();
      privateNode.fadeOut();
    }
  });
  typeNode.trigger('change');

  //添加题目、用户
  var proInputNode = $('.J_proInput');
  var proListNode = $('#J_proList');
  var insert = function(){
    var list = proInputNode.val().match(/(\d+)/g);
      if( list == null )
        return;
      $.each(list, function(i, v) {
        if (v != '')
          proListNode.append(
            '<div class=\"item\">' +
              '<span class=\" J_pro\">' +
                v +
              '</span>' +
              '<div class=\"item-del J_del\">×</div>' +
            '</div>'
          );
      });
      proInputNode.val('');
      proInputNode.parent().removeClass('is-dirty');
  };
  proInputNode.on('keyup', function(e) {
    e.preventDefault();
    e.stopPropagation();
    if (e.keyCode == 13) {
      insert();
    }
  });
  
  //proNode下无节点时拖动插件不起作用，故选择先挂节点，载入时再去掉proNode下的节点
  $(function() {
    proListNode.empty();
    proListNode.show();
    insert();
  });
  //拖动插件
  proListNode.dragsort({
    dragSelector: "div",
    placeHolderTemplate: "<div class=\"replace\"></div>",
    dragBetween: true
  });
  //删除题目、用户
  proListNode.on('click', '.J_del', function() {
    $(this).parent().remove();
  });
  //提交表单
  $('.J_submit').on('click', function(e) {
    e.preventDefault();
    var proList = $('.J_pro').map(function(){
      return $(this).text();
    }).get().join(';');
    console.log(proList);
    var start = new Date(startNode.val());
    var end = new Date(endNode.val());
    $.ajax({
      type: 'post',
      url: '/admin/contests',
      data: {
        title: $('[name=title]').val(),
        startTimeYear: start.getFullYear(),
        startTimeMonth: parseInt(1) + start.getMonth(),
        startTimeDay: start.getDate(),
        startTimeHour: start.getHours(),
        startTimeMinute: start.getMinutes(),
        endTimeYear: end.getFullYear(),
        endTimeMonth: parseInt(1) + end.getMonth(),
        endTimeDay: end.getDate(),
        endTimeHour: end.getHours(),
        endTimeMinute: end.getMinutes(),
        problemList: proList,
        userlist: $('[name=userlist]').val(),
        type: $('[name=type]').val(),
        password: $('[name=password]').val()
      },
      success: function(data){
        warning('成功！', function(){
          location.href = '/admin/contests'
        });
      },
      error: function(err){
        warning(err);
      }
    });
  });
}
if( PATH == '/admin/contests' ){
	var statusNode = $('.J_status');
	statusNode.on('click', function() {
	  var cid = $(this).data('id');
	  var text = $.trim($(this).html());
	  var self = $(this);
		if( text == 'Available' )
			text = 'Reserved';
		else
			text = 'Available';
	  $.ajax({
	    type: 'POST',
	    url: '/admin/contests/'+cid+'/status',
	    data: $(this).serialize(),
	    error: function(err){
	    	console.log(err);
	      warning(err.responseText);
	    },
	    success: function(data){
	    	self.html(text);
	    	warning('修改成功！');
	    }
	  });
	});
	var deleteNode = $('.J_delete');
	deleteNode.on('click', function() {
		var self = $(this);
	  if ( confirm('Delete the contest?') ) {
	    var cid = $(this).data('id');
	    $.ajax({
	      type: 'DELETE',
	      url: '/admin/contests/'+cid,
	      data: $(this).serialize(),
	      error: function(err) {
	        warning(err.responseText);
	      },
	      success: function(data) {
	      	self.parent().parent().remove();
	        warning('删除成功！');
	      }
	    });
	  }
	});
}
if( /^\/admin\/news\/new$/.test(PATH) || /^\/admin\/news\/(\d+)$/.test(PATH) ){

	KindEditor.ready(function(K) {
	  setTimeout(function(){
	    $('.J_load').hide();
	    options.height = '500px';
	    window.editor = K.create('#J_content', options);
	  }, 1000);
	});
}
if( /^\/admin\/news$/.test(PATH) ){
	var statusNode = $('.J_status');
	statusNode.on('click', function() {
		var self = $(this);
		var nid = $(this).data('id');
		var text = $.trim($(this).html());
		if( text == 'Available' )
			text = 'Reserved';
		else
			text = 'Available';
		$.ajax({
			type: 'POST',
			url: '/admin/news/'+nid+'/status',
			data: $(this).serialize(),
			error: function() {
				warning('修改失败！');
			},
			success: function() {
				self.html(text);
				warning('修改成功!');
			}
		});
	});
	var deleteNode = $('.J_delete');
	deleteNode.on('click', function() {
		var self = $(this);
		if( confirm('Delete the News?') ) {
			var nid = $(this).data('id');
			$.ajax({
				type: 'DELETE',
				url: '/admin/news/'+nid,
				data: $(this).serialize(),
				error: function() {			
					warning('删除失败！');
				},
				success: function() {
					warning('删除成功！');
					self.parent().parent().remove();
				}
			});
		}
	});
}
if( '/admin/problems/new' == PATH || /^\/admin\/problems\/(\d+)$/.test(PATH) ){
	console.log(1);
	KindEditor.ready(function(K) {
	  options.height = '250px';
	  setTimeout(function(){
	  	$('.J_load').hide();
	  	window.editor = K.create('#J_description', options);
	  	window.editor = K.create('#J_input', options);
	  	window.editor = K.create('#J_output', options);
	  	window.editor = K.create('#J_hint', options);
	  }, 1000);
	});
	var fontList = {
    false: 'check_box_outline_blank',
    true: 'check_box'
  };
  var font = $('#problem_special').check;
  var labelNode = $('.J_label');
  var specialNode = $('#J_special');
  specialNode.on('change', function(){
  	font = !font;
  	labelNode.html( fontList[font] );
  });
}
if( PATH == '/admin/problems' ){
	var statusNode = $('.J_status');
	statusNode.on('click', function() {
	  var pid = $(this).data('id');
	  var text = $.trim($(this).html());
	  var self = $(this);
		if( text == 'Available' )
			text = 'Reserved';
		else
			text = 'Available';
	  $.ajax({
	    type: 'POST',
	    url: '/admin/problems/'+pid+'/status',
	    data: $(this).serialize(),
	    error: function(){
	      warning('修改失败！');
	    },
	    success: function(){
				self.html(text);
				warning('修改成功！');
	    }
	  });
	});
	var deleteNode = $('.J_delete');
	deleteNode.on('click', function() {
		var self = $(this);
	  if ( confirm('Delete the Problem?') ){
	    var pid = $(this).data('id');
	    $.ajax({
	      type: 'DELETE',
	      url: '/admin/problems/'+pid,
	      data: $(this).serialize(),
	      error: function() {
	        warning('删除失败！');
	      },
	      success: function() {
	        warning('删除成功！');
					self.parent().parent().remove();
	      }
	    });
	  }
	});
}
if( PATH == '/admin/rejudger' ){
	var typeNode = $('.J_type');
	var idNode = $('#J_id');

	idNode.on('input', function(){
		var value = idNode.val();
		idNode.val( value.replace(/[^0-9]/,'' ) );
	});

	$('.J_addForm').on('submit', function(e){
		e.preventDefault();
		if( !idNode.val() || idNode.val()=='' ){
			warning('请输入ID');
			return;
		}
		$.ajax({
			type: 'POST',
			url: '/admin/rejudger?type=' + typeNode.val() + '&id=' + idNode.val(),
			data: $(this).serialize(),
			dataType: 'text',
			success: function(data){
				warning("Rejudge Complete", function(){
					location.href = '/status';
				});
			},
			error: function(err){
				warning(JSON.parse(err.responseText).info);
			}
		});

	});

}
if( PATH == '/admin/users' ){
var formNode = $('.J_addForm');
formNode.submit( function(e) {
	e.preventDefault();
	if( $('[name=uid]').val() == '' ){
		warning('请输入用户名');
		return;
	}
	$.ajax({
		type: 'POST',
		url: '/admin/privilegeset?' + $(this).serialize(),
		data: $(this).serialize(),
		error: function(err){
			warning(JSON.parse(err.responseText).hint);
		},
		success: function(data){
			warning('添加成功');
      location.reload();
		}
	});
});
var delNode = $('.J_delete');
delNode.on('click', function() {
	var self = $(this);
	var uid = $(this).data("id");
	if ( confirm('Delete the user '+uid+'?') ) {
		$.ajax({
			type: 'POST',
			url: '/admin/privilegeset?type=PU&uid=' + uid,
			data: $(this).serialize(),
			error: function(err) {
				warning(JSON.parse(err.responseText).hint);
			},
			success: function(data) {
				warning('删除成功！');
				self.parent().parent().remove();
			}
		});
	}
});
}
if( PATH == '/admin/users/pagepassword' ){
	$('.J_addForm').submit( function(e) {
		e.preventDefault();
		var passwd1 = $('[name="user[newPassword]"]').val();
		console.log(passwd1);
		var passwd2 = $('[name="user[confirmPassword]"]').val();
		console.log(passwd2);
		if( passwd1.length < 6 ){
			warning('密码长度不得小于6位！');
			return;
		}
		else if( passwd1 != passwd2 ){
			warning('两次密码不一致！');
			return;
		}
		$.ajax({
			type: 'post',
			url: '/admin/users/password',
			data: $(this).serialize(),
			error: function(err) {
				var uid = JSON.parse(err.responseText).uid;
				if( uid || uid=='' ){
					warning('该用户不存在！');
				}
			},
			success: function(data) {
				warning('修改成功！', function(){
					location.href='/admin/users';
				});
			}
		});
	});
}
if( /\/contests\/(\d+)\/problems\/(\d+)/.test(PATH) ){
  
  var editor;
  var extendNode = $('.J_extend');
  var editNode = $('#advanced_editor');
  var codeNode = $('#code');
  var subNode = $('.J_submission');
  var labelNode = $('.J_label');
  var fontList = {
    false: 'check_box_outline_blank',
    true: 'check_box'
  };
  var font = true;

  function set_mode() {
    var compiler=$('#compiler_id option:selected').text();
    var modes=[ 
    'Javascript', 
    'Haskell', 
    'Lua', 
    'Pascal', 
    'Python', 
    'Ruby', 
    'Scheme', 
    'Smalltalk', 
    'Clojure',
    ['PHP', 'text/x-php'],
    ['C', 'text/x-csrc'],
    ['C++', 'text/x-c++src'],
    ['Java', 'text/x-java'],
    ['', 'text/plain'] ];
    for( var i in modes ){
      var n=modes[i], m=modes[i];
      if( $.isArray(n) ) { 
        m=n[1]; 
        n=n[0]; 
      }
      if( compiler.indexOf(n) >= 0 ){
        editor.setOption('mode', m.toLowerCase() );
        break;
      }
    }
  };
  function toggle_editor() {
    var cm = $('.CodeMirror');
    if( editNode.prop('checked') ) {
      cm.show();
      editor.setValue(codeNode.val());
      codeNode.hide();
    } 
    else {
      codeNode.val(editor.getValue()).show();
      cm.hide();
    };
    return true;
  }
  extendNode.on('click', function() {
    subNode.show();
    extendNode.hide();
    editor = CodeMirror.fromTextArea(document.getElementById("code"), {
      lineNumbers: true,
    });
    codeNode.on('blur', function(){
      editor.setValue( codeNode.val() );
    });
    $('#compiler_id').on('change', set_mode);
    set_mode();
    toggle_editor();
  });

  editNode.on('change', toggle_editor);

  labelNode.on('click', function(){
    font = !font;
    labelNode.html(fontList[font]);
  });
  var cid, pid;
  $('#problem_submit').on('submit', function(e) {
    e.preventDefault();
    codeNode.val(editor.getValue());
    cid = $(this).attr('data-cid');
    pid = $(this).attr('data-pid');
    $.ajax({
      type: 'POST',
      url: '/contests/' + cid + '/problems/' + pid,
      data: $(this).serialize(),
      success: function( result ) {
        codeNode.val('')
        warning('提交成功', function(){
          location.href = '/contests/' + cid + '/status';
        });
      },
      error: function(err) {
        if( err.status == 401 ){
          warning('请先登录', function(){
            location.href = '/sess';
          });
        }
        else {
          var json = eval('('+err.responseText+')');
          if( json.info != null )
            warning(json.info);
        }
      }
    });
  });
  
}
if( /^\/contests\/(\d+)$/.test(PATH) ) {
  
  var timeNode = $('.J_time');
  var startNode = $('.J_start');
  var endNode = $('.J_end');
  var proNode = $('.J_process');

  var endTime = new Date(endNode.html().split('-').join('/'));
  var startTime = new Date(startNode.html().split('-').join('/'));
  var totalTime = endTime.getTime()-startTime.getTime();
  //渲染时间
  var handleTime = function(){
    var time = new Date();

    if( time.getTime() < startTime.getTime() ){  //start 比赛还未开始
      timeNode.addClass('static-1');
      proNode.val(0);
    }
    else if( time.getTime() < endTime.getTime() ){  //running 比赛进行中
      timeNode.removeClass('static-1').addClass('static-3');
      var pastlTime = time.getTime()-startTime.getTime();
      proNode.val( 100 * pastlTime / totalTime );
    }
    else{ //end 比赛结束
      timeNode.removeClass('static-3').addClass('static-4');
      proNode.val(100);
    }
    time = time.Format('yyyy-MM-dd hh:mm:ss');
    timeNode.html(time);
  };
  $(handleTime);
  //动态渲染当前时间
  setInterval(handleTime, 1000);

}
if( /^\/contests\/(\d+)\/status$/.test(PATH) ) {
  
  $('#search_form').submit( function(e) {
    e.preventDefault();
    var uid = $('[name=search_uid]').val();
    var pid = $('[name=search_pid]').val();
    var judge = $('[name=search_judge]').val();
    var language = $('[name=search_language]').val();
    var url = 'status?';
    if (uid != '')
      url += 'uid=' + uid + "&";
    if (pid != '')
      url += 'pid=' + pid + "&";
    if (judge > 0){
      judge = judge-1;
      url += 'judge=' + judge + "&";
    }
    if (language > 0)
      url += 'language=' + language + "&";
    location.href = url;
  });
}

if( /^\/contests\/(\d+)\/status\/(\d+)\/code$/.test(PATH) ) {
	var codeNode = $('.J_code');
	var sourceNode = $('.J_source');
	var flag = true;

	codeNode.on('dblclick', function(){
		codeNode.hide();
			if( flag ){
				sourceNode.height(codeNode.height());
				falg = !flag;
			}
		sourceNode.show();
	});
	sourceNode.on('blur', function(){
		codeNode.show();
		sourceNode.hide();
	});
}
if( /^\/problems\/(\d+)$/.test(PATH) ){
  var editor;
  var extendNode = $('.J_extend');
  var editNode = $('#advanced_editor');
  var codeNode = $('#code');
  var subNode = $('.J_submission');
  var labelNode = $('.J_label');
  var fontList = {
    false: 'check_box_outline_blank',
    true: 'check_box'
  };
  var font = true;

  function set_mode() {
    var compiler=$('#compiler_id option:selected').text();
    var modes=[ 
    'Javascript', 
    'Haskell', 
    'Lua', 
    'Pascal', 
    'Python', 
    'Ruby', 
    'Scheme', 
    'Smalltalk', 
    'Clojure',
    ['PHP', 'text/x-php'],
    ['C', 'text/x-csrc'],
    ['C++', 'text/x-c++src'],
    ['Java', 'text/x-java'],
    ['', 'text/plain'] ];
    for( var i in modes ){
      var n=modes[i], m=modes[i];
      if( $.isArray(n) ) { 
        m=n[1]; 
        n=n[0]; 
      }
      if( compiler.indexOf(n) >= 0 ){
        editor.setOption('mode', m.toLowerCase() );
        break;
      }
    }
  };
  function toggle_editor() {
    var cm = $('.CodeMirror');
    if( editNode.prop('checked') ) {
      cm.show();
      editor.setValue(codeNode.val());
      codeNode.hide();
    } 
    else {
      codeNode.val(editor.getValue()).show();
      cm.hide();
    };
    return true;
  }
  extendNode.on('click', function() {
    subNode.show();
    extendNode.hide();
    editor = CodeMirror.fromTextArea(document.getElementById("code"), {
      lineNumbers: true,
    });
    codeNode.on('blur', function(){
      editor.setValue( codeNode.val() );
    });
    $('#compiler_id').on('change', set_mode);
    set_mode();
    toggle_editor();
  });

  editNode.on('change', toggle_editor);

  $('#advanced_editor').on('change', function(){
    font = !font;
    labelNode.html(fontList[font]);
  });

  $('#problem_submit').on('submit', function(e) {
    e.preventDefault();
    codeNode.val(editor.getValue());
    $.ajax({
      type: 'POST',
      url: '/problems/' + $(this).attr('data-id'),
      data: $(this).serialize(),
      success: function( result ) {
        codeNode.val('')
        warning('提交成功', function(){
          location.href = '/status';
        });
      },
      error: function(err) {
        if( err.status == 401 ){
          warning('请先登录', function(){
            location.href = '/sess';
          });
        }
        else {
          var json = eval('('+err.responseText+')');
          if( json.info != null )
            warning(json.info);
        }
      }
    });
  });
  
}
if( /^\/problems$/.test(PATH) ){
	console.log(1);
	var inputNode = $('[name=search]');
	var selectNode = $('[name=option]');
	var formNode = $('.J_searchForm');

	var validator = function(){
		var value = inputNode.val();
		console.log(value);
		inputNode.val( value.replace(/[^\.\d]/g,'') );
	};
	selectNode.on('change', function(){
		var value = $(this).val();
		inputNode.val('');
		if( value=='pid' || value=='page' )
			inputNode.on('input', validator);
		else
			inputNode.off('input');
	});
	formNode.on('submit', function(e) {
    e.preventDefault();
    var value = inputNode.val();
    var key = selectNode.val();
    value = encodeURIComponent(value);
    location.href = '/problems?'+key+'='+value;
  });
}
if( PATH == '/status/code' ){

	var codeNode = $('.J_code');
	var sourceNode = $('.J_source');
	var flag = true;

	codeNode.on('dblclick', function(){
		codeNode.hide();
			if( flag ){
				sourceNode.height(codeNode.height());
				falg = !flag;
			}
		sourceNode.show();
	});
	sourceNode.on('blur', function(){
		codeNode.show();
		sourceNode.hide();
	});
}
if( PATH == '/profile' ){
	console.log(1);
	$('.J_addForm').on('submit', function(e){
		e.preventDefault();
		if( $('[name="user[nick]"]').val() ){
			warning('请输入昵称');
			return;
		}
		$.ajax({
			type: 'post',
			url: '/profile',
			data: $(this).serialize(),
			success: function(data) {
				warning('修改成功', function(){
					location.href = '/settings';
				});
			},
			error: function(err) {
				warning(JSON.parse(err.responseText).nick);
			}
		});
	});
}
if( PATH == '/account' ){
	var formNode = $('.J_addForm');
	var oldNode = $('[name="user[oldPassword]"]');
	var newNode = $('[name="user[newPassword]"]');
	var confirmNode = $('[name="user[confirmPassword]"]');
	formNode.on('submit', function(e){
		e.preventDefault();
		var passWd = oldNode.val();
		var newPassWd = newNode.val();
		var confirmPassWd = confirmNode.val();
		if( !passWd || passWd=='' ){
			warning('请输入原密码');
		}
		else if( !newPassWd || newPassWd=='' ){
			warning('请输入新密码');
		}
		else if( !confirmPassWd || confirmPassWd=='' ){
			warning('请再次输入密码');
		}
		else if( newPassWd != confirmPassWd ){
			warning('两次密码不一致');
		}
		else{
			$.ajax({
				url: '/account',
				type: 'post',
				data: $(this).serialize(),
				dataType: 'text',
				success: function(data){
					warning('修改成功', function(){
						location.href = '/settings';
					});
				},
				error: function(err){
					var info = JSON.parse(err.responseText);
					if( info.oldPassword ){
						warning( info.oldPassword );
					}
					else if( info.newPassword ){
						warning( info.newPassword );
					}
					else if( info.confirmPassword ) {
						warning( info.confirmPassword );
					}
					else {
						warning( JSON.stringify(err) );
					}
				}
			});
		}
	});
}
//登陆页
if( /^\/sess$/.test(PATH) ) {

  var formNode = $('.J_addForm');
  var registerBtn = $('.J_register');
  var submitBtn = $('.J_submit');
  var nameNode = $("#user_handle");
  var passwdNode = $('#user_password');
  //前往注册页
  registerBtn.on('click', function(){
    location.href = '/users/new';
  });
  //登陆
  formNode.on('submit', function(e){
    e.preventDefault();
    if( !nameNode.val() ){
      warning('请输入账号');
      return;
    }
    else if( !passwdNode.val() ){
      warning('请输入密码');
      return;
    }
    $.ajax({
      type: 'post',
      url: '/sess',
      data: formNode.serialize(),
      dataType: 'text',
      error: function(err){
        console.log(err);
        warning('用户名或密码错误！');
      },
      success: function(data){
        console.log(data);
        warning('登陆成功', function(){
          if (document.referrer != ""){
            location.href = document.referrer;
          }else{
            location.href = "/";
          }
        });
      }
    });
  });

}