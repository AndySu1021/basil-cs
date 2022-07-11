import Vue from 'vue'

Vue.prototype.$confirmDelete = () => {
  return Vue.prototype.$confirm('確定要刪除嗎？', '系統提醒', {
    confirmButtonText: '確定',
    cancelButtonText: '取消',
    type: 'warning'
  })
}

Vue.prototype.$confirmClose = () => {
  return Vue.prototype.$confirm('確定要關閉嗎？', '系統提醒', {
    confirmButtonText: '確定',
    cancelButtonText: '取消',
    type: 'warning'
  })
}

Vue.prototype.$confirmLogout = () => {
  return Vue.prototype.$confirm('確定要登出嗎？', '系統提醒', {
    confirmButtonText: '確認',
    cancelButtonText: '取消',
    type: 'warning'
  })
}

Vue.prototype.$confirmSendScore = () => {
  return Vue.prototype.$confirm('確定要發送評分嗎？', '系統提醒', {
    confirmButtonText: '確定',
    cancelButtonText: '取消',
    type: 'warning'
  })
}

Vue.prototype.$showSuccessMessage = (message, callback) => {
  return new Promise(resolve => {
    if (callback) {
      Vue.prototype.$message({
        message,
        type: 'success',
        duration: 2000,
        onClose: () => {
          callback()
          resolve()
        }
      })
    } else {
      Vue.prototype.$message({
        message,
        type: 'success',
        duration: 2000
      })
      resolve()
    }
  })
}
