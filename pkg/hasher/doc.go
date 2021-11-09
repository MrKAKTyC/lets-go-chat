/*
Package hasher have functions for hashing strings and comparing string and hash.
For hashing it`s uses sha256 algorithm.

Functions

func HashPassword(password string) (string, error)

This function accepts non zero length string and returns its hash
Algorithm used for hashing is sha256. If string is zero length then
empty string and error is returned

func CheckPasswordHash(password, hash string) bool

This function accepts two strings, password is a string which will be hashed and
compared with hash. Function return true if hash of password and hash parameter is matching.
If hashes missmatching or password is empty false is returned
*/