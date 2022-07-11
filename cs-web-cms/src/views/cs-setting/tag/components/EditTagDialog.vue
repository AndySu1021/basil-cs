<template>
  <AppDialog ref="dialog" title="編輯標籤" width="500px" top="5vh">
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
      <el-button type="primary" plain @click="handleCancel">取消</el-button>
      <el-button type="primary" @click="submitForm">確認</el-button>
    </div>
  </AppDialog>
</template>

<script>
import AppDialog from '@/components/AppDialog'
import {apiGetTagDetail, apiUpdateTag} from "@/api/tag";

export default {
  name: 'EditTagDialog',
  components: {
    AppDialog
  },
  data() {
    return {
      tagId: 0,
      loading: false,
      formData: {
        name: '',
        status: 1
      },
      rules: {
        name: [{ required: true, message: '必填', trigger: 'blur' }],
      },
    }
  },
  methods: {
    show(id) {
      this.tagId = id
      this.fetchData()
      this.$refs.dialog.show()
    },
    async fetchData() {
      try {
        this.loading = true
        this.$refs.dialog.toggleLoadingDialog()
        const { data } = await apiGetTagDetail(this.tagId)
        this.formData.name = data.name
        this.formData.status = data.status
        this.$refs.dialog.toggleLoadingDialog()
      } catch (err) {
        console.log(err)
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
            await apiUpdateTag(this.tagId, this.formData)
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
