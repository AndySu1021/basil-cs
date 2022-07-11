<template>
  <AppDialog ref="dialog" title="編輯職員" width="500px" top="4vh">
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
import AppDialog from '@/components/AppDialog'
import {apiGetRoleDetail, apiUpdateRole} from "@/api/role";
import {permissionTree} from "@/utils/authorities";

export default {
  name: 'EditRoleDialog',
  components: {
    AppDialog
  },
  data() {
    return {
      roleId: 0,
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
    show(id) {
      this.roleId = id
      this.fetchData()
      this.$refs.dialog.show()
    },
    async fetchData() {
      try {
        this.$refs.dialog.toggleLoadingDialog()
        const { data } = await apiGetRoleDetail(this.roleId)
        this.$refs.PermissionTree.setCheckedKeys(data.permissions)
        this.formData.name = data.name
        this.formData.permissions = data.permissions
      } catch (err) {
        console.log(err)
      } finally {
        this.$refs.dialog.toggleLoadingDialog()
      }
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
            await apiUpdateRole(this.roleId, this.formData)
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
    }
  }
}
</script>
