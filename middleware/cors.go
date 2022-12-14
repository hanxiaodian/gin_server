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
			c.Header("Access-Control-Allow-Methods", "GET,HEAD,PUT,POST,DELETE,PATCH")
			// 后续有变动可以增加，用到比较少不做配置了
			c.Header("Access-Control-Expose-Headers", "Authorization")
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
