package util

func Paging(page int, pagesize int) (int32, int32) {
	maxPagesize := 1000

	if pagesize > maxPagesize {
		pagesize = maxPagesize
	}

	if page < 0 {
		page = 0
	}

	if page == 0 && pagesize == 0 {
		return 0, int32(maxPagesize)
	}

	return int32(page * pagesize), int32(pagesize)
}
