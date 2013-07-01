package pgx

import (
	"fmt"
	"io"
	"strconv"
)

type valueTranscoder struct {
	DecodeText   func(*MessageReader, int32) interface{}
	DecodeBinary func(*MessageReader, int32) interface{}
	EncodeTo     func(io.Writer, interface{})
	EncodeFormat int16
}

var valueTranscoders map[oid]*valueTranscoder

func init() {
	valueTranscoders = make(map[oid]*valueTranscoder)

	// bool
	valueTranscoders[oid(16)] = &valueTranscoder{DecodeText: decodeBoolFromText}

	// int8
	valueTranscoders[oid(20)] = &valueTranscoder{DecodeText: decodeInt8FromText}

	// int2
	valueTranscoders[oid(21)] = &valueTranscoder{DecodeText: decodeInt2FromText}

	// int4
	valueTranscoders[oid(23)] = &valueTranscoder{DecodeText: decodeInt4FromText}

	// float4
	valueTranscoders[oid(700)] = &valueTranscoder{DecodeText: decodeFloat4FromText}

	// float8
	valueTranscoders[oid(701)] = &valueTranscoder{DecodeText: decodeFloat8FromText}
}

func decodeBoolFromText(mr *MessageReader, size int32) interface{} {
	s := mr.ReadByteString(size)
	switch s {
	case "t":
		return true
	case "f":
		return false
	default:
		panic(fmt.Sprintf("Received invalid bool: %v", s))
	}
}

func decodeInt8FromText(mr *MessageReader, size int32) interface{} {
	s := mr.ReadByteString(size)
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Received invalid int8: %v", s))
	}
	return n
}

func decodeInt2FromText(mr *MessageReader, size int32) interface{} {
	s := mr.ReadByteString(size)
	n, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		panic(fmt.Sprintf("Received invalid int2: %v", s))
	}
	return int16(n)
}

func decodeInt4FromText(mr *MessageReader, size int32) interface{} {
	s := mr.ReadByteString(size)
	n, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("Received invalid int4: %v", s))
	}
	return int32(n)
}

func decodeFloat4FromText(mr *MessageReader, size int32) interface{} {
	s := mr.ReadByteString(size)
	n, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(fmt.Sprintf("Received invalid float4: %v", s))
	}
	return float32(n)
}

func decodeFloat8FromText(mr *MessageReader, size int32) interface{} {
	s := mr.ReadByteString(size)
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(fmt.Sprintf("Received invalid float8: %v", s))
	}
	return v
}
