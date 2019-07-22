package main

import (
	"context"

	"bitbucket.org/snapmartinc/logger"
	"bitbucket.org/snapmartinc/trace"
)

func main() {
	// Create loggerFactory with level Error above
	loggerFactory := logger.NewLoggerFactory(logger.DebugLevel)

	// fake context,
	// in real cases, it should be current app's context
	ctx := context.TODO()

	// must use userclient.User
	user := userclient.User{
		Id: "2",
	}

	ctx = userclient.ContextWithUser(ctx, &user)
	ctx = trace.ContextWithRequestID(ctx, "123455")

	// withField will be filled into "content" field
	log := loggerFactory.Logger(ctx).WithField("test key", "test")

	log.WithFields(logger.Fields{
		"test1": "key 1",
		"test2": "key 2",
		"obj": logger.Fields{
			"test": "test",
		},
	}).Alert("Test")

	// simple log, no content field
	log.WithField("TestKey 2", "vae").Debug("Testing")

	log.Critical("Testing")
	log.Error("Testing")
	log.Warning("Testing")
	log.Notice("Testing")

	// Global logger
	logger.Debug("global debug")
	logger.Info("global info")
	logger.Error("global error")
	logger.Notice("global notice")
	logger.Emergency("global emergency")
	logger.Warning("global warning")
	logger.Alert("global alert")
	logger.Critical("global critical")

	// log with format
	logger.Infof("%d, %s", 10, "a")
}
