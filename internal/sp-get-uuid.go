package ucsm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/cloud104/tks-go-ucsm-sdk/api"
	"github.com/cloud104/tks-go-ucsm-sdk/mo"
)

const uuidSize = 16

type UUID struct {
	TimeLow          uint32
	TimeMid          uint16
	TimeHiAndVersion uint16
	ClockSeqHiAndRes uint8
	ClockSeqLow      uint8
	Node             [6]byte
}

func UUIDToLittleEndian(uuidBigEndian string) (uuidLittleEndian string, err error) {
	var uuid UUID
	b := bytes.NewBuffer(make([]byte, 0, uuidSize))
	if _, err = fmt.Sscanf(uuidBigEndian, "%x-%x-%x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		&uuid.TimeLow, &uuid.TimeMid, &uuid.TimeHiAndVersion,
		&uuid.ClockSeqHiAndRes, &uuid.ClockSeqLow,
		&uuid.Node[0], &uuid.Node[1], &uuid.Node[2],
		&uuid.Node[3], &uuid.Node[4], &uuid.Node[5]); err != nil {
		return uuidLittleEndian, fmt.Errorf("UUIDToLittleEndian: invalid UUID format %s: %w", uuidBigEndian, err)
	}
	if err = binary.Write(b, binary.LittleEndian, uuid.TimeLow); err != nil {
		return uuidLittleEndian, fmt.Errorf("UUIDToLittleEndian: failed to write TimeLow: %w", err)
	}
	if err = binary.Write(b, binary.LittleEndian, uuid.TimeMid); err != nil {
		return uuidLittleEndian, fmt.Errorf("UUIDToLittleEndian: failed to write TimeMid: %w", err)
	}
	if err = binary.Write(b, binary.LittleEndian, uuid.TimeHiAndVersion); err != nil {
		return uuidLittleEndian, fmt.Errorf("UUIDToLittleEndian: failed to write TimeHiAndVersion: %w", err)
	}
	if err = binary.Write(b, binary.BigEndian, uuid.ClockSeqHiAndRes); err != nil {
		return uuidLittleEndian, fmt.Errorf("UUIDToLittleEndian: failed to write ClockSeqHiAndRes: %w", err)
	}
	if err = binary.Write(b, binary.BigEndian, uuid.ClockSeqLow); err != nil {
		return uuidLittleEndian, fmt.Errorf("UUIDToLittleEndian: failed to write ClockSeqLow: %w", err)
	}
	if err = binary.Write(b, binary.BigEndian, uuid.Node); err != nil {
		return uuidLittleEndian, fmt.Errorf("UUIDToLittleEndian: failed to write Node: %w", err)
	}
	reader := bytes.NewReader(b.Bytes())
	if err = binary.Read(reader, binary.BigEndian, &uuid); err != nil {
		return uuidLittleEndian, fmt.Errorf("UUIDToLittleEndian: failed to read UUID: %w", err)
	}
	uuidLittleEndian = strings.ToUpper(fmt.Sprintf("%x-%x-%x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		uuid.TimeLow, uuid.TimeMid, uuid.TimeHiAndVersion,
		uuid.ClockSeqHiAndRes, uuid.ClockSeqLow,
		uuid.Node[0], uuid.Node[1], uuid.Node[2],
		uuid.Node[3], uuid.Node[4], uuid.Node[5],
	))
	return uuidLittleEndian, nil
}

func SpGetUUID(c *api.Client, spDn string) (assocState string, err error) {
	var lsServers []*mo.LsServer
	if lsServers, err = ServerGet(c, spDn, "instance"); err == nil {
		if len(lsServers) > 0 {
			assocState = lsServers[0].UUID
		} else {
			err = fmt.Errorf("ServerGet: no server %s found", spDn)
		}
	}
	return assocState, fmt.Errorf("SpGetUUID: failed to get UUID for service profile '%s': %w", spDn, err)
}
