import httpRequest from "@/utils/http";

export const apiGetSiteList = async () => {
	return httpRequest({
		url: '/sites',
		method: 'GET',
	})
}

export const apiCreateSite = async (data) => {
	return httpRequest({
		url: '/site',
		method: 'POST',
		data
	})
}

export const apiGetSiteDetail = async (id) => {
	return httpRequest({
		url: '/site/' + id,
		method: 'GET',
	})
}

export const apiUpdateSite = async (id, data) => {
	return httpRequest({
		url: '/site/' + id,
		method: 'PUT',
		data
	})
}

export const apiDeleteSite = async (id) => {
	return httpRequest({
		url: '/site/' + id,
		method: 'DELETE',
	})
}
