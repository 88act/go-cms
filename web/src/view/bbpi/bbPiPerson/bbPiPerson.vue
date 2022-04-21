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
                  <el-form-item label="userid">
                      <el-input placeholder="搜索条件" v-model="searchInfo.userid" clearable />
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
          <el-table-column label="userid" prop="userid" width="120"   sortable="custom"  />
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
          <!--zjhm  BeQuickEdit -->  
        <el-table-column label="身份证件号码" prop="zjhm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zjhm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zjhm',scope.row.ID,scope.row.zjhm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zjhm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zjlbdm  BeQuickEdit -->  
        <el-table-column label="身份证件类别" prop="zjlbdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zjlbdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zjlbdm',scope.row.ID,scope.row.zjlbdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zjlbdm}} </div>
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
          <!--xbmc  BeQuickEdit -->  
        <el-table-column label="性别名称" prop="xbmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xbmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xbmc',scope.row.ID,scope.row.xbmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xbmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--hzsx  BeQuickEdit -->  
        <el-table-column label="患者属性" prop="hzsx" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.hzsx"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('hzsx',scope.row.ID,scope.row.hzsx,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.hzsx}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--hyztdm  BeQuickEdit -->  
        <el-table-column label="婚姻状况代码" prop="hyztdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.hyztdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('hyztdm',scope.row.ID,scope.row.hyztdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.hyztdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--hyztmc  BeQuickEdit -->  
        <el-table-column label="婚姻状态名称" prop="hyztmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.hyztmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('hyztmc',scope.row.ID,scope.row.hyztmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.hyztmc}} </div>
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
          <!--mzdm  BeQuickEdit -->  
        <el-table-column label="民族代码" prop="mzdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.mzdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('mzdm',scope.row.ID,scope.row.mzdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.mzdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--mzmc  BeQuickEdit -->  
        <el-table-column label="民族名称" prop="mzmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.mzmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('mzmc',scope.row.ID,scope.row.mzmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.mzmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--gjdm  BeQuickEdit -->  
        <el-table-column label="国籍代码" prop="gjdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.gjdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('gjdm',scope.row.ID,scope.row.gjdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.gjdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--gjmc  BeQuickEdit -->  
        <el-table-column label="国籍名称" prop="gjmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.gjmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('gjmc',scope.row.ID,scope.row.gjmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.gjmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jgnbdah  BeQuickEdit -->  
        <el-table-column label="机构内部档案号" prop="jgnbdah" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jgnbdah"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jgnbdah',scope.row.ID,scope.row.jgnbdah,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jgnbdah}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--gddhhm  BeQuickEdit -->  
        <el-table-column label="固定电话" prop="gddhhm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.gddhhm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('gddhhm',scope.row.ID,scope.row.gddhhm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.gddhhm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sjhm  BeQuickEdit -->  
        <el-table-column label="手机号码" prop="sjhm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sjhm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sjhm',scope.row.ID,scope.row.sjhm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sjhm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--dzyj  BeQuickEdit -->  
        <el-table-column label="电子邮件" prop="dzyj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.dzyj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('dzyj',scope.row.ID,scope.row.dzyj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.dzyj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--whcddm  BeQuickEdit -->  
        <el-table-column label="文化程度代码" prop="whcddm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.whcddm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('whcddm',scope.row.ID,scope.row.whcddm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.whcddm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--whcdmc  BeQuickEdit -->  
        <el-table-column label="文化程度名称" prop="whcdmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.whcdmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('whcdmc',scope.row.ID,scope.row.whcdmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.whcdmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zylbdm  BeQuickEdit -->  
        <el-table-column label="职业类别代码" prop="zylbdm" width="120"   sortable="custom" >
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
          <!--zylbmc  BeQuickEdit -->  
        <el-table-column label="职业类别名称" prop="zylbmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zylbmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zylbmc',scope.row.ID,scope.row.zylbmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zylbmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--csdxzqhm  BeQuickEdit -->  
        <el-table-column label="出生地行政区划码" prop="csdxzqhm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.csdxzqhm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('csdxzqhm',scope.row.ID,scope.row.csdxzqhm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.csdxzqhm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--csd  BeQuickEdit -->  
        <el-table-column label="出生地" prop="csd" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.csd"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('csd',scope.row.ID,scope.row.csd,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.csd}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jzdxzqhm  BeQuickEdit -->  
        <el-table-column label="居住地行政区划码" prop="jzdxzqhm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jzdxzqhm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jzdxzqhm',scope.row.ID,scope.row.jzdxzqhm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jzdxzqhm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jzdz  BeQuickEdit -->  
        <el-table-column label="居住地址" prop="jzdz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jzdz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jzdz',scope.row.ID,scope.row.jzdz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jzdz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--hkdxzqhm  BeQuickEdit -->  
        <el-table-column label="户口地行政区划码" prop="hkdxzqhm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.hkdxzqhm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('hkdxzqhm',scope.row.ID,scope.row.hkdxzqhm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.hkdxzqhm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--hkdz  BeQuickEdit -->  
        <el-table-column label="户口地址" prop="hkdz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.hkdz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('hkdz',scope.row.ID,scope.row.hkdz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.hkdz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--lxrxm  BeQuickEdit -->  
        <el-table-column label="联系人姓名" prop="lxrxm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.lxrxm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('lxrxm',scope.row.ID,scope.row.lxrxm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.lxrxm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--lxrdh  BeQuickEdit -->  
        <el-table-column label="联系人电话" prop="lxrdh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.lxrdh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('lxrdh',scope.row.ID,scope.row.lxrdh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.lxrdh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--abo  BeQuickEdit -->  
        <el-table-column label="ABO 血型" prop="abo" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.abo"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('abo',scope.row.ID,scope.row.abo,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.abo}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--rh  BeQuickEdit -->  
        <el-table-column label="RH 血型" prop="rh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.rh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('rh',scope.row.ID,scope.row.rh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.rh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cblbdm  BeQuickEdit -->  
        <el-table-column label="参保类别代码" prop="cblbdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cblbdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cblbdm',scope.row.ID,scope.row.cblbdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cblbdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--grdaid  BeQuickEdit -->  
        <el-table-column label="个人档案ID" prop="grdaid" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.grdaid"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('grdaid',scope.row.ID,scope.row.grdaid,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.grdaid}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--yl1  BeQuickEdit -->  
        <el-table-column label="预留一" prop="yl1" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.yl1"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('yl1',scope.row.ID,scope.row.yl1,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.yl1}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--yl2  BeQuickEdit -->  
        <el-table-column label="预留二" prop="yl2" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.yl2"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('yl2',scope.row.ID,scope.row.yl2,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.yl2}} </div>
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
        <el-form-item label="userid:">
                 <el-input v-model.number="formData.userid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="上传:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_upOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="机构标识:"> 
              <el-input v-model="formData.jgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="卡号:"> 
              <el-input v-model="formData.kh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="卡类型:"> 
              <el-input v-model="formData.klx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="身份证件号码:"> 
              <el-input v-model="formData.zjhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="身份证件类别:"> 
              <el-input v-model="formData.zjlbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="姓名:"> 
              <el-input v-model="formData.xm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别代码:"> 
              <el-input v-model="formData.xbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别名称:"> 
              <el-input v-model="formData.xbmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="患者属性:"> 
              <el-input v-model="formData.hzsx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="婚姻状况代码:"> 
              <el-input v-model="formData.hyztdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="婚姻状态名称:"> 
              <el-input v-model="formData.hyztmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出生日期:"> 
              <el-input v-model="formData.csrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="民族代码:"> 
              <el-input v-model="formData.mzdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="民族名称:"> 
              <el-input v-model="formData.mzmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="国籍代码:"> 
              <el-input v-model="formData.gjdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="国籍名称:"> 
              <el-input v-model="formData.gjmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构内部档案号:"> 
              <el-input v-model="formData.jgnbdah" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="固定电话:"> 
              <el-input v-model="formData.gddhhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机号码:"> 
              <el-input v-model="formData.sjhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电子邮件:"> 
              <el-input v-model="formData.dzyj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文化程度代码:"> 
              <el-input v-model="formData.whcddm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文化程度名称:"> 
              <el-input v-model="formData.whcdmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="职业类别代码:"> 
              <el-input v-model="formData.zylbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="职业类别名称:"> 
              <el-input v-model="formData.zylbmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出生地行政区划码:"> 
              <el-input v-model="formData.csdxzqhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出生地:"> 
              <el-input v-model="formData.csd" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="居住地行政区划码:"> 
              <el-input v-model="formData.jzdxzqhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="居住地址:"> 
              <el-input v-model="formData.jzdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="户口地行政区划码:"> 
              <el-input v-model="formData.hkdxzqhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="户口地址:"> 
              <el-input v-model="formData.hkdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="联系人姓名:"> 
              <el-input v-model="formData.lxrxm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="联系人电话:"> 
              <el-input v-model="formData.lxrdh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ABO 血型:"> 
              <el-input v-model="formData.abo" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="RH 血型:"> 
              <el-input v-model="formData.rh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="参保类别代码:"> 
              <el-input v-model="formData.cblbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="个人档案ID:"> 
              <el-input v-model="formData.grdaid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预留一:"> 
              <el-input v-model="formData.yl1" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预留二:"> 
              <el-input v-model="formData.yl2" clearable placeholder="请输入" />
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
  createBbPiPerson, 
  deleteBbPiPersonByIds,
  updateBbPiPerson,
  findBbPiPerson,
  getBbPiPersonList,
  quickEdit,
  excelList
} from '@/api/bbPiPerson' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'BbPiPerson',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getBbPiPersonList,   
      excelListApi: excelList,
      status_upOptions: [],
      formData: {
            userid: 0,
            status: 0,
           jgdm: '',
           kh: '',
           klx: '',
           zjhm: '',
           zjlbdm: '',
           xm: '',
           xbdm: '',
           xbmc: '',
           hzsx: '',
           hyztdm: '',
           hyztmc: '',
           csrq: '',
           mzdm: '',
           mzmc: '',
           gjdm: '',
           gjmc: '',
           jgnbdah: '',
           gddhhm: '',
           sjhm: '',
           dzyj: '',
           whcddm: '',
           whcdmc: '',
           zylbdm: '',
           zylbmc: '',
           csdxzqhm: '',
           csd: '',
           jzdxzqhm: '',
           jzdz: '',
           hkdxzqhm: '',
           hkdz: '',
           lxrxm: '',
           lxrdh: '',
           abo: '',
           rh: '',
           cblbdm: '',
           grdaid: '',
           yl1: '',
           yl2: '',
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
        //this.deleteBbPiPerson(row)
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
     const res = await deleteBbPiPersonByIds({ ids })
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
			this.$router.push({ name:'bbPiPersonForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'bbPiPersonForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findBbPiPerson({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.bbPiPerson 
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
          res = await createBbPiPerson(this.formData);
          break
        case "update": 
          res = await updateBbPiPerson(this.formData);
          break
        default: 
          res = await createBbPiPerson(this.formData);
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
