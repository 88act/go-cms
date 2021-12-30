<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用户id:">
                 <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类型:">
                 <el-input v-model.number="formData.type" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="旧值:">
                 <el-input v-model.number="formData.valOld" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="新值:">
                 <el-input v-model.number="formData.valNew" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import {
  createMemUserSafe,
  updateMemUserSafe,
  findMemUserSafe
} from '@/api/memUserSafe' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑MemUserSafe',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      statusOptions: [],
      formData: {
            userId: 0,
            type: 0,
            valOld: 0,
            valNew: 0,
            status: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findMemUserSafe({ID:id})
      if (res.code === 200) {
        this.formData = res.data.memUserSafe
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
    await this.getDict('status') 
  },
  methods: {
    async save() {  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createMemUserSafe(this.formData)
          break
        case 'update':
          res = await updateMemUserSafe(this.formData)
          break
        default:
          res = await createMemUserSafe(this.formData)
          break
      }
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
         emitter.emit('closeThisPage') 
      }
    },
    back() {
      this.$router.go(-1)
      emitter.emit('closeThisPage') 
    }
  }
}
</script>
<style>
</style>
