package main

import (
	"bytes"

	"github.com/cvgw/proto-stream-exp/proto/proxysql"
	"github.com/gogo/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

func newQueryDigestBinData(hostGroup int) []byte {
	queryDigest := &proxysql.QueryDigest{
		HostGroup: int64(hostGroup),
	}

	binData, err := proto.Marshal(queryDigest)
	if err != nil {
		log.Fatal(err)
	}

	return binData
}

func testUnmarshal(b []byte) proxysql.QueryDigest {
	mQueryDigest := proxysql.QueryDigest{}
	err := proto.Unmarshal(b, &mQueryDigest)
	if err != nil {
		log.Fatal(err)
	}

	log.Info(mQueryDigest)
	log.Infof("proto message size %d", proto.Size(&mQueryDigest))

	return mQueryDigest
}

func main() {
	binData := newQueryDigestBinData(99)

	blob := make([]byte, 0)
	buf := bytes.NewBuffer(blob)
	buf.Write(binData)

	log.Info(buf.Bytes())

	testUnmarshal(buf.Bytes())

	for i := 1; i < 11; i++ {
		binData := newQueryDigestBinData(i)

		buf.Write(binData)
	}

	log.Info(buf.Bytes())

	b := buf.Bytes()
	message := testUnmarshal(b)
	msgSize := proto.Size(&message)
	for len(b) > msgSize {
		b = b[:len(b)-msgSize]
		testUnmarshal(b)
		msgSize = proto.Size(&message)
	}
}
