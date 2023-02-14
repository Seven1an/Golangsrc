package main

import (
    "crypto/rand"
    "flag"
    "fmt"
    "math/big"
    "strings"
)

var (
    // 生成密码时使用的字符集
    letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    digits  = "0123456789"
    symbols = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

// generatePassword 生成一个指定长度的随机密码
func generatePassword(length int) (string, error) {
    // 生成的密码包含的字符集
    charset := []string{letters, digits, symbols}
    // 生成密码时，每种字符集中字符的数量
    counts := []int{(length + 2) / 3, length / 3, length / 3}

    // 随机生成每种字符集中的字符
    var password strings.Builder
    for i, charSet := range charset {
        for j := 0; j < counts[i]; j++ {
            index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
            if err != nil {
                return "", err
            }
            password.WriteByte(charSet[index.Int64()])
        }
    }

    // 如果生成的密码长度不足指定长度，从字符集中随机选择字符填充
    for password.Len() < length {
        index, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters+digits+symbols))))
        if err != nil {
            return "", err
        }
        password.WriteByte((letters + digits + symbols)[index.Int64()])
    }

    return password.String(), nil
}

func main() {
    // 定义一个 length 参数，用于指定密码长度
    length := flag.Int("length", 12, "password length")
    flag.Parse()

    password, err := generatePassword(*length)
    if err != nil {
        fmt.Println("Failed to generate password:", err)
        return
    }
    fmt.Println("Generated password:", password)
}
