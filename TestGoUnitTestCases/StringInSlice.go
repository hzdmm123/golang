package TestGoUnitTestCases

func StringInSlice(a string,list []string) bool  {
	for _,b := range list{
		if a==b {
			return true
		}
	}
	return false
}