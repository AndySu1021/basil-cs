import store from '@/store'

export default {
  inserted(el, binding) {
    const { value } = binding
    const permissionKeys = store.state && store.state.permission.permissionKeys
    if (!!value) {
      const hasPermission = permissionKeys.includes(value)
      if (!hasPermission) {
        el.parentNode && el.parentNode.removeChild(el)
      }
    } else {
      throw new Error('v-permission值不得為空!')
    }
  }
}
