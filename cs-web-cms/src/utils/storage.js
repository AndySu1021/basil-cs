// App
const sidebarStatusKey = 'sidebarStatus'
export const getSidebarStatus = () => localStorage.getItem(sidebarStatusKey)
export const setSidebarStatus = (sidebarStatus) => localStorage.setItem(sidebarStatusKey, sidebarStatus)

const languageKey = 'language'
export const getLanguage = () => localStorage.getItem(languageKey)
export const setLanguage = (language) => localStorage.setItem(languageKey, language)

// User
const tokenKey = 'adminToken'
export const getToken = () => localStorage.getItem(tokenKey)
export const setToken = (token) => localStorage.setItem(tokenKey, `${token}`)
export const removeToken = () => localStorage.removeItem(tokenKey)

const permKey = 'perm'
export const getPermission = () => localStorage.getItem(permKey)
export const setPermission = (perm) => localStorage.setItem(permKey, perm)
export const removePermission = () => localStorage.removeItem(permKey)