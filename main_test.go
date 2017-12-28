package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/pocockn/crypto-compare-go/models"
	"github.com/pocockn/crypto-compare-go/wallet"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockTask = models.Task{
		ID:   123,
		Name: "Tie Shoe Laces",
	}
	bodyJSON          = `{ "name" : "Tie Shoe Laces" }`
	createWalletJSON  = `{"CoinsHeld":{"btc":100}}`
	withdrawJSON      = `{"CoinsHeld":{"BTC":50}}`
	depositWalletJSON = `{"CoinsHeld":{"BTC":200}}`

	mockWallet = wallet.Wallet{
		CoinsHeld: make(map[string]int),
	}
)

func TestCreateWalletHandler(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("coin", "btc")
	q.Set("units", "100")
	req, err := http.NewRequest(echo.GET, "/createWallet?"+q.Encode(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, createWallet(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.True(t, strings.ContainsAny(rec.Body.String(), createWalletJSON))
	}
}

func TestDepositHandler(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("coin", "BTC")
	q.Set("units", "100")
	req, err := http.NewRequest(echo.GET, "/deposit?"+q.Encode(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, depositFunds(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.True(t, strings.ContainsAny(rec.Body.String(), depositWalletJSON))
	}
}

func TestWithdrawHandler(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("coin", "BTC")
	q.Set("units", "50")
	req, err := http.NewRequest(echo.GET, "/withdraw?"+q.Encode(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, withdrawFunds(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.ContainsAny(rec.Body.String(), withdrawJSON))
	}
}

func TestGetWalletHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/wallet")
}
