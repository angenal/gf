package dictpool

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/angenal/gf/internal/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Dict) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "D":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "D")
				return
			}
			if cap(z.D) >= int(zb0002) {
				z.D = (z.D)[:zb0002]
			} else {
				z.D = make([]KV, zb0002)
			}
			for za0001 := range z.D {
				var zb0003 uint32
				zb0003, err = dc.ReadMapHeader()
				if err != nil {
					err = msgp.WrapError(err, "D", za0001)
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						err = msgp.WrapError(err, "D", za0001)
						return
					}
					switch msgp.UnsafeString(field) {
					case "Key":
						z.D[za0001].Key, err = dc.ReadString()
						if err != nil {
							err = msgp.WrapError(err, "D", za0001, "Key")
							return
						}
					case "Value":
						z.D[za0001].Value, err = dc.ReadIntf()
						if err != nil {
							err = msgp.WrapError(err, "D", za0001, "Value")
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							err = msgp.WrapError(err, "D", za0001)
							return
						}
					}
				}
			}
		case "BinarySearch":
			z.BinarySearch, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "BinarySearch")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Dict) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "D"
	err = en.Append(0x82, 0xa1, 0x44)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.D)))
	if err != nil {
		err = msgp.WrapError(err, "D")
		return
	}
	for za0001 := range z.D {
		// map header, size 2
		// write "Key"
		err = en.Append(0x82, 0xa3, 0x4b, 0x65, 0x79)
		if err != nil {
			return
		}
		err = en.WriteString(z.D[za0001].Key)
		if err != nil {
			err = msgp.WrapError(err, "D", za0001, "Key")
			return
		}
		// write "Value"
		err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
		if err != nil {
			return
		}
		err = en.WriteIntf(z.D[za0001].Value)
		if err != nil {
			err = msgp.WrapError(err, "D", za0001, "Value")
			return
		}
	}
	// write "BinarySearch"
	err = en.Append(0xac, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68)
	if err != nil {
		return
	}
	err = en.WriteBool(z.BinarySearch)
	if err != nil {
		err = msgp.WrapError(err, "BinarySearch")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Dict) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "D"
	o = append(o, 0x82, 0xa1, 0x44)
	o = msgp.AppendArrayHeader(o, uint32(len(z.D)))
	for za0001 := range z.D {
		// map header, size 2
		// string "Key"
		o = append(o, 0x82, 0xa3, 0x4b, 0x65, 0x79)
		o = msgp.AppendString(o, z.D[za0001].Key)
		// string "Value"
		o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
		o, err = msgp.AppendIntf(o, z.D[za0001].Value)
		if err != nil {
			err = msgp.WrapError(err, "D", za0001, "Value")
			return
		}
	}
	// string "BinarySearch"
	o = append(o, 0xac, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68)
	o = msgp.AppendBool(o, z.BinarySearch)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Dict) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "D":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "D")
				return
			}
			if cap(z.D) >= int(zb0002) {
				z.D = (z.D)[:zb0002]
			} else {
				z.D = make([]KV, zb0002)
			}
			for za0001 := range z.D {
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "D", za0001)
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "D", za0001)
						return
					}
					switch msgp.UnsafeString(field) {
					case "Key":
						z.D[za0001].Key, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "D", za0001, "Key")
							return
						}
					case "Value":
						z.D[za0001].Value, bts, err = msgp.ReadIntfBytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "D", za0001, "Value")
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							err = msgp.WrapError(err, "D", za0001)
							return
						}
					}
				}
			}
		case "BinarySearch":
			z.BinarySearch, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BinarySearch")
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
func (z *Dict) Msgsize() (s int) {
	s = 1 + 2 + msgp.ArrayHeaderSize
	for za0001 := range z.D {
		s += 1 + 4 + msgp.StringPrefixSize + len(z.D[za0001].Key) + 6 + msgp.GuessSize(z.D[za0001].Value)
	}
	s += 13 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *DictMap) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0003 uint32
	zb0003, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if (*z) == nil {
		(*z) = make(DictMap, zb0003)
	} else if len((*z)) > 0 {
		for key := range *z {
			delete((*z), key)
		}
	}
	for zb0003 > 0 {
		zb0003--
		var zb0001 string
		var zb0002 interface{}
		zb0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		zb0002, err = dc.ReadIntf()
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
		(*z)[zb0001] = zb0002
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z DictMap) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(uint32(len(z)))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0004, zb0005 := range z {
		err = en.WriteString(zb0004)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		err = en.WriteIntf(zb0005)
		if err != nil {
			err = msgp.WrapError(err, zb0004)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DictMap) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, uint32(len(z)))
	for zb0004, zb0005 := range z {
		o = msgp.AppendString(o, zb0004)
		o, err = msgp.AppendIntf(o, zb0005)
		if err != nil {
			err = msgp.WrapError(err, zb0004)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DictMap) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if (*z) == nil {
		(*z) = make(DictMap, zb0003)
	} else if len((*z)) > 0 {
		for key := range *z {
			delete((*z), key)
		}
	}
	for zb0003 > 0 {
		var zb0001 string
		var zb0002 interface{}
		zb0003--
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		zb0002, bts, err = msgp.ReadIntfBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
		(*z)[zb0001] = zb0002
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DictMap) Msgsize() (s int) {
	s = msgp.MapHeaderSize
	if z != nil {
		for zb0004, zb0005 := range z {
			_ = zb0005
			s += msgp.StringPrefixSize + len(zb0004) + msgp.GuessSize(zb0005)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *KV) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Key":
			z.Key, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Key")
				return
			}
		case "Value":
			z.Value, err = dc.ReadIntf()
			if err != nil {
				err = msgp.WrapError(err, "Value")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z KV) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Key"
	err = en.Append(0x82, 0xa3, 0x4b, 0x65, 0x79)
	if err != nil {
		return
	}
	err = en.WriteString(z.Key)
	if err != nil {
		err = msgp.WrapError(err, "Key")
		return
	}
	// write "Value"
	err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return
	}
	err = en.WriteIntf(z.Value)
	if err != nil {
		err = msgp.WrapError(err, "Value")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z KV) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Key"
	o = append(o, 0x82, 0xa3, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Key)
	// string "Value"
	o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o, err = msgp.AppendIntf(o, z.Value)
	if err != nil {
		err = msgp.WrapError(err, "Value")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *KV) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Key":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Key")
				return
			}
		case "Value":
			z.Value, bts, err = msgp.ReadIntfBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Value")
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
func (z KV) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Key) + 6 + msgp.GuessSize(z.Value)
	return
}
