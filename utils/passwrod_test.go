package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	plainText := "abc123"
	hashedPwd, err := GenHashedPwsd(plainText)
	require.NoError(t, err)
	err = VerifyPassword(plainText, hashedPwd)
	require.NoError(t, err)

	err = VerifyPassword("abc1234", hashedPwd)
	require.Error(t, err)
}
