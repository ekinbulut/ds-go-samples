package hash

import "strconv"

// Hash returns a hash of the given string.
func Hash(s string) uint32 {
	h := uint32(2166136261)
	for i := 0; i < len(s); i++ {
		h = h*16777619 ^ uint32(s[i])
	}
	return h
}

// HashString returns a hash of the given string.
func HashString(s string) string {
	return strconv.FormatUint(uint64(Hash(s)), 10)
}
