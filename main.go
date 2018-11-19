package main

import (
	"bytes"

	"github.com/cvgw/proto-stream-exp/proto/proxysql"
	"github.com/gogo/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

func main() {
	queryDigestHg := &proxysql.QueryDigest{
		HostGroup: int64(99),
	}

	binDataHg, err := proto.Marshal(queryDigestHg)
	if err != nil {
		log.Fatal(err)
	}

	queryDigestD := &proxysql.QueryDigest{
		Digest: "meow",
	}

	binDataD, err := proto.Marshal(queryDigestD)
	if err != nil {
		log.Fatal(err)
	}

	blob := make([]byte, 0)
	buf := bytes.NewBuffer(blob)

	buf.Write(binDataHg)
	log.Info(buf.Bytes())

	buf.Write(binDataD)
	log.Info(buf.Bytes())

	b := buf.Bytes()
	message := testUnmarshal(b)
	log.Info(message)
}
