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
                <el-form-item label="状态" prop="status">                
                    <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
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
          <!--userId  BeQuickEdit -->  
        <el-table-column label="用户id" prop="userId" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.userId"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('user_id',scope.row.ID,scope.row.userId,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.userId}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--catId  BeQuickEdit -->  
        <el-table-column label="系统栏目" prop="catId" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.catId"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cat_id',scope.row.ID,scope.row.catId,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.catId}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--catIdUser  BeQuickEdit -->  
        <el-table-column label="用户栏目" prop="catIdUser" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.catIdUser"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cat_id_user',scope.row.ID,scope.row.catIdUser,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.catIdUser}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--beOnline  BeQuickEdit -->
        <el-table-column label="类型" prop="beOnline" width="120"   sortable="custom"  >                        
            <template #default="scope" ><el-switch v-model="scope.row.beOnline" @change="quickEdit_do('be_online',scope.row.ID,scope.row.beOnline,scope)"/></template> 
        </el-table-column>
          <!--actType  BeQuickEdit -->  
        <el-table-column label="活动类型" prop="actType" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.actType"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('act_type',scope.row.ID,scope.row.actType,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.actType}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--beMulti  BeQuickEdit -->
        <el-table-column label="分期活动" prop="beMulti" width="120"   sortable="custom"  >                        
            <template #default="scope" ><el-switch v-model="scope.row.beMulti" @change="quickEdit_do('be_multi',scope.row.ID,scope.row.beMulti,scope)"/></template> 
        </el-table-column>
          <!--phase  BeQuickEdit -->  
        <el-table-column label="分期" prop="phase" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.phase"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('phase',scope.row.ID,scope.row.phase,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.phase}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--title  BeQuickEdit -->  
        <el-table-column label="标题" prop="title" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.title"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('title',scope.row.ID,scope.row.title,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.title}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--desc  BeQuickEdit -->  
        <el-table-column label="简介" prop="desc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.desc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('desc',scope.row.ID,scope.row.desc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.desc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--avater  BeQuickEdit -->  
        <el-table-column label="缩略图" prop="avater" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.avater"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('avater',scope.row.ID,scope.row.avater,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.avater}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--mediaList  BeQuickEdit -->  
        <el-table-column label="媒体列表" prop="mediaList" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.mediaList"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('media_list',scope.row.ID,scope.row.mediaList,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.mediaList}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--address  BeQuickEdit -->  
        <el-table-column label="地址" prop="address" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.address"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('address',scope.row.ID,scope.row.address,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.address}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--areaId  BeQuickEdit -->  
        <el-table-column label="地区id" prop="areaId" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.areaId"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('area_id',scope.row.ID,scope.row.areaId,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.areaId}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--lng  BeQuickEdit -->  
        <el-table-column label="经度" prop="lng" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.lng"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('lng',scope.row.ID,scope.row.lng,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.lng}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--lat  BeQuickEdit -->  
        <el-table-column label="纬度" prop="lat" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.lat"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('lat',scope.row.ID,scope.row.lat,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.lat}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--beginTime  BeQuickEdit -->  
        <el-table-column label="开始时间" prop="beginTime" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.beginTime"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('begin_time',scope.row.ID,scope.row.beginTime,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.beginTime}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--endTime  BeQuickEdit -->  
        <el-table-column label="结束时间" prop="endTime" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.endTime"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('end_time',scope.row.ID,scope.row.endTime,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.endTime}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--liveTime  BeQuickEdit -->  
        <el-table-column label="直播开始" prop="liveTime" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.liveTime"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('live_time',scope.row.ID,scope.row.liveTime,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.liveTime}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--liveEnd  BeQuickEdit -->  
        <el-table-column label="直播结束" prop="liveEnd" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.liveEnd"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('live_end',scope.row.ID,scope.row.liveEnd,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.liveEnd}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--presenter  BeQuickEdit -->  
        <el-table-column label="主持/讲师" prop="presenter" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.presenter"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('presenter',scope.row.ID,scope.row.presenter,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.presenter}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--phoneKefu  BeQuickEdit -->  
        <el-table-column label="客服电话" prop="phoneKefu" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.phoneKefu"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('phone_kefu',scope.row.ID,scope.row.phoneKefu,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.phoneKefu}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--phoneHezu  BeQuickEdit -->  
        <el-table-column label="合作电话" prop="phoneHezu" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.phoneHezu"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('phone_hezu',scope.row.ID,scope.row.phoneHezu,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.phoneHezu}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--wx  BeQuickEdit -->  
        <el-table-column label="微信" prop="wx" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.wx"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('wx',scope.row.ID,scope.row.wx,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.wx}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--qq  BeQuickEdit -->  
        <el-table-column label="QQ" prop="qq" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.qq"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('qq',scope.row.ID,scope.row.qq,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.qq}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--groupId  BeQuickEdit -->  
        <el-table-column label="群组id" prop="groupId" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.groupId"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('group_id',scope.row.ID,scope.row.groupId,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.groupId}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--endJoinTime  BeQuickEdit -->  
        <el-table-column label="结束报名" prop="endJoinTime" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.endJoinTime"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('end_join_time',scope.row.ID,scope.row.endJoinTime,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.endJoinTime}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--price  BeQuickEdit -->  
        <el-table-column label="票价" prop="price" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.price"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('price',scope.row.ID,scope.row.price,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.price}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--priceVip  BeQuickEdit -->  
        <el-table-column label="vip票价" prop="priceVip" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.priceVip"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('price_vip',scope.row.ID,scope.row.priceVip,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.priceVip}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--priceDesc  BeQuickEdit -->  
        <el-table-column label="票价描述" prop="priceDesc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.priceDesc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('price_desc',scope.row.ID,scope.row.priceDesc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.priceDesc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--refundType  BeQuickEdit -->  
        <el-table-column label="退费类型" prop="refundType" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.refundType"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('refund_type',scope.row.ID,scope.row.refundType,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.refundType}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--refundDays  BeQuickEdit -->  
        <el-table-column label="退费天" prop="refundDays" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.refundDays"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('refund_days',scope.row.ID,scope.row.refundDays,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.refundDays}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--beShowjoin  BeQuickEdit -->  
        <el-table-column label="是否显示人数" prop="beShowjoin" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.beShowjoin"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('be_showjoin',scope.row.ID,scope.row.beShowjoin,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.beShowjoin}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--statusAct  BeQuickEdit -->  
        <el-table-column label="活动状态" prop="statusAct" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.statusAct"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('status_act',scope.row.ID,scope.row.statusAct,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.statusAct}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--maxJoin  BeQuickEdit -->  
        <el-table-column label="最大报名人数" prop="maxJoin" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.maxJoin"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('max_join',scope.row.ID,scope.row.maxJoin,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.maxJoin}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--nowJoin  BeQuickEdit -->  
        <el-table-column label="当前报名人数" prop="nowJoin" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.nowJoin"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('now_join',scope.row.ID,scope.row.nowJoin,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.nowJoin}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--beChosen  BeQuickEdit -->
        <el-table-column label="是否精选" prop="beChosen" width="120"   sortable="custom"  >                        
            <template #default="scope" ><el-switch v-model="scope.row.beChosen" @change="quickEdit_do('be_chosen',scope.row.ID,scope.row.beChosen,scope)"/></template> 
        </el-table-column>
          <!--beWeek  BeQuickEdit -->
        <el-table-column label="是否周末" prop="beWeek" width="120"   sortable="custom"  >                        
            <template #default="scope" ><el-switch v-model="scope.row.beWeek" @change="quickEdit_do('be_week',scope.row.ID,scope.row.beWeek,scope)"/></template> 
        </el-table-column>
          <!--beVip  BeQuickEdit -->
        <el-table-column label="是否会员" prop="beVip" width="120"   sortable="custom"  >                        
            <template #default="scope" ><el-switch v-model="scope.row.beVip" @change="quickEdit_do('be_vip',scope.row.ID,scope.row.beVip,scope)"/></template> 
        </el-table-column>
          <!--totalWhole  BeQuickEdit -->  
        <el-table-column label="综合指数" prop="totalWhole" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.totalWhole"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('total_whole',scope.row.ID,scope.row.totalWhole,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.totalWhole}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--totalShare  BeQuickEdit -->  
        <el-table-column label="总分享" prop="totalShare" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.totalShare"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('total_share',scope.row.ID,scope.row.totalShare,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.totalShare}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--totalFav  BeQuickEdit -->  
        <el-table-column label="总收藏" prop="totalFav" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.totalFav"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('total_fav',scope.row.ID,scope.row.totalFav,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.totalFav}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--totalJoin  BeQuickEdit -->  
        <el-table-column label="总报名数" prop="totalJoin" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.totalJoin"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('total_join',scope.row.ID,scope.row.totalJoin,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.totalJoin}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--totalDiscuss  BeQuickEdit -->  
        <el-table-column label="总评论" prop="totalDiscuss" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.totalDiscuss"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('total_discuss',scope.row.ID,scope.row.totalDiscuss,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.totalDiscuss}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--totalClick  BeQuickEdit -->  
        <el-table-column label="总点击" prop="totalClick" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.totalClick"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('total_click',scope.row.ID,scope.row.totalClick,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.totalClick}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--totalStar  BeQuickEdit -->  
        <el-table-column label="总评" prop="totalStar" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.totalStar"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('total_star',scope.row.ID,scope.row.totalStar,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.totalStar}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--totalStar1  BeQuickEdit -->  
        <el-table-column label="总星评1" prop="totalStar1" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.totalStar1"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('total_star_1',scope.row.ID,scope.row.totalStar1,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.totalStar1}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--totalStar2  BeQuickEdit -->  
        <el-table-column label="总星评2" prop="totalStar2" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.totalStar2"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('total_star_2',scope.row.ID,scope.row.totalStar2,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.totalStar2}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--totalStar3  BeQuickEdit -->  
        <el-table-column label="总星评3" prop="totalStar3" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.totalStar3"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('total_star_3',scope.row.ID,scope.row.totalStar3,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.totalStar3}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--status  BeQuickEdit -->
        <el-table-column label="状态" prop="status" width="120"  sortable="custom" >
        <template #default="scope">  
        <el-popover trigger="click" placement="top"  width = "280">  
              <el-select v-model="scope.row.status" placeholder="请选择"  @change="quickEdit_do('status',scope.row.ID,scope.row.status,scope)">
                  <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
              </el-select> 
              <template #reference>
                  <div class="quickEdit" > {{filterDict(scope.row.status,"status")}} </div>
              </template>
            </el-popover>
        </template>  
        </el-table-column>
          <!--statusMsg  BeQuickEdit -->  
        <el-table-column label="审核原因" prop="statusMsg" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.statusMsg"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('status_msg',scope.row.ID,scope.row.statusMsg,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.statusMsg}} </div>
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
        <el-form-item label="用户id:">
                 <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="系统栏目:">
                 <el-input v-model.number="formData.catId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户栏目:">
                 <el-input v-model.number="formData.catIdUser" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类型:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beOnline" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="活动类型:"> 
              <el-input v-model="formData.actType" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="分期活动:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beMulti" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="分期:">
                 <el-input v-model.number="formData.phase" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="标题:"> 
              <el-input v-model="formData.title" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="简介:"> 
              <el-input v-model="formData.desc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="缩略图:"> 
              <el-input v-model="formData.avater" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="媒体列表:"> 
              <el-input v-model="formData.mediaList" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地址:"> 
              <el-input v-model="formData.address" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地区id:">
                 <el-input v-model.number="formData.areaId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="经度:">
                 <el-input v-model.number="formData.lng" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="纬度:">
                 <el-input v-model.number="formData.lat" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="开始时间:">
                <el-date-picker v-model="formData.beginTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="结束时间:">
                <el-date-picker v-model="formData.endTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="直播开始:">
                <el-date-picker v-model="formData.liveTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="直播结束:">
                <el-date-picker v-model="formData.liveEnd" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="主持/讲师:"> 
              <el-input v-model="formData.presenter" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="客服电话:"> 
              <el-input v-model="formData.phoneKefu" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="合作电话:"> 
              <el-input v-model="formData.phoneHezu" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="微信:"> 
              <el-input v-model="formData.wx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="QQ:"> 
              <el-input v-model="formData.qq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="群组id:"> 
              <el-input v-model="formData.groupId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="结束报名:">
                <el-date-picker v-model="formData.endJoinTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="票价:">
                 <el-input v-model.number="formData.price" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="vip票价:">
                 <el-input v-model.number="formData.priceVip" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="票价描述:"> 
              <el-input v-model="formData.priceDesc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="退费类型:">
                 <el-input v-model.number="formData.refundType" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="退费天:">
                 <el-input v-model.number="formData.refundDays" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否显示人数:">
                 <el-input v-model.number="formData.beShowjoin" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="活动状态:">
                 <el-input v-model.number="formData.statusAct" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="最大报名人数:">
                 <el-input v-model.number="formData.maxJoin" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前报名人数:">
                 <el-input v-model.number="formData.nowJoin" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否精选:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beChosen" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="是否周末:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beWeek" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="是否会员:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beVip" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="综合指数:">
                 <el-input v-model.number="formData.totalWhole" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总分享:">
                 <el-input v-model.number="formData.totalShare" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总收藏:">
                 <el-input v-model.number="formData.totalFav" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总报名数:">
                 <el-input v-model.number="formData.totalJoin" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总评论:">
                 <el-input v-model.number="formData.totalDiscuss" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总点击:">
                 <el-input v-model.number="formData.totalClick" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总评:">
                 <el-input v-model.number="formData.totalStar" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评1:">
                 <el-input v-model.number="formData.totalStar1" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评2:">
                 <el-input v-model.number="formData.totalStar2" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评3:">
                 <el-input v-model.number="formData.totalStar3" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="审核原因:"> 
              <el-input v-model="formData.statusMsg" clearable placeholder="请输入" />
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
  createActAct, 
  deleteActActByIds,
  updateActAct,
  findActAct,
  getActActList,
  quickEdit,
  excelList
} from '@/api/actAct' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'ActAct',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getActActList,   
      excelListApi: excelList,
      statusOptions: [],
      formData: {
            userId: 0,
            catId: 0,
            catIdUser: 0,
           beOnline: false,
           actType: '',
           beMulti: false,
            phase: 0,
           title: '',
           desc: '',
           avater: '',
           mediaList: '',
           address: '',
            areaId: 0,
            lng: 0,
            lat: 0,
            beginTime: new Date(),
            endTime: new Date(),
            liveTime: new Date(),
            liveEnd: new Date(),
           presenter: '',
           phoneKefu: '',
           phoneHezu: '',
           wx: '',
           qq: '',
           groupId: '',
            endJoinTime: new Date(),
            price: 0,
            priceVip: 0,
           priceDesc: '',
            refundType: 0,
            refundDays: 0,
            beShowjoin: 0,
            statusAct: 0,
            maxJoin: 0,
            nowJoin: 0,
           beChosen: false,
           beWeek: false,
           beVip: false,
            totalWhole: 0,
            totalShare: 0,
            totalFav: 0,
            totalJoin: 0,
            totalDiscuss: 0,
            totalClick: 0,
            totalStar: 0,
            totalStar1: 0,
            totalStar2: 0,
            totalStar3: 0,
            status: 0,
           statusMsg: '',
            mapData: {}
      } 
    }
  },
  
  async created() {
    await this.getDict('status')
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
        //this.deleteActAct(row)
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
     const res = await deleteActActByIds({ ids })
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
			this.$router.push({ name:'actActForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'actActForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findActAct({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.actAct 
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
          res = await createActAct(this.formData);
          break
        case "update": 
          res = await updateActAct(this.formData);
          break
        default: 
          res = await createActAct(this.formData);
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
