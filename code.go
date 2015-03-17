package main

import "fmt"

type Code struct {
	keyStack [][2]byte
	decryptionMap [26]byte
	encryptionMap [26]byte
}

func NewCode() *Code {
	return new(Code)
}

// Return the newly added mappings, and if successful or not
func (this *Code) TryUpdateWithDecryptedWord(cyphertext, plaintext []byte) (int, bool) {
	keysAdded := 0

	for i := 0; i < len(cyphertext); i++ {
		if decrypted, ok := this.decryptByte(cyphertext[i]); ok {
			if decrypted != plaintext[i] {
				this.RemoveLastNKeys(keysAdded)
				return 0, false
			}
			continue // key already exists in mapping, skip
		} else if encrypted, ok := this.encryptByte(plaintext[i]); ok {
			if encrypted != cyphertext[i] {
				this.RemoveLastNKeys(keysAdded)
				return 0, false
			}
		}

		this.addKey([2]byte{ cyphertext[i], plaintext[i] })
		keysAdded++
	}

	return keysAdded, true
}

func (this *Code) addKey(key [2]byte) {
	this.decryptionMap[byteToMapIndex(key[0])] = key[1]
	this.encryptionMap[byteToMapIndex(key[1])] = key[0]

	this.keyStack = append(this.keyStack, key)
}

func (this *Code) removeLastKey() {
	keyToRemove := this.keyStack[len(this.keyStack) - 1]

	this.removeKey(keyToRemove)

	this.keyStack = this.keyStack[:len(this.keyStack) - 1]
}

func (this *Code) RemoveLastNKeys(n int) {
	for i := 0; i < n; i++ {
		this.removeLastKey()
	}
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