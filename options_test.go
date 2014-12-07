package dropletkit

import (
  . "gopkg.in/check.v1"
)

type OptionsTestSuite struct{}

var _ = Suite(&OptionsTestSuite{})

func (suite *AccountTestSuite) TestDefaultOptionsHasABaseUrl (c *C) {
  options := DefaultOptions()

  c.Assert(options.baseUrl, Equals, "https://api.digitalocean.com")
}

func (suite *AccountTestSuite) TestDefaultOptionsHasAVersion (c *C) {
  options := DefaultOptions()

  c.Assert(options.version, Equals, "v2")
}
