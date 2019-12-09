package ggn_encryption

func Xor(input, key string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i % len(key)])
	}
	return output
}
