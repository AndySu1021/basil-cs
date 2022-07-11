import httpRequest from '@/utils/http'

export const apiUpdateRoomScore = (data) => {
	return httpRequest({
		url: '/room/score',
		method: 'PATCH',
		data
	})
}