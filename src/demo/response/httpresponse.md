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
    //场景1
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
    //如果是调用r.PostForm,则只会得到表单提交的参数 url中拼接的参数将无法获取
    	r.ParseForm()
	    fmt.Fprintln(w, r.PostForm)
    //map[pwd:[dd] user:[wang]]
    //另外可以使用快捷的函数 formValue(key),PostFormValue(key) 来获取对应key的值，这样获取的只能是第一个值，如果有多个键值对的话 比如上面例子url中和表单中同时提交了参数，则会取表单里面的第一个参数，这种情况不需要手动调用parse...方法 方法内部会自动调用
	  fmt.Fprintln(w, r.FormValue("user")) //wang
	  fmt.Fprintln(w, r.PostFormValue("user")) //wang

    //parseMultipartForm 主要用在文件传输
      <form action="http://127.0.0.1:8080/handleFile" method="POST"  enctype="multipart/form-data">
        <label>文件名：</label>
        <input type="text" name="file_name"/><br>
        <label>选择文件：</label>
        <input type="file" name="apk"/><br>
        <input type="submit"/>
      </form>
      func handleFile(w http.ResponseWriter, r *http.Request) {
	      fmt.Println("request here")
	      r.ParseMultipartForm(2048)
	      fmt.Fprintln(w, r.MultipartForm)//&{map[file_name:[ss]] map[apk:[0xc000040050]]
      } //ss参数是文件名 apk参数是文件
    }
    //修改后的
    func handleFile(w http.ResponseWriter, r *http.Request) {
	    fmt.Println("request here")
	    r.ParseMultipartForm(2048)
	    //fmt.Fprintln(w, r.MultipartForm)//&{map[file_name:[ss]] map[apk:[0xc000040050]]}
	    fmt.Println("fileParams", r.MultipartForm.File["apk"])
	    fileHandler := r.MultipartForm.File["apk"][0]
	    file, err := fileHandler.Open()
	    if err == nil {
		    contents, err := ioutil.ReadAll(file)
		    if err == nil {
			  fmt.Fprintln(w, string(contents))
		    }
	    } 
    }
    func handleFile(w http.ResponseWriter, r *http.Request) {
	    fmt.Println("request here")
	    //可以使用简化的formFile("apk")来简化操作，这样只会返回第一个文件的句柄
	    file, _, err := r.FormFile("apk")
	      if err == nil {
		      contents, err := ioutil.ReadAll(file)
		      if err == nil {
			    fmt.Fprintln(w, string(contents))
		    }
	    }
    }
    ```
  + 

##### 二.处理请求体是json的请求
  + 上面介绍的Form系列方法，对于请求主体是json的请求是没有办法处理的，因此需要自己解析
    go对于json有两种方式可以处理：
    ```
    // 方式一：json.Unmarshal()
    b, err := ioutil.ReadAll(response.Body)
      if err != nil {
        log.Println("err=>", err)
    } 
    data = []byte(string(b))
    err = json.Unmarshal(data, &user)
    适用场景：如果要处理的JSON数据已经存在内存中，使用json.Unmarshal

    // 方式二：json.NewDecoder()
    err := json.NewDecoder(response.Body).Decode(&user)
    //适用场景：如果数据来自io.Reader流，或者需要从数据流中解码多个值，使用json.Decoder
    //   http请求的读取，也属于流的读取 
    ```
##### 三.请求和响应
##### 四.请求和响应
##### 五.请求和响应