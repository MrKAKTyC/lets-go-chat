package hasher

import "testing"

var (
	testPassword       = "qwerty"
	testPasswordHash   = "65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5"
	zeroLengthPassword = ""
)

func TestHashPassword_hashMatching(t *testing.T) {
	t.Parallel()
	hashedPassword, hashError := HashPassword(testPassword)
	if testPasswordHash != hashedPassword {
		t.Errorf("Exected: %s got: %s", testPasswordHash, hashedPassword)
	}
	if hashError != nil {
		t.Error("No errors expected")
	}
}

func TestHashPassword_ErrorForZeroLength(t *testing.T) {
	t.Parallel()
	_, hashError := HashPassword(zeroLengthPassword)
	if hashError == nil {
		t.Error("Error is expected for empty string")
	}
}

func TestCheckPasswordHash_CorrectHash(t *testing.T) {
	t.Parallel()
	if !CheckPasswordHash(testPassword, testPasswordHash) {
		t.Error("Must be true for correct password hash")
	}
}

func TestCheckPasswordHash_WrongHash(t *testing.T) {
	t.Parallel()
	if CheckPasswordHash(testPassword, "incorrect password hash") {
		t.Error("Must be false for ncorrect password hash")
	}
}

func TestCheckPasswordHash_IllegalPassword(t *testing.T) {
	t.Parallel()
	if CheckPasswordHash(zeroLengthPassword, testPasswordHash) {
		t.Error("Must be false for zero length password")
	}
}
