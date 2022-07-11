import {getToken} from "@/utils/storage";
import store from '@/store'
const wsUrl = process.env.VUE_APP_WS_URL;
let socket = null;
let limitConnect = 3;
let timeConnect = 0;

export const connectSocket = () => {
	const token = getToken()
	socket = new WebSocket(wsUrl+"?type=staff&sid="+token);
	socket.onopen = function() {
		console.log("websocket connected!!");
		limitConnect = 3
		timeConnect = 0
		store.commit("app/SET_WEBSOCKET_STATUS", true)
	};
	socket.onmessage = function(msg) {
		store.commit("cs/APPEND_MESSAGE", JSON.parse(msg.data))
	};
	socket.onerror = function(err) {
		console.log("error", err);
	};
	socket.onclose = function (event) {
		console.log("websocket closed!!");
		console.log(event.reason)
		store.commit("app/SET_WEBSOCKET_STATUS", false)
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
		}
	}
};

export const sendSocketMessage = msg => {
	if (1 === socket.readyState) socket.send(JSON.stringify(msg));
};

export const socketClosed = () => {
	if(socket === null) {
		return true
	}
	return 3 === socket.readyState
};
