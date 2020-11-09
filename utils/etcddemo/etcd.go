package etcddemo

//import (
//	"go.etcd.io/etcd/clientv3"
//	"log"
//)
//func Lock(key string, cli *clientv3.Client) error {
//	//获取key，判断是否存在锁
//	resp, err := cli.Get(context.Background(), key)
//	if err != nil {
//		return err
//	}
//	//锁存在，返回上锁失败
//	if len(resp.Kvs) > 0 {
//		return errors.New("lock fail")
//	}
//	_, err = cli.Put(context.Background(), key, "lock")
//	if err != nil {
//		return err
//	}
//	return nil
//}
////删除key，解锁
//func UnLock(key string, cli *clientv3.Client) error {
//	_, err := cli.Delete(context.Background(), key)
//	return err
//}
////等待key删除后再竞争锁
//func waitDelete(key string, cli *clientv3.Client) {
//	rch := cli.Watch(context.Background(), key)
//	for wresp := range rch {
//		for _, ev := range wresp.Events {
//			switch ev.Type {
//			case mvccpb.DELETE: //删除
//				return
//			}
//		}
//	}
//}
//
//
///*
//避免死锁版
//#操作没完成，锁被别人占用了，不安全
//#操作完成后，进行解锁，这时候把别人占用的锁解开了
//解决方案：给key添加过期时间后，以Keep leases alive方式延续leases，当client正常持有锁时，锁不会过期；当client程序崩掉后，程序不能执行Keep leases alive，从而让锁过期，避免死锁
//*/
//func Lock2(key string, cli *clientv3.Client) error {
//	//获取key，判断是否存在锁
//	resp, err := cli.Get(context.Background(), key)
//	if err != nil {
//		return err
//	}
//	//锁存在，等待解锁后再竞争锁
//	if len(resp.Kvs) > 0 {
//		waitDelete(key, cli)
//		return Lock(key)
//	}
//	//设置key过期时间
//	resp, err := cli.Grant(context.TODO(), 30)
//	if err != nil {
//		return err
//	}
//	//设置key并绑定过期时间
//	_, err = cli.Put(context.Background(), key, "lock", clientv3.WithLease(resp.ID))
//	if err != nil {
//		return err
//	}
//	//延续key的过期时间
//	_, err = cli.KeepAlive(context.TODO(), resp.ID)
//	if err != nil {
//		return err
//	}
//	return nil
//}
////通过让key值过期来解锁
//func UnLock2(resp *clientv3.LeaseGrantResponse, cli *clientv3.Client) error {
//	_, err := cli.Revoke(context.TODO(), resp.ID)
//	return err
//}
//
//
//
//func ExampleMutex_Lock() {
//	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer cli.Close()
//
//	// create two separate sessions for lock competition
//	s1, err := concurrency.NewSession(cli)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer s1.Close()
//	m1 := concurrency.NewMutex(s1, "/my-lock/")
//
//	s2, err := concurrency.NewSession(cli)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer s2.Close()
//	m2 := concurrency.NewMutex(s2, "/my-lock/")
//
//	// acquire lock for s1
//	if err := m1.Lock(context.TODO()); err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("acquired lock for s1")
//
//	m2Locked := make(chan struct{})
//	go func() {
//		defer close(m2Locked)
//		// wait until s1 is locks /my-lock/
//		if err := m2.Lock(context.TODO()); err != nil {
//			log.Fatal(err)
//		}
//	}()
//
//	if err := m1.Unlock(context.TODO()); err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("released lock for s1")
//
//	<-m2Locked
//	fmt.Println("acquired lock for s2")
//
//	// Output:
//	// acquired lock for s1
//	// released lock for s1
//	// acquired lock for s2
//}