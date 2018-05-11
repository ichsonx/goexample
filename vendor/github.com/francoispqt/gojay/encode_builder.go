package gojay

// grow grows b's capacity, if necessary, to guarantee space for
// another n bytes. After grow(n), at least n bytes can be written to b
// without another allocation. If n is negative, grow panics.
func (enc *Encoder) grow(n int) {
	if cap(enc.buf)-len(enc.buf) < n {
		Buf := make([]byte, len(enc.buf), 2*cap(enc.buf)+n)
		copy(Buf, enc.buf)
		enc.buf = Buf
	}
}

// Write appends the contents of p to b's Buffer.
// Write always returns len(p), nil.
func (enc *Encoder) writeBytes(p []byte) {
	enc.buf = append(enc.buf, p...)
}

// WriteByte appends the byte c to b's Buffer.
// The returned error is always nil.
func (enc *Encoder) writeByte(c byte) {
	enc.buf = append(enc.buf, c)
}

// WriteString appends the contents of s to b's Buffer.
// It returns the length of s and a nil error.
func (enc *Encoder) writeString(s string) {
	enc.buf = append(enc.buf, s...)
}

func (enc *Encoder) writeStringEscape(s string) {
	l := len(s)
	for i := 0; i < l; i++ {
		switch s[i] {
		case '\\', '"':
			enc.writeByte('\\')
			enc.writeByte(s[i])
		case '\n':
			enc.writeByte('\\')
			enc.writeByte('n')
		case '\r':
			enc.writeByte('\\')
			enc.writeByte('r')
		case '\t':
			enc.writeByte('\\')
			enc.writeByte('t')
		default:
			enc.writeByte(s[i])
		}
	}

}
