package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// func main() {
// 	// 模拟数据库中的哈希密码（假设从数据库中查询出的密码哈希）
// 	storedPasswordHash := "$2a$10$jdyfTSJ2BLd3d89ViJVIt.SVvJ83tSPZCH8rM2QukurlGJn1e3jgW" // 数据库中的哈希密码

// 	// 模拟用户输入的密码
// 	inputPassword := "password" // 用户输入的明文密码

// 	// 生成用户输入的密码的哈希值
// 	hashedInputPassword, err := bcrypt.GenerateFromPassword([]byte(inputPassword), bcrypt.DefaultCost)
// 	if err != nil {
// 		log.Fatal("Error generating hash for input password:", err)
// 	}

// 	// 打印生成的哈希密码和数据库中的哈希密码用于调试
// 	log.Println("Generated Hashed Password from Input (Login):", string(hashedInputPassword))
// 	log.Println("Stored Hashed Password from DB:", storedPasswordHash)

//		// 使用 bcrypt 比较数据库中的哈希密码和用户输入密码生成的哈希值
//		err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(inputPassword))
//		if err != nil {
//			log.Println("Password mismatch:", err)
//		} else {
//			log.Println("Password match!")
//		}
//	}

func main() {
	password := "password1"

	// 生成哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generated hashed password:", string(hashedPassword))

	// 使用生成的哈希密码再次进行比较，确保 bcrypt 正常工作
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		log.Println("Password mismatch:", err)
	} else {
		log.Println("Password match!")
	}
}
