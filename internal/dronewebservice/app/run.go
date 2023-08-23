package app

func (a *App) Run() {
	a.log.Info("DroneWebService is starting handle")
	go a.handler.StartHandle(a.cfg.HTTPHost, a.cfg.HTTPPort)
	// go a.handler.StartEnpointUdpServerHandle("127.0.0.1", "5600", a.cfg.Dialect, 10)
}
