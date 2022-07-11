import httpRequest from '@/utils/http'

export const apiGetFAQList = () => {
	return httpRequest({
		url: '/available-faqs',
		method: 'GET',
	})
}