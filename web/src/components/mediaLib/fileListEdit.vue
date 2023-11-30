<template>
	<el-space wrap>
		<div v-for="(obj, index) in myList" :key="obj">
			<!-- <el-image v-if="isImg(obj.url)" :src="obj.url" fit="cover"    :preview-src-list="myListImg" />
			<audio v-else-if="isAudio(obj.url)" :src="obj.url" controls style="width:100px" />
			<video v-else="isVideo(obj.url)" :src="obj.url" controls="false" style="width:100px" />
		    <el-link @click="deleteImg(obj.guid)">删除<i class="el-icon-delete el-icon--right"></i></el-link> -->
			<div class="my-image-icon-edit" v-if="isImg(obj.path)" :style="{
					'background-image': `url(${obj.path})`,
					'background-repeat': 'no-repeat',
          'background-size': '100% 100%',
          'width':getSize+'px',
          'height':getSize+'px'
				  }">
        <div style="display: flex; flex-wrap: wrap; justify-content: center; margin-left: 20px; width:80px ;height: 120px;">
            <div> <span style="font-size: 20px;cursor: pointer;padding-bottom: 40px;"
                @click="deleteImg(obj.guid)">
                <el-icon>
                  <delete />
                </el-icon>删除
              </span></div><div>
             <span style="font-size: 20px;cursor: pointer;" @click="play(obj.path)">
                  <el-icon>
                    <search />
                  </el-icon>查看
                </span></div>
          </div>
			</div>

			<div class="my-image-icon-edit" v-else-if="isAudio(obj.path)" :style="{
									'background-image': `url(/img/audio.png)`,
									'background-repeat': 'no-repeat',
								   'background-size': '100% 100%',
                  'width':getSize+'px',
                  'height':getSize+'px'
								  }">
				<el-row :justify="center" :align="center" class="update">
					<el-col> <span style="font-size: 20px;padding-bottom: 10px;" @click="deleteImg(obj.guid)">
							<el-icon>
								<delete />
							</el-icon>删除
						</span> </el-col>
					<el-col> <span style="font-size: 20px;" @click="play(obj.path)">
							<el-icon><Service /></el-icon>播放
						</span></el-col>
				</el-row>
			</div>
			<div class="my-image-icon-edit" v-else="isVideo(obj.path)" :style="{
								'background-image': `url(/img/video.png)`,
								'background-repeat': 'no-repeat',
							'background-size': '100% 100%',
							'width':getSize+'px',
							'height':getSize+'px'
				 }">
				<el-row :justify="center" :align="center" class="update">
					<el-col> <span style="font-size: 20px;padding-bottom: 10px;" @click="deleteImg(obj.guid)">
							<el-icon>
								<delete />
							</el-icon>删除
						</span> </el-col>
					<el-col> <span style="font-size: 20px;" @click="play(obj.path)">
							<el-icon><VideoPlay /></el-icon>播放
						</span></el-col>
				</el-row>
			</div>
		</div>

	<!-- 	<div v-if="isMulti()" class="my-image-icon-edit" @click="openChooseImg" :style="{
				'background-image': `url(/img/addimg.png)`,
				'background-repeat': 'no-repeat',
				'background-size': 'cover',
			  }">
		</div> -->
  <MediaLibUp v-if="isMulti"  ref="mediaLibUp" :max="maxSize" :mediaType="mediaType" @selectOneImg="selectOneImg" />
	</el-space>

</template>

<script setup>
	import MediaLibUp from '@/components/mediaLib/mediaLibUp.vue'
	// 这里的ref(null) = this.$refs vue3中使用refs需要在底部声明
	const mediaLibUp = ref(null)
	import {
		ref,
		getCurrentInstance,
		reactive,
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

	const getSize = computed(() => {
	 		return props.imgSize;
	})
	const isMulti = computed(() => {

	 		 if (props.multi)
	 		 	return true;

	 		 if (!props.multi && myList.value.length == 0)
	 		 	return true;
	 		 else
	 		 	return false;
	})

	const openChooseImg = () => {
		console.log("openChooseImg  mediaLib.value.open");
		// if (props.beEdit)
		 console.log(currentInstance.refs.mediaLibUp);
		mediaLibUp.value.open()
		// myList.value.push(value.url)
		// this.$refs.mediaLib.open()
	}
	const deleteImg = (guid) => {
		console.log("deleteImg guid=", guid);
		for (var i = 0; i < myList.value.length; i++) {
			if (myList.value[i].guid == guid) {
				myList.value.splice(i, 1)
				return
			}
		}
	}
	const selectOneImg = (obj) => {
		console.log("接收 ----  selectOneImg");
		console.log(obj);
		myList.value.push(obj)
		console.log(myList.value);
		// // this.myUrl = obj.url;
		// // this.myList = this.getImgList(obj.url);
		// // 数组
		// //myList.value =[obj.guid];
		// //myGuid.value = obj.guid;
	}

	const getGuidList = () => {
		let guids = []
		myList.value.forEach(function(value, index) {
			// console.log("forEach :", value)
			guids.push(value.guid)
		})
		return guids
	}

	const play = (path) => {
		path = getImgList(path)
		window.open(path)
	}
	const init = () => {
		myList.value = []
		 console.log("init :", props.objList)
		if (props.objList && props.objList.length > 0) {
			props.objList.forEach(function(value, index) {
				console.log("forEach :", value)
				myList.value.push(value)
			});
		} else {
			noImg.value = true
		}
		if (props.max && props.max > 100) {
			maxSize.value = props.max
		}
		if (props.mediaType ) {
			mediaType.value = props.mediaType
		}
		console.log("init fileListEdit maxSize.value = ", maxSize.value)
	    console.log("init fileListEdit mediaType.value = ", mediaType.value)

	}
	init()

	watch(() => props.objList, (newValue, oldValue) => {
		    // myList.value = []
			 console.log('objList 值变了', '变化后的值是'+newValue, '变化前的值是'+oldValue);
			 if (props.objList && props.objList.length > 0) {
			 	props.objList.forEach(function(value, index) {
			 		//console.log("forEach :", value)
			 		myList.value.push(value)
			 	});
			 } else {
			  	noImg.value = true
			 }
	})

	defineExpose({
		getGuidList
	})
</script>
<style>


</style>
