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
          <!--zzjlid  BeQuickEdit -->  
        <el-table-column label="转诊记录" prop="zzjlid" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zzjlid"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zzjlid',scope.row.ID,scope.row.zzjlid,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zzjlid}} </div>
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
          <!--xb  BeQuickEdit -->  
        <el-table-column label="性别" prop="xb" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xb"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xb',scope.row.ID,scope.row.xb,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xb}} </div>
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
        <el-table-column label="就诊时间" prop="jzrqsj" width="120"   sortable="custom" >
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
          <!--fzysgh  BeQuickEdit -->  
        <el-table-column label="负责医生工号" prop="fzysgh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fzysgh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fzysgh',scope.row.ID,scope.row.fzysgh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fzysgh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fzysxm  BeQuickEdit -->  
        <el-table-column label="负责医生姓名" prop="fzysxm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fzysxm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fzysxm',scope.row.ID,scope.row.fzysxm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fzysxm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fzksbm  BeQuickEdit -->  
        <el-table-column label="负责科室编码" prop="fzksbm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fzksbm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fzksbm',scope.row.ID,scope.row.fzksbm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fzksbm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fzksmc  BeQuickEdit -->  
        <el-table-column label="负责科室名称" prop="fzksmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fzksmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fzksmc',scope.row.ID,scope.row.fzksmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fzksmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zzyy  BeQuickEdit -->  
        <el-table-column label="转诊原因" prop="zzyy" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zzyy"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zzyy',scope.row.ID,scope.row.zzyy,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zzyy}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zzsj  BeQuickEdit -->  
        <el-table-column label="转诊时间" prop="zzsj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zzsj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zzsj',scope.row.ID,scope.row.zzsj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zzsj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zxjgdm  BeQuickEdit -->  
        <el-table-column label="转向机构代码" prop="zxjgdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zxjgdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zxjgdm',scope.row.ID,scope.row.zxjgdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zxjgdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zxjgmc  BeQuickEdit -->  
        <el-table-column label="转向机构名称" prop="zxjgmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zxjgmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zxjgmc',scope.row.ID,scope.row.zxjgmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zxjgmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zxksdm  BeQuickEdit -->  
        <el-table-column label="转向科室代码" prop="zxksdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zxksdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zxksdm',scope.row.ID,scope.row.zxksdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zxksdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zxksmc  BeQuickEdit -->  
        <el-table-column label="转向科室名称" prop="zxksmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zxksmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zxksmc',scope.row.ID,scope.row.zxksmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zxksmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zxysgh  BeQuickEdit -->  
        <el-table-column label="转向医生工号" prop="zxysgh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zxysgh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zxysgh',scope.row.ID,scope.row.zxysgh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zxysgh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zxysxm  BeQuickEdit -->  
        <el-table-column label="转向医生姓名" prop="zxysxm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zxysxm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zxysxm',scope.row.ID,scope.row.zxysxm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zxysxm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zyjws  BeQuickEdit -->  
        <el-table-column label="主要既往史" prop="zyjws" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zyjws"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zyjws',scope.row.ID,scope.row.zyjws,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zyjws}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zljg  BeQuickEdit -->  
        <el-table-column label="治疗经过" prop="zljg" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zljg"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zljg',scope.row.ID,scope.row.zljg,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zljg}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xybzlfa  BeQuickEdit -->  
        <el-table-column label="下一步治疗方案" prop="xybzlfa" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xybzlfa"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xybzlfa',scope.row.ID,scope.row.xybzlfa,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xybzlfa}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zzbz  BeQuickEdit -->  
        <el-table-column label="转诊标志" prop="zzbz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zzbz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zzbz',scope.row.ID,scope.row.zzbz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zzbz}} </div>
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
        <el-form-item label="转诊记录:"> 
              <el-input v-model="formData.zzjlid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="卡号:"> 
              <el-input v-model="formData.kh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="卡类型:"> 
              <el-input v-model="formData.klx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="门诊号:"> 
              <el-input v-model="formData.mzh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="姓名:"> 
              <el-input v-model="formData.xm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别:"> 
              <el-input v-model="formData.xb" clearable placeholder="请输入" />
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
        <el-form-item label="就诊时间:"> 
              <el-input v-model="formData.jzrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="负责医生工号:"> 
              <el-input v-model="formData.fzysgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="负责医生姓名:"> 
              <el-input v-model="formData.fzysxm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="负责科室编码:"> 
              <el-input v-model="formData.fzksbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="负责科室名称:"> 
              <el-input v-model="formData.fzksmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转诊原因:"> 
              <el-input v-model="formData.zzyy" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转诊时间:"> 
              <el-input v-model="formData.zzsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转向机构代码:"> 
              <el-input v-model="formData.zxjgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转向机构名称:"> 
              <el-input v-model="formData.zxjgmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转向科室代码:"> 
              <el-input v-model="formData.zxksdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转向科室名称:"> 
              <el-input v-model="formData.zxksmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转向医生工号:"> 
              <el-input v-model="formData.zxysgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转向医生姓名:"> 
              <el-input v-model="formData.zxysxm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="主要既往史:"> 
              <el-input v-model="formData.zyjws" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="治疗经过:"> 
              <el-input v-model="formData.zljg" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="下一步治疗方案:"> 
              <el-input v-model="formData.xybzlfa" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转诊标志:"> 
              <el-input v-model="formData.zzbz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数据生成时间:"> 
              <el-input v-model="formData.sjscsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="填报日期:"> 
              <el-input v-model="formData.tbrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="密级:"> 
              <el-input v-model="formData.mj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="撤销标志:"> 
              <el-input v-model="formData.cxbz" clearable placeholder="请输入" />
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
  createBbPiTreatmentReferral, 
  deleteBbPiTreatmentReferralByIds,
  updateBbPiTreatmentReferral,
  findBbPiTreatmentReferral,
  getBbPiTreatmentReferralList,
  quickEdit,
  excelList
} from '@/api/bbPiTreatmentReferral' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'BbPiTreatmentReferral',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getBbPiTreatmentReferralList,   
      excelListApi: excelList,
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           fwwddm: '',
           zzjlid: '',
           kh: '',
           klx: '',
           mzh: '',
           xm: '',
           xb: '',
           csrq: '',
           nls: '',
           nly: '',
           jzrqsj: '',
           fzysgh: '',
           fzysxm: '',
           fzksbm: '',
           fzksmc: '',
           zzyy: '',
           zzsj: '',
           zxjgdm: '',
           zxjgmc: '',
           zxksdm: '',
           zxksmc: '',
           zxysgh: '',
           zxysxm: '',
           zyjws: '',
           zljg: '',
           xybzlfa: '',
           zzbz: '',
           sjscsj: '',
           tbrqsj: '',
           mj: '',
           cxbz: '',
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
        //this.deleteBbPiTreatmentReferral(row)
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
     const res = await deleteBbPiTreatmentReferralByIds({ ids })
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
			this.$router.push({ name:'bbPiTreatmentReferralForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'bbPiTreatmentReferralForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findBbPiTreatmentReferral({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.bbPiTreatmentReferral 
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
          res = await createBbPiTreatmentReferral(this.formData);
          break
        case "update": 
          res = await updateBbPiTreatmentReferral(this.formData);
          break
        default: 
          res = await createBbPiTreatmentReferral(this.formData);
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
