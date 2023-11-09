package utils

import (
	"bytes"
	"encoding/binary"
)

// PacketReader is a simple packet reader.
type PacketReader struct {
	buffer *bytes.Reader
}

// NewPacketReader creates a new PacketReader.
func NewPacketReader(data []byte) *PacketReader {
	return &PacketReader{
		buffer: bytes.NewReader(data),
	}
}

// ReadByte reads a single byte from the packet.
func (pr *PacketReader) ReadByte() (byte, error) {
	return pr.buffer.ReadByte()
}

// ReadInt16 reads a 16-bit signed integer from the packet.
func (pr *PacketReader) ReadInt16() (int16, error) {
	var value int16
	err := binary.Read(pr.buffer, binary.LittleEndian, &value)
	return value, err
}

// ReadUint16 reads a 16-bit unsigned integer from the packet.
func (pr *PacketReader) ReadUint16() (uint16, error) {
	var value uint16
	err := binary.Read(pr.buffer, binary.LittleEndian, &value)
	return value, err
}

// ReadInt32 reads a 32-bit signed integer from the packet.
func (pr *PacketReader) ReadInt32() (int32, error) {
	var value int32
	err := binary.Read(pr.buffer, binary.LittleEndian, &value)
	return value, err
}

// ReadUint32 reads a 32-bit unsigned integer from the packet.
func (pr *PacketReader) ReadUint32() (uint32, error) {
	var value uint32
	err := binary.Read(pr.buffer, binary.LittleEndian, &value)
	return value, err
}

// ReadInt64 reads a 64-bit signed integer from the packet.
func (pr *PacketReader) ReadInt64() (int64, error) {
	var value int64
	err := binary.Read(pr.buffer, binary.LittleEndian, &value)
	return value, err
}

// ReadUint64 reads a 64-bit unsigned integer from the packet.
func (pr *PacketReader) ReadUint64() (uint64, error) {
	var value uint64
	err := binary.Read(pr.buffer, binary.LittleEndian, &value)
	return value, err
}

// ReadFloat32 reads a 32-bit floating-point number from the packet.
func (pr *PacketReader) ReadFloat32() (float32, error) {
	var value float32
	err := binary.Read(pr.buffer, binary.LittleEndian, &value)
	return value, err
}

// ReadFloat64 reads a 64-bit floating-point number from the packet.
func (pr *PacketReader) ReadFloat64() (float64, error) {
	var value float64
	err := binary.Read(pr.buffer, binary.LittleEndian, &value)
	return value, err
}

// ReadString reads a string from the packet.
func (pr *PacketReader) ReadString() (string, error) {
	length, err := pr.ReadUint16()
	if err != nil {
		return "", err
	}

	data := make([]byte, length)
	_, err = pr.buffer.Read(data)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
