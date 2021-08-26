j### 什么时候Web服务
### JSON处理
1. 写入json
  bytes, err := json.Marshal(&post)
	bytes, err := json.MarshalIndent(&post, "", "\t\t")
2. 读取j'son

### 创建Web服务