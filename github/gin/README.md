###gin请求数据校验 playground/validator
go get github.com/go-playground/validator/v10 
   
对请求数据的校验

### gin-jwt
go get github.com/appleboy/gin-jwt/v2  

TokenLookup：token检索模式，用于提取token，默认值为header:Authorization  
SigningAlgorithm：签名算法，默认值为HS256  
Timeout：token过期时间，默认值为time.Hour  
TimeFunc：测试或服务器在其他时区可设置该属性，默认值为time.Now  
TokenHeadName：token在请求头时的名称，默认值为Bearer  
IdentityKey：身份验证的key值，默认值为identity  
Realm：可以理解成该中间件的名称，用于展示，默认值为gin jwt  
CookieName：Cookie名称，默认值为jwt  
privKey：私钥  
pubKey：公钥  
Authenticator函数：根据登录信息对用户进行身份验证的回调函数  
PayloadFunc函数：登录期间的回调的函数  
IdentityHandler函数：解析并设置用户身份信息  
Authorizator函数：接收用户信息并编写授权规则，本项目的API权限控制就是通过该函数编写授权规则的  
Unauthorized函数：处理不进行授权的逻辑  
LoginResponse函数：完成登录后返回的信息，用户可自定义返回数据，默认返回  