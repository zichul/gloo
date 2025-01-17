// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/grpc/validation/gloo_validation.proto

package validation

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
	"github.com/mitchellh/hashstructure"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Hash function
func (m *GlooValidationServiceRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	if h, ok := interface{}(m.GetProxy()).(interface {
		Hash(hasher hash.Hash64) (uint64, error)
	}); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetProxy(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GlooValidationServiceResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	if h, ok := interface{}(m.GetProxyReport()).(interface {
		Hash(hasher hash.Hash64) (uint64, error)
	}); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetProxyReport(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *NotifyOnResyncRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *NotifyOnResyncResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ProxyReport) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	for _, v := range m.GetListenerReports() {

		if h, ok := interface{}(v).(interface {
			Hash(hasher hash.Hash64) (uint64, error)
		}); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListenerReport) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	for _, v := range m.GetErrors() {

		if h, ok := interface{}(v).(interface {
			Hash(hasher hash.Hash64) (uint64, error)
		}); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	switch m.ListenerTypeReport.(type) {

	case *ListenerReport_HttpListenerReport:

		if h, ok := interface{}(m.GetHttpListenerReport()).(interface {
			Hash(hasher hash.Hash64) (uint64, error)
		}); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetHttpListenerReport(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	case *ListenerReport_TcpListenerReport:

		if h, ok := interface{}(m.GetTcpListenerReport()).(interface {
			Hash(hasher hash.Hash64) (uint64, error)
		}); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(m.GetTcpListenerReport(), nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HttpListenerReport) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	for _, v := range m.GetErrors() {

		if h, ok := interface{}(v).(interface {
			Hash(hasher hash.Hash64) (uint64, error)
		}); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetVirtualHostReports() {

		if h, ok := interface{}(v).(interface {
			Hash(hasher hash.Hash64) (uint64, error)
		}); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *VirtualHostReport) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	for _, v := range m.GetErrors() {

		if h, ok := interface{}(v).(interface {
			Hash(hasher hash.Hash64) (uint64, error)
		}); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetRouteReports() {

		if h, ok := interface{}(v).(interface {
			Hash(hasher hash.Hash64) (uint64, error)
		}); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RouteReport) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	for _, v := range m.GetErrors() {

		if h, ok := interface{}(v).(interface {
			Hash(hasher hash.Hash64) (uint64, error)
		}); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetWarnings() {

		if h, ok := interface{}(v).(interface {
			Hash(hasher hash.Hash64) (uint64, error)
		}); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *TcpListenerReport) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	for _, v := range m.GetErrors() {

		if h, ok := interface{}(v).(interface {
			Hash(hasher hash.Hash64) (uint64, error)
		}); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetTcpHostReports() {

		if h, ok := interface{}(v).(interface {
			Hash(hasher hash.Hash64) (uint64, error)
		}); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *TcpHostReport) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	for _, v := range m.GetErrors() {

		if h, ok := interface{}(v).(interface {
			Hash(hasher hash.Hash64) (uint64, error)
		}); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListenerReport_Error) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	err = binary.Write(hasher, binary.LittleEndian, m.GetType())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetReason())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HttpListenerReport_Error) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	err = binary.Write(hasher, binary.LittleEndian, m.GetType())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetReason())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *VirtualHostReport_Error) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	err = binary.Write(hasher, binary.LittleEndian, m.GetType())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetReason())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RouteReport_Error) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	err = binary.Write(hasher, binary.LittleEndian, m.GetType())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetReason())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RouteReport_Warning) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	err = binary.Write(hasher, binary.LittleEndian, m.GetType())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetReason())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *TcpListenerReport_Error) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	err = binary.Write(hasher, binary.LittleEndian, m.GetType())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetReason())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *TcpHostReport_Error) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error

	err = binary.Write(hasher, binary.LittleEndian, m.GetType())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetReason())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
