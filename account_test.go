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

func (suite *AccountTestSuite) TestDecodesTheAccountUuid (c *C) {
	account, _ := AccountInfo(suite.options, nil)

	c.Assert(account.Uuid, Equals, "alksdjfhlakjdsfh12983712")
}

func (suite *AccountTestSuite) TestDecodesTheAccountEmail (c *C) {
	account, _ := AccountInfo(suite.options, nil)

	c.Assert(account.Email, Equals, "droplet_kit@digitalocean.com")
}

func (suite *AccountTestSuite) TestDecodesTheAccountEmailVerified (c *C) {
	account, _ := AccountInfo(suite.options, nil)

	c.Assert(account.EmailVerified, Equals, true)
}

func (suite *AccountTestSuite) TestDecodesTheAccountDropletLimit (c *C) {
	account, _ := AccountInfo(suite.options, nil)

	c.Assert(account.DropletLimit, Equals, 200)
}

func (suite *AccountTestSuite) TestReturnsAnUnauthenticatedErrorIfTheErrorIsWrong (c *C) {
	suite.options.Token = "invalid"

	account, error := AccountInfo(suite.options, nil)

	c.Assert(account.Email, Equals, "")
	c.Assert(error, Equals, UnauthenticatedError)
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
