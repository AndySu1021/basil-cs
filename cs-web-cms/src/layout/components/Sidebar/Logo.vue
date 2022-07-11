<template>
  <div class="sidebar-logo-container" :class="{ collapse: collapse }">
    <transition name="sidebarLogoFade">
      <router-link
        v-if="collapse"
        key="collapse"
        class="sidebar-logo-link"
        to="/"
      >
         <img src="@/assets/logo.svg" class="sidebar-logo" />
      </router-link>
      <router-link v-else key="expand" class="sidebar-logo-link" to="/">
        <img src="@/assets/logo.svg" class="sidebar-logo" style="margin-right: 8px;"/>
        <span class="sidebar-logo--text">客服中心</span>
      </router-link>
    </transition>
  </div>
</template>

<script>
export default {
  name: 'SidebarLogo',
  props: {
    collapse: {
      type: Boolean,
      required: true
    }
  }
}
</script>

<style lang="scss" scoped>
.sidebarLogoFade-enter-active {
  transition: opacity 1.5s;
}

.sidebarLogoFade-enter,
.sidebarLogoFade-leave-to {
  opacity: 0;
}

.sidebar-logo-container {
  position: relative;
  width: 100%;
  height: 50px;
  line-height: 50px;
  background: #2b2f3a;
  text-align: center;
  overflow: hidden;

  & .sidebar-logo-link {
    width: 100%;
    padding: 0 10px;

    & .sidebar-logo {
      width: 35px;
      vertical-align: middle;
      padding: 4px;
      display: inline-block;
    }
    & .sidebar-logo--text {
      width: calc(100% - 35px);
      color: #fff;
      font-size: 16px;
      font-weight: bold;
      vertical-align: middle;
    }
  }
}
</style>
