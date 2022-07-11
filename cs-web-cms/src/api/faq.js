import httpRequest from "@/utils/http";

export const apiGetFAQList = async (query) => {
	return httpRequest({
		url: '/faqs?' + query,
		method: 'GET',
	})
}

export const apiCreateFAQ = async (data) => {
	return httpRequest({
		url: '/faq',
		method: 'POST',
		data
	})
}

export const apiGetFAQDetail = async (id) => {
	return httpRequest({
		url: '/faq/' + id,
		method: 'GET',
	})
}

export const apiUpdateFAQ = async (id, data) => {
	return httpRequest({
		url: '/faq/' + id,
		method: 'PUT',
		data
	})
}

export const apiDeleteFAQ = async (id) => {
	return httpRequest({
		url: '/faq/' + id,
		method: 'DELETE',
	})
}
