<template>
  <AppDialog ref="dialog" title="新增通知" width="500px">
    <div slot="body">
      <el-form
        ref="dialogForm"
        :model="formData"
        :rules="rules"
        label-width="100px"
      >
        <el-form-item label="標題:" prop="title">
          <el-input
            v-model="formData.title"
            size="small"
            style="width: 300px"
          />
        </el-form-item>
        <el-form-item label="內容:" prop="content">
          <el-input
              style="width: 300px"
              v-model="formData.content"
              placeholder="請輸入內容"
              size="small"
              :autosize="true"
              type="textarea"
              resize="none"
              maxlength="60"
          />
        </el-form-item>
        <el-form-item label="開始時間:" prop="start_at">
          <el-date-picker
              v-model="formData.start_at"
              type="datetime"
              placeholder="請選擇時間"
              size="small"
              value-format="yyyy-MM-dd HH:mm:ss"
          />
        </el-form-item>
        <el-form-item label="結束時間:" prop="end_at">
          <el-date-picker
              v-model="formData.end_at"
              type="datetime"
              placeholder="請選擇時間"
              size="small"
              value-format="yyyy-MM-dd HH:mm:ss"
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
import {apiCreateNotice} from "@/api/notice";

export default {
  name: 'AddNoticeDialog',
  components: {
    AppDialog
  },
  data() {
    return {
      initialFormData: {
        title: '',
        content: '',
        start_at: '',
        end_at: '',
        status: 1
      },
      formData: {
        title: '',
        content: '',
        start_at: '',
        end_at: '',
        status: 1
      },
      rules: {
        title: [{ required: true, message: '必填', trigger: 'blur' }],
        content: [{ required: true, message: '必填', trigger: 'blur' }],
        start_at: [{ required: true, message: '必填', trigger: 'blur' }],
        end_at: [{ required: true, message: '必填', trigger: 'blur' }],
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
            await apiCreateNotice(this.formData)
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
