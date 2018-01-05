package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/pocockn/crypto-compare-go/persistance"
	"github.com/pocockn/crypto-compare-go/wallet"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func init() {
	persistance.InitDB("crypto_compare_test")
	persistance.DB.Exec("TRUNCATE TABLE wallets;")
	err := persistance.CreateSchema(&wallet.Wallet{})
	if err != nil {
		panic(err)
	}
	wallet.BootstrapWallet()
}

func TestCreateWalletHandler(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("coin", "BTC")
	q.Set("units", "100")
	req, err := http.NewRequest(echo.POST, "/createWallet", strings.NewReader(q.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, CreateWallet(c)) {
		assert.Equal(t, 301, rec.Code)
	}
}

func TestDepositHandler(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("coin", "BTC")
	q.Set("units", "20")
	req, err := http.NewRequest(echo.POST, "/", strings.NewReader(q.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/deposit/:id")
	c.SetParamNames("id")
	c.SetParamValues("1234")
	if assert.NoError(t, DepositCoin(c)) {
		assert.Equal(t, 301, rec.Code)
	}
}

func TestWithdrawHandler(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("coin", "BTC")
	q.Set("units", "20")
	req, err := http.NewRequest(echo.POST, "/", strings.NewReader(q.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/withdraw/:id")
	c.SetParamNames("id")
	c.SetParamValues("1234")
	if assert.NoError(t, WithdrawCoin(c)) {
		assert.Equal(t, 301, rec.Code)
	}
}

func TestGetWalletHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/wallet")
}
