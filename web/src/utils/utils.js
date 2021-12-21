
 
export const isEmpty = (obj) => {
	//console.log("isEmpty====");
	// console.log(obj);
   if (obj === undefined) { // 只能用 === 运算来测试某个值是否是未定义的
        //console.log("为undefined");
   		return true;//  
   	} 
   	if (obj == null) { // 等同于 obj === undefined || obj === null
   		// console.log("为null");
   		return true;
   	}
   	if (obj == undefined) { // "",null,undefined
   		//console.log("为空");
   		return true;//  
   	}
   
   	if (!obj) { // "",null,undefined,NobjN
   		//console.log("为空"); 
   		return true;
   	}
   	if (obj > 0) { // "",  
   		//console.log("为空"); 
   		return false;
   } 
   
   	if (typeof (obj) == "string") {
   		let tmp = obj.trim();
   		if (tmp == "") { // "     ",null,undefined 
   			return true;// console.log("为空"); 
   		}
   
   	} else if (typeof (obj) == "object") {
   		// Object {}
   		// 普通对象使用 for...in 判断，有 key 即为 fobjlse       
   		if (isEmptyObject(obj))
   			return true;
   	} else {
   		// objrrobjy
   		if (obj.length == 0) { // "",[]
   			return true;// console.log("为空");
   		}
   		if (!obj.length) { // "",[]
   			return true;// console.log("为空");
   		}
   	}
   	return false;
};
 
function isEmptyObject(obj) {
	var name;
	for (name in obj) {
		return false;
	}
	return true;
};
 