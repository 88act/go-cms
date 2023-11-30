import { defineStore } from 'pinia'
import { ref } from 'vue'

import {
	pidData,
	pidTreeData,
	getDictData,
	getDictData2,
	getDictTreeData
} from '@/api/common_db'


//获取字典表 新版本的字典 basic_dict  getDictNew
export const getDictNew = async (data) => {
	const res = await getDictData(data)
	if (res.code == 200)
		return res.data.List
	else return []
} 


//本地缓存了dict ，可能需要强制更新 
export const useDictionaryStore = defineStore('dictionary', () => {
  const dictionaryMap = ref({})

  const setDictionaryMap = (dictionaryRes) => {
    dictionaryMap.value = { ...dictionaryMap.value, ...dictionaryRes }
  }

  const getDictionary = async(type) => {
	//console.log("getDictNew , type = ",type) 
	// 可能不是最新的 数据 ，需要清理缓存机制 
    if (dictionaryMap.value[type] && dictionaryMap.value[type].length) {
         return dictionaryMap.value[type]
    } else {   
		let res = await getDictNew({"module":type}) 
		console.log("getDictNew")
		console.log(res)
		if (res?.length > 0 ) {
		  const dictionaryRes = {} 
		  dictionaryRes[type] = res
		  setDictionaryMap(dictionaryRes)
		  return res
		} else return [] 
		  /* 
		  const res = await findSysDictionary({ type })
		  if (res.code === 200) {
			const dictionaryRes = {}
			const dict = []
			res.data.resysDictionary.sysDictionaryDetails && res.data.resysDictionary.sysDictionaryDetails.forEach(item => {
			  dict.push({
				label: item.label,
				value: item.value
			  })
			})
			dictionaryRes[res.data.resysDictionary.type] = dict
			setDictionaryMap(dictionaryRes)
			return dictionaryMap.value[type]
		  }
		  */ 
       }
  }

  return {
    dictionaryMap,
    setDictionaryMap,
    getDictionary
  }
})
