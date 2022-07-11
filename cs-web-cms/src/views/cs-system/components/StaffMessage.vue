<template>
  <div class="staff-message-container">
    <template v-if="contentType === 1">
      <div class="top-section">
        <div class="staff-time">{{getTime(timestamp)}}</div>
        <div class="staff-name">{{name}}</div>
      </div>
      <div class="content-text">{{decryptContent(content)}}</div>
    </template>
    <template v-if="contentType === 2">
      <div class="top-section">
        <div class="staff-time">{{getTime(timestamp)}}</div>
        <div class="staff-name">{{name}}</div>
      </div>
      <div class="content-text">
        <el-image
            style="width: 100%;"
            :src="decryptContent(content)"
            :preview-src-list="[decryptContent(content)]">
        </el-image>
      </div>
    </template>
  </div>
</template>

<script>
import moment from "moment";
import {decrypt} from "@/utils/crypto";

export default {
  name: "StaffMessage",
  props: {
    name: String,
    contentType: Number,
    content: String,
    timestamp: Number,
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
.staff-message-container {
  display: flex;
  flex-direction: column-reverse;
  align-items: flex-end;
  justify-content: flex-end;
  margin-bottom: 20px;
  .top-section{
    margin-bottom: 3px;
    .staff-time {
      display: inline-block;
      margin-right: 3px;
      margin-bottom: 3px;
      font-size: 13px;
      color: #999;
    }
    .staff-name {
      display: inline-block;
      color: #666;
      font-size: 14px;
      font-weight: bold;
      margin-right: 3px;
    }
  }
  .content-text {
    color: white;
    padding: 6px 10px;
    border-radius: 10px;
    position: relative;
    margin-right: 5px;
    margin-left: 120px;
    order: -1;
    background-color: #06BEC9;
    font-weight: 500;
    &::before {
      border-left: 10px solid #06BEC9;
      right: -8px;
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
</style>