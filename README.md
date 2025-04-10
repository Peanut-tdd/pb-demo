## protobuf 多包及同包导入proto文件demo

### 生成 go mod 文件
```
go mod init github.com/pb-demo
```




### 生成.pb文件到指定目录

1. 定义好go_package,并import需要的包
```
syntax = "proto3";
package user;  // 声明所在包
option go_package = "github.com/pb-demo/pb/user";  // 定义 go package,按go module及生成的.pb文件目录来，对后续生成的.pb文件相互引用特别重要，会报错

import "user/message.proto";  // 导入同包内的其他 proto 文件
import "article/article.proto"; //导入不同包proto文件
```



2. 指定目录生成.pb文件（本例生成文件到根目录的pb文件夹下）
```
protoc  -I ./proto --go_out=./pb --go_opt paths=source_relative ./proto/*/*.proto
```




### 生成.pb文件到当前目录

1. 定义好go_package，并import需要的包
```
option go_package="github.com/pb-demo/proto/user"; //生成.pb文件在当前目录下

import "user/message.proto";  // 导入同包内的其他 proto 文件
import "article/article.proto"; //导入不同包proto文件
```



2. 生成.pb文件到.proto文件当前目录
```
protoc -I .  --go_out=paths=source_relative:.   ./*/*.proto 
```





注：
--go_out和--go_opt控制.pb文件生成位置
paths参数有两个选项，分别是 import 和 source_relative，默认为 import，表示按照生成的Go代码的包的全路径去创建目录层级，source_relative 表示按照 proto源文件的目录层级去创建Go代码的目录层级，如果目录已存在则不用创建











