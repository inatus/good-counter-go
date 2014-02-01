$(document).ready(function(){
	$.ajax({
		url: "/count",
		success: function(html){
			try{
				var obj = JSON.parse(html);
				$("#count").text(obj.count);
			}catch(e){
				console.log(e);
			}
		}
	});

	$("#cntButton").click(function() {
		$.ajax({
			type: 'POST',
			url: "/count",
			success: function(html){
				try{
					var obj = JSON.parse(html);
					$("#count").text(obj.count);
				}catch(e){
					console.log(e);
				}
			}
		});
	});
});

