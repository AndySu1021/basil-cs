<template>
  <AppDialog ref="dialog" title="新增標籤" width="500px">
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
      <el-button @click="handleCancel">取消</el-button>
      <el-button type="primary" @click="submitForm">確認</el-button>
    </div>
  </AppDialog>
</template>

<script>
import AppDialog from '@/components/AppDialog'
import { deepCopy } from '@/utils'
import {apiCreateTag} from "@/api/tag";

export default {
  name: 'AddTagDialog',
  components: {
    AppDialog
  },
  data() {
    return {
      initialFormData: {
        name: '',
        status: 1
      },
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
            await apiCreateTag(this.formData)
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
