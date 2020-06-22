package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"golang.org/x/crypto/pbkdf2"
)

const salt = "nacl"
const nonceString = "deadbeefdeadc0dedeaddead" // 12-byte hex encoded nonce

/*ViperUsage - Viper comes ready to use out of the box. There is no configuration or initialization needed to begin using Viper.
Since most applications will want to use a single central repository for their configuration, the viper package provides this. It is similar to a singleton.*/
func ViperUsage() {

	/* Viper can search multiple paths, but currently a single Viper instance only supports a single configuration file.
	Viper does not default to any configuration search paths leaving defaults decision to an application. */
	viper.AddConfigPath("/Users/anwarsahib/go/src/Go-Basics") // look for config in the working directory
	viper.SetConfigName("test")                               // name of config file (without extension)
	viper.SetConfigType("json")                               // REQUIRED if the config file does not have the extension in the name

	//Default EDBRAdmin cred available without password and user name
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config not found:", err)
	} else {
		name := viper.GetString("credName")
		fmt.Println("Config found, name = ", name)
	}

	/* Viper supports the ability to have your application live read a config file while running.
	Make sure you add all of the configPaths prior to calling WatchConfig() */
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name) //This event is called every time we write to config file
	})

	//On set-admin-password, set the user to defautl EDBR user "admin"
	viper.Set("user", "Admin")
	fmt.Println("After setting user to admin, reading user:", viper.GetString("user"))

	/*One important thing to recognize is that each Get function will return a zero value if it’s not found.
	To check if a given key exists, the IsSet() method has been provided.*/
	fmt.Println("Get User : ", viper.GetString("User"))         //Viper configuration keys are case insensitive
	fmt.Println("Check User is present: ", viper.IsSet("User")) //Viper configuration keys are case insensitive, so this will return true
	fmt.Println("Check Usr is present: ", viper.IsSet("Usr"))   //This will return false

	/*WriteConfig - writes the current viper configuration to the predefined path, if exists. Errors if no predefined path. Will overwrite the current config file, if it exists.
	SafeWriteConfig - writes the current viper configuration to the predefined path. Errors if no predefined path. Will not overwrite the current config file, if it exists.
	WriteConfigAs - writes the current viper configuration to the given filepath. Will overwrite the given file, if it exists.
	SafeWriteConfigAs - writes the current viper configuration to the given filepath. Will not overwrite the given file, if it exists*/
	viper.WriteConfig() //The value is not written in file until this function is called
	fmt.Println("Writing the name in file:", viper.GetString("user"))

	//Consider admin password is set as "AdminPwd", save it in json file in encrypted format
	encryptPwd, err := encryptString([]byte("MasterPassword"), "AdminPwd")
	if err != nil {
		fmt.Println("Error encrypting password:", err)
	}
	fmt.Println("Encrypted Admin password that is saved in file:", encryptPwd)
	viper.Set("password", encryptPwd)

	decryptPwd, err := decryptString([]byte("MasterPassword"), viper.GetString("password"))
	if err != nil {
		fmt.Println("Error encrypting password:", err)
	}
	fmt.Println("Password decrypted after reading from file:", decryptPwd)
	viper.WriteConfig()

	//Aliases permit a single value to be referenced by multiple keys
	viper.RegisterAlias("ip", "edbrIp")
	fmt.Println("Current IP :", viper.GetString("ip"))

	viper.Set("edbrIp", "192.168.1.11")
	fmt.Println("Current IP after updating via alias:", viper.GetString("ip"))

}

//==============Just added for POC purposes. Core supports this functions where password can be encrypted by adding it to secure part=====
// helper function that encrypts a string and returns the ciphertext as a hex encoded string
func encryptString(password []byte, plaintext string) (ciphertext string, err error) {
	encrypted, err := encrypt(password, []byte(plaintext))
	ciphertext = hex.EncodeToString(encrypted)
	return
}

// helper function that decrypts a hex encoded string and returns the plaintext as a string
func decryptString(password []byte, ciphertext string) (plaintext string, err error) {
	cipherBytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plain, err := decrypt(password, cipherBytes)
	plaintext = string(plain)
	return
}

// This functions encrypts the given plaintext using the given password.  A 32 byte key is derived from the
// password using PBKDF2 with SHA256 hash.  The key is then used to encrypt the plaintext using AES256 and
// the function then returns the ciphertext.
func encrypt(password, plaintext []byte) (ciphertext []byte, err error) {
	key := pbkdf2.Key(password, []byte(salt), 4096, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce, _ := hex.DecodeString(nonceString)

	ciphertext = aesgcm.Seal(nil, nonce, plaintext, nil)
	//fmt.Printf("ciphertext: %x\n", ciphertext)
	return ciphertext, nil
}

// This functions decrypts the given ciphertext using the given password.  A 32 byte key is derived from the
// password using PBKDF2 with SHA256 hash.  The key is then used to decrypt the ciphertext using AES256 and
// the function then returns the plaintext.
func decrypt(password, ciphertext []byte) (plaintext []byte, err error) {
	key := pbkdf2.Key(password, []byte(salt), 4096, 32, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce, _ := hex.DecodeString(nonceString)

	plaintext, err = aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("plaintext: %s\n", plaintext)
	return plaintext, nil
}
