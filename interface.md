#### 1.创建应用
##### URI: /appMng/v1/apps  
##### METHOD: POST
##### BODY:   
 
``` json
{
	"name": "testapp6",
	"description": "test",
	"lang": "go",
	"type": "app"
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

#### 3. 删除应用
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