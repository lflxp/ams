
(function(){
	var menus = [
		{title:'棣栭〉', href:'index.html'},
		{title:'涓嬭浇', href:'download.html'},
		{title:'API 鏂囨。', href:'api.html'},
		{title:'浜斿垎閽熷叆闂�', href:'introduction-in-5-minutes.html'},
			{title:'鍦ㄧ嚎婕旂ず', href:'demo/helloworld.html'}
		//{title:'鎹愯禒', href:'donate.html', tip:'瑙夊緱杩樹笉閿欙紵鎹愮偣灏忛挶鍎挎敮鎸佹垜鐨勫紑鍙戝惂锛�', target:'_blank'}
		//https://me.alipay.com/jtopo
	];

	function drawMenus(menus){
		var ul = $('#nav_menu').empty();				
		$.each(menus, function(i, e){
			var url = e.href;
			if(location.href.indexOf('/demo/') != -1){
				url = '../'+e.href;
			}
			var li = $('<li>').addClass('menu-item').appendTo(ul);	
			var a = $('<a>').attr('href', url).attr('title', e.tip == null ? '':e.tip).html(e.title).appendTo(li);

			if(location.href.indexOf(e.href) != -1){
				a.addClass('active');
			}		
			if(e.target){
				a.attr('target', e.target);
			}
		});
	}
	
	$(document).ready(function(){
		drawMenus(menus);		
	});

})($ || jQuery);