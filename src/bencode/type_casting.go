package bencode

type (
    BencodeString string
    BencodeInteger int64
    BencodeList []interface{}
    BencodeDictionary map[BencodeString]interface{}
)
