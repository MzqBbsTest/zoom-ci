package server

var manage Manage

func init() {
	manage = Manage{
		serMap: map[int]*client{},
	}
}
