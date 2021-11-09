import tinymce from '@/components/tinymce/tinyEditor.vue'
import ImageView from '@/components/mediaLib/imageView.vue'
import MediaLib  from '@/components/mediaLib/mediaLib.vue'

 //import { uploadFile } from '@/api/common_file'  
 //import tinymce from '@tinymce/tinymce-vue' 
//import tinymce from '@/components/tinymce/tinyEditor.vue'

export default {
   components: {
       editor: tinymce,
	   ImageView:ImageView,
	   MediaLib:MediaLib
   } 
}
