<template>
  <AppDialog ref="dialog" title="新增常見問題" width="500px">
    <div slot="body">
      <el-form
        ref="dialogForm"
        :model="formData"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item label="問題:" prop="question">
          <el-input
            v-model="formData.question"
            size="small"
            style="width: 300px"
          />
        </el-form-item>
        <el-form-item label="答案:" prop="answer">
          <el-input
              v-model="formData.answer"
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
import {apiCreateFAQ} from "@/api/faq";

export default {
  name: 'AddFAQDialog',
  components: {
    AppDialog
  },
  data() {
    return {
      initialFormData: {
        question: '',
        answer: '',
        status: 1
      },
      formData: {
        question: '',
        answer: '',
        status: 1
      },
      rules: {
        question: [{ required: true, message: '必填', trigger: 'blur' }],
        answer: [{ required: true, message: '必填', trigger: 'blur' }],
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
            await apiCreateFAQ(this.formData)
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
