package fake

import (
	"github.com/rantav/go-logger"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("meets the interface", func() {
	var _ log.Logger = &FakeLogger{}
})
