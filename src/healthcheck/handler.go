package healthcheck

func NewHandler(config configPkg.Config) Handler {
	return &theHandler{
		config:         config,
		hriHealthcheck: GetCheck,
	}
}
