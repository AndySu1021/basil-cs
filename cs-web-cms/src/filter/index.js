import Vue from 'vue'

import {format} from "date-fns";
import {decrypt} from "@/utils/crypto";

/**
 * 轉換日期格式
 */
Vue.filter('dateFormat', (date, p_format = 'yyyy-MM-dd') => {
	if (!date) {
		return date
	}
	return format(new Date(date), p_format)
})

/**
 * 轉換時間與日期格式
 */
Vue.filter('dateTimeFormat', (date, p_format = 'yyyy-MM-dd HH:mm:ss') => {
	if (!date) {
		return date
	}
	return format(new Date(date), p_format)
})

/**
 * 將數字加上千分位符號
 * @param {Number} 需要Format的值
 * @returns {String}
 */
Vue.filter('thousandsSeparator', number => {
	const result = parseFloat(number).toLocaleString()
	return isNaN(parseFloat(result)) ? number : result
})

/**
 * 將數字取到小數2位
 * @param {Number} 需要Format的值
 * @returns {Number}
 */
Vue.filter('decimalMark', number => {
	const result = Math.round(number * 100) / 100
	return isNaN(result) ? number : result
})

/**
 * 將金額數字加上千分位符號並取到小數2位
 * @param {Number} 需要Format的值
 * @returns {String}
 */
Vue.filter('moneyFormat', number => {
	const result = thousandsSeparator(decimalMark(number))
	return isNaN(parseFloat(result)) ? number : result
})

Vue.filter('decrypt', content => {
	return decrypt(content)
})
