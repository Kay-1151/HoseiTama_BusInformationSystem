
/* ここで表の情報を変える(多分使わない) */
/*function changeinfo(){
	var table=document.getElementById("timetable");
	var number=1;
	var str;
	var row;
	var cell;
	while(table.rows[1]){
		table.deleteRow(1);
	}
	for(var h=1;h<4;h++){
		ow=table.insertRow(h);
		for(var t=0;t<=2;t++){
			cell=row.insertCell(t);
			str="time"+number;
			var data={{.str}};
			cell.innerHTML=data;
			number+=1;
		}
	}
}*/
function checklocate(location){
	var select=document.getElementById("location");
	var option=document.getElementById("location").options;
	var selected=option.item(select.selectedIndex).value;
	if(selected==="selectlocation"){
		window.alert("場所を選択してください");
	}else{
		/* 実際はここでサーバーに送る情報を変える。
		セレクトボックスで初期値が選ばれたときにエラーのウィンドウを表示させるようにしている。*/
		//Connect(selected);
	}
}

 var canvas;
 var context;
 /* canvasの初期化 */
function init(){
	canvas=document.getElementById("map");
	if(canvas.getContext){
		context=canvas.getContext("2d");
		canvas.addEventListener('click',GetClickAction,true);
		DrawMap();
	}
}

/* 地図を表示 */
function DrawMap(){
	var img=new Image();
 	var canvas1=document.getElementById("map");
 	img.src="../../static/images/tama_map_point.png?"+new Date().getTime();
 	context.style="#000000";
 	img.onload=function(){
  		context.drawImage(img,0,0,canvas1.width,canvas1.height);
 };
}
/* バス停クリックしたときにセレクトボックスの値を変える */
function GetClickAction(event){
	var x=0;
	var y=0;
	var rect=event.target.getBoundingClientRect();
	var selectbox=document.getElementById("location").options;
	var select=null;
	x=event.clientX-Math.floor(rect.left);
	y=event.clientY-Math.floor(rect.top);
	if(x>56&&x<135&&y>3&&y<100){
		select="stop1";
	}else if(x>499&&x<575&&y>129&&y<226){
		select="stop2";
	}else if(x>476&&x<546&&y>242&&y<311){
		select="stop3";
	}else if(x>823&&x<891&&y>353&&y<439){
		select="stop4";
	}
	if(select!==null){
		for(var i=0;i<selectbox.length;i++){
			if(selectbox[i].value===select){
				selectbox[i].selected=true;
				break;
			}
		}
	}
}
 /* ウィンドウサイズが変わったときにバスの位置を修正する関数 */
 function move_resized(){
  var canvas=document.getElementById("map");
  var rect=canvas.getBoundingClientRect();
  /* バスの数だけ定義 */
  var bus1=document.getElementById("bus1");
  bus1.style.marginLeft=String(rect.left)+"px";
  bus1.style.marginTop=String(rect.top +window.pageYOffset)+"px";
  var bus2=document.getElementById("bus2");
  bus2.style.marginLeft=String(rect.left)+"px";
  bus2.style.marginTop=String(rect.top +window.pageYOffset)+"px";
  var bus3=document.getElementById("bus3");
  bus3.style.marginLeft=String(rect.left)+"px";
  bus3.style.marginTop=String(rect.top +window.pageYOffset)+"px";
  var bus4=document.getElementById("bus4");
  bus4.style.marginLeft=String(rect.left)+"px";
  bus4.style.marginTop=String(rect.top +window.pageYOffset)+"px";
  var bus5=document.getElementById("bus5");
  bus5.style.marginLeft=String(rect.left)+"px";
  bus5.style.marginTop=String(rect.top +window.pageYOffset)+"px";
  var bus6=document.getElementById("bus6");
  bus6.style.marginLeft=String(rect.left)+"px";
  bus6.style.marginTop=String(rect.top +window.pageYOffset)+"px";
  var bus7=document.getElementById("bus7");
  bus7.style.marginLeft=String(rect.left)+"px";
  bus7.style.marginTop=String(rect.top +window.pageYOffset)+"px";
  var bus8=document.getElementById("bus8");
  bus8.style.marginLeft=String(rect.left)+"px";
  bus8.style.marginTop=String(rect.top +window.pageYOffset)+"px";
  var bus9=document.getElementById("bus9");
  bus9.style.marginLeft=String(rect.left)+"px";
  bus9.style.marginTop=String(rect.top +window.pageYOffset)+"px";
  var bus10=document.getElementById("bus10");
  bus10.style.marginLeft=String(rect.left)+"px";
  bus10.style.marginTop=String(rect.top +window.pageYOffset)+"px";
  var bus11=document.getElementById("bus11");
  bus11.style.marginLeft=String(rect.left)+"px";
  bus11.style.marginTop=String(rect.top +window.pageYOffset)+"px";
	console.log(rect.left);
	console.log(rect.top);
 }
 window.onresize=function(){
	 move_resized();
 };
 window.
	window.onload=function(){
	 	init();
		// move_resized();
	 	  var target = document.getElementById("output");
	 	  var targetlast=document.getElementById("outlast");
	 	  var targetdesu=document.getElementById("outdesu");
	 	  var path=location.href;
	 	  var element=document.getElementById("location");
	 	  var elements=element.options;
	 	  //console.log(path.match('stop1'));
	 		if(path.indexOf('stop1')>0){
	           target.innerHTML = "体育館の運行時間は";
	 		  targetlast.innerHTML ="最終便は";
	 		  targetdesu.innerHTML="です。";
	 		  elements[1].selected=true;
	 	  }else if(path.indexOf('stop2')>0){
	           target.innerHTML = "正門・総合棟・経済の運行時間は";
	 		  targetlast.innerHTML ="最終便は";
	 		  targetdesu.innerHTML="です。";
	 		  elements[2].selected=true;
	 	  }else if(path.indexOf('stop3')>0){
	           target.innerHTML = "エッグドームの運行時間は";
	 		  targetlast.innerHTML ="最終便は";
	 		  targetdesu.innerHTML="です。";
	 		  elements[3].selected=true;
	 	  }else if(path.indexOf('stop4')>0){
	           target.innerHTML = "スポーツ健康学部の運行時間は";
	 		  targetlast.innerHTML ="最終便は";
	 		  targetdesu.innerHTML="です。";
	 		  elements[4].selected=true;
	 	  }else{
	 		  target.innerHTML ="";
	 		  targetlast.innerHTML ="";
	 		  targetdesu.innerHTML="";
	 		  elements[0].selected=true;
	 	  }
			move_resized();
	 };

function doReloadWithCache() {

    /* キャッシュを利用してリロード */
    /*window.location.reload(true);*/

}
window.addEventListener('load', function () {

    /* ページ表示完了した5秒後にリロード */
    //setTimeout(doReloadWithCache, 35000);
 });
 function button(){
 var selectbox=document.getElementById("location").options;
 var button=document.getElementById("search");
 if(selectbox[0].selected===true){
  button.disabled=true;
 }else{
  button.disabled=false;
}
}
