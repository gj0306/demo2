#进阶操作  
###1 超时处理  
###2 自定义认证  
生成RSA私钥：openssl genrsa -out server.key 2048  
生成RSA私钥，命令的最后一个参数，将指定生成密钥的位数，如果没有指定，默认512  
生成ECC私钥：openssl ecparam -genkey -name secp384r1 -out server.key  
生成ECC私钥，命令为椭圆曲线密钥参数生成及操作，本命令中ECC曲线选择的是secp384r1  
  
生成公钥  
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650  
openssl req：生成自签名证书，-new指生成证书请求、-sha256指使用sha256加密、-key指定私钥文件、-x509指输出证书、-days 3650为有效期  
  
此后则输入证书拥有者信息  
Country Name (2 letter code) [AU]:CN  
State or Province Name (full name) [Some-State]:XxXx  
Locality Name (eg, city) []:XxXx  
Organization Name (eg, company) [Internet Widgits Pty Ltd]:XX Co. Ltd  
Organizational Unit Name (eg, section) []:Dev  
Common Name (e.g. server FQDN or YOUR name) []:go-grpc-example  
Email Address []:xxx@xxx.com  

###3 Token
Authentication：  
GetRequestMetadata  
RequireTransportSecurity  

###4 拦截器
