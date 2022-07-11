<template>
  <div class="footer-container">
    <div class="input">
      <div id="input-area" class="input-area" @compositionstart="isComposing = true; this.haveChanged = true;" @compositionend="isComposing = false" @keyup.enter.exact="onEnter">
        <textarea v-model="message.content" placeholder="請輸入訊息..."></textarea>
      </div>
      <el-upload
          class="upload-btn"
          :action="url()"
          :headers="{'X-Token': sid()}"
          :data="{'type': 'member'}"
          :show-file-list="false"
          accept="image/jpg,image/jpeg,image/png"
          :on-success="successCB"
      >
        <el-tooltip placement="top">
          <template #content>上傳圖片</template>
          <PictureFilled style="width: 25px; height: 25px;color:#888;margin-top: 7px" />
        </el-tooltip>
      </el-upload>
      <div class="send-btn" @click="sendMessage">
        <promotion style="width: 22px; height: 22px; color: white;margin-top: 8px;margin-right: 2px;" />
      </div>
    </div>
  </div>
</template>

<script>
import {PictureFilled, Promotion} from '@element-plus/icons-vue'
import { sendSocketMessage } from "@/utils/ws";
import {ElTooltip, ElUpload} from "element-plus";
import {getSID} from "@/utils/storage";
import {encrypt} from "@/utils/crypto";

export default {
  name: "ChatBoxFooter",
  components: {Promotion, ElUpload, PictureFilled, ElTooltip},
  data() {
    return {
      isComposing: false,
      isOK: true,
      haveChanged: false,
      message: {
        op_type: 2,
        content_type: 1,
        content: "",
      }
    }
  },
  methods: {
    sendMessage() {
      this.message.content = this.message.content.trim()
      this.message.content = this.message.content.replace(/\r\n|\n/g,"");
      if(this.message.content !== "") {
        this.message.content = encrypt(this.message.content);
        sendSocketMessage(this.message)
        this.isOK = true
        this.haveChanged = false
      }
      this.message.content = ""
    },
    onEnter() {
      if(this.isComposing === false && this.haveChanged === false && this.isOK === true) {
        this.sendMessage()
      } else {
        this.haveChanged = false
        this.isOK = true
      }
    },
    url() {
      return process.env.VUE_APP_BASE_URL + "/upload"
    },
    sid() {
      return getSID()
    },
    successCB(response) {
      sendSocketMessage({
        op_type: 2,
        content_type: 2,
        content: encrypt(response.data),
      })
    },
  }
}
</script>

<style lang="scss" scoped>
@import "../assets/global.scss";
.footer-container {
  width: 100%;
  text-align: center;
  .input {
    position: relative;
    .message-input {
      width: 95%;
      padding-bottom: 15px;
    }
    .input-area {
      width: 100%;
      background-color: white;
      border-radius: 0 0 10px 10px;
      text-align: left;
      padding: 10px;
      textarea {
        border: none;
        width: 100%;
        padding: 10px;
        resize: none;
        height: 150px;
        font-size: 16px;
        letter-spacing: 0.5px;
        &:focus {
          outline: none !important;
        }
      }
    }
    .upload-btn {
      position: absolute;
      bottom: 20px;
      right: 70px;
      width: 40px;
      height: 40px;
      border-radius: 50%;
    }
    .send-btn {
      position: absolute;
      bottom: 20px;
      right: 20px;
      width: 40px;
      height: 40px;
      background-color: $primary;
      border-radius: 50%;
      &:hover {
        box-shadow: 0 0 6px #666;
      }
    }
  }
}
</style>