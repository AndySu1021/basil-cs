import httpRequest from '@/utils/http'

export const apiGetMessageList = (query) => {
	return httpRequest({
		url: '/room-messages/member?' + query,
		method: 'GET'
	})
}