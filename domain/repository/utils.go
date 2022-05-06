package repository

import "time"

type UtilsRepo interface {
	CurrentTimeF(format string) (string, error)
	RandomString(length int) string
	ToStruct(v map[string]interface{}, r interface{}) (err error)
	ToMap(v interface{}) (r map[string]interface{})
	GenerateRangeDate(fromYear, fromMonth int) (result []map[string]time.Time, err error)
	Encrypt(passphrase, plaintext string) (cipherText string, err error)
	Decrypt(passphrase, cipherText string) (plainText string, err error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
