# CryptoCompare Go

Keep track of you're cryptocurrency investments and check the prices of the top 10 in the crypto world.

## Prerequisites

You will need to install the following:

* Postgres

### Getting started

Clone repo to your machine

```
https://github.com/pocockn/crypto-compare-go.git
```

To setup on the database on Mac OS X

```
psql --dbname=postgres -f create_database.sql
```

For Linux users run

```
sudo su postgres -c psql < create_database.sql
```


## Running the tests

```
go test ./...
```

## CI

When the code is pushed to Github all tests are run automatically by Circle CI

## TODO:

* Add a config file for the database credentials
* Get dynamic price of coins within wallet
* Get total worth of coins within wallet


## Built With

* [GoLang](https://golang.org/)
* [Echo](https://github.com/labstack/echo)
* [Testify](https://github.com/stretchr/testify)
* [CryptoCompare API](http://cryptocompare.com/)
* [Coin Market Cap API](https://coinmarketcap.com/)
