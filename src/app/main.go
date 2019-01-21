package main

import (
	_ "app/routers"
	"github.com/astaxie/beego"
	// OAuth authentication packages
	"github.com/ausrasul/redisorm" // Redis with resource pool
	"github.com/ausrasul/jwt"        // Web token packages
	"github.com/ausrasul/tim"        // TIM packages
	"github.com/ausrasul/m2mserver"	 // M2M server package
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/gplus"
	"log"
	"strconv"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	beego.SetLogFuncCall(true)
	beego.SessionOn = true
	goth.UseProviders(
		gplus.New(
			beego.AppConfig.String("CLIENT_ID"),
			beego.AppConfig.String("CLIENT_SECRET"),
			beego.AppConfig.String("CLIENT_CALLBACK"),
		),
	)
	SessionTimeout, err := beego.AppConfig.Int("SESSION_TIMEOUT")
	if err != nil {
		beego.Critical(err)
	}
	SessionRefreshInterval, err := beego.AppConfig.Int("SESSION_REFRESH_INTERVAL")
	if err != nil {
		beego.Critical(err)
	}

	jwt.Configure(
		map[string]interface{}{
			"privateKeyFile":         beego.AppConfig.String("PrivateKeyFile"),
			"publicKeyFile":          beego.AppConfig.String("PublicKeyFile"),
			"algorithm":              beego.AppConfig.String("Algorithm"),
			"sessionName":            beego.AppConfig.String("SESSION_NAME"),
			"sessionTimeout":         SessionTimeout,
			"sessionRefreshInterval": SessionRefreshInterval,
		},
	)

	tim.Configure(
		map[string]interface{}{
			"ldap_server": beego.AppConfig.String("Ldap_server"),
			"ldap_port":   beego.AppConfig.String("Ldap_port"),
			"base_dn":     beego.AppConfig.String("Base_dn"),
			"ldap_user":   beego.AppConfig.String("Ldap_user"),
			"ldap_pass":   beego.AppConfig.String("Ldap_pass"),
		},
	)

	poolMaxIdle, err := beego.AppConfig.Int("REDIS_MaxIdle")
	if err != nil {
		beego.Critical(err)
	}
	poolMaxActive, err := beego.AppConfig.Int("REDIS_MaxActive")
	if err != nil {
		beego.Critical(err)
	}

	redisorm.Configure(
		map[string]interface{}{
			"poolMaxIdle":   poolMaxIdle,
			"poolMaxActive": poolMaxActive,
			"port":          beego.AppConfig.String("REDIS_Port"),
		},
	)
	
	m2mConnectionTimeout, err := beego.AppConfig.Int("ComServTimeout")
	if err != nil{
		beego.Critical("Com server timeout is not configured: ", err)
		panic(err)
	}
	
	m2mserver.Configure(
		map[string]interface{}{
			"connectionTimeout":   m2mConnectionTimeout,
			"port": beego.AppConfig.String("ComServPort"),
		},
	)
	
	go m2mserver.Listen()
	go func(){
		
		
		
		for i:= 0; i<10000000; i++{
			c, err := m2mserver.GetClient("000c29620510")
			if err != nil{
				beego.Debug("no client")
				time.Sleep(time.Second)	
				continue
			}
			if !c.HasHandler("test_ack"){
				c.AddHandler("test_ack", Test)
			}
			if !c.IsActive(){
				beego.Debug("client not active")
				time.Sleep(time.Second)	
				continue
			}
			
			c.SendCmd(m2mserver.Cmd{"test", "test " + strconv.Itoa(i)})
		}
	}()
	
	beego.SetStaticPath("/public", "static")
	beego.Run()
}

func Test(c *m2mserver.Client, param string){
	beego.Debug("Tested!!", param, "--", c.IsActive())
}