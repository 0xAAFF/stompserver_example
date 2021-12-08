package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/0xAAFF/stompserver/tools"
)

/*
 * 本节提供Vue生成的Web项目,为前端浏览器访问本地服务提供的html页面
 * 该模块是一个微型的web服务
 * ROOT_PATH 是一个相对的路径,指向Web资源的目录
 */

const (
	Status_Success     = 200
	Status_Redirect    = 302
	Status_NotFound    = 404
	Content_Type_HTML  = "text/html;charset=utf-8"
	Content_Type_CSS   = "text/css"
	Content_Type_JS    = "application/x-javascript"
	Content_Type_IMAGE = "image/png"
	Content_Type_FONT  = "application/x-font-ttf"
	Content_Type_ICO   = "image/vnd.microsoft.icon"

	Content_Type_JSON     = "text/json"
	Content_Type_APP_JSON = "application/json;charset=utf-8"
)

const (
	ROOT_PATH        = "data"           // 本地Web资源的主目录
	CSS_PATH         = "/css/"          // ROOT_PATH+css存放路径
	STATIC_CSS_PATH  = "/static/css/"   // ROOT_PATH+css存放路径
	JS_PATH          = "/js/"           // ROOT_PATH+JS存放路径
	STATIC_JS_PATH   = "/static/js/"    // ROOT_PATH+JS存放路径
	IMG_PATH         = "/img/"          // ROOT_PATH+IMG Vue项目的图片路径
	STATIC_IMG_PATH  = "/static/img/"   // ROOT_PATH+IMG Vue项目的图片路径
	IMAGE_PATH       = "/images/"       // ROOT_PATH+IMAGE 从controller下载的图片 用来动态加载图片的路径
	FONT_PATH        = "/fonts/"        // ROOT_PATH+FONT 字体库
	STATIC_FONT_PATH = "/static/fonts/" // ROOT_PATH+FONT 字体库
	MAIN_PATH        = "/index.html"    // ROOT_PATH+ 主页文件路径
)

// IsAllowAccess  允许访问的资源过滤器(Vue2/3/4 build 输出目录兼容)
//  参数:
//  path	string	指定资源的相对路径
//  return	bool	是否可以访问指定资源
/*
vue cli3+ 打包静态文件目录的配置
vue cli3+ 默认打包生成的文件和cli2.0同样是生成在dist目录中，但静态文件如css、js、img并未像cli2.0一样放在static文件夹下，而是和index.html处在同一目录中，显得十分散乱

要使打包后的文件与cli2.0的相同，需要修改vue.config.js的几项配置变量：

module.exports = {
  publicPath: './',
  outputDir: "dist", // 输出文件目录
  lintOnSave: false, // eslint 是否在保存时检查
  assetsDir: 'static', // 配置js、css静态资源二级目录的位置
}
这样便可以改为和cli2.0一致了
*/
func IsAllowAccess(path string) bool {
	return ((strings.HasPrefix(path, CSS_PATH) || (strings.HasPrefix(path, STATIC_CSS_PATH))) && (strings.HasSuffix(path, ".css") || strings.HasSuffix(path, ".css.map"))) ||
		((strings.HasPrefix(path, JS_PATH) || (strings.HasPrefix(path, STATIC_JS_PATH))) && (strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".js.map"))) ||
		strings.HasPrefix(path, STATIC_IMG_PATH) || (strings.HasPrefix(path, IMG_PATH) || strings.HasPrefix(path, IMAGE_PATH)) || strings.HasPrefix(path, STATIC_FONT_PATH) || strings.HasPrefix(path, FONT_PATH)

}

// GetContentType  获取文件的content-type属性
//  参数:
//  path	string	文件的相对路径
//  return	string	content-type
func GetContentType(path string) string {
	if strings.HasSuffix(path, ".css") || strings.HasSuffix(path, ".css.map") {
		return Content_Type_CSS
	} else if strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".js.map") {
		return Content_Type_JS
	} else if strings.HasPrefix(path, IMG_PATH) || strings.HasPrefix(path, IMAGE_PATH) {
		return Content_Type_IMAGE
	} else if strings.HasPrefix(path, FONT_PATH) {
		return Content_Type_FONT
	} else {
		return ""
	}
}


// HttpVueServer  本地服务对外提供的Web解析服务
//  参数:
//  w		http.ResponseWriter
//  req		*http.Request
//
//  return	返回空
func HttpVueServer(w http.ResponseWriter, request *http.Request) {
	// ./data is Vue build files
	filepath := tools.GetCurrentPathNoError() + ROOT_PATH

	if request.URL.Path == "/" {
		filepath += MAIN_PATH //"/index.html"
		w.Header().Set("Content-Type", Content_Type_HTML)
	} else if request.URL.Path == "/favicon.ico" {
		filepath += "/favicon.ico"
		w.Header().Set("Content-Type", Content_Type_ICO)
	} else if IsAllowAccess(request.URL.Path) {
		filepath += request.URL.Path

		if len(request.RequestURI) > 50 {
			fmt.Println("Web路径访问 路径较长.长度:", len(request.RequestURI), "\t:", request.RequestURI)
		}
		// 路径中的特殊字符不能要(只需要0-9,a-z,-,.,/)
		// 这么做的原因是防止xss注入和非正常路径访问
		// 当然 这里并不一定完全正确,需要各种思路多多测试
		matchOK, _ := regexp.MatchString(`^([0-9a-zA-Z_\-\_\./])+$`, request.RequestURI)
		if !matchOK {
			fmt.Println("Web路径访问 非法字符.<code>\t", request.RequestURI, "\t</code>")
			//filepath = "" // 路径访问非法 则拒绝访问
		}
		w.Header().Set("Content-Type", GetContentType(filepath))
	} else {
		onReflex(w, request)
		w.WriteHeader(Status_NotFound)
		return
	}
	existFile := tools.ExistsFile(filepath)
	if existFile {
		ReadAllBytes(filepath, w)
	} else {
		w.WriteHeader(Status_NotFound)
	}
}

// ReadAllBytes  读取本地资源并使用http.ResponseWriter将资源写入http连接
//  参数:
//  filePath		文件路径
//  w				http.ResponseWriter
//  return	返回无
func ReadAllBytes(filePath string, w http.ResponseWriter) {
	fp, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	for {
		buff := make([]byte, 100)
		lens, err := fp.Read(buff)
		if err == io.EOF || lens < 0 {
			break
		}
		w.Write(buff[:lens])
	}
}

func onReflex(responseW http.ResponseWriter, request *http.Request) {
	/*
	* 如下是示例,用于测试
	 */
	switch request.URL.Path {
	case "/test.do":
		{
			json := "{x:\"Back to Web\"}"
			responseW.WriteHeader(Status_Success)
			responseW.Header().Set("Content-Type", Content_Type_APP_JSON)
			responseW.Write([]byte(json))
			break
		}
	}
}
