<template>
	<el-space wrap>
		<div v-for="(obj, index) in myList" :key="obj">
			<div v-if="isImg(obj.path)" :style="{
									'width':props.imgSize+'px',
									'height':props.imgSize+'px',
									'padding':'4px'
								  }" >

				<el-image :src="obj.path" :style="{
									'width':props.imgSize+'px',
									'height':props.imgSize+'px',
									'border-radius':'4px'
								  }"  :initial-index=index
					:preview-src-list="myListImg" />
			</div>
			<div class="my-image-icon" v-else-if="isAudio(obj.path)" :style="{
									'background-image': `url(/imgadm/audio.png)`,
									'background-repeat': 'no-repeat',
									'background-size': 'cover',
									'width':props.imgSize+'px',
									'height':props.imgSize+'px'
								  }">
				<el-row :justify="center" :align="center" class="update">
					<el-col> </el-col>
					<el-col> <span style="font-size:16px;" @click="play(obj.path,1)">
							<el-icon>
								<Service />
							</el-icon>播放
						</span></el-col>
				</el-row>
			</div>
			<div class="my-image-icon" v-else="isVideo(obj.path)" :style="{
								'background-image': `url(/imgadm/video.png)`,
								'background-repeat': 'no-repeat',
								'background-size': 'cover',
				 }">
				<el-row :justify="center" :align="center" class="update">
					<el-col> </el-col>
					<el-col> <span style="font-size: 16px;" @click="play(obj.path,2)">
							<el-icon>
								<VideoPlay />
							</el-icon>播放
						</span></el-col>
				</el-row>
			</div>
		</div>

	</el-space>
	<!---------- 播放音视频------------------ -->
	<el-dialog width="360px" sappend-to-body="true" v-if="dialogVideo" v-model="dialogVideo"
		:show-close="false">
		<template #header="{ close, titleId, titleClass }">
			<div class="my-dialog-header">
				<span>播放音视频</span>
				<el-icon size="40" @click="closeDialogVideo" class="my-dialog-closeIcon">
					<CircleCloseFilled />
				</el-icon>
			</div>
		</template>
		 <audio v-if="audioVideo==1" controls width="250">
		      <source :src="audioVideoUrl"  type="audio/wav" />
		 </audio>
		 <video v-if="audioVideo==2" controls width="250">
		      <source :src="audioVideoUrl" type="video/mp4" />
		 </video>
	</el-dialog>
</template>

<script setup> 
	import {
		ref,
		watch,
		computed
	} from 'vue'
	import {
		isEmpty,
		obj2Num
	} from '@/utils/utils'
	const currentInstance = getCurrentInstance()

	// const path = import.meta.env.VITE_BASE_PATH+":"+ import.meta.env.VITE_SERVER_PORT "/"
	// 注意属性是只读的，不能修改
	const props = defineProps({
		urlList: {
			type: Array,
			required: false,
			default: []
		},
		objList: {
			type: Array,
			required: false,
			default: []
		},
		multi: {
			type: Boolean,
			required: false,
			default: false
		},
		max: {
			type: Number,
			required: false,
			default: 1024
		},
		mediaType: {
			type: Number,
			required: false,
			default: 0
		},
		imgSize: {
			type: Number,
			required: false,
			default: 60
		},
	})

	const myList = ref([])
	const myListImg = ref([])
	const noImg = ref(false)
	const maxSize = ref(1024)
	const mediaType = ref(0)



	const getImgList = (path) => {
		//更换高清图
		if (isEmpty(path)) {
			console.log("空路径 getImgList path=", path);
			return "";
		}
		path = path.replace(".jpg", "_src.jpg");
		path = path.replace(".png", "_src.png");
		return [path];

	}
	const isImg = (path) => {
		if (isEmpty(path))
			return false;

		if (path.indexOf(".jpg") != -1 || path.indexOf(".jpeg") != -1 || path.indexOf(".png") != -1) {
			myListImg.value.push(getImgList(path))
			return true;
		} else {

			return false;
		}

	}
	const isAudio = (path) => {
		if (isEmpty(path))
			return false;

		if (path.indexOf(".wav") != -1 || path.indexOf(".mp3") != -1)
			return true;
		else
			return false;
	}
	const isVideo = (path) => {
		if (isEmpty(path))
			return false;

		if (path.indexOf("mp4") != -1 || path.indexOf(".mpg") != -1) {
			//this.haveVideo = true;
			return true;
		} else
			return false;
	}



	const getGuidList = () => {
		let guids = []
		myList.value.forEach(function(value, index) {
			// console.log("forEach :", value)
			guids.push(value.guid)
		})
		return guids
	}

    // 音视频播放
	const dialogVideo = ref(false)
	const audioVideo = ref(0)
	const audioVideoUrl = ref("")
	function closeDialogVideo(){
		dialogVideo.value = false
		audioVideo.value = 0
	}
	function  play (path,va)  {
		audioVideoUrl.value = getImgList(path)
		dialogVideo.value = true
		audioVideo.value = va
		//window.open(path)
	}
	const init = () => {
		myList.value = []
		if (props.objList && props.objList.length > 0) {
			props.objList.forEach(function(value, index) {
				myList.value.push(value)
			});
		} else if (props.urlList && props.urlList.length > 0) {
			props.urlList.forEach(function(value, index) {
				let obj = new Object()
				obj.path = value
				obj.guid = index
				myList.value.push(obj)
			});

		} else
			noImg.value = true

		// console.log("myList.value :", myList.value)
	}
	init()

	watch(() => props.objList, (newValue, oldValue) => {
		console.log('objList 值变了', '变化后的值是' + newValue, '变化前的值是' + oldValue);
		myList.value = []
		if (props.objList && props.objList.length > 0) {
			props.objList.forEach(function(value, index) {
				myList.value.push(value)
			});
		}
	})

	watch(() => props.urlList, (newValue, oldValue) => {
		console.log('urlList 值变了', '变化后的值是' + newValue, '变化前的值是' + oldValue);
		myList.value = []
	   if (props.urlList && props.urlList.length > 0) {
			props.urlList.forEach(function(value, index) {
				let obj = new Object()
				obj.path = value
				obj.guid = index
				myList.value.push(obj)
			});
		}
	})


	defineExpose({
		getGuidList
	})
</script>
<style>


</style>
