<template> 
       <!-- <div :class="{fullscreen:fullscreen}" class="tinymce-container" style="width:1000px"> -->
            <textarea :id="tinymceId" v-model="content" :placeholder="placeholder" class="tinymce-textarea" />
       <!-- </div> -->  
	   	 <MediaLib ref="mediaLib" @select-one-img="selectOneImg" /> 
</template> 
<script> 
    import SparkMD5 from 'spark-md5'
	import sha1 from 'sha1-file-web';
	import { findBasicFile } from '@/api/basicFile'  
	import { ref } from 'vue'
    import load from './dynamicLoadScript.js'
	import { uploadFile } from '@/api/common_file'
    import MediaLib from '@/components/mediaLib/mediaLib.vue'
    const tinymce_path = window.location.origin + import.meta.env.VITE_ADMIN_WEB
	const tinymceCDN = tinymce_path+'/tinymce/tinymce.min.js'
	  
	  // const tinymceCDN = 'https://cdn.tiny.cloud/1/no-api-key/tinymce/5/tinymce.min.js'
	  // import plugins from './plugins.js'
	  // import toolbar from './toolbar.js'
    export default {
        name: "Tinymce",
		components: {
			MediaLib
		},
        props : {
            value: {
                type: String,
                default: ''
            },
            placeholder:{
                type: String,
                default: '请输入'
            },
        },
        data() {
            return {
                hasChange: false,
                hasInit: false,
                tinymceId : 'vue-tinymce-' + (Math.random()*1000).toFixed(0),
                fullscreen: false,
              //  mediaLibOpen : 0,
                content : '',
				
            }
        },
       // async created() {
       //      this.init();
       //  },
	   async mounted() {
	       this.init()
	   },
        watch: {
			// mediaLibOpen(val) {
			// 	 console.log("watch=====mediaLibOpen=====",val);  		  
			// 	 if (val ==1 ) this.openChooseImg();
			// },
            value(val) {
				this.content =val;
				if (this.hasInit)
				   this.setContent(val);
				//console.log("watch==========",val);  				
                // if (!this.hasChange && this.hasInit) {
                //     this.$nextTick(() =>
                //     window.tinymce.get(this.tinymceId).setContent(val || ''))
                // }
            } 
        },
        methods: {
            init() {
                // dynamic load tinymce from cdn
                load(tinymceCDN, (err) => {
                    if (err) {
                        console.error(err.message)
                        return
                    }
                    window.tinymce.baseURL = tinymce_path + '/tinymce'
                    this.initTinymce()
                })
            },
			openChooseImg() { 
				this.$refs.mediaLib.open()
			},
			selectOneImg(obj) {
				console.log("selectOneImg");
				console.log(obj);
			    window.tinymce.get(this.tinymceId).execCommand('mceInsertContent', false, '<img alt="Smiley face" height="80" width="80" src="'+obj.url+'"/>');
				
				 // this.myUrl = obj.url;
				 // this.myList =[obj.url];
				 // this.myGuid = obj.guid;
			},
            initTinymce() {
                const _this = this
                window.tinymce.init({
				   menubar: false, // 显示最上方menu菜单
                   language: 'zh_CN',
                   selector: `#${this.tinymceId}`,
                   height: 350, //高度	
                   end_container_on_empty_block: true,
                   powerpaste_word_import: 'clean',
                   code_dialog_height: 450,
                   code_dialog_width: 1000,
                   advlist_bullet_styles: 'square',
                   advlist_number_styles: 'default',
                   default_link_target: '_blank',
                   link_title: false,
                   nonbreaking_force_tab: true, // inserting nonbreaking space &nbsp; need Nonbreaking Space Plugin
				   branding: false, // 去水印
				   statusbar: false, // 隐藏编辑器底部的状态栏
				   elementpath: false, //禁用下角的当前标签路径
				   paste_data_images: true, // 允许粘贴图像  
				   plugin_preview_width:375, // 预览宽度 plugin_preview_height: 668,
				   plugin_preview_height:600,					  
                   toolbar: "undo redo  | formatselect alignleft aligncenter alignright alignjustify  indent outdent , \
                       | link unlink | numlist bullist |ImageLib image media table codesample | fontselect fontsizeselect forecolor backcolor | bold italic underline strikethrough | superscript subscript | removeformat |help code fullscreen preview",
                   //toolbar_drawer: "sliding",
                   quickbars_selection_toolbar: "removeformat | bold italic underline strikethrough | fontsizeselect forecolor backcolor",
                   plugins: ['preview link image media table lists fullscreen quickbars', 
                           'insertdatetime paste code help wordcount codesample'],
                   content_style: 'img {max-width:100% !important }',
                   // body_class: 'panel-body', 
                  // object_resizing: false,  //调整图片大小
                   //width: 1300, //高度
                  // toolbar: this.toolbar.length > 0 ? this.toolbar : toolbar,   	
                  // plugins: plugins,	
                                 
					 init_instance_callback: editor => {
						if (_this.value) {
							editor.setContent(_this.value)
						}
						_this.hasInit = true
						// editor.on('NodeChange Change KeyUp SetContent', () => {
						// 	this.hasChange = true
						// 	this.$emit('input', editor.getContent())
						// })
                    },
                    setup(editor) {
                        editor.on('FullscreenStateChanged', (e) => {
                            _this.fullscreen = e.state
                        });
						// // 注册一个icon
						editor.ui.registry.addIcon(
						  "imageLib",
						  `<svg viewBox="0 0 1024 1024" data-icon="shopping-cart" width="1.5em" height="1.5em" fill="currentColor" aria-hidden="true" focusable="false" class=""><path d="M922.9 701.9H327.4l29.9-60.9 496.8-.9c16.8 0 31.2-12 34.2-28.6l68.8-385.1c1.8-10.1-.9-20.5-7.5-28.4a34.99 34.99 0 0 0-26.6-12.5l-632-2.1-5.4-25.4c-3.4-16.2-18-28-34.6-28H96.5a35.3 35.3 0 1 0 0 70.6h125.9L246 312.8l58.1 281.3-74.8 122.1a34.96 34.96 0 0 0-3 36.8c6 11.9 18.1 19.4 31.5 19.4h62.8a102.43 102.43 0 0 0-20.6 61.7c0 56.6 46 102.6 102.6 102.6s102.6-46 102.6-102.6c0-22.3-7.4-44-20.6-61.7h161.1a102.43 102.43 0 0 0-20.6 61.7c0 56.6 46 102.6 102.6 102.6s102.6-46 102.6-102.6c0-22.3-7.4-44-20.6-61.7H923c19.4 0 35.3-15.8 35.3-35.3a35.42 35.42 0 0 0-35.4-35.2zM305.7 253l575.8 1.9-56.4 315.8-452.3.8L305.7 253zm96.9 612.7c-17.4 0-31.6-14.2-31.6-31.6 0-17.4 14.2-31.6 31.6-31.6s31.6 14.2 31.6 31.6a31.6 31.6 0 0 1-31.6 31.6zm325.1 0c-17.4 0-31.6-14.2-31.6-31.6 0-17.4 14.2-31.6 31.6-31.6s31.6 14.2 31.6 31.6a31.6 31.6 0 0 1-31.6 31.6z"></path></svg>`
						);
						// // 注册一个自定义的按钮
						editor.ui.registry.addButton("ImageLib", {
						  icon: `imageLib`,
						  onAction: function(_) { 
						    _this.openChooseImg(); 
						   }
						});
                    },
					images_upload_handler: (blobInfo, success, failure) => { 
					   let file =  blobInfo.blob();  
					    sha1(file).then(sha1 => { 
							 console.log("文件 sha1 = ", sha1)  
							 findBasicFile({sha1:sha1}).then(res => {
							    console.log("findBasicFile",res)
							 	if (res.code === 200 && res.data.basicFile && res.data.basicFile.path !="" ) { 
									let path = res.data.basicFile.path
									console.log(sha1,"已存在文件 ",path)
									//更换高清图
									path = path.replace(".jpg","_src.jpg")
									path = path.replace(".png","_src.png")
							 		success(path);
							 		return;
							 	}else {
									console.log(sha1,"不存在 开始上传.... ")
									let formData = new FormData()
									formData.append('file',file)   					   
									let ext ="jpg";
									if (file.type === 'image/jpeg')  ext ="jpg";
									if (file.type === 'image/png')  ext ="png"; 		
									formData.append('sha1',sha1)
									formData.append('ext',ext)
									formData.append('name',file.name)
									formData.append('size',file.size)
									formData.append('size',file.size)
									formData.append('media_type',2) 
									uploadFile(formData).then(res => {									  
										if (res.code === 200) {
											let path = res.data.path;
											//更换高清图
											path = path.replace(".jpg","_src.jpg")
											path = path.replace(".png","_src.png")
											success(path);
											return
										}
										console.log("上传出错1")
										console.log(res)
										failure(' 上传失败')
									}).catch(() => {
										console.log("上传出错2")
										failure('上传出错')
									}) 
								} 
							 }).catch(() => {
							 	failure('findBasicFile出错')
							 })
							 
						}).catch(() => {
							console.log('sha1(file) 出错')
						}) 
					  
					} 
                })
            },
			
	/*		
			language: "zh_CN", //语言设置
			height: 350, //高度
			menubar: false, // 显示最上方menu菜单
			toolbar: true, //false禁用工具栏（隐藏工具栏）
			browser_spellcheck: true, // 拼写检查
			branding: false, // 去水印
			statusbar: false, // 隐藏编辑器底部的状态栏
			elementpath: false, //禁用下角的当前标签路径
			paste_data_images: true, // 允许粘贴图像  
			plugin_preview_width:375, // 预览宽度 plugin_preview_height: 668,
			plugin_preview_height:600,  
			//theme_advanced_buttons3_add : "preview",                            
			toolbar: "undo redo  | formatselect alignleft aligncenter alignright alignjustify  indent outdent , \
			     | link unlink | numlist bullist | image media table codesample | fontselect fontsizeselect forecolor backcolor | bold italic underline strikethrough | superscript subscript | removeformat |help code fullscreen preview",
			toolbar_drawer: "sliding",
			quickbars_selection_toolbar: "removeformat | bold italic underline strikethrough | fontsizeselect forecolor backcolor",
			plugins: ['preview link image media table lists fullscreen quickbars', 
			         'insertdatetime paste code help wordcount codesample'],
			content_style: 'img {max-width:100% !important }',  
		*/	
			
            setContent(value) {
				console.log("this.tinymceId = ",this.tinymceId);
				let obj =window.tinymce.get(this.tinymceId);
				if (obj)
                   obj.setContent(value)
            },
            getContent() { 
                return window.tinymce.get(this.tinymceId).getContent();
            },
            destroyTinymce() {
                const tinymce = window.tinymce.get(this.tinymceId)

                if (tinymce) {
                    tinymce.destroy()
                }
            },
        },
    
        activated() {
            if (window.tinymce) {
                this.initTinymce()
            }
        },
        deactivated() {
            this.destroyTinymce()
        },
        destroyed() {
            this.destroyTinymce()
        }
    }
</script>