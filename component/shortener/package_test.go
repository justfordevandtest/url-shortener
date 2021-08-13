package shortener_test

import (
	"shorturl/component/shortener/test"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
