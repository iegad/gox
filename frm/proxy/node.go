package proxy

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
	"github.com/iegad/gox/pb"
	"google.golang.org/protobuf/proto"
)

func NodeUIDToCode(ncode []byte) (string, error) {
	if len(ncode) != 16 {
		log.Fatal("ncode is nil")
	}

	return uuid.UUID(ncode).String(), nil
}

type NodeHandler interface {
	GetMessageID() int32
	OnPackage(in *pb.Package, c *nw.Client)
}

type message struct {
	pack *pb.Package
	c    *nw.Client
}

type Node struct {
	NodeCode   string
	Idempotent int64
	proxyMap   sync.Map
	wg         sync.WaitGroup
	handlers   map[int32]NodeHandler
	msgCh      chan *message
}

func NewNode(nodeUID []byte) *Node {
	if len(nodeUID) != 16 {
		log.Fatal("nodeID is invalid")
	}

	nodeCode, err := NodeUIDToCode(nodeUID)
	if err != nil {
		log.Fatal(err)
	}

	return &Node{
		NodeCode: nodeCode,
		msgCh:    make(chan *message, 10000),
	}
}

func (this_ *Node) Run(n int) {
	if n <= 0 {
		n = runtime.NumCPU()
	}

	for i := 0; i < n; i++ {
		this_.wg.Add(1)
		go func() {
			defer this_.wg.Done()

			for msg := range this_.msgCh {
				if handler, ok := this_.handlers[msg.pack.MessageID]; ok {
					handler.OnPackage(msg.pack, msg.c)
				} else {
					this_.KickPlayer(msg.c, msg.pack.RealAddr)
				}
			}
		}()
	}

	this_.wg.Wait()
}

func (this_ *Node) AddHandler(h NodeHandler) {
	if _, ok := this_.handlers[h.GetMessageID()]; ok {
		log.Fatal("%v is already exists", h.GetMessageID())
	}

	this_.handlers[h.GetMessageID()] = h
}

func (this_ *Node) AddProxy(nodeCode, rep string) {
	if _, ok := this_.proxyMap.Load(rep); ok {
		return
	}

	this_.proxyMap.Store(rep, nil)
	go this_.runProxy(nodeCode, rep)
}

func (this_ *Node) registNode(nodeCode string, c *nw.Client) error {
	nodeUID, err := uuid.Parse(this_.NodeCode)
	if err != nil {
		return err
	}

	this_.Idempotent++
	data, err := proto.Marshal(&pb.RegistNodeReq{
		NodeUID: nodeUID[:],
	})
	if err != nil {
		return err
	}

	proxyUID, err := uuid.Parse(nodeCode)
	if err != nil {
		return err
	}

	data, err = pb.SerializeNodePackage(proxyUID[:], pb.MID_B_RegistNodeReq, this_.Idempotent, data)
	if err != nil {
		return err
	}

	_, err = c.TcpWrite(data)
	if err != nil {
		return err
	}

	data, err = c.TcpRead()
	if err != nil {
		return err
	}

	in, err := pb.ParseNodePackage(data)
	if err != nil {
		return err
	}

	nodeCode, err = NodeUIDToCode(in.NodeUID)
	if err != nil {
		return err
	}

	if nodeCode != this_.NodeCode {
		return fmt.Errorf("%v <> %v", nodeCode, this_.NodeCode)
	}

	if in.MessageID != pb.MID_B_RegistNodeRsp {
		return fmt.Errorf("message_id: %v is invalid", in.MessageID)
	}

	rsp, err := pb.ParseMessage[pb.RegistNodeRsp](in.Data)
	if err != nil {
		return err
	}

	if rsp.Code != 0 {
		return fmt.Errorf("%v", rsp.Error)
	}

	return nil
}

func (this_ *Node) runProxy(nodeCode, rep string) {
	defer this_.wg.Done()

	for {
		log.Info("连接网关[%v:%v]...开始", rep, nodeCode)
		c, err := nw.NewTcpClient("", rep, 0)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		log.Info("连接网关[%v:%v]...成功", rep, nodeCode)
		c.UserData = this_
		if v, ok := this_.proxyMap.LoadAndDelete(rep); ok {
			if tmpc, ok := v.(*nw.Client); ok {
				tmpc.Raw().Close()
			}
		}
		this_.proxyMap.Store(rep, c)
		this_.wg.Add(1)

		log.Info("注册网关[%v:%v]...开始", rep, nodeCode)
		err = this_.registNode(nodeCode, c)
		if err != nil {
			log.Error(err)
			time.Sleep(time.Second)
			continue
		}

		log.Info("注册网关[%v:%v]...成功", rep, nodeCode)
		for {
			data, err := c.TcpRead()
			if err != nil {
				log.Error(err)
				break
			}

			in := &pb.Package{}
			err = proto.Unmarshal(data, in)
			if err != nil {
				log.Error(err)
				break
			}

			if len(in.NodeUID) != 16 {
				log.Error("in.NodeCode is invalid: %v", c.Raw().RemoteAddr().String())
				break
			}

			nstr, err := NodeUIDToCode(in.NodeUID)
			if err != nil {
				log.Error(err)
				break
			}

			if nstr != this_.NodeCode {
				log.Error("in.NodeCode is invalid: %v <> %v", nstr, this_.NodeCode)
				break
			}

			if len(in.RealAddr) == 0 {
				log.Error("real addr is invalid: %v", c.Raw().RemoteAddr().String())
				break
			}

			msg := &message{
				c:    c,
				pack: in,
			}

			this_.msgCh <- msg
		}
	}

}

func (this_ *Node) KickPlayer(c *nw.Client, realAddr string) {
	// todo
}
