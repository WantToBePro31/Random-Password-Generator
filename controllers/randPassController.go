package controllers

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/WantToBePro31/rand-pass/config"
	"github.com/WantToBePro31/rand-pass/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	specialChar       = "!@#$%&*"
	number            = "1234567890"
	uppercase         = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase         = "abcdefghijklmnopqrstuvwxyz"
	charSet           = specialChar + number + uppercase + lowercase
	generatedPassword string
)

func generatePassword(passLength int, minSpeChar int, minNum int, minUpper int) string {
	rand.Seed(time.Now().UnixNano())
	password := ""
	remPassLength := passLength - minSpeChar - minNum - minUpper

	for i := 0; i < minSpeChar; i++ {
		password += string(specialChar[rand.Intn(len(specialChar))])
	}

	for i := 0; i < minNum; i++ {
		password += string(number[rand.Intn(len(number))])
	}

	for i := 0; i < minUpper; i++ {
		password += string(uppercase[rand.Intn(len(uppercase))])
	}

	for i := 0; i < remPassLength; i++ {
		password += string(charSet[rand.Intn(len(charSet))])
	}

	runePass := []rune(password)
	rand.Shuffle(len(runePass), func(i, j int) {
		runePass[i], runePass[j] = runePass[j], runePass[i]
	})

	return string(runePass)
}

func RedirectToRandPassPage(c *gin.Context) {
	c.Redirect(http.StatusFound, "/random-password-generator")
}

func PasswordGeneratorPage(c *gin.Context) {
	c.HTML(http.StatusOK, "passGenerator.html", gin.H{
		"title": "Needify-Random Password Generator",
	})
}

func SubmitRequestHandler(c *gin.Context) {
	pas_len, _ := strconv.Atoi(c.PostForm("password-length"))
	special, _ := strconv.Atoi(c.PostForm("special-char"))
	num, _ := strconv.Atoi(c.PostForm("number"))
	upper, _ := strconv.Atoi(c.PostForm("uppercase"))

	rand_pass := generatePassword(pas_len, special, num, upper)

	password := models.RandPass{
		Id:             primitive.NewObjectID(),
		PasswordLength: pas_len,
		SpecialChar:    special,
		Number:         num,
		Uppercase:      upper,
		Password:       rand_pass,
	}

	_, err := config.Collection.InsertOne(context.Background(), &password)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "passGenerator.html", gin.H{"message": err.Error()})
		return
	}

	generatedPassword = password.Password
	c.Redirect(http.StatusFound, "/random-password-generator/result")
}

func PasswordResultPage(c *gin.Context) {
	c.HTML(http.StatusOK, "passResult.html", gin.H{
		"title":    "Needify-Random Password Generator",
		"password": generatedPassword,
	})
}

func DownloadPassword(c *gin.Context) {
	file, err := os.Create("password.txt")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "passResult.html", gin.H{"message": err.Error()})
		return
	}

	defer file.Close()
	text := "Your new password is " + generatedPassword
	fmt.Fprintf(file, text)

	c.Redirect(http.StatusFound, "/random-password-generator/result/downloaded")
}

func PasswordDownloadedPage(c *gin.Context) {
	c.HTML(http.StatusOK, "passDownloaded.html", gin.H{
		"title":    "Needify-Random Password Generator",
		"password": generatedPassword,
	})
}
