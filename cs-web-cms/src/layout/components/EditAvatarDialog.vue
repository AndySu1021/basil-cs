<template>
  <AppDialog ref="dialog" title="選擇大頭貼" width="600px">
    <div slot="body">
      <div class="avatar-container">
        <div class="avatar-item" v-for="(avatar,idx) in avatars" :key="idx" :class="selected === avatar ? 'is-selected' : ''" @click="handleClick(avatar)">
          <img :src="require('@/assets/'+avatar)"/>
        </div>
      </div>
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
import {apiUpdateStaffAvatar} from "@/api/staff";

export default {
  name: 'EditAvatarDialog',
  components: {
    AppDialog
  },
  data() {
    return {
      selected: "",
      avatars: [
          "avatar_1.jpg",
          "avatar_2.jpg",
          "avatar_3.jpg",
          "avatar_4.jpg",
          "avatar_5.jpg",
      ],
      rules: {
        name: [{ required: true, message: '必填', trigger: 'blur' }],
        username: [{ required: true, message: '必填', trigger: 'blur' }],
        password: [{ required: true, message: '必填', trigger: 'blur' }],
      },
    }
  },
  methods: {
    show() {
      this.$refs.dialog.show()
    },
    handleCancel() {
      this.$refs.dialog.hide()
    },
    async submitForm() {
      this.$refs.dialog.toggleLoadingFullScreen()
      try {
        await apiUpdateStaffAvatar({ avatar: this.selected })
        this.$store.state.user.avatar = this.selected
        this.$showSuccessMessage("設置成功", this.afterSubmit)
      } catch (error) {
        this.$refs.dialog.toggleLoadingFullScreen()
      }
    },
    handleClick(avatar) {
      this.selected = avatar
    },
    afterSubmit() {
      this.$refs.dialog.afterSubmit()
    }
  }
}
</script>

<style lang="scss" scoped>
.avatar-container {
  width: 80%;
  margin: auto;
  display: flex;
  justify-content: space-between;
  .avatar-item {
    display: inline-block;
    padding: 3px 3px 1px 3px;
    border-radius: 50%;
    border: 3px solid transparent;
    img {
      width: 70px;
      height: 70px;
      border-radius: 50%;
      opacity: 0.3;
    }
    &:hover {
      img {
        opacity: 1;
      }
    }
    &.is-selected {
      img {
        opacity: 1;
      }
    }
  }
}
</style>