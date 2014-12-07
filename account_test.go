package dropletkit

import (
	"fmt"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type AccountTestSuite struct{
	mockServer *httptest.Server
	options Options
}

var _ = Suite(&AccountTestSuite{})

func (suite *AccountTestSuite) SetUpTest(c *C) {
	suite.mockServer = createMockServer()
	suite.options = Options{baseUrl: suite.mockServer.URL, version: "v2", Token: "qwertyuiop"}
}

func (suite *AccountTestSuite) TearDownTest(c *C) {
	suite.mockServer.Close()
}

func (suite *AccountTestSuite) TestDecodeTheAccountUuid (c *C) {
	account := AccountInfo(suite.options)

	c.Assert(account.Uuid, Equals, "alksdjfhlakjdsfh12983712")
}

func (suite *AccountTestSuite) TestDecodeTheAccountEmail (c *C) {
	account := AccountInfo(suite.options)

	c.Assert(account.Email, Equals, "droplet_kit@digitalocean.com")
}

func (suite *AccountTestSuite) TestDecodeTheAccountEmailVerified (c *C) {
	account := AccountInfo(suite.options)

	c.Assert(account.EmailVerified, Equals, true)
}

func (suite *AccountTestSuite) TestDecodeTheAccountDropletLimit (c *C) {
	account := AccountInfo(suite.options)

	c.Assert(account.DropletLimit, Equals, 200)
}

func createMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if(request.Header.Get("Authorization") != "Bearer qwertyuiop") {
			writer.WriteHeader(401)
		} else if request.Method == "GET" && request.RequestURI == "/v2/account" {
			data, err := ioutil.ReadFile("./fixtures/account/info.json")
			if(err != nil) { panic("Failed to read the fixture") }
			fmt.Fprintln(writer, string(data))
		} else {
			writer.WriteHeader(404)
		}
	}))
}
