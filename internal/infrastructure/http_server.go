package infrastructure

import (
	"TodoApp_basic/infrastructure/logging"
	"TodoApp_basic/infrastructure/router"
	"TodoApp_basic/infrastructure/storing"
	"TodoApp_basic/routes/logger"
	"database/sql"
	"strconv"
	"time"
)

type Config struct {
	appName       string
	logger        logger.Logger
	sqlStore      *sql.DB
	ctxTimeout    time.Duration
	webServerPort router.Port
	webServer     router.Server
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ContextTimeout(t time.Duration) *Config {
	c.ctxTimeout = t
	return c
}

func (c *Config) AppName(name string) *Config {
	c.appName = name
	return c
}

func (c *Config) Logger(instance int) *Config {
	log, err := logging.NewLoggerFactory(instance)
	if err != nil {
		log.WithError(err)
	}

	c.logger = log
	c.logger.Info("Successfully configured log")
	return c
}

func (c *Config) SqlStore(instance int, dns string, log logger.Logger) *Config {
	db, err := storing.NewStoreFactory(instance, dns, log)
	if err != nil {
		c.logger.Error("Failed to connect to SQL data storing", err)
	}

	c.logger.Info("Successfully connected to the SQL data storing")

	c.sqlStore = db
	return c
}

func (c *Config) WebServerPort(port string) *Config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		c.logger.Error("Error to configuring server port", err)
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *Config) WebServer(instance int) *Config {
	s, err := router.NewWebServerFactory(
		instance,
		c.logger,
		c.sqlStore,
		c.webServerPort,
		c.ctxTimeout,
	)

	if err != nil {
		c.logger.Error("Error to configuring server router", err)
	}

	c.logger.Info("Successfully configured server router")

	c.webServer = s
	return c
}

func (c *Config) Start() {
	c.webServer.Run()
}
