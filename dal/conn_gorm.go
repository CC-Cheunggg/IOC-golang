package dal

import (
	"github.com/alibaba/ioc-golang/extension/db/gorm"
)

//type GORMConn[T interface{}] struct {
//	GORM    gorm.GORMDBIOCInterface `normal:""`
//	queryer T
//	once    sync.Once
//}
//
//func (conn *GORMConn[T]) GetConn() gorm.GORMDBIOCInterface {
//	return conn.GORM
//}
//
//func (conn *GORMConn[T]) GetQueryer() T {
//	return conn.queryer
//}
//
//func (conn *GORMConn[T]) SetQueryer(queryer T) {
//	conn.once.Do(func() {
//		conn.queryer = queryer
//	})
//}

//type Conn[T interface{}] struct {
//	Test T
//}
//
//type Conns struct {
//
//}
//
//type Demo struct {
//	Test Conn[Conns]
//}
//
//func NewConn[T interface{}](test T) *Conn[T] {
//	return &Conn[T]{Test: test}
//}
//
//func main() {
//	conn := NewConn("")
//	strings.Compare("", conn.Test)
//}

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type GORMConn struct {
	GORM gorm.GORMDBIOCInterface `normal:""`
}

func (conn *GORMConn) GetConn() gorm.GORMDBIOCInterface {
	return conn.GORM
}
