package middleware

import (
	"gin_server/routers"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// CorsByRules 按照配置处理跨域请求
func CorsByRules() gin.HandlerFunc {
	// 放行全部
	return func(c *gin.Context) {
		origin := c.GetHeader("origin")

		if origin == "" {
			c.Next()
			return
		}
		path := c.Request.URL.Path

		isAllow := checkCors(origin, path)

		// 通过检查，添加请求头
		if isAllow {
			c.Header("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "GET,HEAD,PUT,POST,DELETE,PATCH")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization,Content-Length,X-CSRF-Token,Token,session,Content-Type,x-kms-project,x-kms-sign,x-kms-stamp,x-kms-user,x-kms-username,x-kms-appid,sentry-trace")
			// 允许浏览器（客户端）可以解析的头部（重要）
			c.Header("Access-Control-Expose-Headers", "Authorization,Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "false")
			// TODO: 设置 cookie 时间
		}

		// 处理请求
		c.Next()
	}
}

func checkCors(currentOrigin string, path string) bool {
	paths := strings.Split(path, "/")
	whitelists := routers.GetAllowOrigins(paths[1], strings.ToUpper(paths[2]))

	// 判断 origin 的正则
	regexps := getOriginRegs(whitelists)

	isAllow := false
	for i := 0; i < len(regexps); i++ {
		result := regexps[i].MatchString(currentOrigin)
		if result {
			isAllow = true
		}
	}

	return isAllow
}

func getOriginRegs(origins []string) []*regexp.Regexp {
	var length = len(origins)
	var originRegs = make([]*regexp.Regexp, length)

	for i := 0; i < length; i++ {
		originRegs[i] = generateRegExp(origins[i])
	}
	return originRegs
}

func generateRegExp(domain string) *regexp.Regexp {
	// var reg1 = regexp.MustCompile(`/\./g`)
	var reg1 = regexp.MustCompile(`\.`)
	var reg2 = regexp.MustCompile(`^\\.`)
	var str = reg1.ReplaceAllString(domain, "\\.")
	str = reg2.ReplaceAllString(str, "(?:.*\\.)?") + "$"
	return regexp.MustCompile("^(?i:https?://|//)?" + str)
}
