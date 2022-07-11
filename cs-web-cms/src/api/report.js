import httpRequest from "@/utils/http";

export const apiGetDailyTagReport = async (query) => {
	return httpRequest({
		url: '/report/daily/tag?' + query ,
		method: 'GET',
	})
}

export const apiGetDailyStaffReport = async (query) => {
	return httpRequest({
		url: '/report/daily/staff?' + query ,
		method: 'GET',
	})
}
