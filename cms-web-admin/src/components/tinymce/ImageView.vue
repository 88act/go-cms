<template>
 <div>  
   <el-image class="image-div"  lazy fit="fill" :src="myUrl" :preview-src-list="myList" hide-on-click-modal/> 
   <el-link v-if="beEdit" icon="el-icon-edit" @click="openChooseImg">重新上传</el-link>
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
		    this.myList =[this.myUrl];
		    this.myGuid = this.guid;
		},
		watch: {
		    url (val) {  
		      this.myUrl = val
			  this.myList =[this.myUrl];
			 console.log("this.myUrl");  console.log(this.myUrl); 
		    }
		  }, 
 		methods: {
 			openChooseImg() {
 				console.log("1111");
 				this.$refs.mediaLib.open()
 			},
 			selectOneImg(obj) {
 				console.log("selectOneImg");
				console.log(obj);
				 this.myUrl = obj.url;
				 this.myList =[obj.url];
				 this.myGuid = obj.guid;
 			} 
 		} 
		//  watch: {
		//    url(val) {
		//      this.myUrl = val;
					 //   this.myList = [val]; 
		//    },
					 // guid(val) {
					 //   this.myGuid = val; 
					 // } 
		//  }, 
		
		// computed: {
		   //    getUrlList() {
		   //      return  [this.url]
		   //    }
		   // },
		// async created() {
		//    console.log("created----------");
		//    this.urllist = [this.url]
		//    console.log(this.urllist);
		// },
		// setup(props, context) {
		//     console.log('props:', {
		//       ...props,
		//     })
			// // urllist = [url]
			// // console.log(" setup urllist  === ");
			// //  console.log(urllist);
		//     // console.log('context.attrs:', {
		//     //   ...context.attrs,
		//     // })
		// },
		
 	}
 </script>
 <style scoped>
 	.image-div {
 		display: flex;
 		width: 80px;
 		height: 80px;
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
