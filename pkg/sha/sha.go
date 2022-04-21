package sha

import (
	"encoding/binary"
	"fmt"
)

type word struct {
	b [4]byte
}

type chunk struct {
	w [80]word
}

func Digest(message []byte) {
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
	lenBytes := make([]byte, 8)
	messageLen := len(message) * 8

	binary.BigEndian.PutUint64(lenBytes, uint64(messageLen))

	paddedMessage = append(paddedMessage, lenBytes...)
	fmt.Printf("%s\n", bits(paddedMessage))
	fmt.Printf("num bits: %d\n", len(paddedMessage)*8)
	//fmt.Printf("%064b\n", messageLen)
	//fmt.Printf("%s\n", bits(lenBytes))

	// break message into 512 bit chuncks
	bitsInChunk := 512
	bitsInByte := 8
	bytesInChunk := bitsInChunk / bitsInByte

	numChunks := (len(paddedMessage) * bitsInByte) / bitsInChunk
	fmt.Printf("num chunks: %v\n", numChunks)

	chunks := make([]chunk, numChunks)
	for i := 0; i < numChunks; i++ {
		chunkSlice := paddedMessage[i*bytesInChunk : (i+1)*bytesInChunk]
		//fmt.Printf("slice %v:\n%v\n", i, bits(chunkSlice))

		var nextChunk chunk
		initialWords := 16
		wordsInChunk := len(nextChunk.w)
		fmt.Printf("words in chunk: %v\n", wordsInChunk)

		// break 512 bit chuncks into 32 bit 'words'
		for j := 0; j < initialWords; j++ {
			var nextWord word
			var wordbytes [4]byte
			bytesInWord := len(nextWord.b)

			copy(wordbytes[:], chunkSlice[j*bytesInWord:(j+1)*bytesInWord])

			nextWord = word{wordbytes}

			// add word to chunk
			nextChunk.w[j] = nextWord
			//fmt.Printf(bits(nextChunk.w[j].b[:]))
		}
		chunks = append(chunks, nextChunk)
	}

	// generate more words until there are 80 words

}

func bits(b []byte) string {
	bitString := ""
	for _, b := range b {
		bitString = fmt.Sprintf("%s%08b ", bitString, b)
	}
	bitString = fmt.Sprintf("%s\n", bitString)
	return bitString
}
