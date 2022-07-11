import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes 靜態路由
 * 不需要權限控管的路由，所有人都可訪問
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/error-page/404'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/index'),
        meta: { title: '首頁', icon: 'dashboard' }
      }
    ]
  },
]

/**
 * asyncRoutes 動態路由
 * 用來與後端接口回傳的路由key做比對，從而根據權限產生對應的路由
 * 路由的key由前端這邊定義，新增的路由記得補上key
 * 底下的按鈕操作，非實際頁面設定 hidden 為 true
 */
export const asyncRoutes = [
  {
    path: '/cs-system',
    component: Layout,
    redirect: '/cs-system/dashboard',
    meta: { title: '客服系統', icon: 'headset', key: 'CsSystem' },
    children: [
      {
        path: '/cs-system/dashboard',
        name: 'CsDashboard',
        component: () => import('@/views/cs-system/CsSystem'),
        meta: { title: '客服系統', icon: 'headset', key: 'CsSystem' }
      },
    ]
  },
  {
    path: '/cs-setting',
    component: Layout,
    redirect: '/cs-setting/tag-list',
    meta: { title: '客服設置', icon: 'setting', key: 'CsSetting' },
    children: [
      {
        path: '/cs-setting/tag-list',
        name: 'TagList',
        component: () => import('@/views/cs-setting/tag/TagList'),
        meta: { title: '標籤管理', icon: 'tag', key: 'ListTag' },
      },
      {
        path: '/cs-setting/fast-reply-list',
        name: 'FastReplyList',
        component: () => import('@/views/cs-setting/fast-reply/FastReplyList'),
        meta: { title: '快捷回覆', icon: 'send', key: 'ListFastReply' }
      },
      {
        path: '/cs-setting/faq-list',
        name: 'FAQList',
        component: () => import('@/views/cs-setting/faq/FAQList'),
        meta: { title: '常見問題', icon: 'faq', key: 'ListFAQ' }
      },
      {
        path: '/cs-setting/notice-list',
        name: 'NoticeList',
        component: () => import('@/views/cs-setting/notice/NoticeList'),
        meta: { title: '通知管理', icon: 'notice', key: 'ListNotice' }
      },
      {
        path: '/cs-setting/config-setting',
        name: 'ConfigSetting',
        component: () => import('@/views/cs-setting/config/ConfigSetting'),
        meta: { title: '參數配置', icon: 'setting', key: 'GetCsConfig' }
      },
    ]
  },
  {
    path: '/history-record',
    component: Layout,
    redirect: '/history-record/room-list',
    meta: { title: '歷史紀錄', icon: 'history', key: 'HistoryRecord' },
    children: [
      {
        path: '/history-record/room-list',
        name: 'RoomList',
        component: () => import('@/views/history-record/RoomList'),
        meta: { title: '歷史諮詢房', icon: 'phone', key: 'ListRoom' }
      },
      {
        path: '/history-record/message-list',
        name: 'MessageList',
        component: () => import('@/views/history-record/MessageList'),
        meta: { title: '歷史訊息', icon: 'message', key: 'ListMessage' }
      },
      {
        path: '/history-record/member-list',
        name: 'MemberList',
        component: () => import('@/views/history-record/MemberList'),
        meta: { title: '會員列表', icon: 'user', key: 'ListMember' }
      },
    ]
  },
  {
    path: '/report-management',
    component: Layout,
    redirect: '/report-management/daily-tag-report',
    meta: { title: '報表管理', icon: 'folder', key: 'Report' },
    children: [
      {
        path: '/report-management/daily-tag-report',
        name: 'DailyTagReport',
        component: () => import('@/views/report/DailyTagReport'),
        meta: { title: '每日標籤報表', icon: 'file', key: 'DailyTagReport' }
      },
      {
        path: '/report-management/daily-staff-report',
        name: 'DailyStaffReport',
        component: () => import('@/views/report/DailyStaffReport'),
        meta: { title: '每日客服報表', icon: 'file', key: 'DailyStaffReport' }
      },
    ]
  },
  {
    path: '/cms-management',
    component: Layout,
    redirect: '/cms-management/remind-list',
    meta: { title: '後台管理', icon: 'tool', key: 'CMS' },
    children: [
      {
        path: '/cms-management/remind-list',
        name: 'RemindList',
        component: () => import('@/views/cms/remind/RemindList'),
        meta: { title: '提醒管理', icon: 'remind', key: 'ListRemind' }
      },
      {
        path: '/cms-management/site-list',
        name: 'SiteList',
        component: () => import('@/views/cms/site/SiteList'),
        meta: { title: '站點管理', icon: 'site', key: 'ListSite' }
      },
    ]
  },
  {
    path: '/staff-management',
    component: Layout,
    redirect: '/staff/list',
    meta: { title: '職員管理', icon: 'bank', key: 'Staff' },
    children: [
      {
        path: '/staff-management/staff-list',
        name: 'StaffList',
        component: () => import('@/views/staff/StaffList'),
        meta: { title: '職員管理', icon: 'user', key: 'ListStaff' },
      },
      {
        path: '/staff-management/role-list',
        name: 'RoleList',
        component: () => import('@/views/staff/RoleList'),
        meta: { title: '角色管理', icon: 'identity', key: 'ListRole' },
      },
    ]
  },
]

const createRouter = () =>
  new Router({
    // mode: 'history', // require service support
    scrollBehavior: () => ({ y: 0 }),
    routes: constantRoutes
  })

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
