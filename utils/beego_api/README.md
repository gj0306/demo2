#1 安装beego bee

go get github.com/astaxie/beego 
go get github.com/beego/bee

#2 加入环境变量
略～～

#3 生成自动化api CRUD
bee api api_crud -conn="root:root1234@tcp(127.0.0.1:3306)/xxx"

#4 生成文档并运行
bee run -downdoc=true -gendoc=true

# 注意
1 models 必须和main.go 同级,否则生成的文档没有model数据.

# 地址
http://127.0.0.1:8082/swagger/index.html