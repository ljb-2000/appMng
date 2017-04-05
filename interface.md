#### 1.创建应用
##### URI: /appMng/v1/apps  
##### METHOD: POST
##### BODY:   
 
``` json
{
	"name": "testapp6",
	"description": "test",
	"lang": "go",
	"type": "app",
	"services": ""
}
```

##### 返回BODY：  
``` json
{
  "code": "0",
  "data": {
    "id": "9d4cd8d2-2353-4fb2-908a-5e4567bf581d",
    "name": "testapp6",
    "user": "18022222222",
    "description": "test",
    "createdTime": "2017-03-21 10:31:30",
    "lang": "go",
    "type": "app",
    "services": "",
    "git": "http://223.202.32.60:8071/gk-test/testapp6.git",
    "state": "created"
  },
  "msg": "OK"
}
```

#### 2. 查看应用列表
##### URI: /appMng/v1/apps  
##### METHOD: GET
##### BODY:  空  

##### 返回BODY：  
``` json
[
  {
    "id": "b0c517d2-b822-4fb8-9150-703f3a1b5530",
    "name": "testapp4",
    "user": "18022222222",
    "description": "test",
    "createdTime": "2017-03-21 10:31:34",
    "lang": "go",
    "type": "app",
    "services": "",
    "git": "http://223.202.32.60:8071/gk-test/testapp4.git",
    "state": "created"
  },
  {
      "id": "b0c517d2-b822-4fb8-9150-703f23243423",
      "name": "testapp5",
      "user": "18022222222",
      "description": "test",
      "createdTime": "2017-03-21 10:31:40",
      "lang": "go",
      "type": "app",
      "services": "",
      "git": "http://223.202.32.60:8071/gk-test/testapp4.git",
      "state": "created"
    }
]
```


#### 3. 查询某个应用
##### URI: /appMng/v1/apps/:appId  
##### METHOD: GET
##### BODY:  空  

##### 返回BODY：  
``` json
{
  "code": "0",
  "data": {
    "id": "9d4cd8d2-2353-4fb2-908a-5e4567bf581d",
    "name": "testapp6",
    "user": "18022222222",
    "description": "test",
    "createdTime": "2017-03-21 10:31:30",
    "lang": "go",
    "type": "app",
    "services": "",
    "git": "http://223.202.32.60:8071/gk-test/testapp6.git",
    "state": "created"
  },
  "msg": "OK"
}
```

#### 4. 删除应用
##### URI: /appMng/v1/apps/:appId  
##### METHOD: DELETE
##### BODY:  空  

##### 返回BODY：  
``` json
{
  "code": "0",
  "msg": "OK"
}
```


#### 5. 生成镜像
##### URI: /appMng/v1/images  
##### METHOD: POST
##### BODY:   
``` json
{
	"name": "testabc",
	"tag": "1.0",
	"appId": "app123456",
	"lang": "go",
	"git": "http://223.202.32.60:8071/gk-test/testapp6.git"
}
```

##### 返回BODY：  
``` json
{
    "code": "0",
    "msg": "OK",
    "img": "223.202.32.59:8080/gk-test/testabc:1.0"
}
```

#### 6. 查询某个app下的所有镜像
##### URI: /appMng/v1/apps/:appId/images  
##### METHOD: GET
##### BODY:    

##### 返回BODY：
``` json
[
  {
    "id": "9d4cd8d2-2353-4fb2-908a-5e4567b123",
    "name": "testabc",
    "tag": "1.0",
    "lang": "go",
    "img": "223.202.32.59:8080/apptest/testabc:1.0",
    "state": ""
  },
  {
    "id": "9d4cd8d2-2353-4fb2-908a432545436346g",
    "name": "test2",
    "tag": "1.0",
    "lang": "go",
    "img": "223.202.32.59:8080/apptest/test2:1.0",
    "state": ""
    }
]
```

#### 7. 删除镜像
##### URI: /appMng/v1/images/:imageId  
##### METHOD: DELETE
##### BODY:    

##### 返回BODY：
``` json
{
  "code": "0",
  "msg": "OK"
}
```

#### 8. 根据镜像部署服务
##### URI: /appMng/v1/services  
##### METHOD: POST
##### BODY:    

##### 返回BODY：
