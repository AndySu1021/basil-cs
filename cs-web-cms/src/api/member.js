import httpRequest from "@/utils/http";

export const apiGetMemberList = async (query) => {
	return httpRequest({
		url: '/members?' + query,
		method: 'GET',
	})
}

export const apiUpdateMemberInfo = async (id, data) => {
	return httpRequest({
		url: '/member/' + id + '/info',
		method: 'PATCH',
		data
	})
}