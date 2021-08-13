package url_test

import (
	"shorturl/repository/url/test"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
