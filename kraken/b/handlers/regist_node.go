package handlers

import (
	"fmt"

	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
	"github.com/iegad/gox/frm/proxy"
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

	nodeCode, err := proxy.NodeUIDToCode(in.NodeUID)
	if err != nil {
		return err
	}

	if nodeCode != conf.Instance.NodeCode {
		return fmt.Errorf("regist node req must be this engine code")
	}

	req, err := pb.ParseMessage[pb.RegistNodeReq](in.Data)
	if err != nil {
		return err
	}

	if len(req.NodeUID) != 16 {
		return registNodeRsp(sess, req, &pb.RegistNodeRsp{
			Code:  -1,
			Error: pb.Err_NodeCodeInvalid.Error(),
		})
	}

	m.Nodes.AddNode(m.NewNode(req.NodeUID, sess))
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
	data, err = pb.SerializeNodePackage(req.NodeUID, pb.MID_B_RegistNodeRsp, nd.Idempotent, data)
	if err != nil {
		log.Fatal(err)
	}

	_, err = sess.TcpWrite(data)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
