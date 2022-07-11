import { deepCopy } from '@/utils'
import {apiGetQueuingRoomList, apiGetStaffRoomList} from "@/api/room";
import {apiGetRoomMessageList} from "@/api/message";
import {encrypt} from "@/utils/crypto";

const initialState = {
	newMessage: false,
	activeRoomId: 0,
	roomList: [],
	activeTab: "Serving",
	messageList: [],
	roomInfo: {
		member_id: 0,
		ip: "",
		source: "",
		site_name: "",
		created_at: "",
		real_name: "",
		email: "",
		mobile: "",
		note: "",
		online_status: 2,
	}
}

const state = deepCopy(initialState)

const mutations = {
	RESET: (state) => {
		state.activeRoomId = 0
		state.roomList = []
		state.activeTab = "Serving"
		state.messageList = []
	},
	SET_ACTIVE_ROOM: (state, activeRoomId) => {
		state.activeRoomId = activeRoomId
	},
	SET_ROOMS: (state, rooms) => {
		state.roomList = []
		for(let i=0; i < rooms.length; i++) {
			state.roomList.push({...rooms[i], isRead: true})
		}
	},
	SET_ACTIVE_TAB: (state, tab) => {
		state.activeTab = tab
	},
	APPEND_MESSAGE(state, message) {
		state.newMessage = true
		// handle member join
		if(message.op_type === 5) {
			const result = state.roomList.filter(item => item.id === message.room_id);
			if(result.length === 0) {
				state.roomList.push({
					id: message.room_id,
					member_id: message.extra.member_id,
					member_name: message.content,
					status: "Serving",
					content: "",
					isRead: false,
				})
			}
		} else {
			let i
			for (i = 0; i < state.roomList.length; i++) {
				if (state.roomList[i].id === message.room_id) {
					if(message.sender_type === 1) {
						if (message.content_type === 2) {
							state.roomList[i].content = encrypt("會員上傳一張圖片")
						} else {
							state.roomList[i].content = message.content
						}
					}
					break
				}
			}
			if(message.room_id === state.activeRoomId) {
				state.roomList[i].isRead = true
				state.messageList.push(message)
			} else {
				state.roomList[i].isRead = false
			}
		}
	},
	SET_MESSAGE_LIST(state, messages) {
		state.messageList = messages
	},
	SET_NEW_MESSAGE(state, value) {
		state.newMessage = value
	},
	SET_MESSAGE_READ(state, roomId) {
		for (let i = 0; i < state.roomList.length; i++) {
			if (state.roomList[i].id === roomId) {
				state.roomList[i].isRead = true
				break
			}
		}
	},
	SET_ROOM_INFO(state, info) {
		state.roomInfo = info
	},
	RESET_ROOM_INFO(state) {
		state.roomInfo = {
			member_id: 0,
			ip: "",
			source: "",
			site_name: "",
			created_at: "",
			user_agent: "",
			real_name: "",
			email: "",
			mobile: "",
			note: "",
			online_status: 2,
		}
	},
}

const actions = {
	getStaffRoomList({ commit }) {
		return new Promise((resolve, reject) => {
			apiGetStaffRoomList()
				.then((response) => {
					commit('SET_ROOMS', response.data)
					resolve()
				})
				.catch((error) => {
					reject(error)
				})
		})
	},
	getQueuingRoomList({ commit }) {
		return new Promise((resolve, reject) => {
			apiGetQueuingRoomList()
				.then((response) => {
					commit('SET_ROOMS', response.data)
					resolve()
				})
				.catch((error) => {
					reject(error)
				})
		})
	},
	getRoomMessageList({ commit }, payload) {
		return new Promise((resolve, reject) => {
			apiGetRoomMessageList(payload)
				.then((response) => {
					commit('SET_MESSAGE_LIST', response.data)
					resolve()
				})
				.catch((error) => {
					reject(error)
				})
		})
	},
}

const getters = {
	newMessage: state => state.newMessage,
	activeRoomId: state => state.activeRoomId === 0 ? "" : state.activeRoomId,
	activeTab: state => state.activeTab,
	roomList: state => state.roomList,
	messageList: state => state.messageList,
	roomInfo: state => state.roomInfo,
}

export default {
	namespaced: true,
	state,
	mutations,
	actions,
	getters
}
