/**
* Created by Michael on 2015/8/5.
*	web服务器与推送服务器通讯的RPC调用
*
*
 */
package rpc

import (
	log "github.com/golang/glog"
	"net"
	"net/rpc"
)

type RPCServer struct {
	rpc  *rpc.Server
	port string
}

var Server *RPCServer

var Client *RPCClient

type RPCClient struct {
	rpc       *rpc.Client
	port      string
	ip        string
	conn      *net.TCPConn
	Connected bool
}

func CreateClient(ip string, port string) {
	Client = &RPCClient{ip: ip, port: port}

}

func CreateServer(port string) {
	Server = &RPCServer{port: port}
	go Server.Listen()
}

func (this *RPCClient) Call(serviceMethod string, args interface{}, reply interface{}) interface{} {
	defer func() {
		if e := recover(); e != nil {
			log.Errorln("rpc call", e)
		}
	}()

	if this.Connected == false {
		e := Client.connect()
		if e != nil {
			return e
		}
		this.Connected = true
	}
	//	defer this.rpc.Close()

	err := this.rpc.Call(serviceMethod, args, reply)
	if err != nil {
		log.Errorln("error:", err)
		e := Client.connect()
		if e == nil {
			this.rpc.Call(serviceMethod, args, reply)
		}
	}
	log.Infoln(reply)

	return nil
}

func (this *RPCClient) connect() (e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln("rpc call", err)
			e = err
		}
	}()
	log.Infoln(this.ip, this.port)
	address, err := net.ResolveTCPAddr("tcp", this.ip+":"+this.port)
	if err != nil {
		log.Errorln(err)
		return err
	}
	if this.conn != nil {
		this.conn.Close()
	}
	this.conn, err = net.DialTCP("tcp", nil, address)
	if err != nil {
		log.Errorln(err)
		return err
	}

	this.rpc = rpc.NewClient(this.conn)
	if err != nil {
		log.Errorln(err)
		return err
	}
	//	defer conn.Close()

	return nil
}

func (this *RPCServer) Register(rcvr interface{}) {
	this.rpc.Register(rcvr)
}

func (this *RPCServer) Listen() {
	this.rpc = rpc.NewServer()
//	this.rpc.Register(new(BroadcastPublic))

	l, e := net.Listen("tcp", ":"+this.port) // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}

	this.rpc.Accept(l)
}
