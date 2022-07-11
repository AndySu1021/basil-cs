<template>
  <div v-if="!item.hidden">
    <template
      v-if="
        hasOneShowingChild(item.children, item) &&
          (!onlyOneChild.children || onlyOneChild.noShowingChildren) &&
          !item.alwaysShow
      "
    >
      <app-link v-if="onlyOneChild.meta" :to="resolvePath(onlyOneChild.path)">
        <el-menu-item
            v-if="onlyOneChild.meta.key === 'CsSystem'"
            :index="resolvePath(onlyOneChild.path)"
            :class="{ 'submenu-title-noDropdown': !isNest, 'red-point': checkNewMessage(onlyOneChild.meta.key) }"
            @click="handleCheckMessage"
        >
          <item
              :icon="onlyOneChild.meta.icon || (item.meta && item.meta.icon)"
              :title="onlyOneChild.meta.title"
          />
        </el-menu-item>
        <el-menu-item
          v-else
          :index="resolvePath(onlyOneChild.path)"
          :class="{ 'submenu-title-noDropdown': !isNest }"
        >
          <item
              :icon="onlyOneChild.meta.icon || (item.meta && item.meta.icon)"
              :title="onlyOneChild.meta.title"
          />
        </el-menu-item>
      </app-link>
    </template>

    <el-submenu
      v-else
      ref="subMenu"
      :index="resolvePath(item.path)"
    >
      <template slot="title">
        <item
          v-if="item.meta"
          :icon="item.meta && item.meta.icon"
          :title="item.meta.title"
        />
      </template>
      <sidebar-item
        v-for="child in item.children"
        :key="child.path"
        :is-nest="true"
        :item="child"
        :base-path="resolvePath(child.path)"
        class="nest-menu"
      />
    </el-submenu>
  </div>
</template>

<script>
import path from 'path'
import { isExternal } from '@/utils/validate'
import Item from './Item'
import AppLink from './Link'
import FixiOSBug from './FixiOSBug'
import {mapGetters} from "vuex";

export default {
  name: 'SidebarItem',
  components: { Item, AppLink },
  mixins: [FixiOSBug],
  props: {
    // route object
    item: {
      type: Object,
      required: true
    },
    isNest: {
      type: Boolean,
      default: false
    },
    basePath: {
      type: String,
      default: ''
    }
  },
  computed: {
    ...mapGetters({
      newMessage: 'cs/newMessage',
    })
  },
  data() {
    // To fix https://github.com/PanJiaChen/vue-admin-template/issues/237
    // TODO: refactor with render function
    this.onlyOneChild = null
    return {}
  },
  methods: {
    hasOneShowingChild(children = [], parent) {
      const showingChildren = children.filter(item => {
        if (item.hidden) {
          return false
        } else {
          // Temp set(will be used if only has one showing child)
          this.onlyOneChild = item
          return true
        }
      })

      // When there is only one child router, the child router is displayed by default
      if (showingChildren.length === 1) {
        return true
      }

      // Show parent if there are no child router to display
      if (showingChildren.length === 0) {
        this.onlyOneChild = { ...parent, path: '', noShowingChildren: true }
        return true
      }

      return false
    },
    resolvePath(routePath) {
      if (isExternal(routePath)) {
        return routePath
      }
      if (isExternal(this.basePath)) {
        return this.basePath
      }
      return path.resolve(this.basePath, routePath)
    },
    checkNewMessage(key) {
      if (key === 'CsSystem' && this.newMessage) {
        if (this.$route.path !== "/cs-system/dashboard") {
          return true
        } else {
          this.$store.commit("cs/SET_NEW_MESSAGE", false)
        }
      }
      return false;
    },
    handleCheckMessage() {
      this.$store.commit("cs/SET_NEW_MESSAGE", false)
    }
  }
}
</script>

<style lang="scss" scoped>
.red-point::after {
  content: " ";
  display: inline-block;
  border: 4px solid #ff0c0c;/*设置红色*/
  border-radius:50%;
  z-index: 10;
  height: 2px;
  margin-left: auto;
  margin-right: 3px;
}
</style>
