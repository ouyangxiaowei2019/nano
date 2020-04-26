// Copyright (c) nano Authors. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package message

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/lonng/nano/internal/env"
	"github.com/lonng/nano/internal/log"
)

// Type represents the type of message, which could be Request/Notify/Response/Push
type Type byte

// Message types
const (
	Request  Type = 0x00
	Notify        = 0x01
	Response      = 0x02
	Push          = 0x03
)

const (
	msgRouteCompressMask = 0x08
	msgTypeMask          = 0x07
	msgRouteLengthMask   = 0xFF
	msgHeadLength        = 0x02
)

var types = map[Type]string{
	Request:  "Request",
	Notify:   "Notify",
	Response: "Response",
	Push:     "Push",
}

var (
	// Routes is default routes for meesage
	Routes map[string]uint16 = make(map[string]uint16)
	// Codes is default codes for meesage
	Codes map[uint16]string = make(map[uint16]string)
)

func (t Type) String() string {
	return types[t]
}

// Errors that could be occurred in message codec
var (
	ErrWrongMessageType  = errors.New("wrong message type")
	ErrInvalidMessage    = errors.New("invalid message")
	ErrRouteInfoNotFound = errors.New("route info not found in dictionary")
)

// Message represents a unmarshaled message or a message which to be marshaled
type Message struct {
	Type       Type   // message type
	ID         uint64 // unique id, zero while notify mode
	Route      string // route for locating service
	Data       []byte // payload
	compressed bool   // is message compressed
}

// New returns a new message instance
func New() *Message {
	return &Message{}
}

// String, implementation of fmt.Stringer interface
func (m *Message) String() string {
	return fmt.Sprintf("%s %s (%dbytes)", types[m.Type], m.Route, len(m.Data))
}

// Encode marshals message to binary format.
func (m *Message) Encode(routes map[string]uint16) ([]byte, error) {
	return Encode(m, routes)
}

func routable(t Type) bool {
	return t == Request || t == Notify || t == Push
}

func invalidType(t Type) bool {
	return t < Request || t > Push

}

// Encode marshals message to binary format. Different message types is corresponding to
// different message header, message types is identified by 2-4 bit of flag field. The
// relationship between message types and message header is presented as follows:
// The figure above indicates that the bit does not affect the type of message.
// See ref: https://github.com/lonnng/nano/blob/master/docs/communication_protocol.md
func Encode(m *Message, routes map[string]uint16) ([]byte, error) {
	if invalidType(m.Type) {
		return nil, ErrWrongMessageType
	}

	buf := make([]byte, 0)
	flag := byte(m.Type)

	code, compressed := routes[m.Route]
	if !compressed {
		flag |= msgRouteCompressMask
	}
	buf = append(buf, flag)

	n := m.ID
	// variant length encode
	for {
		b := byte(n % 128)
		n >>= 7
		if n != 0 {
			buf = append(buf, b+128)
		} else {
			buf = append(buf, b)
			break
		}
	}

	if routable(m.Type) {
		if compressed {
			buf = append(buf, byte((code>>8)&0xFF))
			buf = append(buf, byte(code&0xFF))
		} else {
			buf = append(buf, byte(len(m.Route)))
			buf = append(buf, []byte(m.Route)...)
		}
	}

	buf = append(buf, m.Data...)
	return buf, nil
}

// Decode unmarshal the bytes slice to a message
// See ref: https://github.com/lonnng/nano/blob/master/docs/communication_protocol.md
func Decode(data []byte, codes map[uint16]string) (*Message, error) {
	if len(data) < msgHeadLength {
		return nil, ErrInvalidMessage
	}
	m := New()
	flag := data[0]
	offset := 1
	m.Type = Type(flag & msgTypeMask)

	if invalidType(m.Type) {
		return nil, ErrWrongMessageType
	}

	id := uint64(0)
	// little end byte order
	// WARNING: must can be stored in 64 bits integer
	// variant length encode
	for i := offset; i < len(data); i++ {
		b := data[i]
		id += uint64(b&0x7F) << uint64(7*(i-offset))
		if b < 128 {
			offset = i + 1
			break
		}
	}
	m.ID = id

	if routable(m.Type) {
		if flag&msgRouteCompressMask == 0 {
			m.compressed = true
			code := binary.BigEndian.Uint16(data[offset:(offset + 2)])
			route, ok := codes[code]
			if !ok {
				return nil, ErrRouteInfoNotFound
			}
			m.Route = route
			offset += 2
		} else {
			m.compressed = false
			rl := data[offset]
			offset++
			m.Route = string(data[offset:(offset + int(rl))])
			offset += int(rl)
		}
	}

	m.Data = data[offset:]
	return m, nil
}

// TransformDictionary transfrom user defined dict into routes and codes
func TransformDictionary(dict map[string]uint16) (map[string]uint16, map[uint16]string) {
	routes := make(map[string]uint16)
	codes := make(map[uint16]string)

	for route, code := range dict {
		r := strings.TrimSpace(route)

		// duplication check
		if _, ok := routes[r]; ok {
			log.Println(fmt.Sprintf("duplicated route(route: %s, code: %d)", r, code))
		}

		if _, ok := codes[code]; ok {
			log.Println(fmt.Sprintf("duplicated route(route: %s, code: %d)", r, code))
		}

		// update map, using last value when key duplicated
		routes[r] = code
		codes[code] = r
	}

	return routes, codes
}

// GetDictionary generates routes and codes from env.RouteDict()
// Warning: env.RouteDict must be thread-safe func
func GetDictionary() (map[string]uint16, map[uint16]string) {
	if env.RouteDict == nil {
		return Routes, Codes
	}
	dict := env.RouteDict()
	return TransformDictionary(dict)
}
