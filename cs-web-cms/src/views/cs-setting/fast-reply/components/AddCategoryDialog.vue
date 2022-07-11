<template>
  <AppDialog ref="dialog" title="新增快捷分類" width="500px">
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
        <el-form-item label="現有:">
          <template v-for="(category, idx) in categories">
            <el-tag :key="idx" type="warning" effect="dark" style="display: inline-block;margin-right: 14px;">{{category.value}}</el-tag>
          </template>
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
import {apiCreateCategory, apiGetCategoryList} from "@/api/fast-reply";

export default {
  name: 'AddCategoryDialog',
  components: {
    AppDialog
  },
  data() {
    return {
      categories: [],
      initialFormData: {
        name: "",
      },
      formData: {
        name: "",
      },
      rules: {
        name: [{ required: true, message: '必填', trigger: 'blur' }],
      },
    }
  },
  methods: {
    show() {
      this.formData = deepCopy(this.initialFormData)
      this.fetchCategories()
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
            await apiCreateCategory(this.formData)
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
      } catch (err) {
        console.log(err)
      } finally {
        this.loading = false
      }
    },
  }
}
</script>
