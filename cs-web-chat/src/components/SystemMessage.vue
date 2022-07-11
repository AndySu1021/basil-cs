<template>
<div class="system-message-container">
  <div v-if="opType === 7" class="system-message">諮詢已結束</div>
  <div v-else-if="opType === 9" class="system-message">已轉接客服</div>
  <div v-else class="system-message">{{content}}</div>
</div>
</template>

<script>
export default {
  name: "SystemMessage",
  props: {
    opType: Number,
    content: String
  }
}
</script>

<style lang="scss" scoped>
.system-message-container {
  width: 100%;
  text-align: center;
  .system-message {
    display: inline-block;
    padding: 6px 20px;
    font-size: 16px;
    background-color: #c3c3c3;
    color: #fff;
    font-weight: bold;
    border-radius: 20px;
  }
}
</style>