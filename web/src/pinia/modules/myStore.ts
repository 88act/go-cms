import {
	defineStore
} from 'pinia'
import {
	ref
} from 'vue'


const MyStore = defineStore('MyStore', {
	state: () => ({
		btnList: [],
		name: '小明',
		btnClickObj: {},
		sysUserList:[], 
		prjList:[]    ,
		prjList2:[]
	}),
	actions: {
		 // 系统用户 列表
		setSysUserList(val) {
			 this.sysUserList = val
		},
		//项目列表
		setPrjList(val) {
			 this.prjList = val
		},
		 //待申请项目
		setPrjList2(val) {
			 this.prjList2 = val
		},
		 
		changeName(name, isError) {
			return new Promise((resolve, reject) => {
				this.name = name
				if (isError) {
					resolve(`姓名：${this.name}`)
				} else {
					reject('error')
				}
			})
		},
		changeBtnList(list, isError) {
			return new Promise((resolve, reject) => {
				this.btnList = list
				if (isError) {
					resolve(`changeBtnList：${this.btnList}`)
				} else {
					reject('error')
				}
			})
		},
		changeBtnClick(obj, isError) {
			return new Promise((resolve, reject) => {
				this.btnClickObj = obj
				if (isError) {
					resolve(`changeBtnClick：${this.btnClickObj}`)
				} else {
					reject('error')
				}
			})
		},

	},

})
export default MyStore
