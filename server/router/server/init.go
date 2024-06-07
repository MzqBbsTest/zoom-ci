package server

var manage Manage

func init() {
	manage.serMap = map[int]*client{}
}
