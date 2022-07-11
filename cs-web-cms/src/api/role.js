import httpRequest from "@/utils/http";

export const apiGetRoleList = async (query) => {
	return httpRequest({
		url: '/roles?' + query,
		method: 'GET',
	})
}

export const apiGetAllRoles = async () => {
	return httpRequest({
		url: '/roles/all',
		method: 'GET',
	})
}

export const apiCreateRole = async (data) => {
	return httpRequest({
		url: '/role',
		method: 'POST',
		data
	})
}

export const apiGetRoleDetail = async (id) => {
	return httpRequest({
		url: '/role/' + id,
		method: 'GET',
	})
}

export const apiUpdateRole = async (id, data) => {
	return httpRequest({
		url: '/role/' + id,
		method: 'PUT',
		data
	})
}

export const apiDeleteRole = async (id) => {
	return httpRequest({
		url: '/role/' + id,
		method: 'DELETE',
	})
}
