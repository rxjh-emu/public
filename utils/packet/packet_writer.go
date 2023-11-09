package utils

import (
	"bytes"
	"encoding/binary"
)

// PacketWriter is a simple packet writer.
type PacketWriter struct {
	buffer *bytes.Buffer
}

// NewPacketWriter creates a new PacketWriter.
func NewPacketWriter() *PacketWriter {
	return &PacketWriter{
		buffer: new(bytes.Buffer),
	}
}

// WriteByte writes a single byte to the packet.
func (pw *PacketWriter) WriteByte(value byte) {
	pw.buffer.WriteByte(value)
}

// WriteInt16 writes a 16-bit signed integer to the packet.
func (pw *PacketWriter) WriteInt16(value int16) {
	binary.Write(pw.buffer, binary.LittleEndian, value)
}

// WriteUint16 writes a 16-bit unsigned integer to the packet.
func (pw *PacketWriter) WriteUint16(value uint16) {
	binary.Write(pw.buffer, binary.LittleEndian, value)
}

// WriteInt32 writes a 32-bit signed integer to the packet.
func (pw *PacketWriter) WriteInt32(value int32) {
	binary.Write(pw.buffer, binary.LittleEndian, value)
}

// WriteUint32 writes a 32-bit unsigned integer to the packet.
func (pw *PacketWriter) WriteUint32(value uint32) {
	binary.Write(pw.buffer, binary.LittleEndian, value)
}

// WriteInt64 writes a 64-bit signed integer to the packet.
func (pw *PacketWriter) WriteInt64(value int64) {
	binary.Write(pw.buffer, binary.LittleEndian, value)
}

// WriteUint64 writes a 64-bit unsigned integer to the packet.
func (pw *PacketWriter) WriteUint64(value uint64) {
	binary.Write(pw.buffer, binary.LittleEndian, value)
}

// WriteFloat32 writes a 32-bit floating-point number to the packet.
func (pw *PacketWriter) WriteFloat32(value float32) {
	binary.Write(pw.buffer, binary.LittleEndian, value)
}

// WriteFloat64 writes a 64-bit floating-point number to the packet.
func (pw *PacketWriter) WriteFloat64(value float64) {
	binary.Write(pw.buffer, binary.LittleEndian, value)
}

// WriteString writes a string to the packet.
func (pw *PacketWriter) WriteString(value string) {
	pw.WriteUint16(uint16(len(value)))
	pw.buffer.WriteString(value)
}

// GetData returns the data written to the packet.
func (pw *PacketWriter) GetData() []byte {
	return pw.buffer.Bytes()
}
