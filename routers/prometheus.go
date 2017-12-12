package routers

import (
    "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/lflxp/ams/controllers"
	"github.com/astaxie/beego"
)

func init() {
	prometheus.MustRegister(controllers.HttpRequestCount)
	prometheus.MustRegister(controllers.HttpRequestDuration)
    beego.Handler("/metrics",promhttp.Handler())
}
