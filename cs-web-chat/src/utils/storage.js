const sidKey = 'sid'
export const getSID = () => sessionStorage.getItem(sidKey)
export const setSID = (sid) => sessionStorage.setItem(sidKey, `${sid}`)
export const removeSID = () => sessionStorage.removeItem(sidKey)