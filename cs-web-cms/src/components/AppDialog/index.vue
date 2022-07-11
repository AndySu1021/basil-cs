<template>
  <div class="app-dialog-container">
    <el-dialog
        v-loading.fullscreen.lock="loadingFullScreen"
        :visible.sync="dialogVisible"
        :top="top"
        :title="title"
        :width="width"
        :before-close="hide"
        :close-on-click-modal="false"
        :close-on-press-escape="true"
        :show-close="false"
        @closed="afterClose"
        center
        :modal-append-to-body="false"
    >
      <div v-loading="loadingDialog">
        <div class="app-dialog-body">
          <slot name="body" />
        </div>
        <div class="app-dialog-footer">
          <slot name="footer" />
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'index',
  props: {
    title: {
      type: String,
      required: true,
    },
    top: {
      type: String,
      default: "15vh",
    },
    width: {
      type: String,
      default: '450px',
    },
    beforeClose: {
      type: Function,
      default: () => () => {},
    },
  },
  data() {
    return {
      loadingFullScreen: false,
      loadingDialog: false,
      dialogVisible: false,
      dialogBodyVisible: false,
    };
  },
  methods: {
    show() {
      this.dialogBodyVisible = true;
      this.dialogVisible = true;
    },
    afterClose() {
      this.dialogBodyVisible = false;
    },
    hide() {
      this.dialogVisible = false;
      this.beforeClose();
    },
    toggleLoadingDialog() {
      this.loadingDialog = !this.loadingDialog;
    },
    toggleLoadingFullScreen() {
      this.loadingFullScreen = !this.loadingFullScreen;
    },
    afterSubmit() {
      this.loadingFullScreen = false;
      this.hide();
    },
  },
};
</script>

<style lang="scss" scoped>
.app-dialog-container{
  ::v-deep .el-dialog {
    border-radius: 8px;
    .el-dialog__header {
      background-color: #304156;
      margin-right: 0;
      border-radius: 8px 8px 0 0;
      padding: 10px;
      .el-dialog__title {
        color: #f4f5f6;
        font-weight: bolder;
        font-size: 16px;
      }
    }
    .el-dialog__body {
      padding: 25px;
    }
  }
  .app-dialog-footer {
    margin-top: 30px;
    text-align: center;
  }
}
</style>
