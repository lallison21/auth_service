package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/lallison21/auth_service/internal/app_errors"
	"golang.org/x/crypto/argon2"
	"strings"
)

type Password struct {
	Memory      uint32 `env:"PASSWORD_MEMORY" required:"true" default:"65536"`
	Iterations  uint32 `env:"PASSWORD_ITERATIONS" required:"true" default:"3"`
	Parallelism uint8  `env:"PASSWORD_PARALLELISM" required:"true" default:"4"`
	SaltLength  uint32 `env:"PASSWORD_SALT_LENGTH" required:"true" default:"16"`
	KeyLength   uint32 `env:"PASSWORD_KEY_LENGTH" required:"true" default:"32"`
}

type Utils struct {
	cfg *Password
}

func New(cfg *Password) *Utils {
	return &Utils{
		cfg: cfg,
	}
}

func (u *Utils) GeneratePassword(password string) (string, error) {
	salt := make([]byte, u.cfg.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, u.cfg.Iterations, u.cfg.Memory, u.cfg.Parallelism, u.cfg.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		u.cfg.Memory,
		u.cfg.Iterations,
		u.cfg.Parallelism,
		b64Salt,
		b64Hash,
	)

	return encodedHash, nil
}

func (u *Utils) ComparePassword(password, hash string) (bool, error) {
	vals := strings.Split(hash, "$")
	if len(vals) != 6 {
		return false, fmt.Errorf("invalid hash format")
	}

	var version int
	if _, err := fmt.Sscanf(vals[2], "v=%d", &version); err != nil {
		return false, app_errors.ErrInvalidHash
	}
	if version != argon2.Version {
		return false, app_errors.ErrIncompatibleVersion
	}

	p := &Password{}
	if _, err := fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism); err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return false, err
	}
	p.SaltLength = uint32(len(salt))

	decodedHash, err := base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return false, err
	}
	p.KeyLength = uint32(len(decodedHash))

	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	return subtle.ConstantTimeCompare(decodedHash, otherHash) == 1, nil
}
