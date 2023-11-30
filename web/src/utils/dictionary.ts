import {
	useDictionaryStore
} from '@/pinia/modules/dictionary'
import {
	pidData,
	pidTreeData,
	getDictData,
	getDictData2,
	getDictTreeData
} from '@/api/common_db'

import {isEmpty,obj2Num} from '@/utils/utils' 


//  获取字典方法 旧版本 使用示例 getDict('sex').then(res)  或者 async函数下 const res = await getDict('sex')
export const getDict = async (type) => {
	const dictionaryStore = useDictionaryStore()
	await dictionaryStore.getDictionary(type)
	return dictionaryStore.dictionaryMap[type]
}

//获取  PidData 任意数据表
export const getPidData = async (data) => {
	const res = await pidData(data)
	//console.log(res)
	if (res.code == 200)
		return res.data.List
	else return []
}

//获取字典PidTreeData 任意数据表
export const getPidTreeData = async (data) => {
	const res = await pidTreeData(data)
	if (res.code == 200)
		return res.data.List
	else return []
}
//获取字典表 新版本的字典 basic_dict  getDictNew
export const getDictNew = async (data) => {
	const res = await getDictData(data)
	if (res.code == 200)
		return res.data.List
	else return []
} 
//获取字典表 新版本的字典 basic_dict   比 getDictNew 多了 type module
export const getDictNew2 = async (data) => {
	const res = await getDictData2(data)
	if (res.code == 200)
		return res.data.List
	else return []
} 
//获取字典树 PidTreeData basic_dict
export const getDictTreeNew = async (data) => {
	const res = await getDictTreeData(data)
	if (res.code == 200)
		return res.data.List
	else return []
}

//根据 value 获取 label
export  const getOptLabel =  (value, options) => {  
	 if (isEmpty(value)) 
		  return "";	 
	 if (isEmpty(options)) 
		 return "";   
	 let label =""
	 for (var i=0;i<options.length;i++) { 
	     if (options[i].value == value) {
			 label =  options[i].label; 
			 break; 
		 }
	 } 
	// console.log('label=' ,label)
	 return label;   
 }  
 
 //获取Tree的全路径 
export  const getTreeFullPath = (treeData, val) => {
 	var arr = []; //定义一个空数组
 	var out = false; //定义一个标识（是否找到对应的id） 
 	//console.log("1111")			
 	fun(treeData); //调用封装好的回显函数				
 	function fun(childrenArr) { //封装的回显函数
 		//console.log("22222")	
 		//console.log(childrenArr)	
 		var bg = 0; //定义一个标杆（标识循环到当前数组的第几条）
 		for (var item of childrenArr) {
 			//console.log(item)
 			//console.log(val)	
 			if (item.value === val) { //判断所在数组（层级），是否有与之匹配的id
 				arr.push(item.value); //存在则返回其id
 				out = true; //存在就将标识设置为true
 				return; //递归出口（出口1）
 			} else if (item.children && item.children.length > 0) {
 				//判断id不匹配的层级是否具有下级目录（children）
 				arr.push(item.value);
 				fun(item.children); //存在下级目录就将下一级目录回调
 			}
 			bg++;
 			if (out == true) return //如果找到对应的id 就执行返回 （出口2）
 			if (bg === childrenArr.length && arr && arr.length > 0) {
 				//找不到的时候清除不属于的父类Id
 				arr.pop()
 			}
 		}
 	}
 //	console.log("getTreeFullPath === ")
 //	console.log(arr)
 	return arr
 }
 
 //获取tree value对应的label 
 export const getTreeName = (treeData, val) => {
 	if (isEmpty(val))
 		return "暂无"
 	var arr = []; //定义一个空数组
 	var out = false; //定义一个标识（是否找到对应的id） 
 	//console.log("1111")			
 	fun(treeData); //调用封装好的回显函数				
 	function fun(childrenArr) { //封装的回显函数
 		//console.log("22222")	
 		//console.log(childrenArr)	
 		var bg = 0; //定义一个标杆（标识循环到当前数组的第几条）
 		for (var item of childrenArr) {
 			//console.log(item)
 			//console.log(val)	
 			if (item.value === val) { //判断所在数组（层级），是否有与之匹配的id
 				//arr.push(item.value); //存在则返回其id
 				arr.push(item.label); //存在则返回其id
 				out = true; //存在就将标识设置为true
 				return; //递归出口（出口1）
 			} else if (item.children && item.children.length > 0) {
 				//判断id不匹配的层级是否具有下级目录（children）
 				//arr.push(item.value);
 				arr.push(item.label); //存在则返回其id
 				fun(item.children); //存在下级目录就将下一级目录回调
 			}
 			bg++;
 			if (out == true) return //如果找到对应的id 就执行返回 （出口2）
 			if (bg === childrenArr.length && arr && arr.length > 0) {
 				//找不到的时候清除不属于的父类Id
 				arr.pop()
 			}
 		}
 	}
 	//console.log(arr)
 	if (arr.length > 0)
 		return arr[arr.length - 1]
 	else
 		return "暂无"
 }
 
  
 export default {
 	getTreeName,
 	getTreeFullPath,
 	getOptLabel,
 	getDictTreeNew,
	getDictNew,
	getPidData,
	getDict 
 };
