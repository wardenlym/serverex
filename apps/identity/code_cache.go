package identity

import (
	"fmt"
	"math/rand"
	"time"

	cache "github.com/patrickmn/go-cache"
)

type CodeCache interface {
	Add(phone string) string
	// Get(phone string) (string, bool)
	Validate(phone, code string) bool
}

func NewCodeVerify() CodeCache {
	return &code_cache{
		c: cache.New(5*time.Minute, 10*time.Minute),
	}
}

type code_cache struct {
	c *cache.Cache
}

func init() {

	rand.Seed(time.Now().UnixNano())
}

func gen_code() string {
	n := rand.Intn(1000000)
	return fmt.Sprintf("%06d", n)
}

func (p *code_cache) Add(phone string) string {
	code := gen_code()
	p.c.Set(phone, code, cache.DefaultExpiration)
	return code
}

// func (p *code_cache) Get(phone string) (string, bool) {
// 	i, ok := p.c.Get(phone)
// 	if ok {
// 		return i.(string), true
// 	}
// 	return "", false
// }

func (p *code_cache) Validate(phone, code string) bool {

	i, ok := p.c.Get(phone)
	if !ok {
		return false
	}

	if code == i.(string) {
		return true
	}

	return false
}
