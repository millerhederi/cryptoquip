package main

import (
	"testing"
	"reflect"
)

func TestShouldHandleUpdateNewCode(t *testing.T) {
	c := NewCode()

	keys, ok := c.TryUpdateWithDecryptedWord([]byte("NT"), []byte("ON"))

	if !ok {
		t.Errorf("ok == %v, expected %v", ok, true)
	}

	expectedKeys := [][2]byte{ [2]byte{ 'N', 'O' }, [2]byte { 'T', 'N' }}
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("keys == %v, expected %v", keys, expectedKeys)
	}
}

func TestShouldHandleUpdateWithDuplicateKeys(t *testing.T) {
	c := NewCode()
	c.TryUpdateWithDecryptedWord([]byte("NT"), []byte("ON"))

	keys, ok := c.TryUpdateWithDecryptedWord([]byte("NT"), []byte("ON"))

	if !ok {
		t.Errorf("ok == %v, expected %v", ok, true)
	}

	expectedKeys := [][2]byte{}
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("keys == %v, expected %v", keys, expectedKeys)
	}
}

func TestShouldHandleUpdateWithNewKeys(t *testing.T) {
	c := NewCode()
	c.TryUpdateWithDecryptedWord([]byte("NT"), []byte("ON"))

	keys, ok := c.TryUpdateWithDecryptedWord([]byte("NR"), []byte("OF"))

	if !ok {
		t.Errorf("ok == %v, expected %v", ok, true)
	}

	expectedKeys := [][2]byte{ [2]byte{ 'R', 'F' }}
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("keys == %v, expected %v", keys, expectedKeys)
	}
}

func TestShouldFailWhenUpdatingWithConflictingDecryptionKey(t *testing.T) {
	c := NewCode()
	c.TryUpdateWithDecryptedWord([]byte("NT"), []byte("ON"))

	keys, ok := c.TryUpdateWithDecryptedWord([]byte("KT"), []byte("IF"))

	if ok {
		t.Errorf("ok == %v, expected %v", ok, false)
	}

	expectedKeys := [][2]byte{ }
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("keys == %v, expected %v", keys, expectedKeys)
	}
}

func TestShouldFailWhenUpdatingWithConflictingEncryptionKey(t *testing.T) {
	c := NewCode()
	c.TryUpdateWithDecryptedWord([]byte("NT"), []byte("ON"))

	keys, ok := c.TryUpdateWithDecryptedWord([]byte("PZ"), []byte("IN"))

	if ok {
		t.Errorf("ok == %v, expected %v", ok, false)
	}

	expectedKeys := [][2]byte{ }
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("keys == %v, expected %v", keys, expectedKeys)
	}
}