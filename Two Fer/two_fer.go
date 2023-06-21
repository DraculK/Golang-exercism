package twofer

func ShareWith(name string) string {
    cookieStr := ""
    if(name != ""){
        cookieStr = "One for " +name+ ", one for me."
    }else{
    	cookieStr = "One for you, one for me."
    }
	return cookieStr
}
