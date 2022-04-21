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
          <!--zzjgdm  BeQuickEdit -->  
        <el-table-column label="统一信用代码" prop="zzjgdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zzjgdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zzjgdm',scope.row.ID,scope.row.zzjgdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zzjgdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jgmc  BeQuickEdit -->  
        <el-table-column label="机构名称" prop="jgmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jgmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jgmc',scope.row.ID,scope.row.jgmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jgmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xzqhdm  BeQuickEdit -->  
        <el-table-column label="行政区划代码" prop="xzqhdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xzqhdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xzqhdm',scope.row.ID,scope.row.xzqhdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xzqhdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jglx  BeQuickEdit -->  
        <el-table-column label="机构类型" prop="jglx" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jglx"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jglx',scope.row.ID,scope.row.jglx,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jglx}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jgclrq  BeQuickEdit -->  
        <el-table-column label="机构成立日期" prop="jgclrq" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jgclrq"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jgclrq',scope.row.ID,scope.row.jgclrq,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jgclrq}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--dwlsgxdm  BeQuickEdit -->  
        <el-table-column label="单位隶属关系代码" prop="dwlsgxdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.dwlsgxdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('dwlsgxdm',scope.row.ID,scope.row.dwlsgxdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.dwlsgxdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jgflgllbdm  BeQuickEdit -->  
        <el-table-column label="机构分类代码" prop="jgflgllbdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jgflgllbdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jgflgllbdm',scope.row.ID,scope.row.jgflgllbdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jgflgllbdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jgfldm  BeQuickEdit -->  
        <el-table-column label="机构分类代码" prop="jgfldm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jgfldm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jgfldm',scope.row.ID,scope.row.jgfldm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jgfldm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jjlxdm  BeQuickEdit -->  
        <el-table-column label="经济类型代码" prop="jjlxdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jjlxdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jjlxdm',scope.row.ID,scope.row.jjlxdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jjlxdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--dz  BeQuickEdit -->  
        <el-table-column label="地址" prop="dz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.dz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('dz',scope.row.ID,scope.row.dz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.dz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--styyzzjgdm  BeQuickEdit -->  
        <el-table-column label="组织机构代码" prop="styyzzjgdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.styyzzjgdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('styyzzjgdm',scope.row.ID,scope.row.styyzzjgdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.styyzzjgdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--styymc  BeQuickEdit -->  
        <el-table-column label="医院名称" prop="styymc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.styymc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('styymc',scope.row.ID,scope.row.styymc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.styymc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--styljgjb  BeQuickEdit -->  
        <el-table-column label="医疗机构级别" prop="styljgjb" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.styljgjb"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('styljgjb',scope.row.ID,scope.row.styljgjb,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.styljgjb}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--styljgdj  BeQuickEdit -->  
        <el-table-column label="医院机构等级" prop="styljgdj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.styljgdj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('styljgdj',scope.row.ID,scope.row.styljgdj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.styljgdj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--hlwyywz  BeQuickEdit -->  
        <el-table-column label="医院网址" prop="hlwyywz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.hlwyywz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('hlwyywz',scope.row.ID,scope.row.hlwyywz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.hlwyywz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xkzhm  BeQuickEdit -->  
        <el-table-column label="许可证号码" prop="xkzhm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xkzhm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xkzhm',scope.row.ID,scope.row.xkzhm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xkzhm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xkxmmc  BeQuickEdit -->  
        <el-table-column label="许可项目名称" prop="xkxmmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xkxmmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xkxmmc',scope.row.ID,scope.row.xkxmmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xkxmmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xkzyxq  BeQuickEdit -->  
        <el-table-column label="许可证有效期" prop="xkzyxq" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xkzyxq"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xkzyxq',scope.row.ID,scope.row.xkzyxq,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xkzyxq}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xxaqdjbh  BeQuickEdit -->  
        <el-table-column label="信息安等级保护" prop="xxaqdjbh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xxaqdjbh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xxaqdjbh',scope.row.ID,scope.row.xxaqdjbh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xxaqdjbh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xxaqdjbhbh  BeQuickEdit -->  
        <el-table-column label="信息安等保编号" prop="xxaqdjbhbh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xxaqdjbhbh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xxaqdjbhbh',scope.row.ID,scope.row.xxaqdjbhbh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xxaqdjbhbh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--kbzjjes  BeQuickEdit -->  
        <el-table-column label="开办资金金额数" prop="kbzjjes" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.kbzjjes"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('kbzjjes',scope.row.ID,scope.row.kbzjjes,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.kbzjjes}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--frdb  BeQuickEdit -->  
        <el-table-column label="法人代表" prop="frdb" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.frdb"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('frdb',scope.row.ID,scope.row.frdb,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.frdb}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jgszdmzzzdfbz  BeQuickEdit -->  
        <el-table-column label="机构民族自治标志" prop="jgszdmzzzdfbz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jgszdmzzzdfbz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jgszdmzzzdfbz',scope.row.ID,scope.row.jgszdmzzzdfbz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jgszdmzzzdfbz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sffzjg  BeQuickEdit -->  
        <el-table-column label="是否分支机构" prop="sffzjg" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sffzjg"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sffzjg',scope.row.ID,scope.row.sffzjg,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sffzjg}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jgdemc  BeQuickEdit -->  
        <el-table-column label="机构第二名称" prop="jgdemc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jgdemc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jgdemc',scope.row.ID,scope.row.jgdemc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jgdemc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jgms  BeQuickEdit -->  
        <el-table-column label="机构描述" prop="jgms" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jgms"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jgms',scope.row.ID,scope.row.jgms,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jgms}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--yzbm  BeQuickEdit -->  
        <el-table-column label="邮政编码" prop="yzbm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.yzbm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('yzbm',scope.row.ID,scope.row.yzbm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.yzbm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--dhhm  BeQuickEdit -->  
        <el-table-column label="电话号码" prop="dhhm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.dhhm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('dhhm',scope.row.ID,scope.row.dhhm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.dhhm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--dwdzyx  BeQuickEdit -->  
        <el-table-column label="电子信箱" prop="dwdzyx" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.dwdzyx"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('dwdzyx',scope.row.ID,scope.row.dwdzyx,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.dwdzyx}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--dsfjgdm  BeQuickEdit -->  
        <el-table-column label="第三方信用代码" prop="dsfjgdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.dsfjgdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('dsfjgdm',scope.row.ID,scope.row.dsfjgdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.dsfjgdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--dsfjgmc  BeQuickEdit -->  
        <el-table-column label="第三方机构名称" prop="dsfjgmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.dsfjgmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('dsfjgmc',scope.row.ID,scope.row.dsfjgmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.dsfjgmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xyyxq  BeQuickEdit -->  
        <el-table-column label="协议有效期" prop="xyyxq" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xyyxq"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xyyxq',scope.row.ID,scope.row.xyyxq,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xyyxq}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--bspt  BeQuickEdit -->  
        <el-table-column label="部署平台" prop="bspt" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.bspt"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('bspt',scope.row.ID,scope.row.bspt,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.bspt}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jgmsxx  BeQuickEdit -->  
        <el-table-column label="架构描述" prop="jgmsxx" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jgmsxx"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jgmsxx',scope.row.ID,scope.row.jgmsxx,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jgmsxx}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jfmj  BeQuickEdit -->  
        <el-table-column label="机房面积" prop="jfmj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jfmj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jfmj',scope.row.ID,scope.row.jfmj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jfmj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sfslgd  BeQuickEdit -->  
        <el-table-column label="是否双路发电" prop="sfslgd" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sfslgd"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sfslgd',scope.row.ID,scope.row.sfslgd,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sfslgd}} </div>
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
        <el-form-item label="统一信用代码:"> 
              <el-input v-model="formData.zzjgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构名称:"> 
              <el-input v-model="formData.jgmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="行政区划代码:"> 
              <el-input v-model="formData.xzqhdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构类型:"> 
              <el-input v-model="formData.jglx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构成立日期:"> 
              <el-input v-model="formData.jgclrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="单位隶属关系代码:"> 
              <el-input v-model="formData.dwlsgxdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构分类代码:"> 
              <el-input v-model="formData.jgflgllbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构分类代码:"> 
              <el-input v-model="formData.jgfldm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="经济类型代码:"> 
              <el-input v-model="formData.jjlxdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地址:"> 
              <el-input v-model="formData.dz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="组织机构代码:"> 
              <el-input v-model="formData.styyzzjgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医院名称:"> 
              <el-input v-model="formData.styymc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医疗机构级别:"> 
              <el-input v-model="formData.styljgjb" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医院机构等级:"> 
              <el-input v-model="formData.styljgdj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医院网址:"> 
              <el-input v-model="formData.hlwyywz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="许可证号码:"> 
              <el-input v-model="formData.xkzhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="许可项目名称:"> 
              <el-input v-model="formData.xkxmmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="许可证有效期:"> 
              <el-input v-model="formData.xkzyxq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="信息安等级保护:"> 
              <el-input v-model="formData.xxaqdjbh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="信息安等保编号:"> 
              <el-input v-model="formData.xxaqdjbhbh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="开办资金金额数:"> 
              <el-input v-model="formData.kbzjjes" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="法人代表:"> 
              <el-input v-model="formData.frdb" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构民族自治标志:"> 
              <el-input v-model="formData.jgszdmzzzdfbz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否分支机构:"> 
              <el-input v-model="formData.sffzjg" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构第二名称:"> 
              <el-input v-model="formData.jgdemc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构描述:"> 
              <el-input v-model="formData.jgms" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮政编码:"> 
              <el-input v-model="formData.yzbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电话号码:"> 
              <el-input v-model="formData.dhhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电子信箱:"> 
              <el-input v-model="formData.dwdzyx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="第三方信用代码:"> 
              <el-input v-model="formData.dsfjgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="第三方机构名称:"> 
              <el-input v-model="formData.dsfjgmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="协议有效期:"> 
              <el-input v-model="formData.xyyxq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="部署平台:"> 
              <el-input v-model="formData.bspt" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="架构描述:"> 
              <el-input v-model="formData.jgmsxx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机房面积:"> 
              <el-input v-model="formData.jfmj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否双路发电:"> 
              <el-input v-model="formData.sfslgd" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:"> 
              <el-input v-model="formData.bz" clearable placeholder="请输入" />
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
  createBbPiInstitution, 
  deleteBbPiInstitutionByIds,
  updateBbPiInstitution,
  findBbPiInstitution,
  getBbPiInstitutionList,
  quickEdit,
  excelList
} from '@/api/bbPiInstitution' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'BbPiInstitution',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getBbPiInstitutionList,   
      excelListApi: excelList,
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           zzjgdm: '',
           jgmc: '',
           xzqhdm: '',
           jglx: '',
           jgclrq: '',
           dwlsgxdm: '',
           jgflgllbdm: '',
           jgfldm: '',
           jjlxdm: '',
           dz: '',
           styyzzjgdm: '',
           styymc: '',
           styljgjb: '',
           styljgdj: '',
           hlwyywz: '',
           xkzhm: '',
           xkxmmc: '',
           xkzyxq: '',
           xxaqdjbh: '',
           xxaqdjbhbh: '',
           kbzjjes: '',
           frdb: '',
           jgszdmzzzdfbz: '',
           sffzjg: '',
           jgdemc: '',
           jgms: '',
           yzbm: '',
           dhhm: '',
           dwdzyx: '',
           dsfjgdm: '',
           dsfjgmc: '',
           xyyxq: '',
           bspt: '',
           jgmsxx: '',
           jfmj: '',
           sfslgd: '',
           bz: '',
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
        //this.deleteBbPiInstitution(row)
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
     const res = await deleteBbPiInstitutionByIds({ ids })
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
			this.$router.push({ name:'bbPiInstitutionForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'bbPiInstitutionForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findBbPiInstitution({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.bbPiInstitution 
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
          res = await createBbPiInstitution(this.formData);
          break
        case "update": 
          res = await updateBbPiInstitution(this.formData);
          break
        default: 
          res = await createBbPiInstitution(this.formData);
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
