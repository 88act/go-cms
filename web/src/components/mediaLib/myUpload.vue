<template>
　　　　　<el-upload :action="`${path}/commFile/upload`" :data="uploadData" :headers="{'Authorization': token }"
							:show-file-list="false" :on-success="handleImageSuccess" :before-upload="beforeImageUpload"
							:multiple="false" :auto-upload="true" :accept="acceptType">
							<el-button type="primary" class="el-btn-save"  style="width: 120px;"  >上传图片</el-button>
		　</el-upload>
</template>



<script setup>
	import {
		getDict,
		getPidData,
		getPidTreeData,
		getDictNew,
		getDictTreeNew
	} from '@/utils/dictionary'

	import {
		isEmpty,
		obj2Num,
		removeNullAttr
	} from '@/utils/utils'
	import {
		formatDate,
		filterDict,
 
		formatBoolean
	} from '@/utils/format'
	import {
		updatePidSort,
	} from '@/api/common_db'


	import {
		ref,
		defineEmits
	} from 'vue'
	import {
		ElMessage,
		ElMessageBox
	} from 'element-plus'

	import {
		createBasicFile,
		deleteBasicFile,
		deleteBasicFileByIds,
		updateBasicFile,
		findBasicFile,
		getBasicFileList,
		quickEdit
	} from '@/api/basicFile'


	import {
		createBasicFileCat,
		deleteBasicFileCatByIds,
		updateBasicFileCat,
		findBasicFileCat,
		getBasicFileCatList,
	} from '@/api/basicFileCat'

	import {
		getFileByKey
	} from '@/api/common_file'


	import MyStore from "@/pinia/modules/myStore"
	const myStore = MyStore()

	const page = ref(1)
	const total = ref(0)
	const pageSize = ref(20)
	const tableData = ref([])
	const searchInfo = ref({})
	const showAddCat = ref(false)

	const catOptions = ref([]) //文件目录

	const mediaType = ref(0) // 强制为　图片　／视频　／音频　／文件　
	const acceptType = ref("") // .jpg,.jpeg,.png,.pdf,.PDF,.doc.docx


	import ImageCompress from '@/utils/image'
	import SparkMD5 from 'spark-md5'
	import sha1 from 'sha1-file-web'
	const path = ref(
		import.meta.env.VITE_BASE_API) //import.meta.env.VITE_BASE_API // process.env.VUE_APP_BASE_API
	const emit = defineEmits(['selectOneImg'])
	//emits: ['selectOneImg'],
	const props = defineProps({
		mediaType: {
			type: Number,
			required: false,
			default: 0
		},
		max: {
			type: Number,
			required: false,
			default: 1024
		},
	})

	const defaultProps = ref({
		checkStrictly: true,
		expandTrigger: "hover"
	})

	const el_tree_props = ref({
		children: 'children',
		label: 'label',
	})

	//------图片上传--------------------
	const maxSize = ref(1024) //上传最大文件大小
	const fileSize = 2048 // 不需压缩的文件大小 2048k
	const maxWH = 1920 //压缩最大宽高
	const uploadData = ref({
		catId: 0,
		module: 0,
		media_type: 0,
		size: 0,
		ext: "",
		md5: "",
		sha1: ""
	})

	//----------------------
	const showDialog = ref(false)
	//---------------

	// dialogFormVisible: false,
	// type: '',
	// deleteVisible: false,

	//const multipleSelection = ref([])
	//const moduleOptions = ref([])
	const media_typeOptions = ref([])
	//const driverOptions = ref([])

	const formData = ref({
		guid: '',
		userIdSys: 0,
		userId: 0,
		catId: 0,
		name: '',
		module: 0,
		mediaType: 0,
		driver: 0,
		path: '',
		ext: '',
		fileType: '',
		size: 0,
		md5: '',
		sha1: '',
		sort: 0,
		download: 0,
		usedTime: 0,
		mapData: {}
	})

	const selectOneFile = (path, guid, newMediaType, mediaType) => {
		console.log("选择一个 selectImg  mediaType = ", path, guid, newMediaType, mediaType)
		if (mediaType > 0 && mediaType != newMediaType) {
			let msg = "请选择 图片"
			if (mediaType == 1)
				msg = "请选择 图片";
			else if (mediaType == 2)
				msg = "请选择 视频";
			else if (mediaType == 3)
				msg = "请选择 音频";
			else if (mediaType == 4)
				msg = "请选择 文档";
			ElMessage.error(msg + " 类型")
			return;
		}
		var obj = new Object();
		obj["path"] = path;
		obj["guid"] = guid;

		emit('selectOneImg', obj)
		showDialog.value = false
	}


	const beforeImageUpload = async (file) => {
		console.log(file)
		console.log("file.type =  ", file.type)
		searchInfo.value.guid = "";
		//console.log("this.uploadData.module", this.uploadData.module)
		//console.log(" Promise media_type", this.uploadData.media_type)
		// if (  this.uploadData.module ==-1)
		// {
		// 	this.$message.error('请选择上传的模块')
		// 	return reject("error")
		// }
		// if ( this.uploadData.media_type ==-1 )
		// {
		// 	this.$message.error('请选择文件类型')
		// 	return reject("error")
		// }

		if (searchInfo.value.pidList && searchInfo.value.pidList.length > 0)
			uploadData.value.catId = searchInfo.value.pidList[searchInfo.value.pidList.length - 1]
		else
			uploadData.value.catId = 0

		let isJPG = file.type === 'image/jpeg'
		let isPng = file.type === 'image/png'
		if (mediaType.value == 1 && !isJPG && !isPng) {
			ElMessage.error('上传图片只能是 jpg或png 格式!')
			return reject("error")
		}
		uploadData.value.sha1 = await getFileSha1(file);
		console.log("this.uploadData.sha1 = ", uploadData.value.sha1)
		await getFileMd5(file);
		console.log("this.uploadData.md5 = ", uploadData.value.md5)

		uploadData.value.size = Math.round(file.size / 1024); //四舍五入（小数部分）
		console.log(" uploadData.value.size = ", uploadData.value.size, isJPG, isPng);
		if (uploadData.value.size > maxSize.value) {
			ElMessage.error("上传失败，文件已超过 " + maxSize.value + "KB")

			return reject("error")
		}

		//判断是否有重复的 sha1 文件
		let res = await getFileByKey({
			sha1: uploadData.value.sha1
		})
		console.log(" res = ", res)
		if (res.code === 200 && res.data.basicFile && res.data.basicFile.path != "") {
			console.log(" 已存在相同的文件  = ", res.data.basicFile)
			let path = res.data.basicFile.path
			var obj = new Object();
			obj["path"] = path;
			obj["guid"] = res.data.basicFile.guid;
			emit('selectOneImg', obj)
			//Elmessage.awa('已存在相同的文件')
			//handleGetByGuid(res.data.basicFile.guid);
			return reject("error")
		} else {
			console.log(" 开始新上传。。。 ");

			if (isJPG || isPng) {
				let isRightSize = uploadData.value.size < fileSize
				console.log(" this.uploadData.size = ", uploadData.value.size, " this.fileSize = ", fileSize)
				if (!isRightSize) {
					// 压缩图片
					console.log("需要压缩图片 isRightSize")
					let compress = new ImageCompress(file, fileSize, maxWH)
					return compress.compress()
				}
				console.log(" resolve =ok ")
			}
			return new Promise(resolve => {
				return resolve(file)
			});
		}
	}
	//获取文件sha1值
	const getFileSha1 = async (file) => {
		let hash = await sha1(file);
		console.log("文件 sha1 = ", hash)
		//this.uploadData.sha1 = hash;
		return hash;
	}
	const getFileMd5 = async (file) => {

		var fileReader = new FileReader();
		var dataFile = file;
		let blobSlice = File.prototype.slice || File.prototype.mozSlice || File.prototype.webkitSlice
		var spark = new SparkMD5.ArrayBuffer();
		//获取文件二进制数据
		fileReader.readAsArrayBuffer(dataFile)
		// 下面要注意的是 fileReader.onload 回调方法是异步的，
		// 需要用到await 把它变成同步的，不然文件上传的时候是拿不到md5的值的
		await new Promise((resolve, reject) => {
			fileReader.onload = function(e) {
				spark.append(e.target.result.slice(0, 10 * 1024 * 1024));
				const md5 = spark.end()
				uploadData.value.md5 = md5;
				console.log(" 文件 md5 = ", md5);
				console.log(" 文件 uploadData.value.md5 = ", uploadData.value.md5);
				resolve()
			}
		})
	}
	// const handleGetByGuid = (guid) => {
	// 	page.value = 1
	// 	pageSize.value = 50
	// 	searchInfo.value.guid = guid
	// 	searchInfo.value.mediaType = undefined
	// 	searchInfo.value.module = undefined
	// 	getTableData()
	// }

	const handleImageSuccess = (res) => {
		console.log(" handleImageSuccess")
		console.log(res)

		var obj = new Object();
		obj["path"] = res.data.path;
		obj["guid"] = res.data.guid;
		emit('selectOneImg', obj)
	}

　


	const isImg = (path) => {
		if (!path || path == "")
			return false;

		if (path.indexOf(".jpg") != -1 || path.indexOf(".png") != -1) {
			//myListImg.value.push(getImgList(path))
			return true;
		} else
			return false;
	}
	const isAudio = (path) => {
		if (!path || path == "")
			return false;

		if (path.indexOf(".wav") != -1 || path.indexOf(".mp3") != -1)
			return true;
		else
			return false;
	}
	const isVideo = (path) => {
		if (!path || path == "")
			return false;

		if (path.indexOf("mp4") != -1 || path.indexOf(".mpg") != -1) {
			//this.haveVideo = true;
			return true;
		} else
			return false;
	}

　　

	const init = () => {
		console.log("init mediaLib props.max = ", props.max)
		console.log("init mediaLib props.mediaType = ", props.mediaType)

		if (props.max && props.max > 100) {
			maxSize.value = props.max
		}
		if (props.mediaType) {
			mediaType.value = props.mediaType
		}
		acceptType.value = ".jpg,.jpeg,.png,.mp4.mp3,.wav,.pdf,.zip,doc,docx"
		if (mediaType.value == 1) {
			acceptType.value = ".jpg,.jpeg,.png"
		} else if (mediaType.value == 2) {
			acceptType.value = ".mp4"
		} else if (mediaType.value == 3) {
			acceptType.value = ".mp3,.wav"
		} else if (mediaType.value == 3) {
			acceptType.value = ".pdf,.zip,doc,docx"
		}
		console.log("init mediaLib mediaType.value ", mediaType.value)
		console.log("init mediaLib acceptType.value ", acceptType.value)

		if (mediaType.value > 0) {
			searchInfo.value.mediaType = mediaType.value
		}

	}
	init()
	const open = async () => {
		　console.log("open----------")
		// showDialog.value = true
		// page.value = 1
		// pageSize.value = 50
		// getOptionsData()
		// getTreeData()
		// getTableData()
	}
	defineExpose({
		open
	})


</script>

<style scoped lang="scss">

</style>
