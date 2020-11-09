#程序监控  
block：goroutine的阻塞信息，本例就截取自一个goroutine阻塞的demo，但block为0，没掌握block的用法  
goroutine：所有goroutine的信息，full goroutine stack dump是输出所有goroutine的调用栈，是goroutine的debug=2  
heap：堆内存的信息  
mutex：锁的信息  
threadcreate：线程信息  
##流程
1 运行domo  
2 浏览器打开地址 http://127.0.0.1:6060/debug/pprof/  
查看goroutine  
http://127.0.0.1:6060/debug/pprof/goroutine?debug=1  
查看goroutine信息 可以看到每个goroutine信息 和上面的互补  
http://127.0.0.1:6060/debug/pprof/goroutine?debug=2