package sha

import (
	"encoding/binary"
	"fmt"
	"math/bits"
)

type word struct {
	b [4]byte
}

type chunk struct {
	w [80]word
}

func Digest(message []byte) []byte {
	// *** pre-process ***
	// append a 1 at the end of the message
	paddedMessage := append(message, 0x80)
	//fmt.Printf("%s", byteString(paddedMessage))

	numBits := len(paddedMessage) * 8

	// pad until number of bits is congruent 448 mod 512
	for numBits%512 != 448 {
		paddedMessage = append(paddedMessage, 0x00)
		numBits = len(paddedMessage) * 8
	}
	//fmt.Printf("%s\n", byteString(paddedMessage))
	//fmt.Printf("num bits: %d\n", numBits)

	// get message length and append to the padded message
	lenBytes := make([]byte, 8)
	messageLen := len(message) * 8

	binary.BigEndian.PutUint64(lenBytes, uint64(messageLen))

	paddedMessage = append(paddedMessage, lenBytes...)
	//fmt.Printf("%s\n", byteString(paddedMessage))
	//fmt.Printf("num bits: %d\n", len(paddedMessage)*8)
	//fmt.Printf("%064b\n", messageLen)
	//fmt.Printf("%s\n", bits(lenBytes))

	// break message into 512 bit chuncks
	bitsInChunk := 512
	bitsInByte := 8
	bytesInChunk := bitsInChunk / bitsInByte

	numChunks := (len(paddedMessage) * bitsInByte) / bitsInChunk
	//fmt.Printf("num chunks: %v\n", numChunks)

	h0 := uint32(0x67452301)
	h1 := uint32(0xEFCDAB89)
	h2 := uint32(0x98BADCFE)
	h3 := uint32(0x10325476)
	h4 := uint32(0xC3D2E1F0)

	//chunks := make([]chunk, numChunks)
	for i := 0; i < numChunks; i++ {
		chunkSlice := paddedMessage[i*bytesInChunk : (i+1)*bytesInChunk]
		//fmt.Printf("slice %v:\n%v\n", i, bits(chunkSlice))

		var nextChunk chunk
		initialWords := 16
		wordsInChunk := len(nextChunk.w)
		//fmt.Printf("words in chunk: %v\n", wordsInChunk)

		// break 512 bit chuncks into 32 bit 'words'
		for j := 0; j < initialWords; j++ {
			var nextWord word
			var wordbytes [4]byte
			bytesInWord := len(nextWord.b)

			copy(wordbytes[:], chunkSlice[j*bytesInWord:(j+1)*bytesInWord])

			nextWord = word{wordbytes}

			// add word to chunk
			nextChunk.w[j] = nextWord
			//fmt.Printf("%v: %v\n", j, byteString(nextChunk.w[j].b[:]))
		}

		// generate more words until there are 80 words
		for k := initialWords; k < wordsInChunk; k++ {
			word1 := binary.BigEndian.Uint32(nextChunk.w[k-3].b[:])
			word2 := binary.BigEndian.Uint32(nextChunk.w[k-8].b[:])
			word3 := binary.BigEndian.Uint32(nextChunk.w[k-14].b[:])
			word4 := binary.BigEndian.Uint32(nextChunk.w[k-16].b[:])

			wordXOR := word1 ^ word2 ^ word3 ^ word4
			wordRotate := bits.RotateLeft32(wordXOR, 1)

			var newWordBytes [4]byte
			binary.BigEndian.PutUint32(newWordBytes[:], wordRotate)

			newWord := word{newWordBytes}
			nextChunk.w[k] = newWord
			//fmt.Printf("%v: %v\n", k, byteString(nextChunk.w[k].b[:]))
		}

		// put new chunk in list
		//chunks = append(chunks, nextChunk)

		// *** end pre-process ***

		// *** create the digest ***

		a := h0
		b := h1
		c := h2
		d := h3
		e := h4

		for l := 0; l < wordsInChunk; l++ {
			var f uint32
			var k uint32

			if l < 20 {
				f = (b & c) | (not(b) & d)
				k = 0x5A827999
			} else if l < 40 {
				f = b ^ c ^ d
				k = 0x6ED9EBA1
			} else if l < 60 {
				f = (b & c) | (b & d) | (c & d)
				k = 0x8F1BBCDC
			} else if l < 80 {
				f = b ^ c ^ d
				k = 0xCA62C1D6
			}

			curWord := binary.BigEndian.Uint32(nextChunk.w[l].b[:])

			aRot := bits.RotateLeft32(a, 5)
			sumF, _ := bits.Add32(aRot, f, 0)
			sumE, _ := bits.Add32(sumF, e, 0)
			sumK, _ := bits.Add32(sumE, k, 0)
			temp, _ := bits.Add32(sumK, curWord, 0)

			e = d
			d = c
			c = bits.RotateLeft32(b, 30)
			b = a
			a = temp

		}

		h0 = h0 + a
		h1 = h1 + b
		h2 = h2 + c
		h3 = h3 + d
		h4 = h4 + e

	}

	var digest []byte
	var h0Bytes [4]byte
	var h1Bytes [4]byte
	var h2Bytes [4]byte
	var h3Bytes [4]byte
	var h4Bytes [4]byte

	binary.BigEndian.PutUint32(h0Bytes[:], h0)
	binary.BigEndian.PutUint32(h1Bytes[:], h1)
	binary.BigEndian.PutUint32(h2Bytes[:], h2)
	binary.BigEndian.PutUint32(h3Bytes[:], h3)
	binary.BigEndian.PutUint32(h4Bytes[:], h4)

	digest = append(digest, h0Bytes[:]...)
	digest = append(digest, h1Bytes[:]...)
	digest = append(digest, h2Bytes[:]...)
	digest = append(digest, h3Bytes[:]...)
	digest = append(digest, h4Bytes[:]...)

	return digest
}

func not(x uint32) uint32 {
	mask := uint32(0xFFFFFFFF)
	return x ^ mask
}

func byteString(b []byte) string {
	bitString := ""
	for _, b := range b {
		bitString = fmt.Sprintf("%s%08b ", bitString, b)
	}
	bitString = fmt.Sprintf("%s\n", bitString)
	return bitString
}
