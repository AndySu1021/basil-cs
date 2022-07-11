<template>
  <div class="staff-message-container remote">
    <div class="avatar">
      <div class="pic">
        <img src="https://picsum.photos/100/100?random=12" />
      </div>
    </div>
    <template v-if="contentType === 1">
      <div class="content-section">
        <div class="content-desc">
          <span class="staff-name">{{name}}</span>
          <span class="staff-time">{{getTime(timestamp)}}</span>
        </div>
        <div v-if="opType === 0" class="content-text">{{content}}</div>
        <div v-else class="content-text">{{decryptContent(content)}}</div>
      </div>
    </template>
    <template v-if="contentType === 2">
      <div class="content-section">
        <div class="content-desc">
          <span class="staff-name">{{name}}</span>
          <span class="staff-time">{{getTime(timestamp)}}</span>
        </div>
        <div class="content-text">
          <el-image
              style="width: 100%;"
              :src="decryptContent(content)"
              :preview-src-list="[decryptContent(content)]">
          </el-image>
        </div>
      </div>
    </template>
    <template v-if="contentType === 99">
      <div class="content-section">
        <div class="content-desc">
          <span class="staff-name">{{name}}</span>
          <span class="staff-time">{{getTime(timestamp)}}</span>
        </div>
        <div class="content-text">
          <div>常見問題</div>
          <div v-for="(item, idx) in extra" :key="idx" class="question-text" @click="handleFAQ(item)">{{idx+1}}. {{item.question}}</div>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import moment from 'moment';
import {ElImage} from "element-plus";
import store from "@/store";
import {mapGetters} from "vuex";
import {decrypt, encrypt} from "@/utils/crypto";

export default {
  name: "StaffMessage",
  components: {ElImage},
  props: {
    name: String,
    opType: Number,
    contentType: Number,
    content: String,
    timestamp: Number,
    extra: null,
  },
  computed: {
    ...mapGetters({
      roomId: 'ws/roomId',
      memberName: "ws/memberName",
    })
  },
  methods: {
    getTime(timestamp) {
      return moment(timestamp*1000).format('HH:mm');
    },
    decryptContent(content) {
      return decrypt(content)
    },
    handleFAQ(item) {
      store.commit("ws/APPEND_MESSAGE", {
        content: encrypt(item.question),
        content_type: 1,
        op_type: 2,
        sender_type: 1,
        ts: moment().unix(),
      });
      store.commit("ws/APPEND_MESSAGE", {
        content: encrypt(item.answer),
        content_type: 1,
        extra: null,
        op_type: 2,
        sender_name: "FAQ",
        sender_type: 2,
        ts: moment().unix(),
      });
    },
  }
}
</script>

<style lang="scss" scoped>
.staff-message-container {
  display: flex;
  align-items: flex-start;
  margin-bottom: 20px;
  .avatar {
    width: 50px;
    text-align: center;
    flex-shrink: 0;
    margin-top: 4px;
    .pic {
      border-radius: 50%;
      overflow: hidden;
      img {
        width: 100%;
        vertical-align: middle;
      }
    }
  }
  .content-section {
    margin-left: 12px;
    .content-desc {
      .staff-name {
        display: inline-block;
        color: black;
        font-weight: bold;
        margin-right: 8px;
      }
      .staff-time {
        display: inline-block;
        font-size: 14px;
        color: #888888;
      }
    }
    .content-text {
      background-color: white;
      padding: 6px 10px;
      border-radius: 10px;
      position: relative;
      margin-right: 80px;
      color: #333333;
      font-weight: 500;
      &::before {
        border-right: 10px solid white;
        left: -8px;
        content: "";
        position: absolute;
        top: 9px;
        border-top: 8px solid transparent;
        border-bottom: 8px solid transparent;
      }
      .el-image {
        max-width: 270px;
        margin-top: 3px;
        border-radius: 10px;
      }
      .question-text {
        margin-top: 5px;
        color: dodgerblue;
        text-decoration: underline;
        &:hover {
          cursor: pointer;
        }
      }
    }
  }
}
</style>