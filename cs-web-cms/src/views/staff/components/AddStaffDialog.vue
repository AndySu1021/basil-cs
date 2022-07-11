<template>
  <AppDialog ref="dialog" title="新增職員" width="500px">
    <div slot="body">
      <el-form
        ref="dialogForm"
        :model="formData"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item label="角色:" prop="roleID">
          <el-select v-model="formData.role_id" placeholder="請選擇角色">
            <el-option
                v-for="(role,idx) in roles"
                :key="idx"
                :label="role.name"
                :value="role.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="名稱:" prop="name">
          <el-input
            v-model="formData.name"
            size="small"
            style="width: 300px"
          />
        </el-form-item>
        <el-form-item label="用戶名:" prop="username">
          <el-input
              v-model="formData.username"
              size="small"
              style="width: 300px"
          />
        </el-form-item>
        <el-form-item label="密碼:" prop="password">
          <el-input
              v-model="formData.password"
              size="small"
              style="width: 300px"
          />
        </el-form-item>
        <el-form-item label="狀態:" prop="status">
          <el-switch
            v-model="formData.status"
            :active-value="1"
            :inactive-value="2"
          />
        </el-form-item>
      </el-form>
    </div>
    <div slot="footer">
      <el-button @click="handleCancel">取消</el-button>
      <el-button type="primary" @click="submitForm">確認</el-button>
    </div>
  </AppDialog>
</template>

<script>
import { deepCopy } from '@/utils'
import {apiCreateStaff} from "@/api/staff";
import AppDialog from "@/components/AppDialog";
import {apiGetAllRoles} from "@/api/role";

export default {
  name: 'AddStaffDialog',
  components: {
    AppDialog
  },
  data() {
    return {
      initialFormData: {
        role_id: null,
        name: '',
        username: '',
        password: '',
        status: 1
      },
      formData: {
        role_id: null,
        name: '',
        username: '',
        password: '',
        status: 1
      },
      rules: {
        role_id: [{ required: true, message: '必填', trigger: 'blur' }],
        name: [{ required: true, message: '必填', trigger: 'blur' }],
        username: [{ required: true, message: '必填', trigger: 'blur' }],
        password: [{ required: true, message: '必填', trigger: 'blur' }],
      },
      roles: [],
    }
  },
  methods: {
    show() {
      this.formData = deepCopy(this.initialFormData)
      this.fetchRoles()
      this.$refs.dialog.show()
    },
    handleCancel() {
      this.$refs.dialog.hide()
    },
    submitForm() {
      this.$refs.dialogForm.validate(async (valid) => {
        if (valid) {
          this.$refs.dialog.toggleLoadingFullScreen()
          try {
            await apiCreateStaff(this.formData)
            this.$showSuccessMessage("操作成功", this.afterSubmit)
          } catch (error) {
            this.$refs.dialog.toggleLoadingFullScreen()
          }
        }
      })
    },
    afterSubmit() {
      this.$refs.dialog.afterSubmit()
      this.$emit('flush-parent')
    },
    async fetchRoles() {
      try {
        this.loading = true
        const { data } = await apiGetAllRoles()
        this.roles = data
        this.loading = false
      } catch (err) {
        console.log(err)
        this.loading = false
      }
    },
  }
}
</script>
