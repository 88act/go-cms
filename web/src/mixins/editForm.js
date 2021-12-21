import { getDict } from '@/utils/dictionary'
import { isEmpty } from '@/utils/utils'
import { formatTimeToStr } from '@/utils/date' 
 export default {
	 
  data() {
    return {	 
      dialogFormVisible: false,//编辑对话框是否显示
	  editType:"create",
	  //type:"create",
	  deleteVisible:false,//删除时是否提示
	}
  },
  methods: {
      //统一关闭编辑对话框
	  closeDialog() {		  
		  this.dialogFormVisible = false
	      this.formData = new Object(); 
	   }
  }
	
}
