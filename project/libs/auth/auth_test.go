package auth

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestAuthenticator_Authenticate(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	publicKey := &privateKey.PublicKey

	cl := &jwt.RegisteredClaims{
		Issuer:    "Test",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 3600)),
		Subject:   "test",
	}
	validToken, _ := jwt.NewWithClaims(jwt.SigningMethodRS256, cl).SignedString(privateKey)

	hmacKey := hmac.New(sha256.New, []byte("test"))
	invalidSignatureToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, cl).SignedString(hmacKey.Sum([]byte("test")))
	require.NoError(t, err)

	cl = &jwt.RegisteredClaims{
		Issuer:    "Test",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * -3600)),
		Subject:   "test",
	}
	invalidToken, _ := jwt.NewWithClaims(jwt.SigningMethodRS256, cl).SignedString(privateKey)

	t.Run("should return error if secret key is invalid", func(t *testing.T) {
		secretKey := "test"

		_, err := NewAuthenticator(secretKey)

		require.Error(t, err)
	})

	t.Run("should return no error for correct secret key", func(t *testing.T) {
		pubPEM := pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PUBLIC KEY",
				Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
			},
		)

		_, err := NewAuthenticator(string(pubPEM))

		require.NoError(t, err)
	})

	tests := []struct {
		name    string
		token   string
		wantErr bool
	}{
		{
			name:    "should return error if failed to parse token",
			token:   "test",
			wantErr: true,
		},
		{
			name:    "should return error if token is expired",
			token:   invalidToken,
			wantErr: true,
		},
		{
			name:    "should return error if token signature is invalid",
			token:   invalidSignatureToken,
			wantErr: true,
		},
		{
			name:    "should return no error for correct token",
			token:   validToken,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Authenticator{
				secretKey: publicKey,
			}
			_, err = a.Authenticate(tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestFromContext(t *testing.T) {
	cl := &jwt.RegisteredClaims{
		Issuer:    "Test",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 3600)),
		Subject:   "test",
	}
	ctx := ToContext(context.Background(), jwt.NewWithClaims(jwt.SigningMethodRS256, cl))
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want jwt.Claims
	}{
		{
			name: "should return nil if no token found",
			args: args{
				ctx: context.Background(),
			},
			want: nil,
		},
		{
			name: "should return valid claims",
			args: args{
				ctx: ctx,
			},
			want: cl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromContext(tt.args.ctx)
			require.Equal(t, tt.want, got)
		})
	}
}
