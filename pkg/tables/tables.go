package tables

var (
	App = [][]string{
		{"NAME", "Name"},
		{"IMAGE", "Spec.Image"},
		{"HEALTHY", "Status.Columns.Healthy"},
		{"UPTODATE", "Status.Columns.UpToDate"},
		{"CREATED", "{{ago .Created}}"},
		{"ENDPOINTS", "Status.Columns.Endpoints"},
		{"MESSAGE", "Status.Columns.Message"},
	}
	AppConverter = MustConverter(App)

	Volume = [][]string{
		{"NAME", "Name"},
		{"APPNAME", "Status.AppName"},
		{"BOUNDVOLUME", "Status.VolumeName"},
		{"CAPACITY", "Spec.Capacity"},
		{"STATUS", "Status.Status"},
		{"ACCESSMODES", "Status.Columns.AccessModes"},
		{"CREATED", "{{ago .Created}}"},
		{"MESSAGE", "Status.Message"},
	}
	VolumeConverter = MustConverter(Volume)

	Image = [][]string{
		{"REPOSITORY", "{{if eq .Repository \"\"}}<none>{{else}}{{.Repository}}{{end}}"},
		{"TAG", "{{if eq .Tag \"\"}}<none>{{else}}{{.Tag}}{{end}}"},
		{"IMAGEID", "{{trunc .Name}}"},
	}
	ImageConverter = MustConverter(Image)

	Container = [][]string{
		{"NAME", "Name"},
		{"APP", "Status.Columns.App"},
		{"IMAGE", "Status.Image"},
		{"STATE", "Status.Columns.State"},
		{"RESTARTCOUNT", "Status.RestartCount"},
		{"CREATED", "{{ago .Created}}"},
		{"MESSAGE", "Status.PodMessage"},
	}
	ContainerConverter = MustConverter(Container)
)
