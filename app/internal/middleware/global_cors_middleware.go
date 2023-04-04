package middleware

import (
	"net/http"
)

func CorsHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// 必须，设置服务器支持的所有跨域请求的方法
		w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
		//返回数据格式是json
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// Passthrough to next handler if need
		next(w, r)
	}
}
