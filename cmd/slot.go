package cmd

// Redis cluster slot [https://redis.io/topics/cluster-spec].
type cslot uint16

const (
	cslotInit = cslot(1 << 14) // can be sent to any redis node in cluster
	cslotNo   = cslot(1 << 15) // has no key slot specified
)

func (s cslot) GetSlot() uint16 {
	return uint16(s)
}

// set new slot value.
// Causes Panic if the key was not InitSlot.
func (s *cslot) set(new cslot) {
	switch {
	case *s == new:
		break
	case *s == cslotInit:
		*s = new
	default:
		panic(multiKeySlotErr)
	}
}

const multiKeySlotErr = "multi key command with different key slots are not allowed"

func getSlot(key string) cslot {
	var s, e int
	for ; s < len(key); s++ {
		if key[s] == '{' {
			break
		}
	}
	if s == len(key) {
		return cslot(crc16(key) & 16383)
	}
	for e = s + 1; e < len(key); e++ {
		if key[e] == '}' {
			break
		}
	}
	if e == len(key) || e == s+1 {
		return cslot(crc16(key) & 16383)
	}
	return cslot(crc16(key[s+1:e]) & 16383)
}
