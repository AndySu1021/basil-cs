<template>
  <div class="app-panel-container">
    <div class="app-panel-title">{{title}}</div>
    <div v-if="hasFilter" class="app-panel-filter">
      <slot name="filter" />
    </div>
    <div class="app-panel-action">
      <slot name="action" />
    </div>
    <div class="app-panel-body">
      <slot name="body" />
    </div>
  </div>
</template>

<script>
export default {
  name: "AppPanel",
  props: {
    title: String,
    hasFilter: {
      type: Boolean,
      default: true,
    },
  }
}
</script>

<style lang="scss" scoped>
.app-panel-container {
  .app-panel-title {
    font-size: 28px;
    font-weight: bold;
  }
  .app-panel-filter{
    margin-top: 20px;
    background-color: white;
    padding: 25px;
    border-radius: 6px;
    & > * {
      margin-left: 10px;
    }
    & > *:first-child {
      margin-left: 0;
    }
  }
  .app-panel-action {
    margin-top: 20px;
  }
  .app-panel-body {
    margin-top: 20px;
    border-radius: 6px;
    padding: 20px;
    background-color: white;
    & > * {
      margin-top: 20px;
    }
    & > *:first-child {
      margin-top: 0;
    }
  }
}
</style>