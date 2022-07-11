<template>
  <div class="member-message-container">
    <template v-if="contentType === 1">
      <div class="content-section">
        <div class="content-desc">
          <span class="member-name">{{name}}</span>
          <span class="member-time">{{getTime(timestamp)}}</span>
        </div>
        <div class="content-text">{{content | decrypt}}</div>
      </div>
    </template>
    <template v-if="contentType === 2">
      <div class="content-section">
        <div class="content-desc">
          <span class="member-name">{{name}}</span>
          <span class="member-time">{{getTime(timestamp)}}</span>
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
  </div>
</template>

<script>
import moment from 'moment';
import {decrypt} from "@/utils/crypto";

export default {
  name: "MemberMessage",
  props: {
    name: String,
    contentType: Number,
    content: String,
    timestamp: Number
  },
  methods: {
    getTime(timestamp) {
      return moment(timestamp*1000).format('HH:mm');
    },
    decryptContent(content) {
      return decrypt(content)
    },
  }
}
</script>

<style lang="scss" scoped>
.member-message-container {
  display: flex;
  align-items: flex-start;
  margin-bottom: 20px;
  .content-section {
    .content-desc {
      margin-bottom: 3px;
      .member-name {
        display: inline-block;
        color: #666;
        font-size: 14px;
        font-weight: bold;
        margin-right: 3px;
      }
      .member-time {
        display: inline-block;
        font-size: 13px;
        color: #999;
      }
    }
    .content-text {
      background-color: white;
      padding: 6px 10px;
      border-radius: 10px;
      position: relative;
      margin-right: 120px;
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
        max-width: 300px;
        margin-top: 3px;
        border-radius: 10px;
      }
    }
  }
}
</style>