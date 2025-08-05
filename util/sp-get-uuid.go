package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/gfalves87/tks-go-ucsm-sdk/api"
	"github.com/gfalves87/tks-go-ucsm-sdk/mo"
)

const uuidSize = 16

type Uuid struct {
	TimeLow          uint32
	TimeMid          uint16
	TimeHiAndVersion uint16
	ClockSeqHiAndRes uint8
	ClockSeqLow      uint8
	Node             [6]byte
}

func UuidToLittleEndian(uuidBigEndian string) (uuidLittleEndian string, err error) {
	var uuid Uuid
	b := bytes.NewBuffer(make([]byte, 0, uuidSize))
	if _, err = fmt.Sscanf(uuidBigEndian, "%x-%x-%x-%02x%02x-%02x%02x%02x%02x%02x%02x", &uuid.TimeLow, &uuid.TimeMid, &uuid.TimeHiAndVersion, &uuid.ClockSeqHiAndRes, &uuid.ClockSeqLow, &uuid.Node[0], &uuid.Node[1], &uuid.Node[2], &uuid.Node[3], &uuid.Node[4], &uuid.Node[5]); err != nil {
		return
	}
	if err = binary.Write(b, binary.LittleEndian, uuid.TimeLow); err != nil {
		return
	}
	if err = binary.Write(b, binary.LittleEndian, uuid.TimeMid); err != nil {
		return
	}
	if err = binary.Write(b, binary.LittleEndian, uuid.TimeHiAndVersion); err != nil {
		return
	}
	if err = binary.Write(b, binary.BigEndian, uuid.ClockSeqHiAndRes); err != nil {
		return
	}
	if err = binary.Write(b, binary.BigEndian, uuid.ClockSeqLow); err != nil {
		return
	}
	if err = binary.Write(b, binary.BigEndian, uuid.Node); err != nil {
		return
	}
	reader := bytes.NewReader(b.Bytes())
	if err = binary.Read(reader, binary.BigEndian, &uuid); err != nil {
		return
	}
	uuidLittleEndian = strings.ToUpper(fmt.Sprintf("%x-%x-%x-%02x%02x-%02x%02x%02x%02x%02x%02x", uuid.TimeLow, uuid.TimeMid, uuid.TimeHiAndVersion, uuid.ClockSeqHiAndRes, uuid.ClockSeqLow, uuid.Node[0], uuid.Node[1], uuid.Node[2], uuid.Node[3], uuid.Node[4], uuid.Node[5]))
	return
}

func SpGetUuid(c *api.Client, spDn string) (assocState string, err error) {
	var lsServers []*mo.LsServer
	if lsServers, err = ServerGet(c, spDn, "instance"); err == nil {
		if len(lsServers) > 0 {
			assocState = lsServers[0].Uuid
		} else {
			err = fmt.Errorf("ServerGet: no server %s found", spDn)
		}
	}
	return
}
