package raindrops
import "strconv"

func Convert(number int) string {
	numb := ""
    if(number%3 == 0){
        numb+="Pling"
    }
    if(number%5 == 0){
        numb+="Plang"
    }
    if(number%7 == 0){
        numb+="Plong"
    }
    if(numb == ""){
    	numb = strconv.Itoa(number)
    }
	return numb
}
