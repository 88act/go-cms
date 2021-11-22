import { getDict } from '@/utils/dictionary'
import { isEmpty } from '@/utils/utils'
import { formatTimeToStr } from '@/utils/date'
import { toSQLLine } from '@/utils/stringFun'
import { ref } from 'vue'
export default {   
  data() {
    return {
      page: 1,
      total: 20,
      pageSize: 20,
      tableData: [],
      searchInfo: {}, //查询
	  multipleSelection: [],//多选	 
	  shortcuts: [
	            {
	              text: '最近一周',
	              value: () => {
	                const end = new Date()
	                const start = new Date()
	                start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
	                return [start, end]
	              },
	            },
	            {
	              text: '最近一个月',
	              value: () => {
	                const end = new Date()
	                const start = new Date()
	                start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
	                return [start, end]
	              },
	            },
	            {
	              text: '最近三个月',
	              value: () => {
	                const end = new Date()
	                const start = new Date()
	                start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
	                return [start, end]
	              },
	       },
	  ],
    }
  },
  methods: {
	  //add by ljd 20210929 这里的代码混合到vue3 里面 ，跟js代码互不能访问
	 getMapData:function(key,map){  
		 let s = "";
        // console.log(map);
		 if (isEmpty(map)) return s;
		 if (isEmpty(key)) return s;
	     s = map[key];
	  	return  s;
	  }, 
    formatBoolean: function(bool) {
      if (bool !== null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    },
    formatDate: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    },
    filterDict(value, type) { 
	  //const rowLabel = this[type + 'Options'] && this[type + 'Options'].filter(item => item.value === value) 
     // by ljd 20210914   === 号 存在问题 ， 
	  //console.log(this[type + 'Options']);
	  const rowLabel = this[type + 'Options'] && this[type + 'Options'].filter(item => item.value == value) 
      return rowLabel && rowLabel[0] && rowLabel[0].label
    }, 
    async getDict(type) {
      const dicts = await getDict(type)
      this[type + 'Options'] = dicts
      return dicts
    },
    handleSizeChange(val) {
      this.pageSize = val
      this.getTableData()
    },
    handleCurrentChange(val) {
      this.page = val
      this.getTableData()
    },
     // 判断一个查询对象是否有空属性  add by ljd 20210718
	 obj_attr_is_null(obj){      
		  for (const key in obj) {
		    if (obj.hasOwnProperty(key)) {
		      if (obj[key] === null || obj[key] === '') {
		        console.log('为空='+key)
				delete obj[key]  
		      }else{
		        console.log('不为空')
		      }
		    }
		  }
	 },
    // @params beforeFunc function 请求发起前执行的函数 默认为空函数
    // @params afterFunc function 请求完成后执行的函数 默认为空函数
    async getTableData(beforeFunc = () => {}, afterFunc = () => {}) {
      beforeFunc()
	  //判断一个查询对象是否有空属性    add by ljd 20210718  
	  this.obj_attr_is_null(this.searchInfo);  
      const table = await this.listApi({ page: this.page, pageSize: this.pageSize, ...this.searchInfo })
      if (table.code === 200) {
        this.tableData = table.data.list
        this.total = table.data.total
        this.page = table.data.page
        this.pageSize = table.data.pageSize
      }
      afterFunc()
    },
    
    async getExcelList(page,pageSize) {  
      this.obj_attr_is_null(this.searchInfo);  
      const excel = await  this.excelListApi({ page: page, pageSize:pageSize, ...this.searchInfo }) 
	  if (excel.code === 200) {
	    this.downloadWithUrl(excel.data.url,excel.data.filename)
	  }
    },
	
	sortChange({ prop, order }) {
	  if (prop) {
	    this.searchInfo.orderKey = toSQLLine(prop)
	    this.searchInfo.orderDesc = order === 'descending'
	  }
	  this.getTableData()
	},
	 /**
	    * @param {string} url
	    * @param {string} fileName
	    * 通过链接下载文件
	    */
	   downloadWithUrl (url, fileName) {
	     const aLink = document.createElement('a');
	     aLink.style.display = 'none';
	     aLink.href = url;
	     aLink.download = fileName;
	     //aLink.target = '_parent';
		 aLink.target = '_blank';
	     document.body.appendChild(aLink);
	     aLink.click();
	     document.body.removeChild(aLink); // 下载完成移除元素
	   } 
	
  }
}
