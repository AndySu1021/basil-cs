import {apiAuthLogin, apiAuthLogout, apiGetStaffInfo} from '@/api/auth'
import {setToken, removeToken, setPermission, removePermission} from '@/utils/storage'
import { resetRouter } from '@/router'
import { deepCopy } from '@/utils'

const initialState = {
  id: 0,
  name: "",
  username: "",
  servingStatus: "",
  avatar: "",
}

const state = deepCopy(initialState)

const mutations = {
  RESET_STATE: (state) => {
    state = deepCopy(initialState)
  },
  SET_INFO: (state, staffInfo) => {
    state.id = staffInfo.id
    state.name = staffInfo.name
    state.username = staffInfo.username
    state.servingStatus = staffInfo.serving_status
    state.avatar = staffInfo.avatar
  },
}

const actions = {
  getStaffInfo({ commit }) {
    return new Promise((resolve, reject) => {
      apiGetStaffInfo()
        .then((response) => {
          commit('SET_INFO', response.data)
          resolve()
        })
        .catch((error) => {
          reject(error)
        })
    })
  },
  login({ commit }, userInfo) {
    return new Promise((resolve, reject) => {
      apiAuthLogin(userInfo)
        .then((response) => {
          commit('RESET_STATE')
          setToken(response.data.token)
          setPermission(response.data.permissions)
          resolve()
        })
        .catch((error) => {
          console.log(error)
          reject(error)
        })
    })
  },
  logout({ commit }) {
    return new Promise((resolve, reject) => {
      apiAuthLogout()
        .then((response) => {
          localStorage.clear()
          removeToken()
          removePermission()
          resetRouter()
          commit('RESET_STATE')
          commit('permission/RESET_STATE', null, { root: true })
          location.reload()
          resolve()
        })
        .catch((error) => {
          reject(error)
        })
    })
  },
  resetToken({ commit }) {
    localStorage.clear()
    removeToken()
    removePermission()
    resetRouter()
    commit('RESET_STATE')
    commit('permission/RESET_STATE', null, { root: true })
  },
}

const getters = {
  id: state => state.id,
  name: state => state.name,
  username: state => state.username,
  servingStatus: state => state.servingStatus === 1 ? 'Closed' : 'Serving',
  avatar: state => state.avatar,
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}
