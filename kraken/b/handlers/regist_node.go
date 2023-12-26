package handlers

import (
	"fmt"

	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/node"
	"github.com/iegad/gox/frm/nw"
	"github.com/iegad/gox/kraken/conf"
	"github.com/iegad/gox/kraken/m"
	"github.com/iegad/gox/pb"
	"google.golang.org/protobuf/proto"
)

func RegistNode(sess *nw.Sess, in *pb.Package) error {
	if sess == nil {
		log.Fatal("sess is nil")
	}

	if in == nil {
		log.Fatal("in is nil")
	}

	nodeCode, err := node.NodeCodeToString(in.NodeCode)
	if err != nil {
		return err
	}

	if nodeCode != conf.Instance.NodeCode {
		return fmt.Errorf("regist node req must be this engine code")
	}

	req := &pb.RegistNodeReq{}
	err = proto.Unmarshal(in.Data, req)
	if err != nil {
		return err
	}

	if req.NodeID <= 0 {
		return registNodeRsp(sess, req, &pb.RegistNodeRsp{
			Code:  -1,
			Error: m.Err_B_NodeIDInvalid.Error(),
		})
	}

	if len(req.NodeCode) != 16 {
		return registNodeRsp(sess, req, &pb.RegistNodeRsp{
			Code:  -1,
			Error: pb.Err_NodeCodeInvalid.Error(),
		})
	}

	m.Nodes.AddNode(m.NewNode(req.NodeCode, sess))
	return registNodeRsp(sess, req, &pb.RegistNodeRsp{Code: 0})
}

func registNodeRsp(sess *nw.Sess, req *pb.RegistNodeReq, rsp *pb.RegistNodeRsp) error {

	nd, ok := sess.UserData.(*m.Node)
	if !ok {
		log.Fatal("sess.user_data is not *m.Node")
	}

	data, err := proto.Marshal(rsp)
	if err != nil {
		log.Fatal(err)
	}

	nd.Idempotent++
	out := &pb.Package{
		NodeCode:   req.NodeCode,
		MessageID:  pb.MID_B_RegistNodeRsp,
		Idempotent: nd.Idempotent,
		Data:       data,
	}

	data, err = proto.Marshal(out)
	if err != nil {
		log.Fatal(err)
	}

	_, err = sess.Write(data)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
