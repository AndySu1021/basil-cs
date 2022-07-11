<template>
  <AppDialog ref="dialog" title="編輯站點" width="500px">
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
              style="width: 350px"
              placeholder="請輸入名稱"
          />
        </el-form-item>
        <el-form-item label="代號:" prop="code">
          <el-input
              v-model="formData.code"
              size="small"
              style="width: 350px"
              placeholder="請輸入代號"
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
import {apiGetSiteDetail, apiUpdateSite} from "@/api/site";

export default {
  name: 'EditSiteDialog',
  components: {
    AppDialog
  },
  data() {
    return {
      siteId: 0,
      formData: {
        name: '',
        code: '',
        status: 1
      },
      rules: {
        name: [{ required: true, message: '必填', trigger: 'blur' }],
        code: [{ required: true, message: '必填', trigger: 'blur' }],
      },
    }
  },
  methods: {
    show(id) {
      this.siteId = id
      this.fetchData()
      this.$refs.dialog.show()
    },
    async fetchData() {
      try {
        this.$refs.dialog.toggleLoadingDialog()
        const { data } = await apiGetSiteDetail(this.siteId)
        this.formData.name = data.name
        this.formData.code = data.code
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
            await apiUpdateSite(this.siteId, this.formData)
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
