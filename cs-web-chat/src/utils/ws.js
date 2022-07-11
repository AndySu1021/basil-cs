import {getSID} from "@/utils/storage";
import { ElMessage } from 'element-plus'
import store from '@/store';

const wsUrl = process.env.VUE_APP_API_URL;
let socket = null;
let limitConnect = 3;
let timeConnect = 0;

export const connectSocket = () => {
	const sid = getSID()
	socket = new WebSocket(wsUrl+"?type=member&sid="+sid);
	socket.onopen = function() {
		console.log("websocket connected!!");
		limitConnect = 3
		timeConnect = 0
	};
	socket.onmessage = function(msg) {
		store.commit("ws/APPEND_MESSAGE", JSON.parse(msg.data ?? {}));
	};
	socket.onerror = function(event) {
		console.log(event);
	};
	socket.onclose = function (event) {
		console.log("websocket closed!!");
		console.log(event.reason)
		reconnect();
	}

	function reconnect() {
		if(limitConnect>0){
			limitConnect--;
			timeConnect++;
			console.log("第"+timeConnect+"次重連");
			setTimeout(function(){
				connectSocket();
			},3000);
		}else{
			console.log("TCP連線超時");
			ElMessage({
				message: '連線失敗',
				type: 'error',
			})
		}
	}
};

export const sendSocketMessage = msg => {
	if (1 === socket.readyState) socket.send(JSON.stringify(msg));
};