import httpRequest from '@/utils/http'

export const apiAuthLogin = async (data) => {
	return httpRequest({
		url: '/auth/login',
		method: 'POST',
		data
	})
}

export const apiGetStaffInfo = async () => {
	return httpRequest({
		url: '/auth/staff-info',
		method: 'GET',
	})
}

export const apiAuthLogout = async () => {
	return httpRequest({
		url: '/auth/logout',
		method: 'POST',
	})
}