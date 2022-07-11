<template>
  <div class="app-container" v-loading.fullscreen.lock="loading">
    <AppPanel title="參數配置" :hasFilter="false">
      <template v-slot:body>
        <el-form
            ref="dialogForm"
            :model="form"
            :rules="rules"
            label-width="130px"
            label-position="left"
        >
          <el-form-item label="服務上限人數:" prop="max_member">
            <el-input
                v-model="form.max_member"
                style="width: 140px"
            >
              <template slot="append">人</template>
            </el-input>
          </el-form-item>
          <el-form-item label="用戶閒置時間:" prop="member_pending_expire">
            <el-input
                v-model="form.member_pending_expire"
                style="width: 140px"
            >
              <template slot="append">分鐘</template>
            </el-input>
          </el-form-item>
          <el-form-item label="問候語:" prop="greeting_text">
            <el-input
                type="textarea"
                v-model="form.greeting_text"
                style="width: 300px"
                rows="5"
            />
          </el-form-item>
        </el-form>
        <el-button type="primary" @click="submitForm">保存</el-button>
      </template>
    </AppPanel>
  </div>
</template>

<script>
import AppPanel from '@/components/AppPanel'
import {apiGetCsConfig, apiUpdateCsConfig} from "@/api/cs-config";

export default {
  name: 'ConfigSetting',
  components: {
    AppPanel,
  },
  data() {
    return {
      testNumber: 1000000,
      loading: false,
      form: {
        max_member: 0,
        member_pending_expire: 0,
        greeting_text: "",
      },
      rules: {
        max_member: [{ required: true, message: '必填', trigger: 'blur' }],
        member_pending_expire: [{ required: true, message: '必填', trigger: 'blur' }],
        greeting_text: [{ required: true, message: '必填', trigger: 'blur' }],
      },
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    submitForm() {
      this.$refs.dialogForm.validate(async (valid) => {
        if (valid) {
          this.loading = true
          try {
            this.form.max_member = parseInt(this.form.max_member, 10)
            this.form.member_pending_expire = parseInt(this.form.member_pending_expire, 10)
            await apiUpdateCsConfig(this.form)
            this.$showSuccessMessage("保存成功")
            await this.fetchData()
            this.loading = false
          } catch (err) {
            console.log(err)
            this.loading = false
          }
        }
      })
    },
    async fetchData() {
      try {
        this.loading = true
        const { data } = await apiGetCsConfig()
        this.form = {...data}
        this.loading = false
      } catch (err) {
        console.log(err)
        this.loading = false
      }
    },
  },
}
</script>

<style lang="scss" scoped>
.app-container {
  ::v-deep .el-form-item__label {
    font-size: 16px;
  }
}
</style>
