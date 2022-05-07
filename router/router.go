package router

import (
	"github.com/gin-gonic/gin"
	"lj-chain-center/handler"
	"lj-chain-center/pkg"
	"lj-chain-center/router/middle"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middle.NoCache)
	g.Use(middle.Options)
	g.Use(middle.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "the incorrect api route")
	})

	check := g.Group("gcc/api/check")
	{
		check.GET("/health", handler.Check_Health)
	}

	coreBizAuth := g.Group("gcc/api/core/biz/auth")
	coreBizAuth.Use(middle.Auth(false, pkg.CORE_TYPE))
	{
		coreBizAuth.POST("/createCompanyDID", handler.Core_CreateCompanyDID)
		coreBizAuth.POST("/createPersonDID", handler.Core_CreatePersonDID)
		coreBizAuth.POST("/queryCompanyDIDList", handler.Core_QueryCompanyDIDList)
		coreBizAuth.POST("/queryPersonDIDList", handler.Core_QueryPersonDIDList)
	}

	antBassBizCommon := g.Group("gcc/api/antBass/biz/common")
	{
		antBassBizCommon.GET("/queryTx", handler.AntBass_QueryTx)
	}

	lubanBizCommon := g.Group("gcc/api/luban/biz/common")
	{
		lubanBizCommon.GET("/queryTx", handler.Luban_QueryTx)
	}

	lubanBizAuth := g.Group("gcc/api/luban/biz/auth")
	lubanBizAuth.Use(middle.Auth(true, pkg.LUBAN_CHAIN_TYPE))
	{
		lubanBizAuth.POST("/depositData", handler.Luban_DepositData)
	}

	return g
}
