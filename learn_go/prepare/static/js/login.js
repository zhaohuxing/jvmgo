function login(){
	var number = document.getElementById("phoneNumber").value;
	var pwd = document.getElementById("password").value;
	//4.提交到后台
	$.ajax({
		type:"post",//请求方式 post/get
		url: "http://127.0.0.1:9000/login",
		data: {
		    phoneNumber: number,
		    password: pwd
		},
		async:false,//是否同步
		timeout:5000,//设置超时时间
		dataType:"json",
		success:function(data){
			//判断data是否为空
			alert(data.message)
		},
		error:function(xhr,textState){
			alert("数据请求失败");
		}
	});
}
