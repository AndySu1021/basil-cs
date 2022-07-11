import httpRequest from "@/utils/http";

export const apiGetTagList = async (query) => {
	return httpRequest({
		url: '/tags?' + query,
		method: 'GET',
	})
}

export const apiCreateTag = async (data) => {
	return httpRequest({
		url: '/tag',
		method: 'POST',
		data
	})
}

export const apiGetTagDetail = async (id) => {
	return httpRequest({
		url: '/tag/' + id,
		method: 'GET',
	})
}

export const apiUpdateTag = async (id, data) => {
	return httpRequest({
		url: '/tag/' + id,
		method: 'PUT',
		data
	})
}

export const apiDeleteTag = async (id) => {
	return httpRequest({
		url: '/tag/' + id,
		method: 'DELETE',
	})
}

export const apiGetAvailableTags = async () => {
	return httpRequest({
		url: '/available-tags',
		method: 'GET',
	})
}
