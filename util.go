package main

func textToLong(ip string) uint64 {
	bytes := textToNumericFormatV4(ip)
	var address uint64
	for _, b := range bytes {
		address = (address << 8) | uint64(b&0xFF)
	}
	return address
}

func textToNumericFormatV4(ip string) []byte {
	var res [4]byte
	tmpValue := uint64(0)
	currByte := 0
	newOctet := true

	for i := 0; i < len(ip); i++ {
		c := ip[i]
		if c == '.' {
			if newOctet || tmpValue > 0xff || currByte == 3 {
				return nil
			}
			res[currByte] = byte(tmpValue & 0xff)
			currByte++
			tmpValue = 0
			newOctet = true
		} else {
			digit := digit(c, 10)
			if digit < 0 {
				return nil
			}
			tmpValue = tmpValue*10 + uint64(digit)
			newOctet = false
		}
	}
	if newOctet || tmpValue >= 1<<((4-uint64(currByte))*8) {
		return nil
	}

	switch currByte {
	case 0:
		res[0] = byte((tmpValue >> 24) & 0xff)
		fallthrough
	case 1:
		res[1] = byte((tmpValue >> 16) & 0xff)
		fallthrough
	case 2:
		res[2] = byte((tmpValue >> 8) & 0xff)
		fallthrough
	case 3:
		res[3] = byte(tmpValue & 0xff)
	}
	return res[:]
}

func digit(c byte, radix int) int {
	val := int(c - '0')
	if val < 0 || val >= radix {
		return -1
	}
	return val
}
