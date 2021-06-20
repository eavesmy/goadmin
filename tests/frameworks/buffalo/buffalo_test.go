package buffalo

import (
	"net/http"
	"testing"

	"github.com/eavesmy/goadmin/tests/common"
	"github.com/gavv/httpexpect"
)

func TestBuffalo(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(newHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
