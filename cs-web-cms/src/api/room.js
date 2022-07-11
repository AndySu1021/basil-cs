import httpRequest from "@/utils/http";

export const apiGetStaffRoomList = async () => {
	return httpRequest({
		url: '/staff-rooms',
		method: 'GET',
	})
}

export const apiGetQueuingRoomList = async () => {
	return httpRequest({
		url: '/queuing-rooms',
		method: 'GET',
	})
}

export const apiGetRoomList = async (query) => {
	return httpRequest({
		url: '/rooms?' + query,
		method: 'GET',
	})
}

export const apiAcceptRoom = async (id) => {
	return httpRequest({
		url: '/room/accept/' + id,
		method: 'PATCH',
	})
}

export const apiCloseRoom = async (id, data) => {
	return httpRequest({
		url: '/room/close/' + id,
		method: 'PATCH',
		data
	})
}

export const apiTransferRoom = async (id, data) => {
	return httpRequest({
		url: '/room/transfer/' + id,
		method: 'PATCH',
		data
	})
}

export const apiGetRoomInfo = async (id) => {
	return httpRequest({
		url: '/room/' + id + '/info',
		method: 'GET',
	})
}
