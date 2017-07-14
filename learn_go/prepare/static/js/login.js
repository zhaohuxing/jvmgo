function login(){
	var number = document.getElementById("phoneNumber").value;
	var pwd = document.getElementById("password").value;
	var data = {
		phoneNumber: $("#phoneNumber").val,
		password: $("#password").val
	}
	//4.提交到后台
	$.ajax({
		type:"post",//请求方式 post/get
		//url:"http://127.0.0.1:8080/bmi/bmi?method=add_bmi",//请求地址
		url: "http://127.0.0.1:9000/login",
		//变量名：变量值 参数与参数之间使用逗号隔开
//		data:{height:height,weight:weight,bmi:bmiNum,dateTime:dateTime},//传递参数到服务器上
		data:JSON.stringify(data),
		async:false,//是否同步
		timeout:5000,//设置超时时间
		dataType:"json",
		success:function(data){
			//判断data是否为空
			if(!jQuery.isEmptyObject(data)){
				var table2=document.getElementById("table2");
				for(var i=0;i<data.length;i++){
					var tr1=document.createElement("tr");
					table2.appendChild(tr1);
					addBmiHistory(data[i]);
				}
			} 
		},
		error:function(xhr,textState){
			alert("数据请求失败");
		}
	});
}

$(document).ready(function(e){
	$.ajax({
		type:"post",//请求方式 post/get
		//url:"http://127.0.0.1:8080/bmi/bmi?method=start_bmi",//请求地址
		url: "http://127.0.0.1:9000/login",	
		async:false,//是否同步
		timeout:5000,//设置超时时间
		dataType:"json",
		success:function(data){
			alert(data.message)
			alert(data);
			var table2=document.getElementById("table2");
			for(var i in data){
			}
		},
		error:function(xhr,textState){
			alert("请求错误");
		}
	});
});
