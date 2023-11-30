
	export const  ipToNumber=(ip)=> {
		var num = 0;
		if(ip == "") {
			return num;
		}
	    var aNum = ip.split(".");
	    if(aNum.length != 4) {
	    	return num;
	    }
	    num += parseInt(aNum[0]) << 24;
	    num += parseInt(aNum[1]) << 16;
	    num += parseInt(aNum[2]) << 8;
	    num += parseInt(aNum[3]) << 0;
	    num = num >>> 0;//这个很关键，不然可能会出现负数的情况
	    return num;
	}

	export const  numberToIp=(number)=> {
	    var ip = "";
	    if(number <= 0) {
	    	return ip;
	    }
	    var ip3 = (number << 0 ) >>> 24;
	    var ip2 = (number << 8 ) >>> 24;
	    var ip1 = (number << 16) >>> 24;
	    var ip0 = (number << 24) >>> 24

	    ip += ip3 + "." + ip2 + "." + ip1 + "." + ip0;

	    return ip;
	}

	export  const getFileByGuidList=(gList, list)=> {
		let newList  = []
		if (!list || list.length == 0) {
			return newList
		}

		if (gList && gList.length > 0) {
			for (var j = 0; j < gList.length; j++) {
				let guid = gList[j]
				//console.log("gList  uid =", uid)
				if (guid != "") {
					for (var i = 0; i < list.length; i++) {
						if (guid == list[i].guid) {
							//console.log("gList  uid =", uid)
							let oneFile = { path: list[i].path, guid: guid }
							newList.push(oneFile)
							break;
						}
					}
				}
			}
		}
		return newList
	}

	export const getFileByGuidStr =(guidStr, list)=> {
		let newList  = []
		if (!list || list.length == 0) {
			return newList
		}

		if (guidStr == "") return newList
		let gList = guidStr.split(",")
		//console.log("gList  =", gList)
		for (var j = 0; j < gList.length; j++) {
			let guid = gList[j]
			if(guid.substr(0, 4) == "http") {
				let oneFile = { path: guid, guid: guid }
				newList.push(oneFile)
			}
			else {
			    //console.log("gList  uid =", guid)
				if (guid != "") {
					for (var i = 0; i < list.length; i++) {
						if (guid == list[i].guid) {
							//console.log("gList  uid =", guid)
							let oneFile = { path: list[i].path, guid: guid }
							newList.push(oneFile)
							break;
						}
					}
				}
			}
		}
		return newList
	}

	export   const getFileByGuid=(guid , list ) =>{
	   let newList = []
	   if (!list || list.length == 0) {
	   	return newList
	   }

		for (var i = 0; i < list.length; i++) {
			if (guid == list[i].guid) {
				 let oneFile = { path: list[i].path, guid: guid }
				 newList.push(oneFile)
				break;
			}
		}
		return newList
	}


// 移除一个查询对象 空属性  add by ljd 20210718
export const removeNullAttr=(obj)=> {
		for (const key in obj) {
			if (obj.hasOwnProperty(key)) {
				if (obj[key] === null || obj[key] === '') {
					console.log('为空=' + key)
					delete obj[key]
				} else {
					//console.log('不为空'+key)
				}
			}
		}
}


