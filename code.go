package main

import "fmt"

type Code struct {
	decryptionMap [26]byte
	encryptionMap [26]byte
}

func NewCode() *Code {
	return new(Code)
}

// Return the newly added mappings, and if successful or not
func (this *Code) TryUpdateWithDecryptedWord(cyphertext, plaintext []byte) ([][2]byte, bool) {
	keysToAdd := make([][2]byte, 0, len(cyphertext))

	for i := 0; i < len(cyphertext); i++ {
		if decrypted, ok := this.decryptByte(cyphertext[i]); ok {
			if decrypted != plaintext[i] {
				return [][2]byte{ }, false
			}
			continue // key already exists in mapping, skip
		} else if encrypted, ok := this.encryptByte(plaintext[i]); ok {
			if encrypted != cyphertext[i] {
				return [][2]byte{ }, false
			}
		}

		keysToAdd = append(keysToAdd, [2]byte{ cyphertext[i], plaintext[i] })
	}

	for _, key := range keysToAdd {
		this.addKey(key)
	}

	return keysToAdd, true
}

func (this *Code) addKey(key [2]byte) {
	this.decryptionMap[byteToMapIndex(key[0])] = key[1]
	this.encryptionMap[byteToMapIndex(key[1])] = key[0]
}

func (this *Code) removeKey(key [2]byte) {
	this.decryptionMap[byteToMapIndex(key[0])] = 0
	this.encryptionMap[byteToMapIndex(key[1])] = 0
}

func (this *Code) RemoveKeys(keys [][2]byte) {
	for _, key := range keys {
		this.removeKey(key)
	}
}

func (this *Code) decryptByte(b byte) (byte, bool) {
	decrypted := this.decryptionMap[byteToMapIndex(b)]
	return decrypted, decrypted != 0
}

func (this *Code) encryptByte(b byte) (byte, bool) {
	encrypted := this.encryptionMap[byteToMapIndex(b)]
	return encrypted, encrypted != 0
}

func byteToMapIndex(b byte) byte {
	result := b - 'A'

	if result < 0 || result >= 26 {
		panic(fmt.Sprintf("byteToMapIndex(%c)=%d", b, result))
	}

	return result
}