package httphandler

import (
	"fmt"
	"net/http"

	"github.com/alexcesaro/log/stdlog"
	// _ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"

	Conf "orders/config"
	Model "orders/model"

	logger "github.com/sirupsen/logrus"
)

func GetWallet(c *gin.Context) {
	logkoe := stdlog.GetFromFlags()

	env, errenv := Conf.Environment()
	if errenv != nil {
		logger.Println(errenv)
		logkoe.Info(errenv)
	} else {
		if env == "production" {
			gin.SetMode(gin.ReleaseMode)
			// router := gin.New()
		} else if env == "development" {
			gin.SetMode(gin.DebugMode)
		}
	}
	dbwallet := Conf.Init()

	rows, err := dbwallet.Query("select id, user_id, balance, status, created_at, updated_at, is_active, currency_id from wallets")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var wallets []Model.Wallet
	for rows.Next() {
		var wlt Model.Wallet
		err = rows.Scan(&wlt.ID, &wlt.UserID, &wlt.Balance, &wlt.Status, &wlt.CreatedAt, &wlt.UpdatedAt, &wlt.IsActive, &wlt.CurrencyID)
		if err != nil {
			panic(err)
		}
		wallets = append(wallets, wlt)
	}

	if wallets == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": wallets,
		})
	}
}

func GetWalletByUserId(c *gin.Context) {
	logkoe := stdlog.GetFromFlags()

	env, errenv := Conf.Environment()
	if errenv != nil {
		logger.Println(errenv)
		logkoe.Info(errenv)
	} else {
		if env == "production" {
			gin.SetMode(gin.ReleaseMode)
			// router := gin.New()
		} else if env == "development" {
			gin.SetMode(gin.DebugMode)
		}
	}

	var walletUserID Model.WalletByUserID

	err := c.ShouldBind(&walletUserID)
	if err != nil {
		logkoe.Info(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userID string = walletUserID.UserID
	fmt.Println(userID)
	dbwallet := Conf.Init()
	var wlt Model.Wallet

	var query = dbwallet.QueryRow("select id, user_id, balance, status, created_at, updated_at, is_active, currency_id from wallets where user_id = $1", userID).Scan(&wlt.ID, &wlt.UserID, &wlt.Balance, &wlt.Status, &wlt.CreatedAt, &wlt.UpdatedAt, &wlt.IsActive, &wlt.CurrencyID)
	if query != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": wlt,
		})
	}

}