// 对象转数值
export const obj2Num = (obj) => {
	//console.log("isEmpty====");
		// console.log(obj);
	if (obj === undefined) { // 只能用 === 运算来测试某个值是否是未定义的
	     //console.log("为undefined");
			return 0;//
	  }
		if (obj == null) { // 等同于 obj === undefined || obj === null
			// console.log("为null");
			return 0;
		}
		if (obj == undefined) { // "",null,undefined
			//console.log("为空");
			return 0;//
		}

		if (!obj) { // "",null,undefined,NobjN
			//console.log("为空");
			return 0;
		}

		if (typeof (obj) == "string") {
			let tmp = obj.trim();
			if (tmp == "") { // "     ",null,undefined
				return 0;// console.log("为空");
			}

		} else if (typeof (obj) == "object") {
			// Object {}
			// 普通对象使用 for...in 判断，有 key 即为 fobjlse
			if (isEmptyObject(obj))
				return 0;
		} else {
			// objrrobjy
			if (obj.length == 0) { // "",[]
				return 0;// console.log("为空");
			}
			if (!obj.length) { // "",[]
				return 0;// console.log("为空");
			}
		}
		return Number(obj);
}
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




 // 默认时间 7天,自动展示前一周到今天的日期
 export const get7days = () => {
 	let date = new Date()
 	// 通过时间戳计算
 	let defalutStartTime = date.getTime() - 7 * 24 * 3600 * 1000 // 转化为时间戳
 	let defalutEndTime = date.getTime()
 	let startDateNs = new Date(defalutStartTime)
 	let endDateNs = new Date(defalutEndTime)
 	// 月，日 不够10补0
 	defalutStartTime = startDateNs.getFullYear() + '-' + ((startDateNs.getMonth() + 1) >= 10 ? (startDateNs.getMonth() + 1) : '0' + (startDateNs.getMonth() + 1)) + '-' + (startDateNs.getDate() >= 10 ? startDateNs.getDate() : '0' + startDateNs.getDate())
 	defalutEndTime = endDateNs.getFullYear() + '-' + ((endDateNs.getMonth() + 1) >= 10 ? (endDateNs.getMonth() + 1) : '0' + (endDateNs.getMonth() + 1)) + '-' + (endDateNs.getDate() >= 10 ? endDateNs.getDate() : '0' + endDateNs.getDate())
 	defalutStartTime = defalutStartTime+ " 00:00:00"
	defalutEndTime = defalutEndTime+ " 00:00:00"
	return [defalutStartTime, defalutEndTime]
  };


   export const  getShortcuts =()=>{

	 let shortcuts = [{
  		text: '上周',
  		value: () => {
  			const end = new Date()
  			const start = new Date()
  			start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
  			return [start, end]
  		},
  	},
  	{
  		text: '上月',
  		value: () => {
  			const end = new Date()
  			const start = new Date()
  			start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
  			return [start, end]
  		},
  	},
  	{
  		text: '上3个月',
  		value: () => {
  			const end = new Date()
  			const start = new Date()
  			start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
  			return [start, end]
  		},
  	},
	{
		text: '上6个月',
		value: () => {
			const end = new Date()
			const start = new Date()
			start.setTime(start.getTime() - 3600 * 1000 * 24 * 180)
			return [start, end]
		},
	},
	{
		text: '一年内',
		value: () => {
			const end = new Date()
			const start = new Date()
			start.setTime(start.getTime() - 3600 * 1000 * 24 * 360)
			return [start, end]
		},
	},
  ]
  return shortcuts;
  }


  function padTo2Digits(num) {
    return num.toString().padStart(2, '0');
  }

  // 获取当前时间 2016-01-02 03:14:15
  export const getDataTimeStr = (date,isTime) => {
  	 if (!isTime)
  	 {
  		return ([
  		     date.getFullYear(),
  		     padTo2Digits(date.getMonth() + 1),
  		     padTo2Digits(date.getDate()),
  		   ].join('-') );
  	 }else{
  		 return ([
  		     date.getFullYear(),
  		     padTo2Digits(date.getMonth() + 1),
  		     padTo2Digits(date.getDate()),
  		   ].join('-') + ' ' + [
  		     padTo2Digits(date.getHours()),
  		     padTo2Digits(date.getMinutes()),
  		     padTo2Digits(date.getSeconds()),
  		   ].join(':')
  		 );
  	 }
   }


  // 截取字符串
   export const substr =(str,m)=> {
	  if (str.length > m) {
		  return str.substring(0,m) + "..."
	  }else {
		   return str
	  }
 }

 // 获取数组的名称
 export const getNameById = (id,list) => {
	 // console.log(' id =' ,id )
	 //   console.log(' list =' ,list)
 	 if (id<=0) return ""
 	 for (let obj of list ) {
		 //console.log("obj =",obj)
 		 if (obj.value == id ){
 			 return obj.label
 		 }
      }
 	 return "暂无"
 }

 //获取城市路径
 export const getCityCode = (id , lev ) => {
 	let idStr  = id + ""
 	let numStr  = ""
 	let num  = 0
 	if (lev == 1) {  // 省
 		numStr = idStr.slice(0, 2)
 		num = Number(numStr) * 10000
 		return num
 	} else if (lev == 2) { //市
 		numStr = idStr.slice(0, 4)
 		num = Number(numStr) * 100
 		return num
 	}
 	else if (lev == 3) { //区
 		return id
 		//  numStr = idStr.slice(0,6)
 		//  num = Number(numStr)*10
 		//  return num
 	}
 }

 //获取城市name
 export const getCityName = (id , lev , areaCityData ) => {
 	let idStr  = ""
 	let pid = getCityCode(id, 1)
 	let cid  = 0
 	let did   = 0
 	if (lev >= 2)
 		cid = getCityCode(id, 2)
 	if (lev >= 3)
 		did = id // getCityCode(id, 3)

 	//console.log("pid,cid  did== ", pid, cid, did)
 	for (let i = 0; i < areaCityData.length; i++) {
 		if (pid == areaCityData[i].id) {
 			//console.log(" pid == ",pid,areaCityData[i].text)
 			idStr = areaCityData[i].text
 			if (lev >= 2) {
 				let child2 = areaCityData[i].children
 				for (let j = 0; j < child2.length; j++) {
 					//console.log(" cid == ",cid, child2[j].text)
 					if (cid == child2[j].id) {
 						//console.log(" cid == ",cid, child2[j].text)
 						idStr = idStr + "," + child2[j].text
 						if (lev >= 3) {
 							let child3 = child2[j].children
 							for (let k = 0; k < child3.length; k++) {
 								//console.log(" cid == ",cid, child2[j].text)
 								if (did == child3[k].id) {
 									//console.log(" cid == ",cid, child2[j].text)
 									idStr = idStr + "," + child3[k].text
 									break
 								}
 							}
 						}
 						break
 					}
 				}
 			}
 			break
 		}
 	}
 	//console.log(" idStr == ", idStr)
 	return idStr
 }
 
 
 
 // 对Date的扩展，将 Date 转化为指定格式的String
 // 月(M)、日(d)、小时(h)、分(m)、秒(s)、季度(q) 可以用 1-2 个占位符，
 // 年(y)可以用 1-4 个占位符，毫秒(S)只能用 1 个占位符(是 1-3 位的数字)
 // (new Date()).Format("yyyy-MM-dd hh:mm:ss.S") ==> 2006-07-02 08:09:04.423
 // (new Date()).Format("yyyy-M-d h:m:s.S")      ==> 2006-7-2 8:9:4.18
 // eslint-disable-next-line no-extend-native
 Date.prototype.Format = function(fmt) {
   var o = {
     'M+': this.getMonth() + 1, // 月份
     'd+': this.getDate(), // 日
     'h+': this.getHours(), // 小时
     'm+': this.getMinutes(), // 分
     's+': this.getSeconds(), // 秒
     'q+': Math.floor((this.getMonth() + 3) / 3), // 季度
     'S': this.getMilliseconds() // 毫秒
   }
   if (/(y+)/.test(fmt)) { fmt = fmt.replace(RegExp.$1, (this.getFullYear() + '').substr(4 - RegExp.$1.length)) }
   for (var k in o) {
     if (new RegExp('(' + k + ')').test(fmt)) { fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? (o[k]) : (('00' + o[k]).substr(('' + o[k]).length))) }
   }
   return fmt
 }
 
 export function formatTimeToStr(times, pattern) {
   var d = new Date(times).Format('yyyy-MM-dd hh:mm:ss')
   if (pattern) {
     d = new Date(times).Format(pattern)
   }
   return d.toLocaleString()
 }
 
 
export const formatBoolean = (bool) => {
  if (bool !== null) {
    return bool ? '是' : '否'
  } else {
    return ''
  }
}
export const formatDate = (time,type) => {
  if (time !== null && time !== '') {
    var date = new Date(time)
	if (type==undefined || type ==0){
		 return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
	} else if (type == 1){
		 return formatTimeToStr(date, 'yyyy-MM-dd')
	}else if (type == 2){
		 return formatTimeToStr(date, 'hh:mm:ss')
	}else { 
		 return formatTimeToStr(date, 'yyyy-MM-dd') 
	} 
  } else {
     return ''
  }
} 

export const filterDict = (value, options) => {
  const rowLabel = options && options.filter(item => item.value === value)
  return rowLabel && rowLabel[0] && rowLabel[0].label
}


 export default {
 	isEmpty,
 	obj2Num,
 	removeNullAttr,
	get7days,
	getShortcuts,
	getDataTimeStr,
	getFileByGuid,
	getFileByGuidList,
	getFileByGuidStr,
	substr,
	getCityCode,
	getCityName,
	getNameById,
 
  filterDict,
  formatDate,
  formatBoolean 
 };
