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
                <el-form-item label="用户名">
                  <el-input placeholder="搜索条件" v-model="searchInfo.username" clearable />
                </el-form-item> 
                <el-form-item label="邮件">
                  <el-input placeholder="搜索条件" v-model="searchInfo.email" clearable />
                </el-form-item> 
                <el-form-item label="手机">
                  <el-input placeholder="搜索条件" v-model="searchInfo.mobile" clearable />
                </el-form-item> 
                <el-form-item label="昵称">
                  <el-input placeholder="搜索条件" v-model="searchInfo.nickname" clearable />
                </el-form-item> 
                <el-form-item label="真实名">
                  <el-input placeholder="搜索条件" v-model="searchInfo.realname" clearable />
                </el-form-item> 
                <el-form-item label="身份证">
                  <el-input placeholder="搜索条件" v-model="searchInfo.cardId" clearable />
                </el-form-item>
                <el-form-item label="性别" prop="sex">                
                    <el-select v-model="searchInfo.sex" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in sexOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
              <el-form-item label="验证手机" prop="mobileValidated">
              <el-select v-model="searchInfo.mobileValidated" clearable placeholder="请选择">
                  <el-option
                      key="true"
                      label="是"
                      value="true">
                  </el-option>
                  <el-option
                      key="false"
                      label="否"
                      value="false">
                  </el-option>
             </el-select>
          </el-form-item>
              <el-form-item label="验证邮箱" prop="emailValidated">
              <el-select v-model="searchInfo.emailValidated" clearable placeholder="请选择">
                  <el-option
                      key="true"
                      label="是"
                      value="true">
                  </el-option>
                  <el-option
                      key="false"
                      label="否"
                      value="false">
                  </el-option>
             </el-select>
          </el-form-item>
              <el-form-item label="验证实名" prop="cardidValidated">
              <el-select v-model="searchInfo.cardidValidated" clearable placeholder="请选择">
                  <el-option
                      key="true"
                      label="是"
                      value="true">
                  </el-option>
                  <el-option
                      key="false"
                      label="否"
                      value="false">
                  </el-option>
             </el-select>
          </el-form-item> 
                <el-form-item label="推荐码16位（自己的）">
                  <el-input placeholder="搜索条件" v-model="searchInfo.recommendCode" clearable />
                </el-form-item>
                <el-form-item label="状态" prop="status">                
                    <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="安全状态" prop="statusSafe">                
                    <el-select v-model="searchInfo.statusSafe" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in usersafe_typeOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
                  <el-form-item label="注册ip">
                      <el-input placeholder="搜索条件" v-model="searchInfo.regIp" clearable />
                  </el-form-item>
                  <el-form-item label="登录ip">
                      <el-input placeholder="搜索条件" v-model="searchInfo.loginIp" clearable />
                  </el-form-item>
              <el-form-item label="最后登录时间"> 
                <el-date-picker
                v-model="formData.loginTime"  
                type="datetimerange"
                format="YYYY-MM-DD HH:mm:ss"
                :shortcuts="shortcuts"
                range-separator="至"
                start-placeholder="开始日期"
                end-placeholder="结束日期"
              /> 
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
          <!--username  BeQuickEdit -->  
        <el-table-column label="用户名" prop="username" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.username"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('username',scope.row.ID,scope.row.username,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.username}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column> 
      <!--password BeHide --> 
      <!--passwordSlat BeHide -->
          <!--email  BeQuickEdit -->  
        <el-table-column label="邮件" prop="email" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.email"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('email',scope.row.ID,scope.row.email,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.email}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--mobile  BeQuickEdit -->  
        <el-table-column label="手机" prop="mobile" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.mobile"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('mobile',scope.row.ID,scope.row.mobile,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.mobile}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--nickname  BeQuickEdit -->  
        <el-table-column label="昵称" prop="nickname" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.nickname"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('nickname',scope.row.ID,scope.row.nickname,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.nickname}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--realname  BeQuickEdit -->  
        <el-table-column label="真实名" prop="realname" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.realname"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('realname',scope.row.ID,scope.row.realname,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.realname}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cardId  BeQuickEdit -->  
        <el-table-column label="身份证" prop="cardId" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cardId"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('card_id',scope.row.ID,scope.row.cardId,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cardId}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sex  BeQuickEdit -->
        <el-table-column label="性别" prop="sex" width="120"  sortable="custom" >
        <template #default="scope">  
        <el-popover trigger="click" placement="top"  width = "280">  
              <el-select v-model="scope.row.sex" placeholder="请选择"  @change="quickEdit_do('sex',scope.row.ID,scope.row.sex,scope)">
                  <el-option v-for="(item,key) in sexOptions" :key="key" :label="item.label" :value="item.value"></el-option>
              </el-select> 
              <template #reference>
                  <div class="quickEdit" > {{filterDict(scope.row.sex,"sex")}} </div>
              </template>
            </el-popover>
        </template>  
        </el-table-column>
          <!--birthday  BeQuickEdit -->  
        <el-table-column label="生日" prop="birthday" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.birthday"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('birthday',scope.row.ID,scope.row.birthday,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.birthday}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--avatar  BeQuickEdit -->  
        <el-table-column label="头像" prop="avatar" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.avatar"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('avatar',scope.row.ID,scope.row.avatar,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.avatar}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--mobileValidated  BeQuickEdit -->
        <el-table-column label="验证手机" prop="mobileValidated" width="120"   sortable="custom"  >                        
            <template #default="scope" ><el-switch v-model="scope.row.mobileValidated" @change="quickEdit_do('mobile_validated',scope.row.ID,scope.row.mobileValidated,scope)"/></template> 
        </el-table-column>
          <!--emailValidated  BeQuickEdit -->
        <el-table-column label="验证邮箱" prop="emailValidated" width="120"   sortable="custom"  >                        
            <template #default="scope" ><el-switch v-model="scope.row.emailValidated" @change="quickEdit_do('email_validated',scope.row.ID,scope.row.emailValidated,scope)"/></template> 
        </el-table-column>
          <!--cardidValidated  BeQuickEdit -->
        <el-table-column label="验证实名" prop="cardidValidated" width="120"   sortable="custom"  >                        
            <template #default="scope" ><el-switch v-model="scope.row.cardidValidated" @change="quickEdit_do('cardid_validated',scope.row.ID,scope.row.cardidValidated,scope)"/></template> 
        </el-table-column>
          <!--info  BeQuickEdit -->  
        <el-table-column label="备注" prop="info" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.info"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('info',scope.row.ID,scope.row.info,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.info}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--recommendCode  BeQuickEdit -->  
        <el-table-column label="推荐码16位（自己的）" prop="recommendCode" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.recommendCode"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('recommend_code',scope.row.ID,scope.row.recommendCode,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.recommendCode}} </div>
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
          <!--statusSafe  BeQuickEdit -->
        <el-table-column label="安全状态" prop="statusSafe" width="120"  sortable="custom" >
        <template #default="scope">  
        <el-popover trigger="click" placement="top"  width = "280">  
              <el-select v-model="scope.row.statusSafe" placeholder="请选择"  @change="quickEdit_do('status_safe',scope.row.ID,scope.row.statusSafe,scope)">
                  <el-option v-for="(item,key) in usersafe_typeOptions" :key="key" :label="item.label" :value="item.value"></el-option>
              </el-select> 
              <template #reference>
                  <div class="quickEdit" > {{filterDict(scope.row.statusSafe,"usersafe_type")}} </div>
              </template>
            </el-popover>
        </template>  
        </el-table-column>
          <!--regIp  BeQuickEdit -->  
        <el-table-column label="注册ip" prop="regIp" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.regIp"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('reg_ip',scope.row.ID,scope.row.regIp,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.regIp}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--loginIp  BeQuickEdit -->  
        <el-table-column label="登录ip" prop="loginIp" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.loginIp"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('login_ip',scope.row.ID,scope.row.loginIp,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.loginIp}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--loginTotal  BeQuickEdit -->  
        <el-table-column label="登录次数" prop="loginTotal" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.loginTotal"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('login_total',scope.row.ID,scope.row.loginTotal,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.loginTotal}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--loginTime  BeQuickEdit -->  
        <el-table-column label="最后登录时间" prop="loginTime" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.loginTime"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('login_time',scope.row.ID,scope.row.loginTime,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.loginTime}} </div>
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
        <el-form-item label="用户名:"> 
              <el-input v-model="formData.username" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="密码:"> 
              <el-input v-model="formData.password" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="密码盐:"> 
              <el-input v-model="formData.passwordSlat" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮件:"> 
              <el-input v-model="formData.email" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机:"> 
              <el-input v-model="formData.mobile" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="昵称:"> 
              <el-input v-model="formData.nickname" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="真实名:"> 
              <el-input v-model="formData.realname" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="身份证:"> 
              <el-input v-model="formData.cardId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别:">
                 <el-select v-model="formData.sex" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in sexOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="生日:">
                <el-date-picker v-model="formData.birthday" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="头像:"> 
              <el-input v-model="formData.avatar" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="验证手机:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.mobileValidated" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="验证邮箱:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.emailValidated" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="验证实名:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.cardidValidated" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="备注:"> 
              <el-input v-model="formData.info" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="推荐码16位（自己的）:"> 
              <el-input v-model="formData.recommendCode" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="安全状态:">
                 <el-select v-model="formData.statusSafe" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in usersafe_typeOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="注册ip:">
                 <el-input v-model.number="formData.regIp" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="登录ip:">
                 <el-input v-model.number="formData.loginIp" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="登录次数:">
                 <el-input v-model.number="formData.loginTotal" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="最后登录时间:">
                <el-date-picker v-model="formData.loginTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
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
  createMemUser, 
  deleteMemUserByIds,
  updateMemUser,
  findMemUser,
  getMemUserList,
  quickEdit,
  excelList
} from '@/api/memUser' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'MemUser',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getMemUserList,   
      excelListApi: excelList,
      sexOptions: [],
      statusOptions: [],
      usersafe_typeOptions: [],
      formData: {
           username: '',
           password: '',
           passwordSlat: '',
           email: '',
           mobile: '',
           nickname: '',
           realname: '',
           cardId: '',
            sex: 0,
            birthday: new Date(),
           avatar: '',
           mobileValidated: false,
           emailValidated: false,
           cardidValidated: false,
           info: '',
           recommendCode: '',
            status: 0,
            statusSafe: 0,
            regIp: 0,
            loginIp: 0,
            loginTotal: 0,
            loginTime: new Date(),
            mapData: {}
      } 
    }
  },
  
  async created() {
    await this.getDict('sex')
    await this.getDict('status')
    await this.getDict('usersafe_type')
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
        //this.deleteMemUser(row)
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
     const res = await deleteMemUserByIds({ ids })
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
			this.$router.push({ name:'memUserForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'memUserForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findMemUser({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.memUser 
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
          res = await createMemUser(this.formData);
          break
        case "update": 
          res = await updateMemUser(this.formData);
          break
        default: 
          res = await createMemUser(this.formData);
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
