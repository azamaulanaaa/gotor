package bencode

type (
    String string
    Integer int64
    List []interface{}
    Dictionary map[String]interface{}
)
