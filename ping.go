package main

func CheckSum(buf []byte) uint16 {
	var sum uint32
	for i := 0; i < len(buf)-1; i += 2 {
		sum += uint32(buf[i+1])<<8 | uint32(buf[i])
	}

	// Take care of left over byte
	if len(buf)%2 != 0 {
		sum += uint32(buf[len(buf)-1])
	}

	sum = (sum >> 16) + (sum & 0xffff)
	sum = sum + (sum >> 16)
	return uint16(^sum)
}

func MakeEchoRequest(seqNo int, length int, id1 byte, id2 byte) [512]byte {
	var msg [512]byte
	// ICMP v4 echo request is 8
	msg[0] = 8
	// Code is always 0
	msg[1] = 0
	// Zero checksum (will be overwritten shortly)
	msg[2] = 0
	msg[3] = 0
	// Id
	msg[4] = id1
	msg[5] = id2
	// Not sure about this one..
	msg[6] = 0
	msg[7] = byte(seqNo)
	// Generate the checksum and insert into message
	check := CheckSum(msg[0:length])
	msg[2] = byte(check & 0xff)
	msg[3] = byte(check >> 8)
	return msg
}
