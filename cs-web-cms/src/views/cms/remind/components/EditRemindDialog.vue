<template>
  <AppDialog ref="dialog" title="編輯提醒" width="500px">
    <div slot="body">
      <el-form
          ref="dialogForm"
          :model="formData"
          :rules="rules"
          label-width="80px"
      >
        <el-form-item label="標題:" prop="title">
          <el-input
              v-model="formData.title"
              size="small"
              style="width: 350px"
              placeholder="請輸入標題"
          />
        </el-form-item>
        <el-form-item label="內容:" prop="content">
          <el-input
              style="width: 350px"
              v-model="formData.content"
              size="small"
              type="textarea"
              placeholder="請輸入內容"
              resize="none"
              :maxlength="200"
              show-word-limit
              :rows="7"
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
import AppDialog from '@/components/AppDialog'
import {apiGetStaffDetail, apiUpdateStaff} from "@/api/staff";
import {apiGetRoleList} from "@/api/role";
import {apiGetRemindDetail, apiUpdateRemind} from "@/api/remind";

export default {
  name: 'EditRemindDialog',
  components: {
    AppDialog
  },
  data() {
    return {
      remindId: 0,
      formData: {
        title: '',
        content: '',
        status: 1
      },
      rules: {
        title: [{ required: true, message: '必填', trigger: 'blur' }],
        content: [{ required: true, message: '必填', trigger: 'blur' }],
      },
    }
  },
  methods: {
    show(id) {
      this.remindId = id
      this.fetchData()
      this.$refs.dialog.show()
    },
    async fetchData() {
      try {
        this.$refs.dialog.toggleLoadingDialog()
        const { data } = await apiGetRemindDetail(this.remindId)
        this.formData.title = data.title
        this.formData.content = data.content
        this.formData.status = data.status
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
            await apiUpdateRemind(this.remindId, this.formData)
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
