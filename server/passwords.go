package server

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var cached_passwords [][]string
var passwords_file string = ".shadow"

// checking password for correctness
func checkPassword(username, password string) (bool, error) {
	// get user's hashed password
	file, err := os.OpenFile(passwords_file, os.O_RDWR, 0600)
	if err != nil {
		return false, err
	}
	defer file.Close()
	file_scanner := bufio.NewScanner(file)
	file_scanner.Split(bufio.ScanLines)
	// password_data index 0 - username, index 1 - password hash
	for file_scanner.Scan() {
		password_data := strings.Split(file_scanner.Text(), ":")
		if password_data[0] == username {
			return CheckPasswordHash(password, password_data[1]), nil
		}
	}
	return false, errors.New("user not exists")
}

// check user exists
func userExists(username string) (bool, error) {
	// get file data
	file, err := os.OpenFile(passwords_file, os.O_RDWR, 0600)
	if err != nil {
		return false, err
	}
	defer file.Close()
	file_scanner := bufio.NewScanner(file)
	file_scanner.Split(bufio.ScanLines)
	// password_data index 0 - username, index 1 - password hash
	for file_scanner.Scan() {
		password_data := strings.Split(file_scanner.Text(), ":")
		if password_data[0] == username {
			return true, nil
		}
	}
	return false, nil
}

// update user password by username
func updateUserPassword(username, new_password string) error {
	// hash password
	new_password_hash, err := HashPassword(new_password)
	if err != nil {
		return err
	}
	// get passwords from file
	file, err := os.OpenFile(passwords_file, os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	file_scanner := bufio.NewScanner(file)
	file_scanner.Split(bufio.ScanLines)
	var passwords []string
	for file_scanner.Scan() {
		passwords = append(passwords, file_scanner.Text())
	}
	// password_data index 0 - username, index 1 - password hash
	var configured_string string
	for password_s := range passwords {
		password_data := strings.Split(passwords[password_s], ":")
		if password_data[0] == username {
			password_data[1] = new_password_hash
		}
		configured_string = fmt.Sprintf("%s%s:%s\n", configured_string, password_data[0], password_data[1])
	}
	fmt.Println(configured_string)
	// write new data
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = file.WriteString(configured_string)
	if err != nil {
		return err
	}
	err = file.Sync()
	return err
}

// add user password
func writePasswordToFile(username, password string) error {
	// check what username is correct
	err := checkUsername(username)
	if err != nil {
		return err
	}
	// hash password
	password_hash, err := HashPassword(password)
	if err != nil {
		return err
	}
	// write to file
	pass_string := fmt.Sprintf("%s:%s\n", username, password_hash)
	file, err := os.OpenFile(passwords_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()
	file.Write([]byte(pass_string))
	return err
}

// Password hashing
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// Compare hash and password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// check that username not contains forbidden symbols
func checkUsername(name string) error {
	if strings.ContainsAny(name, ":\"'`/\\") {
		return errors.New("Incorrect username. Not use symbols :\"'`/\\")
	}
	return nil
}
