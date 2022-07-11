<template>
  <div>
    <AppDialog
        ref="Dialog"
        title="新增管理組"
        top="12vh"
    >
      <div slot="body">
        <el-form
            ref="dialogForm"
            label-width="80px"
        >
          <el-form-item label="名稱:" prop="name">
            <el-input
                size="small"
                style="width: 300px"
            />
          </el-form-item>
          <el-form-item label="用戶名:" prop="username">
            <el-input
                size="small"
                style="width: 300px"
            />
          </el-form-item>
          <el-form-item label="密碼:" prop="password">
            <el-input
                size="small"
                style="width: 300px"
            />
          </el-form-item>
          <el-form-item label="狀態:" prop="status">
            <el-switch
                :active-value="'Enabled'"
                :inactive-value="'Disabled'"
            />
          </el-form-item>
          <el-form-item label="密碼:" prop="password">
            <el-input
                size="small"
                style="width: 300px"
            />
          </el-form-item>
        </el-form>
      </div>
      <div slot="footer">
        <el-button type="primary" plain @click="handleCancel"> 取消 </el-button>
        <el-button type="primary" @click="submitForm"> 確認 </el-button>
      </div>
    </AppDialog>
  </div>
</template>

<script>
import AppDialog from "@/components/AppDialog";

export default {
  name: "AddDialog",
  components: {
    AppDialog
  },
  methods: {
    show() {
      this.$refs.Dialog.show()
    }
  }
}
</script>

<style scoped>

</style>