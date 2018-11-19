package main

import (
	"bytes"
	"encoding/binary"

	"github.com/cvgw/proto-stream-exp/proto/proxysql"
	"github.com/gogo/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

func main() {
	prefixSize := 4

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

	b := make([]byte, prefixSize)
	binary.BigEndian.PutUint32(b, uint32(len(binDataHg)))

	buf.Write(b)
	buf.Write(binDataHg)
	log.Info(buf.Bytes())

	b = make([]byte, prefixSize)
	binary.BigEndian.PutUint32(b, uint32(len(binDataD)))

	log.Infof("prefix length %d", len(b))

	buf.Write(b)
	buf.Write(binDataD)
	log.Info(buf.Bytes())

	b = buf.Bytes()
	s := b[:prefixSize]
	b = b[prefixSize:]

	size := binary.BigEndian.Uint32(s)

	pBin := b[:size]
	b = b[size:]

	unmarshalDigest := &proxysql.QueryDigest{}

	err = proto.Unmarshal(pBin, unmarshalDigest)
	log.Info(unmarshalDigest)
}
