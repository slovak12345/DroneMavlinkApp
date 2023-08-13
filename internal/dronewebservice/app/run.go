package app

func (a *App) Run() {
	a.log.Info("DroneWebService is starting handle")
	a.handler.StartHandle(a.cfg.HTTPHost, a.cfg.HTTPPort)
}
