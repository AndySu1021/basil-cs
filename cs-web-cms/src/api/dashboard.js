import httpRequest from "@/utils/http";

export const apiGetPanelData = async () => {
	return httpRequest({
		url: '/dashboard/panel-data',
		method: 'GET',
	})
}