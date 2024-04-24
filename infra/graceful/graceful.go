package graceful

// func GraceFullApiStart(httpServer *http.Server) {
// 	GracefulWrapper(func() {
// 		httpServer.TLSConfig = pice_tls.GetTlsConfigFromEnv()

// 		zap.L().Info("HTTPS " + string(environment.APP_VARIANT) + " Server Starting : " + fmt.Sprint(httpServer.Addr))
// 		if err := httpServer.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
// 			log.Printf("HTTP Server Start Error: %s\n", err)
// 		}
// 	}, func(ctx *context.Context) {
// 		if err := httpServer.Shutdown(*ctx); err != nil {
// 			logger.Error(ctx, "HTTPS Server Shutdown Failed: "+err.Error())
// 		}
// 		slack.PostShutdown()
// 		logger.Info(ctx, "HTTPS Server Shutdown Successfully")
// 	}, 5)
// }

// func GracefulWrapper(target TargetFunc, closeFunc CloseFunc, waitSeconds int) {
// 	ctx := context.Background()

// 	go target()
// 	logger.Info(&ctx, "Graceful Start")

// 	waitForInterruptOrError(&ctx)
// 	logger.Info(&ctx, "Waiting for server to complete existing requests")

// 	if waitSeconds <= 0 {
// 		waitSeconds = GRACEFUL_WAIT_DELAY
// 	}
// 	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(waitSeconds)*time.Second)
// 	defer func() {
// 		cancel()
// 	}()

// 	closeFunc(&timeoutCtx)
// 	logger.Info(&ctx, "Gracefully Stopped")
// }

// func waitForInterruptOrError(ctx *context.Context) {
// 	interruptChan := make(chan os.Signal, 1)
// 	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
// 	manualErrorChan := GetManualErrorChan()

// 	logger.Info(ctx, "Waiting for interrupt...")
// 	select {
// 	case osSignal := <-interruptChan:
// 		logger.Info(ctx, "Interrupt : "+fmt.Sprint(osSignal))
// 	case manualSignal := <-*manualErrorChan:
// 		logger.Info(ctx, "Manual Interrupt : "+fmt.Sprint(manualSignal))
// 	}
// }

// func GetManualErrorChan() *chan string {
// 	if manualErrorChan == nil {
// 		newChan := make(chan string)
// 		manualErrorChan = &newChan
// 	}

// 	return manualErrorChan
// }

// func IsGracefulErrorHandled() bool {
// 	return manualErrorChan != nil
// }

// func GracefulPanic(target func()) {
// 	defer func() {
// 		err := recover()
// 		if err != nil {
// 			if IsGracefulErrorHandled() {
// 				gracefulError := GetManualErrorChan()
// 				*gracefulError <- util.GetMsgFromError(err)
// 			} else {
// 				util.RethrowErr(err)
// 			}
// 		}
// 	}()
// 	target()
// }
