import httpRequest from "@/utils/http";

export const apiGetNoticeList = async (query) => {
	return httpRequest({
		url: '/notices?' + query,
		method: 'GET',
	})
}

export const apiCreateNotice = async (data) => {
	return httpRequest({
		url: '/notice',
		method: 'POST',
		data
	})
}

export const apiGetNoticeDetail = async (id) => {
	return httpRequest({
		url: '/notice/' + id,
		method: 'GET',
	})
}

export const apiUpdateNotice = async (id, data) => {
	return httpRequest({
		url: '/notice/' + id,
		method: 'PUT',
		data
	})
}

export const apiDeleteNotice = async (id) => {
	return httpRequest({
		url: '/notice/' + id,
		method: 'DELETE',
	})
}
