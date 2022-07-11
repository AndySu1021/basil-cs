<template>
  <AppDialog ref="dialog" title="編輯快捷消息" width="500px">
    <div slot="body">
      <el-form
          ref="dialogForm"
          :model="formData"
          :rules="rules"
          label-width="80px"
      >
        <el-form-item label="分類:" prop="category_id">
          <el-select v-model="formData.category_id" placeholder="請選擇">
            <el-option
                v-for="(category,idx) in categories"
                :key="idx"
                :label="category.value"
                :value="category.id">
            </el-option>
          </el-select>
        </el-form-item>
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
              maxlength="250"
              show-word-limit
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
import {apiGetCategoryList, apiGetFastReplyDetail, apiUpdateFastReply} from "@/api/fast-reply";

export default {
  name: 'EditFastReplyDialog',
  components: {
    AppDialog
  },
  data() {
    return {
      id: 0,
      categories: [],
      formData: {
        category_id: null,
        title: '',
        content: '',
        status: 1
      },
      rules: {
        category_id: [{ required: true, message: '必填', trigger: 'blur' }],
        title: [{ required: true, message: '必填', trigger: 'blur' }],
        content: [{ required: true, message: '必填', trigger: 'blur' }],
      },
    }
  },
  methods: {
    show(id) {
      this.id = id
      this.fetchCategories()
      this.fetchData()
      this.$refs.dialog.show()
    },
    async fetchData() {
      try {
        this.$refs.dialog.toggleLoadingDialog()
        const { data } = await apiGetFastReplyDetail(this.id)
        this.formData.category_id = data.category_id
        this.formData.title = data.title
        this.formData.content = data.content
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
            await apiUpdateFastReply(this.id, this.formData)
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
    async fetchCategories() {
      try {
        this.loading = true
        const { data } = await apiGetCategoryList()
        this.categories = data
        this.loading = false
      } catch (err) {
        console.log(err)
        this.loading = false
      }
    },
  }
}
</script>
