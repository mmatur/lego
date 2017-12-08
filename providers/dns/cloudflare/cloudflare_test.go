package cloudflare

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	cflareLiveTest       bool
	cflareEmail          string
	cflareAPIKey         string
	cflareUserServiceKey string
	cflareDomain         string
)

func init() {
	cflareEmail = os.Getenv("CLOUDFLARE_EMAIL")
	cflareAPIKey = os.Getenv("CLOUDFLARE_API_KEY")
	cflareUserServiceKey = os.Getenv("CLOUDFLARE_USER_SERVICE_KEY")
	cflareDomain = os.Getenv("CLOUDFLARE_DOMAIN")
	if (len(cflareEmail) > 0 && len(cflareAPIKey) > 0 || len(cflareUserServiceKey) > 0) && len(cflareDomain) > 0 {
		cflareLiveTest = true
	}
}

func restoreCloudFlareEnv() {
	os.Setenv("CLOUDFLARE_EMAIL", cflareEmail)
	os.Setenv("CLOUDFLARE_API_KEY", cflareAPIKey)
	os.Setenv("CLOUDFLARE_USER_SERVICE_KEY", "")
}

func TestNewDNSProviderValid(t *testing.T) {
	os.Setenv("CLOUDFLARE_EMAIL", "")
	os.Setenv("CLOUDFLARE_API_KEY", "")
	os.Setenv("CLOUDFLARE_USER_SERVICE_KEY", "")
	_, err := NewDNSProviderCredentials("123", "123", "")
	assert.NoError(t, err)
	restoreCloudFlareEnv()
}

func TestNewDNSProviderValidEnv(t *testing.T) {
	os.Setenv("CLOUDFLARE_EMAIL", "test@example.com")
	os.Setenv("CLOUDFLARE_API_KEY", "123")
	os.Setenv("CLOUDFLARE_USER_SERVICE_KEY", "v1-123")
	_, err := NewDNSProvider()
	assert.NoError(t, err)
	os.Setenv("CLOUDFLARE_EMAIL", "")
	os.Setenv("CLOUDFLARE_API_KEY", "123")
	os.Setenv("CLOUDFLARE_USER_SERVICE_KEY", "v1-123")
	_, err = NewDNSProvider()
	os.Setenv("CLOUDFLARE_EMAIL", "test@example.com")
	os.Setenv("CLOUDFLARE_API_KEY", "")
	os.Setenv("CLOUDFLARE_USER_SERVICE_KEY", "v1-123")
	_, err = NewDNSProvider()
	assert.NoError(t, err)
	os.Setenv("CLOUDFLARE_EMAIL", "")
	os.Setenv("CLOUDFLARE_API_KEY", "")
	os.Setenv("CLOUDFLARE_USER_SERVICE_KEY", "v1-123")
	_, err = NewDNSProvider()
	assert.NoError(t, err)
	restoreCloudFlareEnv()
}

func TestNewDNSProviderMissingCredErr(t *testing.T) {
	os.Setenv("CLOUDFLARE_EMAIL", "")
	os.Setenv("CLOUDFLARE_API_KEY", "")
	os.Setenv("CLOUDFLARE_USER_SERVICE_KEY", "")
	_, err := NewDNSProvider()
	assert.EqualError(t, err, "CloudFlare credentials missing")
	os.Setenv("CLOUDFLARE_EMAIL", "test@example.com")
	os.Setenv("CLOUDFLARE_API_KEY", "")
	os.Setenv("CLOUDFLARE_USER_SERVICE_KEY", "")
	_, err = NewDNSProvider()
	assert.EqualError(t, err, "CloudFlare credentials missing")
	os.Setenv("CLOUDFLARE_EMAIL", "")
	os.Setenv("CLOUDFLARE_API_KEY", "123")
	os.Setenv("CLOUDFLARE_USER_SERVICE_KEY", "")
	_, err = NewDNSProvider()
	assert.EqualError(t, err, "CloudFlare credentials missing")
	restoreCloudFlareEnv()
}

func TestCloudFlarePresent(t *testing.T) {
	if !cflareLiveTest {
		t.Skip("skipping live test")
	}

	provider, err := NewDNSProviderCredentials(cflareEmail, cflareAPIKey, cflareUserServiceKey)
	assert.NoError(t, err)

	err = provider.Present(cflareDomain, "", "123d==")
	assert.NoError(t, err)
}

func TestCloudFlareCleanUp(t *testing.T) {
	if !cflareLiveTest {
		t.Skip("skipping live test")
	}

	time.Sleep(time.Second * 2)

	provider, err := NewDNSProviderCredentials(cflareEmail, cflareAPIKey, cflareUserServiceKey)
	assert.NoError(t, err)

	err = provider.CleanUp(cflareDomain, "", "123d==")
	assert.NoError(t, err)
}
