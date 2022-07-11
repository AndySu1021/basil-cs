import router from './router'
import store from './store'
import { Message } from 'element-ui'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import { getToken } from '@/utils/storage' // get token from cookie
import {connectSocket, socketClosed} from "@/utils/ws";

NProgress.configure({ showSpinner: false }) // NProgress Configuration

const whiteList = ['/login'] // 不需重新導向的白名單

const getPageTitle = (title) => {
  if (title) {
    return `${title} - 客服中心`
  }
  return '客服中心'
}

router.beforeEach(async (to, from, next) => {
  // 開啟 progress bar
  NProgress.start()

  // 設定每頁的頁籤title
  document.title = getPageTitle(to.meta.title)

  const hasToken = getToken()

  // 有token，但要前往login頁面 => 直接導向至首頁
  if (hasToken) {
    if (to.path === '/login') {
      next({ path: '/' })
      return
    }

    try {
      await store.dispatch("user/getStaffInfo")
    } catch (error) {
      console.log(error)
    }

    if(socketClosed()) {
      connectSocket();
    }
    const hasRouteKeys = store.getters["permission/hasRouteKeys"]

    // 已經初始化過動態路由則直接前往
    if (hasRouteKeys) {
      next()
      return
    }

    try {
      // 尚未初始化過動態路由則需初始化
      await store.dispatch('permission/getRouteKeys')
      // 根據權限篩選出可訪問的路由
      const accessRoutes = await store.dispatch('permission/generateRoutes')
      // 動態加上可訪問路由
      router.addRoutes(accessRoutes)

      // 最後再補上含有通配符的匹配路由 (必須放到最後)
      router.addRoutes([{ path: '*', redirect: '/404', hidden: true }])
      // 加上replace，讓它不會向 history 添加新紀錄，而是直接替換當前的 history 紀錄
      accessRoutes.length === 0 ? next() : next({ ...to, replace: true }) // 首次登入後會有 have undefined bug
    } catch (error) {
      await store.dispatch('user/resetToken')
      Message.error(error || 'Has Error')
      next({ path: '/login' })
      NProgress.done()
    }
  } else {
    if (whiteList.includes(to.path)) {
      next()
      return
    }
    store.dispatch('user/resetToken').then(() => {
      next({ path: '/login' })
    })
  }
})

router.afterEach(() => {
  // 關閉 progress bar
  NProgress.done()
})
