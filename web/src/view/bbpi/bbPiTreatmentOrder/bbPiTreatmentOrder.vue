<template>
  <div>  
  <!----------查询form------------------ -->
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline"> 
      <el-form-item label="创建时间">
            <el-date-picker 
                  v-model="searchInfo.createdAtBetween" 
                  type="datetimerange"
                  format="YYYY-MM-DD HH:mm:ss"
                  :shortcuts="shortcuts"
                  range-separator="至"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
                />
              </el-form-item> 
        <el-form-item label="ID">
            <el-input placeholder="搜索ID" v-model="searchInfo.ID" />
        </el-form-item>
                <el-form-item label="上传" prop="status">                
                    <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_upOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSearch">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="goEditForm(0)">新增</el-button>
           <el-button size="mini" type="success"  @click="onExcel">导出当前</el-button>
         <el-button size="mini" type="success"  @click="onExcelAll">导出全部</el-button>
          <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
              <el-button icon="el-icon-delete" size="mini" type="danger" style="margin-left: 10px;">批量删除</el-button>
            </template>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
   <!----------数据表------------------ -->
    <el-table
      ref="multipleTable"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      :data="tableData"
      @selection-change="handleSelectionChange"
      @sort-change="sortChange" 
    >
      <el-table-column type="selection" width="55" />      
       <el-table-column label="ID" min-width="60" prop="ID" sortable="custom" />
          <!--status  BeQuickEdit -->
        <el-table-column label="上传" prop="status" width="120"  sortable="custom" >
        <template #default="scope">  
        <el-popover trigger="click" placement="top"  width = "280">  
              <el-select v-model="scope.row.status" placeholder="请选择"  @change="quickEdit_do('status',scope.row.ID,scope.row.status,scope)">
                  <el-option v-for="(item,key) in status_upOptions" :key="key" :label="item.label" :value="item.value"></el-option>
              </el-select> 
              <template #reference>
                  <div class="quickEdit" > {{filterDict(scope.row.status,"status_up")}} </div>
              </template>
            </el-popover>
        </template>  
        </el-table-column>
          <!--jgdm  BeQuickEdit -->  
        <el-table-column label="机构标识" prop="jgdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jgdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jgdm',scope.row.ID,scope.row.jgdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jgdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fwwddm  BeQuickEdit -->  
        <el-table-column label="服务网点代码" prop="fwwddm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fwwddm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fwwddm',scope.row.ID,scope.row.fwwddm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fwwddm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfbh  BeQuickEdit -->  
        <el-table-column label="处方编号" prop="cfbh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfbh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfbh',scope.row.ID,scope.row.cfbh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfbh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfmxid  BeQuickEdit -->  
        <el-table-column label="处方明细ID" prop="cfmxid" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfmxid"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfmxid',scope.row.ID,scope.row.cfmxid,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfmxid}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--kh  BeQuickEdit -->  
        <el-table-column label="卡号" prop="kh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.kh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('kh',scope.row.ID,scope.row.kh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.kh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--klx  BeQuickEdit -->  
        <el-table-column label="卡类型" prop="klx" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.klx"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('klx',scope.row.ID,scope.row.klx,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.klx}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfklsj  BeQuickEdit -->  
        <el-table-column label="处方开立日期" prop="cfklsj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfklsj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfklsj',scope.row.ID,scope.row.cfklsj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfklsj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfyxts  BeQuickEdit -->  
        <el-table-column label="处方有效天数" prop="cfyxts" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfyxts"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfyxts',scope.row.ID,scope.row.cfyxts,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfyxts}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfklksbm  BeQuickEdit -->  
        <el-table-column label="处方科室编码" prop="cfklksbm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfklksbm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfklksbm',scope.row.ID,scope.row.cfklksbm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfklksbm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfklksmc  BeQuickEdit -->  
        <el-table-column label="处方科室名称" prop="cfklksmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfklksmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfklksmc',scope.row.ID,scope.row.cfklksmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfklksmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--mzh  BeQuickEdit -->  
        <el-table-column label="门诊号" prop="mzh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.mzh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('mzh',scope.row.ID,scope.row.mzh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.mzh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xm  BeQuickEdit -->  
        <el-table-column label="姓名" prop="xm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xm',scope.row.ID,scope.row.xm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xbdm  BeQuickEdit -->  
        <el-table-column label="性别代码" prop="xbdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xbdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xbdm',scope.row.ID,scope.row.xbdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xbdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--csrq  BeQuickEdit -->  
        <el-table-column label="出生日期" prop="csrq" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.csrq"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('csrq',scope.row.ID,scope.row.csrq,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.csrq}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--nls  BeQuickEdit -->  
        <el-table-column label="年龄岁" prop="nls" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.nls"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('nls',scope.row.ID,scope.row.nls,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.nls}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--nly  BeQuickEdit -->  
        <el-table-column label="年龄月" prop="nly" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.nly"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('nly',scope.row.ID,scope.row.nly,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.nly}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jzrqsj  BeQuickEdit -->  
        <el-table-column label="就诊日期时间" prop="jzrqsj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jzrqsj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jzrqsj',scope.row.ID,scope.row.jzrqsj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jzrqsj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--yzsm  BeQuickEdit -->  
        <el-table-column label="医嘱说明" prop="yzsm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.yzsm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('yzsm',scope.row.ID,scope.row.yzsm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.yzsm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--pxh  BeQuickEdit -->  
        <el-table-column label="排序号" prop="pxh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.pxh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('pxh',scope.row.ID,scope.row.pxh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.pxh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--yzxmlxdm  BeQuickEdit -->  
        <el-table-column label="医嘱项目类型代码" prop="yzxmlxdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.yzxmlxdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('yzxmlxdm',scope.row.ID,scope.row.yzxmlxdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.yzxmlxdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--ypcfsx  BeQuickEdit -->  
        <el-table-column label="药品处方属性" prop="ypcfsx" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ypcfsx"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ypcfsx',scope.row.ID,scope.row.ypcfsx,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ypcfsx}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zylbdm  BeQuickEdit -->  
        <el-table-column label="中药类别代码" prop="zylbdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zylbdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zylbdm',scope.row.ID,scope.row.zylbdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zylbdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfmxxmbm  BeQuickEdit -->  
        <el-table-column label="处方明细医保编码" prop="cfmxxmbm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfmxxmbm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfmxxmbm',scope.row.ID,scope.row.cfmxxmbm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfmxxmbm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--ypid  BeQuickEdit -->  
        <el-table-column label="药监药品ID" prop="ypid" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ypid"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ypid',scope.row.ID,scope.row.ypid,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ypid}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfmxmc  BeQuickEdit -->  
        <el-table-column label="处方明细名称" prop="cfmxmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfmxmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfmxmc',scope.row.ID,scope.row.cfmxmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfmxmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--ywmc  BeQuickEdit -->  
        <el-table-column label="药物名称" prop="ywmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ywmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ywmc',scope.row.ID,scope.row.ywmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ywmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--ypgg  BeQuickEdit -->  
        <el-table-column label="药品规格" prop="ypgg" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ypgg"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ypgg',scope.row.ID,scope.row.ypgg,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ypgg}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--dddz  BeQuickEdit -->  
        <el-table-column label="DDD值" prop="dddz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.dddz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('dddz',scope.row.ID,scope.row.dddz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.dddz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--ywjxdm  BeQuickEdit -->  
        <el-table-column label="药物剂型代码" prop="ywjxdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ywjxdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ywjxdm',scope.row.ID,scope.row.ywjxdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ywjxdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--ywsycjl  BeQuickEdit -->  
        <el-table-column label="药物使用次剂量" prop="ywsycjl" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ywsycjl"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ywsycjl',scope.row.ID,scope.row.ywsycjl,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ywsycjl}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--ywsyjldw  BeQuickEdit -->  
        <el-table-column label="药物使用剂量单位" prop="ywsyjldw" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ywsyjldw"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ywsyjldw',scope.row.ID,scope.row.ywsyjldw,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ywsyjldw}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--ywsypcdm  BeQuickEdit -->  
        <el-table-column label="药物使用频次代码" prop="ywsypcdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ywsypcdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ywsypcdm',scope.row.ID,scope.row.ywsypcdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ywsypcdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--ywsypcmc  BeQuickEdit -->  
        <el-table-column label="药物使用频次名称" prop="ywsypcmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ywsypcmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ywsypcmc',scope.row.ID,scope.row.ywsypcmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ywsypcmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--yytjdm  BeQuickEdit -->  
        <el-table-column label="用药途径代码" prop="yytjdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.yytjdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('yytjdm',scope.row.ID,scope.row.yytjdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.yytjdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--yytjmc  BeQuickEdit -->  
        <el-table-column label="用药途径名称" prop="yytjmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.yytjmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('yytjmc',scope.row.ID,scope.row.yytjmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.yytjmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--ywsyzjl  BeQuickEdit -->  
        <el-table-column label="药物使用总剂量" prop="ywsyzjl" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ywsyzjl"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ywsyzjl',scope.row.ID,scope.row.ywsyzjl,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ywsyzjl}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfypzh  BeQuickEdit -->  
        <el-table-column label="处方药品组号" prop="cfypzh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfypzh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfypzh',scope.row.ID,scope.row.cfypzh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfypzh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zyypcf  BeQuickEdit -->  
        <el-table-column label="中药饮片处方" prop="zyypcf" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zyypcf"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zyypcf',scope.row.ID,scope.row.zyypcf,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zyypcf}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zyypjs  BeQuickEdit -->  
        <el-table-column label="中药饮片剂数" prop="zyypjs" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zyypjs"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zyypjs',scope.row.ID,scope.row.zyypjs,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zyypjs}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zyypjzf  BeQuickEdit -->  
        <el-table-column label="中药饮片煎煮法" prop="zyypjzf" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zyypjzf"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zyypjzf',scope.row.ID,scope.row.zyypjzf,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zyypjzf}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zyyyff  BeQuickEdit -->  
        <el-table-column label="中药用药方法" prop="zyyyff" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zyyyff"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zyyyff',scope.row.ID,scope.row.zyyyff,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zyyyff}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fyjl  BeQuickEdit -->  
        <el-table-column label="发药剂量" prop="fyjl" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fyjl"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fyjl',scope.row.ID,scope.row.fyjl,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fyjl}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fyjldw  BeQuickEdit -->  
        <el-table-column label="发药剂量单位" prop="fyjldw" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fyjldw"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fyjldw',scope.row.ID,scope.row.fyjldw,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fyjldw}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--dj  BeQuickEdit -->  
        <el-table-column label="单价" prop="dj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.dj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('dj',scope.row.ID,scope.row.dj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.dj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zje  BeQuickEdit -->  
        <el-table-column label="总金额" prop="zje" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zje"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zje',scope.row.ID,scope.row.zje,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zje}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfklysgh  BeQuickEdit -->  
        <el-table-column label="处方开立医师工号" prop="cfklysgh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfklysgh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfklysgh',scope.row.ID,scope.row.cfklysgh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfklysgh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfklysqm  BeQuickEdit -->  
        <el-table-column label="处方开立医师签名" prop="cfklysqm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfklysqm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfklysqm',scope.row.ID,scope.row.cfklysqm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfklysqm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfshyjsgh  BeQuickEdit -->  
        <el-table-column label="处方审核工号" prop="cfshyjsgh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfshyjsgh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfshyjsgh',scope.row.ID,scope.row.cfshyjsgh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfshyjsgh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfshyjsqm  BeQuickEdit -->  
        <el-table-column label="处方审核签名" prop="cfshyjsqm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfshyjsqm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfshyjsqm',scope.row.ID,scope.row.cfshyjsqm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfshyjsqm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cftpyjsgh  BeQuickEdit -->  
        <el-table-column label="处方调配工号" prop="cftpyjsgh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cftpyjsgh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cftpyjsgh',scope.row.ID,scope.row.cftpyjsgh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cftpyjsgh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cftpyjsqm  BeQuickEdit -->  
        <el-table-column label="处方调配签名" prop="cftpyjsqm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cftpyjsqm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cftpyjsqm',scope.row.ID,scope.row.cftpyjsqm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cftpyjsqm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfhdyjsgh  BeQuickEdit -->  
        <el-table-column label="处方核对工号" prop="cfhdyjsgh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfhdyjsgh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfhdyjsgh',scope.row.ID,scope.row.cfhdyjsgh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfhdyjsgh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cfhdyjsqm  BeQuickEdit -->  
        <el-table-column label="处方核对签名" prop="cfhdyjsqm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cfhdyjsqm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cfhdyjsqm',scope.row.ID,scope.row.cfhdyjsqm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cfhdyjsqm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cffyyjsgh  BeQuickEdit -->  
        <el-table-column label="处方发药工号" prop="cffyyjsgh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cffyyjsgh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cffyyjsgh',scope.row.ID,scope.row.cffyyjsgh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cffyyjsgh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cffyyjsqm  BeQuickEdit -->  
        <el-table-column label="处方发药签名" prop="cffyyjsqm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cffyyjsqm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cffyyjsqm',scope.row.ID,scope.row.cffyyjsqm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cffyyjsqm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zxjg  BeQuickEdit -->  
        <el-table-column label="执行结果" prop="zxjg" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zxjg"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zxjg',scope.row.ID,scope.row.zxjg,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zxjg}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--bz  BeQuickEdit -->  
        <el-table-column label="备注" prop="bz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.bz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('bz',scope.row.ID,scope.row.bz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.bz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--qyjgdm  BeQuickEdit -->  
        <el-table-column label="取药机构代码" prop="qyjgdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.qyjgdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('qyjgdm',scope.row.ID,scope.row.qyjgdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.qyjgdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--qyjgmc  BeQuickEdit -->  
        <el-table-column label="取药机构名称" prop="qyjgmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.qyjgmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('qyjgmc',scope.row.ID,scope.row.qyjgmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.qyjgmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sjscsj  BeQuickEdit -->  
        <el-table-column label="数据生成时间" prop="sjscsj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sjscsj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sjscsj',scope.row.ID,scope.row.sjscsj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sjscsj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--tbrqsj  BeQuickEdit -->  
        <el-table-column label="填报日期" prop="tbrqsj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.tbrqsj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('tbrqsj',scope.row.ID,scope.row.tbrqsj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.tbrqsj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cxbz  BeQuickEdit -->  
        <el-table-column label="撤销标志" prop="cxbz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cxbz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cxbz',scope.row.ID,scope.row.cxbz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cxbz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--mj  BeQuickEdit -->  
        <el-table-column label="密级" prop="mj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.mj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('mj',scope.row.ID,scope.row.mj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.mj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column> 
      <el-table-column label="日期" width="180" prop="created_at" sortable="custom" >
        <template #default="scope">{{ formatDate(scope.row.CreatedAt)}}</template>
      </el-table-column>
      
      <el-table-column label="操作">
        <template #default="scope">
          <el-button plain size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="goEditForm(scope.row.ID)">编辑</el-button>
          <el-button plain size="mini" type="danger" icon="el-icon-delete"  @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10,20,50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <!---------- 编辑弹窗------------------ -->
    <el-dialog  v-if="dialogFormVisible"  :before-close="closeDialog" v-model="dialogFormVisible" title="编辑资料">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="上传:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_upOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="机构标识:"> 
              <el-input v-model="formData.jgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务网点代码:"> 
              <el-input v-model="formData.fwwddm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方编号:"> 
              <el-input v-model="formData.cfbh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方明细ID:"> 
              <el-input v-model="formData.cfmxid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="卡号:"> 
              <el-input v-model="formData.kh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="卡类型:"> 
              <el-input v-model="formData.klx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方开立日期:"> 
              <el-input v-model="formData.cfklsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方有效天数:"> 
              <el-input v-model="formData.cfyxts" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方科室编码:"> 
              <el-input v-model="formData.cfklksbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方科室名称:"> 
              <el-input v-model="formData.cfklksmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="门诊号:"> 
              <el-input v-model="formData.mzh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="姓名:"> 
              <el-input v-model="formData.xm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别代码:"> 
              <el-input v-model="formData.xbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出生日期:"> 
              <el-input v-model="formData.csrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="年龄岁:"> 
              <el-input v-model="formData.nls" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="年龄月:"> 
              <el-input v-model="formData.nly" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="就诊日期时间:"> 
              <el-input v-model="formData.jzrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医嘱说明:"> 
              <el-input v-model="formData.yzsm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="排序号:"> 
              <el-input v-model="formData.pxh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医嘱项目类型代码:"> 
              <el-input v-model="formData.yzxmlxdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药品处方属性:"> 
              <el-input v-model="formData.ypcfsx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中药类别代码:"> 
              <el-input v-model="formData.zylbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方明细医保编码:"> 
              <el-input v-model="formData.cfmxxmbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药监药品ID:"> 
              <el-input v-model="formData.ypid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方明细名称:"> 
              <el-input v-model="formData.cfmxmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物名称:"> 
              <el-input v-model="formData.ywmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药品规格:"> 
              <el-input v-model="formData.ypgg" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="DDD值:"> 
              <el-input v-model="formData.dddz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物剂型代码:"> 
              <el-input v-model="formData.ywjxdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物使用次剂量:"> 
              <el-input v-model="formData.ywsycjl" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物使用剂量单位:"> 
              <el-input v-model="formData.ywsyjldw" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物使用频次代码:"> 
              <el-input v-model="formData.ywsypcdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物使用频次名称:"> 
              <el-input v-model="formData.ywsypcmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用药途径代码:"> 
              <el-input v-model="formData.yytjdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用药途径名称:"> 
              <el-input v-model="formData.yytjmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物使用总剂量:"> 
              <el-input v-model="formData.ywsyzjl" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方药品组号:"> 
              <el-input v-model="formData.cfypzh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中药饮片处方:"> 
              <el-input v-model="formData.zyypcf" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中药饮片剂数:"> 
              <el-input v-model="formData.zyypjs" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中药饮片煎煮法:"> 
              <el-input v-model="formData.zyypjzf" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中药用药方法:"> 
              <el-input v-model="formData.zyyyff" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="发药剂量:"> 
              <el-input v-model="formData.fyjl" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="发药剂量单位:"> 
              <el-input v-model="formData.fyjldw" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="单价:"> 
              <el-input v-model="formData.dj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总金额:"> 
              <el-input v-model="formData.zje" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方开立医师工号:"> 
              <el-input v-model="formData.cfklysgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方开立医师签名:"> 
              <el-input v-model="formData.cfklysqm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方审核工号:"> 
              <el-input v-model="formData.cfshyjsgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方审核签名:"> 
              <el-input v-model="formData.cfshyjsqm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方调配工号:"> 
              <el-input v-model="formData.cftpyjsgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方调配签名:"> 
              <el-input v-model="formData.cftpyjsqm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方核对工号:"> 
              <el-input v-model="formData.cfhdyjsgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方核对签名:"> 
              <el-input v-model="formData.cfhdyjsqm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方发药工号:"> 
              <el-input v-model="formData.cffyyjsgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方发药签名:"> 
              <el-input v-model="formData.cffyyjsqm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执行结果:"> 
              <el-input v-model="formData.zxjg" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:"> 
              <el-input v-model="formData.bz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="取药机构代码:"> 
              <el-input v-model="formData.qyjgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="取药机构名称:"> 
              <el-input v-model="formData.qyjgmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数据生成时间:"> 
              <el-input v-model="formData.sjscsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="填报日期:"> 
              <el-input v-model="formData.tbrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="撤销标志:"> 
              <el-input v-model="formData.cxbz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="密级:"> 
              <el-input v-model="formData.mj" clearable placeholder="请输入" />
       </el-form-item>
     </el-form>
      <div slot="footer" class="el-dialog__footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="saveEditForm">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createBbPiTreatmentOrder, 
  deleteBbPiTreatmentOrderByIds,
  updateBbPiTreatmentOrder,
  findBbPiTreatmentOrder,
  getBbPiTreatmentOrderList,
  quickEdit,
  excelList
} from '@/api/bbPiTreatmentOrder' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'BbPiTreatmentOrder',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getBbPiTreatmentOrderList,   
      excelListApi: excelList,
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           fwwddm: '',
           cfbh: '',
           cfmxid: '',
           kh: '',
           klx: '',
           cfklsj: '',
           cfyxts: '',
           cfklksbm: '',
           cfklksmc: '',
           mzh: '',
           xm: '',
           xbdm: '',
           csrq: '',
           nls: '',
           nly: '',
           jzrqsj: '',
           yzsm: '',
           pxh: '',
           yzxmlxdm: '',
           ypcfsx: '',
           zylbdm: '',
           cfmxxmbm: '',
           ypid: '',
           cfmxmc: '',
           ywmc: '',
           ypgg: '',
           dddz: '',
           ywjxdm: '',
           ywsycjl: '',
           ywsyjldw: '',
           ywsypcdm: '',
           ywsypcmc: '',
           yytjdm: '',
           yytjmc: '',
           ywsyzjl: '',
           cfypzh: '',
           zyypcf: '',
           zyypjs: '',
           zyypjzf: '',
           zyyyff: '',
           fyjl: '',
           fyjldw: '',
           dj: '',
           zje: '',
           cfklysgh: '',
           cfklysqm: '',
           cfshyjsgh: '',
           cfshyjsqm: '',
           cftpyjsgh: '',
           cftpyjsqm: '',
           cfhdyjsgh: '',
           cfhdyjsqm: '',
           cffyyjsgh: '',
           cffyyjsqm: '',
           zxjg: '',
           bz: '',
           qyjgdm: '',
           qyjgmc: '',
           sjscsj: '',
           tbrqsj: '',
           cxbz: '',
           mj: '',
            mapData: {}
      } 
    }
  },
  
  async created() {
    await this.getDict('status_up')
    await this.getTableData()
  },
  methods: { 
    onSearch() {
      this.page = 1
      this.pageSize = 20 
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
         const ids = [row.ID] 
         this.doDelete(ids); 
        //this.deleteBbPiTreatmentOrder(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      this.doDelete(ids); 
    },
  	async doDelete(ids) { 
     const res = await deleteBbPiTreatmentOrderByIds({ ids })
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      } 
	},  
   async goEditForm(id) { 
	  if (this.beNewWindow) {
		  if (id >0) {
			this.$router.push({ name:'bbPiTreatmentOrderForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'bbPiTreatmentOrderForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findBbPiTreatmentOrder({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.bbPiTreatmentOrder 
		 }else
		 {
			this.editType = 'create' 
		 }
		  this.dialogFormVisible = true
	  }
	}, 
    async saveEditForm() {  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case "create":         
          res = await createBbPiTreatmentOrder(this.formData);
          break
        case "update": 
          res = await updateBbPiTreatmentOrder(this.formData);
          break
        default: 
          res = await createBbPiTreatmentOrder(this.formData);
          break
      }
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },   
    quickEdit_do(field,id,value,scope) {    
      let value2 = value;
      if (typeof(value)==="boolean")
         value2 = value?"1":"0"
      value2 =  value2+"";   
      let obj = {field:field,id:id,value:value2}	 
      this.quickEdit_do2(obj);	  
	    // if (scope._self.$refs[`popover-${scope.$index}`])
		  // scope._self.$refs[`popover-${scope.$index}`].doClose();
    },
    async quickEdit_do2(obj) {  
      const res =  await quickEdit(obj)	  
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '修改成功'
        }) 
        // this.getTableData()
      }
    },
   onExcel(){
        this.getExcelList(this.page,this.pageSize)    
    },
    onExcelAll(){
        this.getExcelList(1,1000)  
    }
  },
}
</script>
<style>
</style>
