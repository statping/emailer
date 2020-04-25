package main

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

var (
	config *Config
)

type Config struct {
	Host     string
	Username string
	Password string
	Port     int
}

func InitConfig() {
	config = &Config{
		Host:     os.Getenv("HOST"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Port:     toInt(os.Getenv("PORT")),
	}
}

func SendEmail(u *User) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "info@betatude.com")
	m.SetHeader("To", u.Email)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	fmt.Printf("Sending email to %s...\n", u.Email)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Sent email to %s\n", u.Email)

	return nil
}

func toInt(s string) int {
	num, _ := strconv.ParseInt(s, 10, 64)
	return int(num)
}
