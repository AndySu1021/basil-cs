<template>
  <AppDialog ref="dialog" title="新增角色" width="500px" top="4vh">
    <div slot="body">
      <el-form
          ref="dialogForm"
          :model="formData"
          :rules="rules"
          label-width="80px"
      >
        <el-form-item label="名稱:" prop="name">
          <el-input
              v-model="formData.name"
              size="small"
              style="width: 300px"
          />
        </el-form-item>
        <el-divider content-position="center">權限</el-divider>
        <el-tree
            ref="PermissionTree"
            :data="data"
            show-checkbox
            node-key="id"
            :render-after-expand="true"
            :default-expand-all="true"
        >
        </el-tree>
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
import {permissionTree} from "@/utils/authorities"
import {apiCreateRole} from "@/api/role";

export default {
  name: 'AddRoleDialog',
  components: {
    AppDialog
  },
  data() {
    return {
      initialFormData: {
        name: '',
        permissions: [],
      },
      formData: {
        name: '',
        permissions: [],
      },
      rules: {
        name: [{ required: true, message: '必填', trigger: 'blur' }],
      },
      data: permissionTree,
    }
  },
  methods: {
    show() {
      this.formData = deepCopy(this.initialFormData)
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
            this.formData.permissions = this.$refs.PermissionTree.getCheckedKeys()
            let tmp = this.$refs.PermissionTree.getHalfCheckedKeys()
            for (let i = 0; i < tmp.length; i++) {
              tmp[i] = "-" + tmp[i]
            }
            this.formData.permissions = this.formData.permissions.concat(tmp)
            await apiCreateRole(this.formData)
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
  }
}
</script>
