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
          <!--yhrygh  BeQuickEdit -->  
        <el-table-column label="医护人员账号" prop="yhrygh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.yhrygh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('yhrygh',scope.row.ID,scope.row.yhrygh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.yhrygh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--yhryxm  BeQuickEdit -->  
        <el-table-column label="医护人员姓名" prop="yhryxm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.yhryxm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('yhryxm',scope.row.ID,scope.row.yhryxm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.yhryxm}} </div>
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
          <!--sfzh  BeQuickEdit -->  
        <el-table-column label="身份证号" prop="sfzh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sfzh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sfzh',scope.row.ID,scope.row.sfzh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sfzh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zjlbdm  BeQuickEdit -->  
        <el-table-column label="身份证件类别代码" prop="zjlbdm" width="120"   sortable="custom" >
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
          <!--ksdm  BeQuickEdit -->  
        <el-table-column label="所属科室代码" prop="ksdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ksdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ksdm',scope.row.ID,scope.row.ksdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ksdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zyjszwdm  BeQuickEdit -->  
        <el-table-column label="专业技术职务代码" prop="zyjszwdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zyjszwdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zyjszwdm',scope.row.ID,scope.row.zyjszwdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zyjszwdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zyjszwlbdm  BeQuickEdit -->  
        <el-table-column label="专业技术职务类别代码" prop="zyjszwlbdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zyjszwlbdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zyjszwlbdm',scope.row.ID,scope.row.zyjszwlbdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zyjszwlbdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zzlbmc  BeQuickEdit -->  
        <el-table-column label="资质类别名称" prop="zzlbmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zzlbmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zzlbmc',scope.row.ID,scope.row.zzlbmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zzlbmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zgzsbh  BeQuickEdit -->  
        <el-table-column label="资格证书编号" prop="zgzsbh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zgzsbh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zgzsbh',scope.row.ID,scope.row.zgzsbh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zgzsbh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zghdsj  BeQuickEdit -->  
        <el-table-column label="资格获得时间" prop="zghdsj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zghdsj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zghdsj',scope.row.ID,scope.row.zghdsj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zghdsj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zyzsbm  BeQuickEdit -->  
        <el-table-column label="执业证书编码" prop="zyzsbm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zyzsbm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zyzsbm',scope.row.ID,scope.row.zyzsbm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zyzsbm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fzrq  BeQuickEdit -->  
        <el-table-column label="发证日期" prop="fzrq" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fzrq"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fzrq',scope.row.ID,scope.row.fzrq,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fzrq}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zydd  BeQuickEdit -->  
        <el-table-column label="执业地点" prop="zydd" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zydd"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zydd',scope.row.ID,scope.row.zydd,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zydd}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zyfw  BeQuickEdit -->  
        <el-table-column label="执业范围" prop="zyfw" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zyfw"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zyfw',scope.row.ID,scope.row.zyfw,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zyfw}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zyzyyljgdm  BeQuickEdit -->  
        <el-table-column label="主要执业医疗机构代码" prop="zyzyyljgdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zyzyyljgdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zyzyyljgdm',scope.row.ID,scope.row.zyzyyljgdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zyzyyljgdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zyzyyymc  BeQuickEdit -->  
        <el-table-column label="主要执业医院名称" prop="zyzyyymc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zyzyyymc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zyzyyymc',scope.row.ID,scope.row.zyzyyymc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zyzyyymc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--qkysbz  BeQuickEdit -->  
        <el-table-column label="全科医生标志" prop="qkysbz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.qkysbz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('qkysbz',scope.row.ID,scope.row.qkysbz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.qkysbz}} </div>
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
          <!--zc  BeQuickEdit -->  
        <el-table-column label="专长" prop="zc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zc',scope.row.ID,scope.row.zc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--mz  BeQuickEdit -->  
        <el-table-column label="民族" prop="mz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.mz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('mz',scope.row.ID,scope.row.mz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.mz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cjgzrq  BeQuickEdit -->  
        <el-table-column label="参加工作日期" prop="cjgzrq" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cjgzrq"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cjgzrq',scope.row.ID,scope.row.cjgzrq,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cjgzrq}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zcsj  BeQuickEdit -->  
        <el-table-column label="注册日期时间" prop="zcsj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zcsj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zcsj',scope.row.ID,scope.row.zcsj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zcsj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xl  BeQuickEdit -->  
        <el-table-column label="学历" prop="xl" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xl"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xl',scope.row.ID,scope.row.xl,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xl}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--grzpcfdz  BeQuickEdit -->  
        <el-table-column label="个人照片存放地址" prop="grzpcfdz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.grzpcfdz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('grzpcfdz',scope.row.ID,scope.row.grzpcfdz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.grzpcfdz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zgzcfdz  BeQuickEdit -->  
        <el-table-column label="资格证存放地址" prop="zgzcfdz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zgzcfdz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zgzcfdz',scope.row.ID,scope.row.zgzcfdz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zgzcfdz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zyzcfdz  BeQuickEdit -->  
        <el-table-column label="执业证存放地址" prop="zyzcfdz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zyzcfdz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zyzcfdz',scope.row.ID,scope.row.zyzcfdz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zyzcfdz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sjscsj  BeQuickEdit -->  
        <el-table-column label="数据生成日期时间" prop="sjscsj" width="120"   sortable="custom" >
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
        <el-form-item label="医护人员账号:"> 
              <el-input v-model="formData.yhrygh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医护人员姓名:"> 
              <el-input v-model="formData.yhryxm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别:"> 
              <el-input v-model="formData.xb" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出生日期:"> 
              <el-input v-model="formData.csrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="身份证号:"> 
              <el-input v-model="formData.sfzh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="身份证件类别代码:"> 
              <el-input v-model="formData.zjlbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="所属科室代码:"> 
              <el-input v-model="formData.ksdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="专业技术职务代码:"> 
              <el-input v-model="formData.zyjszwdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="专业技术职务类别代码:"> 
              <el-input v-model="formData.zyjszwlbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="资质类别名称:"> 
              <el-input v-model="formData.zzlbmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="资格证书编号:"> 
              <el-input v-model="formData.zgzsbh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="资格获得时间:"> 
              <el-input v-model="formData.zghdsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执业证书编码:"> 
              <el-input v-model="formData.zyzsbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="发证日期:"> 
              <el-input v-model="formData.fzrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执业地点:"> 
              <el-input v-model="formData.zydd" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执业范围:"> 
              <el-input v-model="formData.zyfw" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="主要执业医疗机构代码:"> 
              <el-input v-model="formData.zyzyyljgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="主要执业医院名称:"> 
              <el-input v-model="formData.zyzyyymc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="全科医生标志:"> 
              <el-input v-model="formData.qkysbz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机号码:"> 
              <el-input v-model="formData.sjhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="专长:"> 
              <el-input v-model="formData.zc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="民族:"> 
              <el-input v-model="formData.mz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="参加工作日期:"> 
              <el-input v-model="formData.cjgzrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="注册日期时间:"> 
              <el-input v-model="formData.zcsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="学历:"> 
              <el-input v-model="formData.xl" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="个人照片存放地址:"> 
              <el-input v-model="formData.grzpcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="资格证存放地址:"> 
              <el-input v-model="formData.zgzcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执业证存放地址:"> 
              <el-input v-model="formData.zyzcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数据生成日期时间:"> 
              <el-input v-model="formData.sjscsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="填报日期:"> 
              <el-input v-model="formData.tbrqsj" clearable placeholder="请输入" />
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
  createBbPiStaff, 
  deleteBbPiStaffByIds,
  updateBbPiStaff,
  findBbPiStaff,
  getBbPiStaffList,
  quickEdit,
  excelList
} from '@/api/bbPiStaff' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'BbPiStaff',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getBbPiStaffList,   
      excelListApi: excelList,
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           yhrygh: '',
           yhryxm: '',
           xb: '',
           csrq: '',
           sfzh: '',
           zjlbdm: '',
           ksdm: '',
           zyjszwdm: '',
           zyjszwlbdm: '',
           zzlbmc: '',
           zgzsbh: '',
           zghdsj: '',
           zyzsbm: '',
           fzrq: '',
           zydd: '',
           zyfw: '',
           zyzyyljgdm: '',
           zyzyyymc: '',
           qkysbz: '',
           sjhm: '',
           zc: '',
           mz: '',
           cjgzrq: '',
           zcsj: '',
           xl: '',
           grzpcfdz: '',
           zgzcfdz: '',
           zyzcfdz: '',
           sjscsj: '',
           tbrqsj: '',
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
        //this.deleteBbPiStaff(row)
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
     const res = await deleteBbPiStaffByIds({ ids })
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
			this.$router.push({ name:'bbPiStaffForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'bbPiStaffForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findBbPiStaff({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.bbPiStaff 
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
          res = await createBbPiStaff(this.formData);
          break
        case "update": 
          res = await updateBbPiStaff(this.formData);
          break
        default: 
          res = await createBbPiStaff(this.formData);
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
