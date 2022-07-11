<template>
  <div class="footer-container">
    <div class="input">
      <div id="input-area" class="input-area" @compositionstart="isComposing = true; this.haveChanged = true;" @compositionend="isComposing = false" @keyup.enter.exact="onEnter">
        <textarea v-model="message.content" placeholder="請輸入訊息..." :disabled="taDisabled"></textarea>
        <el-upload
            class="upload-btn"
            :action="url()"
            :headers="{'X-Token': sid()}"
            :data="{'type': 'staff'}"
            :show-file-list="false"
            accept="image/jpg,image/jpeg,image/png"
            :on-success="successCB"
        >
          <el-tooltip placement="top">
            <template #content>上傳圖片</template>
            <el-button icon="el-icon-picture" circle></el-button>
          </el-tooltip>
        </el-upload>
        <el-tooltip class="send-btn" content="發送" placement="top">
          <el-button type="success" icon="el-icon-s-promotion" circle @click="sendMessage"></el-button>
        </el-tooltip>
      </div>
    </div>
  </div>
</template>

<script>
import { sendSocketMessage } from "@/utils/ws";
import {mapGetters} from "vuex";
import {getToken} from "@/utils/storage";
import {encrypt} from "@/utils/crypto";

export default {
  name: "ChatBoxFooter",
  computed: {
    ...mapGetters({
      activeRoomId: 'cs/activeRoomId',
    }),
  },
  watch: {
    activeRoomId() {
      this.message.room_id = this.activeRoomId
      if(this.activeRoomId > 0) {
        this.taDisabled = false
      } else {
        this.taDisabled = true
      }
    }
  },
  data() {
    return {
      taDisabled: true,
      isComposing: false,
      isOK: true,
      haveChanged: false,
      message: {
        op_type: 2,
        room_id: this.activeRoomId,
        content_type: 1,
        content: ""
      }
    }
  },
  methods: {
    sendMessage() {
      this.message.content = this.message.content.trim()
      this.message.content = this.message.content.replace(/\r\n|\n/g,"");
      if(this.message.content !== "" && this.message.room_id > 0) {
        this.message.content = encrypt(this.message.content)
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
      return process.env.VUE_APP_HTTP_URL + "/upload"
    },
    sid() {
      return getToken()
    },
    successCB(response) {
      sendSocketMessage({
        op_type: 2,
        room_id: this.activeRoomId,
        content_type: 2,
        content: encrypt(response.data),
      })
    },
  }
}
</script>

<style lang="scss" scoped>
@import "../../../assets/global.scss";
.footer-container {
  width: 100%;
  text-align: center;
  .input {
    position: relative;
    .input-area {
      position: relative;
      background-color: white;
      border-radius: 0 0 6px 6px;
      text-align: left;
      textarea {
        border-radius: 10px;
        border: none;
        width: 100%;
        padding: 20px;
        resize: none;
        height: 20vh;
        font-size: 16px;
        letter-spacing: 0.3px;
        &:focus {
          outline: none !important;
        }
      }
      .upload-btn {
        position: absolute;
        bottom: 14px;
        right: 60px;
        border: none;
      }
      .send-btn {
        position: absolute;
        bottom: 14px;
        right: 16px;
        background-color: $primary;
        border: none;
      }
    }
  }
}
</style>