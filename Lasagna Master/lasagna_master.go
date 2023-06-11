package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int{
    layersPrepTime := 0
    if(time==0){
        time = 2
    }
	layersPrepTime = time * len(layers)
    return layersPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodleAmnt int, sauceAmnt float64){
    noodleAmnt = 0
    sauceAmnt = 0
    for i:=0; i<len(layers); i++{
        if(layers[i] == "noodles"){
            noodleAmnt += 50;
        }
    	if(layers[i] == "sauce"){
            sauceAmnt += 0.2
        }
    }
	return noodleAmnt, sauceAmnt
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string){
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(portionAmnt []float64, portionNmb int) []float64{
    newPortion := make([]float64, len(portionAmnt))
	for i := 0; i < len(portionAmnt); i++ {
		newPortion[i] = portionAmnt[i] * float64(portionNmb) / 2
	}
	return newPortion
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
