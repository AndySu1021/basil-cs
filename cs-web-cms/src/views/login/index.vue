<template>
  <div class="login-container">
    <div class="login-text">客服後台</div>
    <el-form
      ref="loginForm"
      :model="loginForm"
      :rules="rules"
      class="login-form"
      autocomplete="off"
      label-position="left"
    >
      <el-form-item prop="username">
        <span class="svg-container">
          <svg-icon icon-class="user" />
        </span>
        <el-input
          v-model="loginForm.username"
          placeholder="帳號"
          name="username"
          type="text"
          tabindex="1"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item prop="password">
        <span class="svg-container">
          <svg-icon icon-class="password" />
        </span>
        <el-input
          :key="passwordType"
          ref="password"
          v-model="loginForm.password"
          :type="passwordType"
          placeholder="密碼"
          tabindex="2"
          autocomplete="off"
          @keyup.enter.native="handleLogin"
        />
        <span class="show-pwd" @click="showPwd">
          <svg-icon
            :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'"
          />
        </span>
      </el-form-item>
      <el-button
        :loading="loading"
        type="primary"
        style="width: 100%; margin-bottom: 30px"
        @click.native.prevent="handleLogin"
      >
        登入
      </el-button>
    </el-form>
    <div class="footer">
      <p>Copyright ©2022 BasilCS. All rights reserved.</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      loginForm: {
        username: '',
        password: '',
      },
      rules: {
        username: [{ required: true, message: '必填', trigger: 'blur' }],
        password: [{ required: true, message: '必填', trigger: 'blur' }]
      },
      loading: false,
      passwordType: 'password',
    }
  },
  methods: {
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.$store.dispatch('user/login', this.loginForm)
            .then(() => {
              this.$router.push({ path: '/' })
              this.loading = false
            })
            .catch(() => {
              this.loading = false
            })
        } else {
          return false
        }
      })
    },
  }
}
</script>

<style lang="scss">
$bg: #283443;
$light_gray: #fff;
$cursor: #fff;

@supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
  .login-container .el-input input {
    color: $cursor;
  }
}

/* reset element-ui css */
.login-container {
  .el-input {
    display: inline-block;
    height: 47px;
    width: 85%;

    input {
      background: transparent;
      border: 0px;
      -webkit-appearance: none;
      border-radius: 0px;
      padding: 12px 5px 12px 15px;
      color: $light_gray;
      height: 47px;
      caret-color: $cursor;

      &:-webkit-autofill {
        -webkit-box-shadow: none;
        transition: background-color 5000s ease-in-out 0s;
        -webkit-text-fill-color: #fff;
      }
    }
  }

  .el-form-item {
    background: rgba(255, 255, 255, 0.1);
    border-radius: 5px;
    color: #454545;
  }

  .el-form-item__error {
    left: 60px;
  }

  .el-button--primary {
    color: #454545;
    background-color: #fff;
    border-color: #fff;
    font-size: 16px;
  }
}
</style>

<style lang="scss" scoped>
$bg: #2d3a4b;
$dark_gray: #889aa4;
$light_gray: #eee;

.login-container {
  min-height: 100%;
  width: 100%;
  background-color: $bg;
  overflow: hidden;

  .login-form {
    position: relative;
    width: 520px;
    max-width: 100%;
    padding: 20px 35px 0;
    margin: 0 auto;
    overflow: hidden;
  }

  .tips {
    font-size: 14px;
    color: #fff;
    margin-bottom: 10px;

    span {
      &:first-of-type {
        margin-right: 16px;
      }
    }
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $dark_gray;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }

  .show-pwd {
    position: absolute;
    right: 10px;
    top: 7px;
    font-size: 16px;
    color: $dark_gray;
    cursor: pointer;
    user-select: none;
  }

  .login-text {
    margin-top: 10%;
    margin-bottom: 30px;
    color: white;
    font-weight: bold;
    font-size: 40px;
    text-align: center;
  }

  .footer {
    color: white;
    font-size: 12px;
    font-weight: 400;
    text-align: center;
    letter-spacing: 2px;
  }

  .show-pwd {
    cursor: pointer;
  }
}
</style>
