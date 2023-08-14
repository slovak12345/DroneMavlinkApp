package app

func (a *App) Run() {
	a.log.Info("DroneWebService is starting handle")
	go a.handler.StartEnpointUdpClientHandle("127.0.0.1", "5600", a.cfg.Dialect, 10)
}
