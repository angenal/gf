package utils

// ToLowerBytes is the equivalent of bytes.ToLower
func ToLowerBytes(b []byte) []byte {
	for i := 0; i < len(b); i++ {
		b[i] = toLowerTable[b[i]]
	}
	return b
}

// ToUpperBytes is the equivalent of bytes.ToUpper
func ToUpperBytes(b []byte) []byte {
	for i := 0; i < len(b); i++ {
		b[i] = toUpperTable[b[i]]
	}
	return b
}

// TrimRightBytes is the equivalent of bytes.TrimRight
func TrimRightBytes(b []byte, cut byte) []byte {
	lenStr := len(b)
	for lenStr > 0 && b[lenStr-1] == cut {
		lenStr--
	}
	return b[:lenStr]
}

// TrimLeftBytes is the equivalent of bytes.TrimLeft
func TrimLeftBytes(b []byte, cut byte) []byte {
	lenStr, start := len(b), 0
	for start < lenStr && b[start] == cut {
		start++
	}
	return b[start:]
}

// TrimBytes is the equivalent of bytes.Trim
func TrimBytes(b []byte, cut byte) []byte {
	i, j := 0, len(b)-1
	for ; i <= j; i++ {
		if b[i] != cut {
			break
		}
	}
	for ; i < j; j-- {
		if b[j] != cut {
			break
		}
	}

	return b[i : j+1]
}

// EqualFoldBytes the equivalent of bytes.EqualFold
func EqualFoldBytes(b, s []byte) bool {
	n := len(b)
	ok := n == len(s)
	if ok {
		for i := 0; i < n; i++ {
			if ok = b[i]|0x20 == s[i]|0x20; !ok {
				break
			}
		}
	}
	return ok
}
