package sha

import (
	"fmt"
)

func Hash(message []byte) {
	// pre-process
	// append a 1 at the end of the message
	paddedMessage := append(message, 0x80)
	fmt.Printf("%s", bits(paddedMessage))

	numBits := len(paddedMessage) * 8

	// pad until number of bits is congruent 448 mod 512
	for numBits%512 != 448 {
		paddedMessage = append(paddedMessage, 0x00)
		numBits = len(paddedMessage) * 8
	}
	fmt.Printf("%s\n", bits(paddedMessage))
	fmt.Printf("num bits: %d\n", numBits)

	// get message length and append to the padded message
	messageLen := len(message) * 8
	paddedMsgLen := int64(messageLen)
	fmt.Printf("%064b", paddedMsgLen)

	// break message into 512 bit chuncks

}

func bits(b []byte) string {
	bitString := ""
	for _, b := range b {
		bitString = fmt.Sprintf("%s%08b ", bitString, b)
	}
	bitString = fmt.Sprintf("%s\n", bitString)
	return bitString
}
