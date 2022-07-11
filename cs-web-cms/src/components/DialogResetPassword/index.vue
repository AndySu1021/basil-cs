<template>
  <Dialog
    ref="dialog"
    :title="dialogTitle"
    width="600px"
    :modal-append-to-body="true"
  >
    <div slot="dialog-body">
      <el-form
        ref="dialogForm"
        :model="formData"
        :rules="rules"
        label-width="120px"
      >
        <el-form-item label="新密码:" prop="password">
          <el-input
            :key="passwordType"
            ref="password"
            v-model="formData.password"
            :type="passwordType"
            autocomplete="off"
            size="mini"
            style="width: 250px"
          />
          <span class="show-pwd" @click="showPassword">
            <svg-icon
              :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'"
            />
          </span>
        </el-form-item>
        <el-form-item label="确认新密码:" prop="rePassword">
          <el-input
            :key="rePasswordType"
            ref="rePassword"
            v-model="formData.rePassword"
            :type="rePasswordType"
            autocomplete="off"
            size="mini"
            style="width: 250px"
          />
          <span class="show-pwd" @click="showRePassword">
            <svg-icon
              :icon-class="rePasswordType === 'password' ? 'eye' : 'eye-open'"
            />
          </span>
        </el-form-item>
      </el-form>
    </div>
    <div slot="dialog-footer">
      <el-button type="primary" plain @click="handleCancel"> 取消 </el-button>
      <el-button type="primary" @click="submitForm"> 确认 </el-button>
    </div>
  </Dialog>
</template>

<script>
import Dialog from '@/components/Dialog'
import { deepCopy } from '@/utils'

export default {
  name: 'DialogResetPassword',
  components: {
    Dialog
  },
  props: {
    apiFunction: {
      type: Function,
      required: true
    }
  },
  data() {
    const validatePassword = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('必填'))
      } else if (
        this.formData.password.length < 6 ||
        this.formData.password.length > 16
      ) {
        callback(new Error('密码长度为6~16位'))
      } else {
        callback()
      }
    }
    const validateRePassword = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('必填'))
      } else if (this.formData.password !== value) {
        callback(new Error('请确认输入一致'))
      } else {
        callback()
      }
    }
    return {
      dialogTitle: '重置密码',
      initialFormData: {
        id: '',
        password: '',
        rePassword: '',
        pass: ''
      },
      formData: {},
      rules: {
        password: [
          { required: true, validator: validatePassword, trigger: 'blur' }
        ],
        rePassword: [
          { required: true, validator: validateRePassword, trigger: 'blur' }
        ]
      },
      passwordType: 'password',
      rePasswordType: 'password'
    }
  },
  methods: {
    show(payload) {
      const { id } = payload
      this.formData = deepCopy(this.initialFormData)
      this.formData.id = id
      this.$refs.dialog.show()
    },
    showPassword() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
    showRePassword() {
      if (this.rePasswordType === 'password') {
        this.rePasswordType = ''
      } else {
        this.rePasswordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.rePassword.focus()
      })
    },
    handleCancel() {
      this.$refs.dialog.hide()
    },
    submitForm() {
      this.$refs.dialogForm.validate(async (valid) => {
        if (valid) {
          this.$refs.dialog.toggleLoadingFullScreen()
          this.formData.pass = this.formData.password
          try {
            const { message } = await this.apiFunction(this.formData)
            this.$showSuccessMessage(message, this.$refs.dialog.afterSubmit)
          } catch (err) {
            this.$refs.dialog.toggleLoadingFullScreen()
          }
        }
      })
    }
  },
  created() {
    this.formData = deepCopy(this.initialFormData)
  }
}
</script>
<style lang="scss" scoped>
.show-pwd {
  margin-left: 10px;
  cursor: pointer;
}
</style>
