##### 一.请求和响应
请求和响应的格式
(1)请求行或者相应行
(2)0个或多个首部(headers)
(3)一个空行
(4)一个可选的报文主体
1. Request结构
+ URL字段
+ Header字段
+ Body字段
+ Form字段、PostForm字段、MultipartForm字段
  + 首部
    ```
    func hanleHeaders(w http.ResponseWriter, r *http.Request) {
      //h := r.Header //取得所有首部
	    //h := r.Header["Accept-Encoding"] //获取某个首部 output:[gzip, deflate, br] 数组
	    h := r.Header.Get("Accept-Encoding") //获取某个首部 是,分割的 gzip, deflate, br
	    fmt.Fprintln(w, h)
    }
    ```
  + 请求主体
    ```
    func handleBody(w http.ResponseWriter, r *http.Request) {
      contentLen := r.ContentLength
	    body := make([]byte, contentLen)
	    r.Body.Read(body)
	    fmt.Println(contentLen)
	    fmt.Fprintln(w, string(body))
    }
    ```
  + Go和Html表单
    html表单的content-type确定了post请求在发送键值对时采用何种格式:
    ```
    <form action="/process" method="POST" enctype="application/x-www-form-urlencoded">
      <input type="text" name="user_account"/>
      <input type="password" name="pwd"/>
      <input type="submit"/>
    </form>
    ```
    Content-Type至少要支持 x-www-form-urlencoded multipart/form-data,
    H5还支持text/plain,默认是x-www-form-urlencoded，Get请求也是可以传参的，
    但是Get没有body所以只能使用x-www-form-urlencoded，一般数据比较少的时候使用
    x-www-form-urlencoded，大量数据使用multipart/form-data，如文件上传等        
  + Form字段
    提交的数据存放在Form、PostForm、MultipartForm中，使用Request结构获取数据的
    一般步骤是:
    (1)ParseForm()或者ParseMultipartForm()解析
    (2)访问Form、PostForm、MultipartForm字段
    ```
    .html
    <form action="http://127.0.0.1:8080/process?user=hui&jsj=89" method="POST"  enctype="application/x-www-form-urlencoded">
      <input type="text" name="user" value="wang"/>
      <input type="password" name="pwd"/>
      <input type="submit"/>
    </form>
    .go文件
    r.ParseForm()
	  fmt.Fprintln(w, r.Form) 
    //output:map[jsj:[89] pwd:[sss] user:[wang hui]] wang来自form表单 hui来自url拼接的参数
    ```
  + 

##### 二.请求和响应
##### 三.请求和响应
##### 四.请求和响应
##### 五.请求和响应