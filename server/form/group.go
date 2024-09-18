package form

type FormGroup struct {
	ID      int    `form:"id"`
	Name    string `form:"name"`
	Servers []int  `form:"servers" `
	Offset  int    `form:"offset"`
	Limit   int    `form:"limit,default=20" binding:"gte=1,lte=999" `
}
