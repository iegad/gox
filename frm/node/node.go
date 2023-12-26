package node

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/iegad/gox/frm/game"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
	"github.com/iegad/gox/pb"
	"google.golang.org/protobuf/proto"
)

func NodeCodeToString(ncode []byte) (string, error) {
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
	NodeID        int32
	NodeCode      string
	Idempotent    int64
	PlayerManager game.PlayerManager
	cMap          sync.Map
	wg            sync.WaitGroup
	handlers      map[int32]NodeHandler
	msgCh         chan *message
}

func NewNode(nodeID int32, nodeCode []byte) *Node {
	nstr, err := NodeCodeToString(nodeCode)
	if err != nil {
		log.Fatal(err)
	}

	return &Node{
		NodeID:   nodeID,
		NodeCode: nstr,
		msgCh:    make(chan *message, 10000),
	}
}

func (this_ *Node) Run(n int) {
	for i := 0; i < n; i++ {
		this_.wg.Add(1)
		go func() {
			defer this_.wg.Done()

			for msg := range this_.msgCh {
				if handler, ok := this_.handlers[msg.pack.MessageID]; ok {
					handler.OnPackage(msg.pack, msg.c)
				} else {
					this_.KickSession(msg.c, msg.pack.RealAddr)
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

func (this_ *Node) AddProxy(rep string) error {
	c, err := nw.NewTcpClient("", rep, 0)
	if err != nil {
		return err
	}

	if _, ok := this_.cMap.Load(rep); ok {
		return fmt.Errorf("%v is already exists", rep)
	}

	c.UserData = this_
	this_.cMap.Store(rep, c)
	this_.wg.Add(1)
	go this_.runProxy(c)
	return nil
}

func (this_ *Node) registNode(c *nw.Client) error {
	nodeUID, err := uuid.Parse(this_.NodeCode)
	if err != nil {
		return err
	}

	this_.Idempotent++
	req := &pb.RegistNodeReq{
		NodeID:   this_.NodeID,
		NodeCode: nodeUID[:],
	}

	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	krakenUID, err := uuid.Parse("ae906ff9-552b-4f63-b807-95a1195deddf") // TODO
	if err != nil {
		return err
	}

	out := &pb.Package{
		NodeCode:   krakenUID[:],
		MessageID:  pb.MID_B_RegistNodeReq,
		Idempotent: this_.Idempotent,
		Data:       data,
	}

	data, err = proto.Marshal(out)
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

	in := &pb.Package{}
	err = proto.Unmarshal(data, in)
	if err != nil {
		return err
	}

	if len(in.NodeCode) != 16 {
		return fmt.Errorf("in.NodeCode is invalid")
	}

	nodeCode, err := NodeCodeToString(in.NodeCode)
	if err != nil {
		return err
	}

	if nodeCode != this_.NodeCode {
		return fmt.Errorf("%v <> %v", nodeCode, this_.NodeCode)
	}

	if in.MessageID != pb.MID_B_RegistNodeRsp {
		return fmt.Errorf("message_id: %v is invalid", in.MessageID)
	}

	rsp := &pb.RegistNodeRsp{}
	err = proto.Unmarshal(in.Data, rsp)
	if err != nil {
		return err
	}

	if rsp.Code != 0 {
		return fmt.Errorf("%v", rsp.Error)
	}

	return nil
}

func (this_ *Node) runProxy(c *nw.Client) {
	defer this_.wg.Done()

	err := this_.registNode(c)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Info("注册网关成功")
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

		if len(in.NodeCode) != 16 {
			log.Error("in.NodeCode is invalid: %v", c.Raw().RemoteAddr().String())
			break
		}

		nstr, err := NodeCodeToString(in.NodeCode)
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

func (this_ *Node) KickSession(c *nw.Client, realAddr string) {
	// todo
}
