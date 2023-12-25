package node

import (
	"fmt"
	"sync"

	"github.com/iegad/gox/frm/game"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
	"github.com/iegad/gox/pb"
	"google.golang.org/protobuf/proto"
)

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
	PlayerManager game.PlayerManager
	cMap          sync.Map
	wg            sync.WaitGroup
	handlers      map[int32]NodeHandler
	msgCh         chan *message
}

func NewNode(nodeID int32) *Node {
	return &Node{
		NodeID: nodeID,
		msgCh:  make(chan *message, 10000),
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

func (this_ *Node) runProxy(c *nw.Client) {
	defer this_.wg.Done()

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

		if in.NodeID != this_.NodeID {
			log.Error("%v is invalid", in.NodeID)
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
