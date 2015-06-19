package matasano_test

// Challenge 1 const's
const HexString string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
const Base64String string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

// Challenge 2 const's
const XorString1 string = "1c0111001f010100061a024b53535009181c"
const XorString2 string = "686974207468652062756c6c277320657965"
const XorAnswer string = "746865206b696420646f6e277420706c6179"

// Challenge 3 const's
const EncodedString string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
const DecodedString string = "Cooking MC's like a pound of bacon"
const DecodedCipher byte = 0x58

// Challenge 4 const's
const FilePath string = "fixtures/4.txt"
const DecodedFileString string = "Now that the party is jumping\n"
const DecodedFileCipher byte = 0x35
const DecodedFileLine int = 171

// Challenge 5 const's
const OpeningStanza string = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
const RepeatingKeyCipher string = "ICE"
const RepeatingXorResult string = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
