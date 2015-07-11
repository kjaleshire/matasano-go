package matasano_test

// Challenge 1 consts
const C1HexString string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
const C1Base64String string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

// Challenge 2 consts
const C2RawString1 string = "1c0111001f010100061a024b53535009181c"
const C2RawString2 string = "686974207468652062756c6c277320657965"
const C2XorResult string = "746865206b696420646f6e277420706c6179"

// Challenge 3 consts
const C3EncodedString string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
const C3DecodedString string = "Cooking MC's like a pound of bacon"
const C3Key byte = 0x58

// Challenge 4 consts
const C4FilePath string = "fixtures/4.txt"
const C4DecodedString string = "Now that the party is jumping\n"
const C4Key byte = 0x35
const C4FileLine int = 171

// Challenge 5 consts
const C5UnencodedString string = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
const C5RepeatingKey string = "ICE"
const C5XorResult string = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

// Challenge 6 consts
const C6HammingString1 string = "this is a test"
const C6HammingString2 string = "wokka wokka!!!"
const C6HammingDistance int = 37

const C6FilePath string = "fixtures/6.txt"
const C6Key string = "Terminator X: Bring the noise"

// Challenge 7 consts
const C7FilePath string = "fixtures/7.txt"
const C7Key string = "YELLOW SUBMARINE"
const C7DecodedFirstLine string = "I'm back and I'm ringin' the bell"

// Challenge 8 consts
const C8FilePath string = "fixtures/8.txt"
const C8LineNumber int = 133
