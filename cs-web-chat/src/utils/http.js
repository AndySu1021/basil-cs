import axios from 'axios'
import { getSID } from '@/utils/storage'
import {ElMessage} from "element-plus";

// create an axios instance
const request = axios.create({
	baseURL: process.env.VUE_APP_BASE_URL, // url = base url + request url
	withCredentials: false, // send cookies when cross-domain requests
	timeout: 10000 // request timeout
})

// request interceptor
request.interceptors.request.use(
	(config) => {
		// do something before request is sent

		const token = getSID()

		if (token) {
			// 將每個request都帶上token
			config.headers['X-Token'] = token
		}
		return config
	},
	(error) => {
		// do something with request error
		console.log(error) // for debug
		return Promise.reject(error)
	}
)

// response interceptor
request.interceptors.response.use(
	// HTTP Status Code 為 200
	(response) => {
		const res = response.data
		// 若 res.code 不為 0 則直接跳錯誤訊息
		if (res.code !== 0) {
			ElMessage({
				message: res.message || 'Error',
				type: 'error',
				duration: 5 * 1000
			})
			return Promise.reject(new Error(res.message || 'Error'))
		} else {
			return res
		}
	},
	// HTTP Status Code 不為 200 時的error處理
	(error) => {
		ElMessage({
			message: error.response.data.message,
			type: 'error',
			duration: 5 * 1000
		})
		return Promise.reject(error)
	}
)
export default request
