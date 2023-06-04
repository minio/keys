package msgp

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *Commit) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "0"
	o = append(o, 0x83, 0xa1, 0x30)
	o = msgp.AppendUint64(o, z.N)
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendUint(o, z.Type)
	// string "2"
	o = append(o, 0xa1, 0x32)
	o = msgp.AppendBytes(o, z.Event)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Commit) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			z.N, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "N")
				return
			}
		case "1":
			z.Type, bts, err = msgp.ReadUintBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Type")
				return
			}
		case "2":
			z.Event, bts, err = msgp.ReadBytesBytes(bts, z.Event)
			if err != nil {
				err = msgp.WrapError(err, "Event")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Commit) Msgsize() (s int) {
	s = 1 + 2 + msgp.Uint64Size + 2 + msgp.UintSize + 2 + msgp.BytesPrefixSize + len(z.Event)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CreateEnclaveCmd) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "0"
	o = append(o, 0x84, 0xa1, 0x30)
	o = msgp.AppendString(o, z.Name)
	// string "1"
	o = append(o, 0xa1, 0x31)
	o, err = z.Key.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Key")
		return
	}
	// string "2"
	o = append(o, 0xa1, 0x32)
	o = msgp.AppendTime(o, z.CreatedAt)
	// string "3"
	o = append(o, 0xa1, 0x33)
	o = msgp.AppendString(o, z.CreatedBy)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CreateEnclaveCmd) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "1":
			bts, err = z.Key.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Key")
				return
			}
		case "2":
			z.CreatedAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedAt")
				return
			}
		case "3":
			z.CreatedBy, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedBy")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CreateEnclaveCmd) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.Name) + 2 + z.Key.Msgsize() + 2 + msgp.TimeSize + 2 + msgp.StringPrefixSize + len(z.CreatedBy)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CreateIdentityCmd) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "0"
	o = append(o, 0x88, 0xa1, 0x30)
	o = msgp.AppendString(o, z.Enclave)
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendString(o, z.Identity)
	// string "2"
	o = append(o, 0xa1, 0x32)
	o = msgp.AppendString(o, z.Policy)
	// string "3"
	o = append(o, 0xa1, 0x33)
	o = msgp.AppendBool(o, z.IsAdmin)
	// string "4"
	o = append(o, 0xa1, 0x34)
	o = msgp.AppendDuration(o, z.TTL)
	// string "5"
	o = append(o, 0xa1, 0x35)
	o = msgp.AppendTime(o, z.ExpiresAt)
	// string "6"
	o = append(o, 0xa1, 0x36)
	o = msgp.AppendTime(o, z.CreatedAt)
	// string "7"
	o = append(o, 0xa1, 0x37)
	o = msgp.AppendString(o, z.CreatedBy)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CreateIdentityCmd) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			z.Enclave, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Enclave")
				return
			}
		case "1":
			z.Identity, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Identity")
				return
			}
		case "2":
			z.Policy, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Policy")
				return
			}
		case "3":
			z.IsAdmin, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IsAdmin")
				return
			}
		case "4":
			z.TTL, bts, err = msgp.ReadDurationBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TTL")
				return
			}
		case "5":
			z.ExpiresAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ExpiresAt")
				return
			}
		case "6":
			z.CreatedAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedAt")
				return
			}
		case "7":
			z.CreatedBy, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedBy")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CreateIdentityCmd) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.Enclave) + 2 + msgp.StringPrefixSize + len(z.Identity) + 2 + msgp.StringPrefixSize + len(z.Policy) + 2 + msgp.BoolSize + 2 + msgp.DurationSize + 2 + msgp.TimeSize + 2 + msgp.TimeSize + 2 + msgp.StringPrefixSize + len(z.CreatedBy)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CreatePolicyCmd) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "0"
	o = append(o, 0x86, 0xa1, 0x30)
	o = msgp.AppendString(o, z.Enclave)
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendString(o, z.Name)
	// string "2"
	o = append(o, 0xa1, 0x32)
	o = msgp.AppendMapHeader(o, uint32(len(z.Allow)))
	for za0001 := range z.Allow {
		o = msgp.AppendString(o, za0001)
		// map header, size 0
		o = append(o, 0x80)
	}
	// string "3"
	o = append(o, 0xa1, 0x33)
	o = msgp.AppendMapHeader(o, uint32(len(z.Deny)))
	for za0003 := range z.Deny {
		o = msgp.AppendString(o, za0003)
		// map header, size 0
		o = append(o, 0x80)
	}
	// string "4"
	o = append(o, 0xa1, 0x34)
	o = msgp.AppendTime(o, z.CreatedAt)
	// string "5"
	o = append(o, 0xa1, 0x35)
	o = msgp.AppendString(o, z.CreatedBy)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CreatePolicyCmd) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			z.Enclave, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Enclave")
				return
			}
		case "1":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "2":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Allow")
				return
			}
			if z.Allow == nil {
				z.Allow = make(map[string]struct{}, zb0002)
			} else if len(z.Allow) > 0 {
				for key := range z.Allow {
					delete(z.Allow, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 struct{}
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Allow")
					return
				}
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Allow", za0001)
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "Allow", za0001)
						return
					}
					switch msgp.UnsafeString(field) {
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							err = msgp.WrapError(err, "Allow", za0001)
							return
						}
					}
				}
				z.Allow[za0001] = za0002
			}
		case "3":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Deny")
				return
			}
			if z.Deny == nil {
				z.Deny = make(map[string]struct{}, zb0004)
			} else if len(z.Deny) > 0 {
				for key := range z.Deny {
					delete(z.Deny, key)
				}
			}
			for zb0004 > 0 {
				var za0003 string
				var za0004 struct{}
				zb0004--
				za0003, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Deny")
					return
				}
				var zb0005 uint32
				zb0005, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Deny", za0003)
					return
				}
				for zb0005 > 0 {
					zb0005--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "Deny", za0003)
						return
					}
					switch msgp.UnsafeString(field) {
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							err = msgp.WrapError(err, "Deny", za0003)
							return
						}
					}
				}
				z.Deny[za0003] = za0004
			}
		case "4":
			z.CreatedAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedAt")
				return
			}
		case "5":
			z.CreatedBy, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedBy")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CreatePolicyCmd) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.Enclave) + 2 + msgp.StringPrefixSize + len(z.Name) + 2 + msgp.MapHeaderSize
	if z.Allow != nil {
		for za0001, za0002 := range z.Allow {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + 1
		}
	}
	s += 2 + msgp.MapHeaderSize
	if z.Deny != nil {
		for za0003, za0004 := range z.Deny {
			_ = za0004
			s += msgp.StringPrefixSize + len(za0003) + 1
		}
	}
	s += 2 + msgp.TimeSize + 2 + msgp.StringPrefixSize + len(z.CreatedBy)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CreateSecretCmd) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "0"
	o = append(o, 0x86, 0xa1, 0x30)
	o = msgp.AppendString(o, z.Enclave)
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendString(o, z.Name)
	// string "2"
	o = append(o, 0xa1, 0x32)
	o = msgp.AppendBytes(o, z.Secret)
	// string "3"
	o = append(o, 0xa1, 0x33)
	o = msgp.AppendUint(o, z.SecretType)
	// string "4"
	o = append(o, 0xa1, 0x34)
	o = msgp.AppendTime(o, z.CreatedAt)
	// string "5"
	o = append(o, 0xa1, 0x35)
	o = msgp.AppendString(o, z.CreatedBy)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CreateSecretCmd) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			z.Enclave, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Enclave")
				return
			}
		case "1":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "2":
			z.Secret, bts, err = msgp.ReadBytesBytes(bts, z.Secret)
			if err != nil {
				err = msgp.WrapError(err, "Secret")
				return
			}
		case "3":
			z.SecretType, bts, err = msgp.ReadUintBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SecretType")
				return
			}
		case "4":
			z.CreatedAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedAt")
				return
			}
		case "5":
			z.CreatedBy, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedBy")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CreateSecretCmd) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.Enclave) + 2 + msgp.StringPrefixSize + len(z.Name) + 2 + msgp.BytesPrefixSize + len(z.Secret) + 2 + msgp.UintSize + 2 + msgp.TimeSize + 2 + msgp.StringPrefixSize + len(z.CreatedBy)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CreateSecretKeyRingCmd) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "0"
	o = append(o, 0x85, 0xa1, 0x30)
	o = msgp.AppendString(o, z.Enclave)
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendString(o, z.Name)
	// string "2"
	o = append(o, 0xa1, 0x32)
	o, err = z.Key.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Key")
		return
	}
	// string "3"
	o = append(o, 0xa1, 0x33)
	o = msgp.AppendTime(o, z.CreatedAt)
	// string "4"
	o = append(o, 0xa1, 0x34)
	o = msgp.AppendString(o, z.CreatedBy)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CreateSecretKeyRingCmd) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			z.Enclave, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Enclave")
				return
			}
		case "1":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "2":
			bts, err = z.Key.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Key")
				return
			}
		case "3":
			z.CreatedAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedAt")
				return
			}
		case "4":
			z.CreatedBy, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedBy")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CreateSecretKeyRingCmd) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.Enclave) + 2 + msgp.StringPrefixSize + len(z.Name) + 2 + z.Key.Msgsize() + 2 + msgp.TimeSize + 2 + msgp.StringPrefixSize + len(z.CreatedBy)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DeleteEnclaveCmd) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "0"
	o = append(o, 0x81, 0xa1, 0x30)
	o = msgp.AppendString(o, z.Name)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DeleteEnclaveCmd) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DeleteEnclaveCmd) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.Name)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DeleteIdentityCmd) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "0"
	o = append(o, 0x82, 0xa1, 0x30)
	o = msgp.AppendString(o, z.Enclave)
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendString(o, z.Identity)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DeleteIdentityCmd) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			z.Enclave, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Enclave")
				return
			}
		case "1":
			z.Identity, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Identity")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DeleteIdentityCmd) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.Enclave) + 2 + msgp.StringPrefixSize + len(z.Identity)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DeletePolicyCmd) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "0"
	o = append(o, 0x82, 0xa1, 0x30)
	o = msgp.AppendString(o, z.Enclave)
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendString(o, z.Name)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DeletePolicyCmd) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			z.Enclave, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Enclave")
				return
			}
		case "1":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DeletePolicyCmd) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.Enclave) + 2 + msgp.StringPrefixSize + len(z.Name)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DeleteSecretCmd) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "0"
	o = append(o, 0x82, 0xa1, 0x30)
	o = msgp.AppendString(o, z.Enclave)
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendString(o, z.Name)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DeleteSecretCmd) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			z.Enclave, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Enclave")
				return
			}
		case "1":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DeleteSecretCmd) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.Enclave) + 2 + msgp.StringPrefixSize + len(z.Name)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DeleteSecretKeyRingCmd) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "0"
	o = append(o, 0x82, 0xa1, 0x30)
	o = msgp.AppendString(o, z.Enclave)
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendString(o, z.Name)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DeleteSecretKeyRingCmd) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			z.Enclave, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Enclave")
				return
			}
		case "1":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DeleteSecretKeyRingCmd) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.Enclave) + 2 + msgp.StringPrefixSize + len(z.Name)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Enclave) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "0"
	o = append(o, 0x84, 0xa1, 0x30)
	o, err = z.Key.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Key")
		return
	}
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Admins)))
	for za0001 := range z.Admins {
		o = msgp.AppendString(o, z.Admins[za0001])
	}
	// string "2"
	o = append(o, 0xa1, 0x32)
	o = msgp.AppendTime(o, z.CreatedAt)
	// string "3"
	o = append(o, 0xa1, 0x33)
	o = msgp.AppendString(o, z.CreatedBy)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Enclave) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			bts, err = z.Key.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Key")
				return
			}
		case "1":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Admins")
				return
			}
			if cap(z.Admins) >= int(zb0002) {
				z.Admins = (z.Admins)[:zb0002]
			} else {
				z.Admins = make([]string, zb0002)
			}
			for za0001 := range z.Admins {
				z.Admins[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Admins", za0001)
					return
				}
			}
		case "2":
			z.CreatedAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedAt")
				return
			}
		case "3":
			z.CreatedBy, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedBy")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Enclave) Msgsize() (s int) {
	s = 1 + 2 + z.Key.Msgsize() + 2 + msgp.ArrayHeaderSize
	for za0001 := range z.Admins {
		s += msgp.StringPrefixSize + len(z.Admins[za0001])
	}
	s += 2 + msgp.TimeSize + 2 + msgp.StringPrefixSize + len(z.CreatedBy)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *EnclaveInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "0"
	o = append(o, 0x83, 0xa1, 0x30)
	o, err = z.Key.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Key")
		return
	}
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendTime(o, z.CreatedAt)
	// string "2"
	o = append(o, 0xa1, 0x32)
	o = msgp.AppendString(o, z.CreatedBy)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *EnclaveInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			bts, err = z.Key.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Key")
				return
			}
		case "1":
			z.CreatedAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedAt")
				return
			}
		case "2":
			z.CreatedBy, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreatedBy")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *EnclaveInfo) Msgsize() (s int) {
	s = 1 + 2 + z.Key.Msgsize() + 2 + msgp.TimeSize + 2 + msgp.StringPrefixSize + len(z.CreatedBy)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *EncryptedRootKey) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "0"
	o = append(o, 0x81, 0xa1, 0x30)
	o = msgp.AppendMapHeader(o, uint32(len(z.Ciphertexts)))
	for za0001, za0002 := range z.Ciphertexts {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendBytes(o, za0002)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *EncryptedRootKey) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Ciphertexts")
				return
			}
			if z.Ciphertexts == nil {
				z.Ciphertexts = make(map[string][]byte, zb0002)
			} else if len(z.Ciphertexts) > 0 {
				for key := range z.Ciphertexts {
					delete(z.Ciphertexts, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 []byte
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Ciphertexts")
					return
				}
				za0002, bts, err = msgp.ReadBytesBytes(bts, za0002)
				if err != nil {
					err = msgp.WrapError(err, "Ciphertexts", za0001)
					return
				}
				z.Ciphertexts[za0001] = za0002
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *EncryptedRootKey) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.Ciphertexts != nil {
		for za0001, za0002 := range z.Ciphertexts {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.BytesPrefixSize + len(za0002)
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *JoinClusterCmd) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "0"
	o = append(o, 0x82, 0xa1, 0x30)
	o = msgp.AppendMapHeader(o, uint32(len(z.Cluster)))
	for za0001, za0002 := range z.Cluster {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendString(o, za0002)
	}
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendString(o, z.Node)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *JoinClusterCmd) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Cluster")
				return
			}
			if z.Cluster == nil {
				z.Cluster = make(map[string]string, zb0002)
			} else if len(z.Cluster) > 0 {
				for key := range z.Cluster {
					delete(z.Cluster, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 string
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Cluster")
					return
				}
				za0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Cluster", za0001)
					return
				}
				z.Cluster[za0001] = za0002
			}
		case "1":
			z.Node, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Node")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *JoinClusterCmd) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.Cluster != nil {
		for za0001, za0002 := range z.Cluster {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	s += 2 + msgp.StringPrefixSize + len(z.Node)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *LeaveClusterCmd) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "0"
	o = append(o, 0x82, 0xa1, 0x30)
	o = msgp.AppendMapHeader(o, uint32(len(z.Cluster)))
	for za0001, za0002 := range z.Cluster {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendString(o, za0002)
	}
	// string "1"
	o = append(o, 0xa1, 0x31)
	o = msgp.AppendString(o, z.Node)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LeaveClusterCmd) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "0":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Cluster")
				return
			}
			if z.Cluster == nil {
				z.Cluster = make(map[string]string, zb0002)
			} else if len(z.Cluster) > 0 {
				for key := range z.Cluster {
					delete(z.Cluster, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 string
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Cluster")
					return
				}
				za0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Cluster", za0001)
					return
				}
				z.Cluster[za0001] = za0002
			}
		case "1":
			z.Node, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Node")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *LeaveClusterCmd) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.Cluster != nil {
		for za0001, za0002 := range z.Cluster {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	s += 2 + msgp.StringPrefixSize + len(z.Node)
	return
}
