package lecturer

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"
	_ "github.com/linmadan/agile-journey-demo/pkg/infrastructure/pg"
	_ "github.com/linmadan/agile-journey-demo/pkg/port/beego"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLecturer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Beego Port Lecturer Correlations Test Case Suite")
}

var handler http.Handler
var server *httptest.Server

var _ = BeforeSuite(func() {
	handler = beego.BeeApp.Handlers
	server = httptest.NewServer(handler)
})

var _ = AfterSuite(func() {
	server.Close()
})
