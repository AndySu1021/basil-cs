import httpRequest from '@/utils/http'

export const apiGetRemindList = async (query) => {
	return httpRequest({
		url: '/reminds?' + query,
		method: 'GET',
	})
}

export const apiCreateRemind = async (data) => {
	return httpRequest({
		url: '/remind',
		method: 'POST',
		data
	})
}

export const apiGetRemindDetail = async (id) => {
	return httpRequest({
		url: '/remind/' + id,
		method: 'GET',
	})
}

export const apiUpdateRemind = async (id, data) => {
	return httpRequest({
		url: '/remind/' + id,
		method: 'PUT',
		data
	})
}

export const apiDeleteRemind = async (id) => {
	return httpRequest({
		url: '/remind/' + id,
		method: 'DELETE',
	})
}

export const apiGetActiveReminds = async () => {
	return httpRequest({
		url: '/active-reminds',
		method: 'GET',
	})
}
