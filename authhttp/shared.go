package authentication

import (
    "crypto/hmac"
    "crypto/md5"
    "crypto/sha1"
    "encoding/base64"
    "fmt"
    "hash"
    "io"
    "log"
    "time"
)

//type Credentials map[string]string

func GenerateHmacSignature(key string, cannonicalString string) hash.Hash {
    var h hash.Hash = hmac.New(sha1.New, []byte(key))
    h.Write([]byte(cannonicalString))
    return h
}

func EncodeHashToBase64(h hash.Hash) string {
    return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func EncodeStringToBase64(s string) string {
    toEncode := []byte(s)
    return base64.StdEncoding.EncodeToString(toEncode)
}

func DecodeBase64String(s string) (string, error) {
    decoded, err := base64.StdEncoding.DecodeString(s)
    if err != nil {
      return "", err
    }
    return string(decoded), nil
}

func GetCurrentGMTTime() string {
    timezone, err := time.LoadLocation("GMT")
    if err != nil {
        log.Fatal(err)
    }
    return time.Now().In(timezone).Format(time.RFC1123)
}

func GenerateMD5(rack string) string {
    if rack == "" {
        return ""
    }
    md5String := md5.New()
    io.WriteString(md5String, rack)
    return fmt.Sprintf("%x", md5String.Sum(nil))
}
