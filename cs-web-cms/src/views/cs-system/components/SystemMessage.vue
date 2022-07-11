<template>
<div class="system-message-container">
  <div v-if="opType === 3" class="system-message">已發送評分請求</div>
  <div v-else-if="opType === 4" class="system-message">已完成評分</div>
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
    padding: 10px 20px;
    font-size: 15px;
    background-color: #c4c4c4;
    color: #fff;
    font-weight: bold;
    border-radius: 30px;
  }
}
</style>