<template>
 <div>  
   <el-image class="image-div"  fit="fill" :src="myUrl" :preview-src-list="myList" hide-on-click-modal/> 
    <el-space :size="2" spacer=" "> 
         <el-link v-if="beEdit" icon="el-icon-edit"  @click="openChooseImg">上传</el-link>    
	     <el-link v-if="beEdit"  @click="deleteImg">删除<i class="el-icon-delete el-icon--right"></i></el-link> 
     </el-space>  
 </div>
 <template  v-if="beEdit">
 	 <MediaLib ref="mediaLib" @select-one-img="selectOneImg" />
 </template> 
 </template> 
 <script>
 	import {
 		ref
 	} from 'vue'
 	import {
 		mapGetters
 	} from 'vuex'
 	import MediaLib from '@/components/mediaLib/mediaLib.vue'
 	import {
 		isEmpty
 	} from '@/utils/utils'
 	// const path = import.meta.env.VITE_BASE_PATH+":"+ import.meta.env.VITE_SERVER_PORT "/"
 	const path = import.meta.env.VITE_BASE_API
 	export default {
 		name: 'ImageView',
 		components: {
 			MediaLib
 		},
 		props: {
 			type: {
 				type: String,
 				required: false,
 				default: 'image'
 			},
 			url: {
 				type: String,
 				required: false,
 				default: ''
 			},
 			guid: {
 				type: String,
 				required: false,
 				default: ''
 			},
 			beEdit: {
 				type: Boolean,
 				required: false,
 				default: false
 			}
 		},
 		data() {
 			return { 
 				myUrl: "",
 				myGuid: "",
 				myList: [] 
 			}
 		},
		async created() {		   
		    this.myUrl = this.url;
		    this.myList = this.getImgList(this.myUrl); 
		    this.myGuid = this.guid;
		},
		watch:{ 
			url(val){
				this.myUrl = val;
				this.myList = this.getImgList(val); 
			},
			guid(val){  
			   this.myGuid = val  
			}  
			/*
		    url:{ 
				handler(newVal, oldVal) {   
				  this.myUrl = newVal
				  this.myList =[newVal];
				  //console.log("this.myUrl");  console.log(this.myUrl); 
				},
				//immediate: true//初始化绑定时就会执行handler方法
			},
			guid:{ 
				handler(newVal, oldVal) {  
				  //console.log("guid -newVal: ", newVal, "guid -oldVal: ",oldVal);
				  this.myGuid = newVal 
				  //console.log("this.myGuid");  console.log(this.myGuid); 
			  },
			 // immediate: true//初始化绑定时就会执行handler方法
			} 
			*/
		  }, 
 		methods: {
 			openChooseImg() {
 				//console.log("1111");
 				this.$refs.mediaLib.open()
 			},
			deleteImg()
			{
				this.myUrl ="";
				this.myList =[];
				this.myGuid ="";	
			},
 			selectOneImg(obj) {
 				//console.log("selectOneImg");
				//console.log(obj);
				 this.myUrl = obj.url; 
				 this.myList = this.getImgList(obj.url);
				 this.myGuid = obj.guid;
 			},
			getImgList(path) {
			   //更换高清图
			   path = path.replace(".jpg","_src.jpg");
			   path = path.replace(".png","_src.png");
			   return [path];
			}
 		} 		 
		// async created() {
		//    console.log("created----------");
		//    this.urllist = [this.url]
		//    console.log(this.urllist);
		// },
 	}
 </script>
 <style scoped>
 	.image-div {
 		display: flex;
 		width: 90px;
 		height: 90px;
 	}
 	.imgtxt {
 		display: flex;
 		font-size: 14px;
 		margin-top: 0px;
 		cursor: pointer;
 		text-decoration: underline;
 		color: #409EFF
 	}
 </style>
