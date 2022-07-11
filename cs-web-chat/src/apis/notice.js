import httpRequest from '@/utils/http'

export const apiGetLatestNotice = () => {
	return httpRequest({
		url: '/notice/latest',
		method: 'GET'
	})
}