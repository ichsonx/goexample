package gojay

// EncodeArray encodes an implementation of MarshalerArray to JSON
func (enc *Encoder) EncodeArray(v MarshalerArray) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, _ = enc.encodeArray(v)
	_, err := enc.Write()
	if err != nil {
		enc.err = err
		return err
	}
	return nil
}
func (enc *Encoder) encodeArray(v MarshalerArray) ([]byte, error) {
	enc.grow(200)
	enc.writeByte('[')
	v.MarshalArray(enc)
	enc.writeByte(']')
	return enc.buf, enc.err
}

// AddArray adds an implementation of MarshalerArray to be encoded, must be used inside a slice or array encoding (does not encode a key)
// value must implement Marshaler
func (enc *Encoder) AddArray(v MarshalerArray) {
	if v.IsNil() {
		enc.grow(3)
		r := enc.getPreviousRune()
		if r != '[' {
			enc.writeByte(',')
		}
		enc.writeByte('[')
		enc.writeByte(']')
		return
	}
	enc.grow(100)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('[')
	v.MarshalArray(enc)
	enc.writeByte(']')
}

// AddArrayOmitEmpty adds an array or slice to be encoded, must be used inside a slice or array encoding (does not encode a key)
// value must implement Marshaler
func (enc *Encoder) AddArrayOmitEmpty(v MarshalerArray) {
	if v.IsNil() {
		return
	}
	enc.grow(4)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('[')
	v.MarshalArray(enc)
	enc.writeByte(']')
}

// AddArrayKey adds an array or slice to be encoded, must be used inside an object as it will encode a key
// value must implement Marshaler
func (enc *Encoder) AddArrayKey(key string, v MarshalerArray) {
	if v.IsNil() {
		enc.grow(2 + len(key))
		r := enc.getPreviousRune()
		if r != '{' {
			enc.writeByte(',')
		}
		enc.writeByte('"')
		enc.writeStringEscape(key)
		enc.writeBytes(objKeyArr)
		enc.writeByte(']')
		return
	}
	enc.grow(5 + len(key))
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKeyArr)
	v.MarshalArray(enc)
	enc.writeByte(']')
}

// AddArrayKeyOmitEmpty adds an array or slice to be encoded and skips it if it is nil.
// Must be called inside an object as it will encode a key.
func (enc *Encoder) AddArrayKeyOmitEmpty(key string, v MarshalerArray) {
	if v.IsNil() {
		return
	}
	enc.grow(5 + len(key))
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKeyArr)
	v.MarshalArray(enc)
	enc.writeByte(']')
}
