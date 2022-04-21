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
          <!--jzlx  BeQuickEdit -->  
        <el-table-column label="就诊类型" prop="jzlx" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jzlx"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jzlx',scope.row.ID,scope.row.jzlx,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jzlx}} </div>
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
          <!--ksbm  BeQuickEdit -->  
        <el-table-column label="科室编码" prop="ksbm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ksbm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ksbm',scope.row.ID,scope.row.ksbm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ksbm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--ksmc  BeQuickEdit -->  
        <el-table-column label="科室名称" prop="ksmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ksmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ksmc',scope.row.ID,scope.row.ksmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ksmc}} </div>
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
          <!--gmsbs  BeQuickEdit -->  
        <el-table-column label="过敏史标识" prop="gmsbs" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.gmsbs"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('gmsbs',scope.row.ID,scope.row.gmsbs,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.gmsbs}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--gms  BeQuickEdit -->  
        <el-table-column label="过敏史" prop="gms" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.gms"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('gms',scope.row.ID,scope.row.gms,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.gms}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cblb  BeQuickEdit -->  
        <el-table-column label="参保类别" prop="cblb" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cblb"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cblb',scope.row.ID,scope.row.cblb,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cblb}} </div>
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
          <!--jzzdsm  BeQuickEdit -->  
        <el-table-column label="就诊诊断说明" prop="jzzdsm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jzzdsm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jzzdsm',scope.row.ID,scope.row.jzzdsm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jzzdsm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--czbzdm  BeQuickEdit -->  
        <el-table-column label="初诊标志代码" prop="czbzdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.czbzdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('czbzdm',scope.row.ID,scope.row.czbzdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.czbzdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zs  BeQuickEdit -->  
        <el-table-column label="主诉" prop="zs" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zs"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zs',scope.row.ID,scope.row.zs,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zs}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xbs  BeQuickEdit -->  
        <el-table-column label="现病史" prop="xbs" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xbs"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xbs',scope.row.ID,scope.row.xbs,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xbs}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jws  BeQuickEdit -->  
        <el-table-column label="既往史" prop="jws" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jws"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jws',scope.row.ID,scope.row.jws,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jws}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fzjcxm  BeQuickEdit -->  
        <el-table-column label="辅助检查项目" prop="fzjcxm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fzjcxm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fzjcxm',scope.row.ID,scope.row.fzjcxm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fzjcxm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fzjcjg  BeQuickEdit -->  
        <el-table-column label="辅助检查结果" prop="fzjcjg" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fzjcjg"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fzjcjg',scope.row.ID,scope.row.fzjcjg,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fzjcjg}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--mzzzzddm  BeQuickEdit -->  
        <el-table-column label="门诊症状-诊断代码" prop="mzzzzddm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.mzzzzddm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('mzzzzddm',scope.row.ID,scope.row.mzzzzddm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.mzzzzddm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--mzzzzdmc  BeQuickEdit -->  
        <el-table-column label="门诊症状-诊断名称" prop="mzzzzdmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.mzzzzdmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('mzzzzdmc',scope.row.ID,scope.row.mzzzzdmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.mzzzzdmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--mzzzzdbmlx  BeQuickEdit -->  
        <el-table-column label="门诊症状诊断编码类型" prop="mzzzzdbmlx" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.mzzzzdbmlx"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('mzzzzdbmlx',scope.row.ID,scope.row.mzzzzdbmlx,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.mzzzzdbmlx}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zzms  BeQuickEdit -->  
        <el-table-column label="症状描述" prop="zzms" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zzms"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zzms',scope.row.ID,scope.row.zzms,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zzms}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--bzyj  BeQuickEdit -->  
        <el-table-column label="辨证依据" prop="bzyj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.bzyj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('bzyj',scope.row.ID,scope.row.bzyj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.bzyj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--zzzf  BeQuickEdit -->  
        <el-table-column label="治则治法" prop="zzzf" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zzzf"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zzzf',scope.row.ID,scope.row.zzzf,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zzzf}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jkzd  BeQuickEdit -->  
        <el-table-column label="健康指导" prop="jkzd" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jkzd"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jkzd',scope.row.ID,scope.row.jkzd,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jkzd}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--czjh  BeQuickEdit -->  
        <el-table-column label="处置计划" prop="czjh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.czjh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('czjh',scope.row.ID,scope.row.czjh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.czjh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--yzysgh  BeQuickEdit -->  
        <el-table-column label="应诊医生工号" prop="yzysgh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.yzysgh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('yzysgh',scope.row.ID,scope.row.yzysgh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.yzysgh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--yzysqm  BeQuickEdit -->  
        <el-table-column label="应诊医师签名" prop="yzysqm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.yzysqm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('yzysqm',scope.row.ID,scope.row.yzysqm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.yzysqm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--czylwsjgdm  BeQuickEdit -->  
        <el-table-column label="初诊医疗卫生机构代码" prop="czylwsjgdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.czylwsjgdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('czylwsjgdm',scope.row.ID,scope.row.czylwsjgdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.czylwsjgdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--czylwsjgmc  BeQuickEdit -->  
        <el-table-column label="初诊机构名称" prop="czylwsjgmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.czylwsjgmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('czylwsjgmc',scope.row.ID,scope.row.czylwsjgmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.czylwsjgmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jzjssj  BeQuickEdit -->  
        <el-table-column label="就诊结束时间" prop="jzjssj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jzjssj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jzjssj',scope.row.ID,scope.row.jzjssj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jzjssj}} </div>
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
          <!--hzmyd  BeQuickEdit -->  
        <el-table-column label="患者满意度" prop="hzmyd" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.hzmyd"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('hzmyd',scope.row.ID,scope.row.hzmyd,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.hzmyd}} </div>
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
          <!--dzcfwjcfdz  BeQuickEdit -->  
        <el-table-column label="电子处方地址" prop="dzcfwjcfdz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.dzcfwjcfdz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('dzcfwjcfdz',scope.row.ID,scope.row.dzcfwjcfdz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.dzcfwjcfdz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--ysdspwjcfdz  BeQuickEdit -->  
        <el-table-column label="医生端视频地址" prop="ysdspwjcfdz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.ysdspwjcfdz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('ysdspwjcfdz',scope.row.ID,scope.row.ysdspwjcfdz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.ysdspwjcfdz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--hzdspwjcfdz  BeQuickEdit -->  
        <el-table-column label="患者端视频地址" prop="hzdspwjcfdz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.hzdspwjcfdz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('hzdspwjcfdz',scope.row.ID,scope.row.hzdspwjcfdz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.hzdspwjcfdz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jlypcfdz  BeQuickEdit -->  
        <el-table-column label="交流音频地址" prop="jlypcfdz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jlypcfdz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jlypcfdz',scope.row.ID,scope.row.jlypcfdz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jlypcfdz}} </div>
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
        <el-form-item label="就诊类型:"> 
              <el-input v-model="formData.jzlx" clearable placeholder="请输入" />
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
        <el-form-item label="科室编码:"> 
              <el-input v-model="formData.ksbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="科室名称:"> 
              <el-input v-model="formData.ksmc" clearable placeholder="请输入" />
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
        <el-form-item label="过敏史标识:"> 
              <el-input v-model="formData.gmsbs" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="过敏史:"> 
              <el-input v-model="formData.gms" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="参保类别:"> 
              <el-input v-model="formData.cblb" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="就诊日期时间:"> 
              <el-input v-model="formData.jzrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="就诊诊断说明:"> 
              <el-input v-model="formData.jzzdsm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="初诊标志代码:"> 
              <el-input v-model="formData.czbzdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="主诉:"> 
              <el-input v-model="formData.zs" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="现病史:"> 
              <el-input v-model="formData.xbs" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="既往史:"> 
              <el-input v-model="formData.jws" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="辅助检查项目:"> 
              <el-input v-model="formData.fzjcxm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="辅助检查结果:"> 
              <el-input v-model="formData.fzjcjg" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="门诊症状-诊断代码:"> 
              <el-input v-model="formData.mzzzzddm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="门诊症状-诊断名称:"> 
              <el-input v-model="formData.mzzzzdmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="门诊症状诊断编码类型:"> 
              <el-input v-model="formData.mzzzzdbmlx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="症状描述:"> 
              <el-input v-model="formData.zzms" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="辨证依据:"> 
              <el-input v-model="formData.bzyj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="治则治法:"> 
              <el-input v-model="formData.zzzf" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="健康指导:"> 
              <el-input v-model="formData.jkzd" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处置计划:"> 
              <el-input v-model="formData.czjh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="应诊医生工号:"> 
              <el-input v-model="formData.yzysgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="应诊医师签名:"> 
              <el-input v-model="formData.yzysqm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="初诊医疗卫生机构代码:"> 
              <el-input v-model="formData.czylwsjgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="初诊机构名称:"> 
              <el-input v-model="formData.czylwsjgmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="就诊结束时间:"> 
              <el-input v-model="formData.jzjssj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转诊标志:"> 
              <el-input v-model="formData.zzbz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="患者满意度:"> 
              <el-input v-model="formData.hzmyd" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预留一:"> 
              <el-input v-model="formData.yl1" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预留二:"> 
              <el-input v-model="formData.yl2" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电子处方地址:"> 
              <el-input v-model="formData.dzcfwjcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医生端视频地址:"> 
              <el-input v-model="formData.ysdspwjcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="患者端视频地址:"> 
              <el-input v-model="formData.hzdspwjcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="交流音频地址:"> 
              <el-input v-model="formData.jlypcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数据生成日期时间:"> 
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
  createBbPiTreatmentRecord, 
  deleteBbPiTreatmentRecordByIds,
  updateBbPiTreatmentRecord,
  findBbPiTreatmentRecord,
  getBbPiTreatmentRecordList,
  quickEdit,
  excelList
} from '@/api/bbPiTreatmentRecord' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'BbPiTreatmentRecord',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getBbPiTreatmentRecordList,   
      excelListApi: excelList,
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           fwwddm: '',
           jzlx: '',
           kh: '',
           klx: '',
           mzh: '',
           ksbm: '',
           ksmc: '',
           xm: '',
           xbdm: '',
           csrq: '',
           nls: '',
           nly: '',
           gmsbs: '',
           gms: '',
           cblb: '',
           jzrqsj: '',
           jzzdsm: '',
           czbzdm: '',
           zs: '',
           xbs: '',
           jws: '',
           fzjcxm: '',
           fzjcjg: '',
           mzzzzddm: '',
           mzzzzdmc: '',
           mzzzzdbmlx: '',
           zzms: '',
           bzyj: '',
           zzzf: '',
           jkzd: '',
           czjh: '',
           yzysgh: '',
           yzysqm: '',
           czylwsjgdm: '',
           czylwsjgmc: '',
           jzjssj: '',
           zzbz: '',
           hzmyd: '',
           yl1: '',
           yl2: '',
           dzcfwjcfdz: '',
           ysdspwjcfdz: '',
           hzdspwjcfdz: '',
           jlypcfdz: '',
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
        //this.deleteBbPiTreatmentRecord(row)
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
     const res = await deleteBbPiTreatmentRecordByIds({ ids })
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
			this.$router.push({ name:'bbPiTreatmentRecordForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'bbPiTreatmentRecordForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findBbPiTreatmentRecord({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.bbPiTreatmentRecord 
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
          res = await createBbPiTreatmentRecord(this.formData);
          break
        case "update": 
          res = await updateBbPiTreatmentRecord(this.formData);
          break
        default: 
          res = await createBbPiTreatmentRecord(this.formData);
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
