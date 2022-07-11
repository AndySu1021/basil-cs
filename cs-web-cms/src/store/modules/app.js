import {
  setLanguage,
  setSidebarStatus,
  getSidebarStatus
} from '@/utils/storage'

const state = {
  sidebar: {
    opened: getSidebarStatus() ? !!+getSidebarStatus() : true,
    withoutAnimation: false
  },
  device: 'desktop',
  websocketStatus: false,
}

const mutations = {
  TOGGLE_SIDEBAR: (state) => {
    state.sidebar.opened = !state.sidebar.opened
    state.sidebar.withoutAnimation = false
    if (state.sidebar.opened) {
      setSidebarStatus(1)
    } else {
      setSidebarStatus(0)
    }
  },
  CLOSE_SIDEBAR: (state, withoutAnimation) => {
    setSidebarStatus(0)
    state.sidebar.opened = false
    state.sidebar.withoutAnimation = withoutAnimation
  },
  OPEN_SIDEBAR: (state) => {
    setSidebarStatus(1)
    state.sidebar.opened = true
    state.sidebar.withoutAnimation = false
  },
  TOGGLE_DEVICE: (state, device) => {
    state.device = device
  },
  SET_LANGUAGE: (state, language) => {
    state.language = language
    setLanguage(language)
  },
  SET_WEBSOCKET_STATUS: (state, websocketStatus) => {
    state.websocketStatus = websocketStatus
  }
}

const actions = {
  toggleSideBar({ commit }) {
    commit('TOGGLE_SIDEBAR')
  },
  closeSideBar({ commit }, { withoutAnimation }) {
    commit('CLOSE_SIDEBAR', withoutAnimation)
  },
  openSideBar({ commit }) {
    commit('OPEN_SIDEBAR')
  },
  toggleDevice({ commit }, device) {
    commit('TOGGLE_DEVICE', device)
  },
  SetLanguage({ commit }, language) {
    commit('SET_LANGUAGE', language)
  }
}

const getters = {
  websocketStatus: state => state.websocketStatus,
  sidebar: state => state.sidebar,
  device: state => state.device,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}
