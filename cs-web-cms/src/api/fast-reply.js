import httpRequest from "@/utils/http";

export const apiGetFastReplyList = async (query) => {
	return httpRequest({
		url: '/fast-replies?' + query,
		method: 'GET',
	})
}

export const apiCreateFastReply = async (data) => {
	return httpRequest({
		url: '/fast-reply',
		method: 'POST',
		data
	})
}

export const apiGetFastReplyDetail = async (id) => {
	return httpRequest({
		url: '/fast-reply/' + id,
		method: 'GET',
	})
}

export const apiUpdateFastReply = async (id, data) => {
	return httpRequest({
		url: '/fast-reply/' + id,
		method: 'PUT',
		data
	})
}

export const apiDeleteFastReply = async (id) => {
	return httpRequest({
		url: '/fast-reply/' + id,
		method: 'DELETE',
	})
}

export const apiCreateCategory = async (data) => {
	return httpRequest({
		url: '/fast-reply/category',
		method: 'POST',
		data
	})
}

export const apiGetCategoryList = async () => {
	return httpRequest({
		url: '/fast-reply/categories',
		method: 'GET',
	})
}

export const apiGetFastReplyGroup = async () => {
	return httpRequest({
		url: '/fast-reply/group',
		method: 'GET',
	})
}
