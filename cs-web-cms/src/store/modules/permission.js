import { constantRoutes, asyncRoutes } from '@/router'
import {getPermissionKeys, permissionKeys} from "@/utils/authorities";
import {getPermission} from "@/utils/storage";

const state = {
  routes: [],
  permissionKeys: [],
}

const mutations = {
  RESET_STATE: (state) => {
    state.routes = []
    state.permissionKeys = []
  },
  SET_ROUTES: (state, routes) => {
    state.routes = constantRoutes.concat(routes)
  },
  SET_ROUTE_KEYS: (state, permissionKeys) => {
    state.permissionKeys = permissionKeys
  },
}

const resolveChildren = function(result, item, keys) {
  let temp = {
    path: item.path,
    component: item.component,
    redirect: item.redirect,
    meta: item.meta,
  }
  for (let i = 0; i < keys.length; i++) {
    if (keys[i].startsWith(item.meta.key) || keys[i].startsWith("-" + item.meta.key)) {
      result.push(temp)
      break
    }
  }
  if (!item.children || item.children.length === 0) {
    return
  }
  temp.children = []
  for (let i = 0; i < item.children.length; i++) {
    resolveChildren(temp.children, item.children[i], keys)
  }
}

const actions = {
  getRouteKeys({ commit }) {
    return new Promise(async (resolve, reject) => {
      let perms = getPermission().split(",")
      if (perms.includes("*")) {
        commit('SET_ROUTE_KEYS', getPermissionKeys())
      } else {
        commit('SET_ROUTE_KEYS', perms)
      }
      resolve()
    })
  },
  generateRoutes({ commit, state }) {
    return new Promise((resolve) => {
      const accessedRoutes = asyncRoutes.reduce((acc, cur) => {
        // if (!state.permissionKeys.some((key) => key === cur.meta.key)) return acc
        if (!cur.children || cur.children.length === 0) return acc
        let result = []
        resolveChildren(result, cur, state.permissionKeys)
        if (result.length === 0) {
          return acc
        }
        if (result[0].children.length > 0) {
          acc.push(result[0])
        }
        return acc
      }, [])
      commit('SET_ROUTES', accessedRoutes)
      resolve(accessedRoutes)
    })
  }
}

const getters = {
  permissionRoutes: state => state.routes,
  permissionKeys: state => state.permissionKeys,
  hasRouteKeys: state => state.permissionKeys.length > 0,
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}
