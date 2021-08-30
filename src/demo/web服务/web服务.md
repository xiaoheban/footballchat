j### 什么时候Web服务
### JSON处理
1. 写入json
  + 数据如果是在内存 
  bytes, err := json.Marshal(&post)
	bytes, err := json.MarshalIndent(&post, "", "\t\t")
  + 数据如果是在流中如文件流/http流中
    ```
    file, _ := os.Create("test2.json")
	  defer file.Close()
	  jsonEncoder := json.NewEncoder(file)
	  if jsonEncoder != nil {
		  jsonEncoder.Encode(&post)
	  }
    ```
2. 读取json
   + 从内存中
    ```
    bytes, err := ioutil.ReadFile("test.json")
	  if err != nil {
		  panic(err)
	  }
	  var post1 Post
	  err = json.Unmarshal(bytes, &post1)
    ```
   + 从文件流/输入流中
    ```
    file, err := os.Open("test2.json")
	  if err != nil {
		  panic(err)
	  }
	  defer file.Close()
	  decorder := json.NewDecoder(file)
	  if decorder != nil {
		  var post2 Post
		  err := decorder.Decode(&post2)
		  if err == nil {
			  fmt.Println("get post from file use json.NewDecoder", post2)
		  }
	  }
    ```
方法的选择就是看数据是在内存还是在文件或其他输入流
### 创建Web服务
创建REST风格的web服务