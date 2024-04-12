package main

import (
	"bytes"
	sha256 "crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	blake2b "golang.org/x/crypto/blake2b"
	keccak256 "golang.org/x/crypto/sha3"
)

func Main() {
	var name string
	fmt.Println("Enter string for hashing: ")
	fmt.Scanf("%s", &name)
	var size int
	fmt.Println("Enter size customize size for blake2b hashing default size will be 256")
	fmt.Scanf("%d", &size)
	fmt.Println("Value", name)

	hashValue := getHash(size, name)
	fmt.Println("blake2b value  ", hashValue)
	hashValue16 := getHash(16, name)

	fmt.Println("blake2b value hash size 16  ", hashValue16)

	sha256HashValue := getHashSha256(name)
	fmt.Println("sha256 value ", sha256HashValue)
	keccakValue := getKeccakValue(name)
	fmt.Println("keccakValue :", keccakValue)
	jsonMarshal, _ := getJsonMarshal(name)
	fmt.Println("JSON marshal value", jsonMarshal)
	txJsonMarshal, _ := getTxMarshal(name)
	fmt.Println("TX JSON marshal", txJsonMarshal)
	hexEncodingToString := hexCodeEncodingString(name)
	fmt.Println("HEX CODE encoding to string", hexEncodingToString)

}

func getHash(hashedSize int, hashedName string) []byte {
	if hashedSize == 0 {
		hashed, _ := blake2b.New256(nil)
		_, _ = hashed.Write([]byte(hashedName))
		return hashed.Sum(nil)
	}
	hashed, _ := blake2b.New(hashedSize, nil)
	_, _ = hashed.Write([]byte(hashedName))
	return hashed.Sum(nil)
}

func getHashSha256(s string) []byte {
	h := sha256.New()
	_, _ = h.Write([]byte(s))
	return h.Sum(nil)
}
func getKeccakValue(name string) []byte {
	h := keccak256.NewLegacyKeccak256()
	_, _ = h.Write([]byte(name))
	return h.Sum(nil)

}

func getJsonMarshal(name string) ([]byte, error) {
	return json.Marshal(name)

}
func getTxMarshal(name string) ([]byte, error) {
	bytesBuffer := new(bytes.Buffer)
	jsonEncoder := json.NewEncoder(bytesBuffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(name)
	if err != nil {
		return nil, err
	}

	encodedResult := bytesBuffer.Bytes()
	return trimLineFeed(encodedResult), nil
}

func trimLineFeed(bytes []byte) []byte {
	// this should be replaced, but for some reason, bytes.TrimRight(b, "\r") does not work
	lastByte := bytes[len(bytes)-1:]
	if lastByte[0] == byte(10) { // hardcoded for now
		return bytes[:len(bytes)-1]
	}

	return bytes
}

func hexCodeEncodingString(txt string) string {
	src := []byte("Hello")
	return hex.EncodeToString(src)

}

func main() {
	Main()
}
