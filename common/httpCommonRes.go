package common

type CommonResult struct {
	Res      int
	Mes      string
	PageInfo struct {
		Page       int
		PageNum    int
		TotalCount int64
	}
	Data interface{}
}
