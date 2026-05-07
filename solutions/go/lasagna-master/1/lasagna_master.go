package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) (totalPrep int){
    if avg == 0{
        avg = 2
    }
    totalPrep = len(layers) * avg
    return
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodleQuant int , sauceQuant float64) {
    
    for _ , layer := range layers {
        if layer == "sauce"{sauceQuant +=0.2}
        if layer == "noodles"{noodleQuant += 50}
    }
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string){
    myList[len(myList)-1] =  friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (desiredQuantity []float64){
    desiredQuantity = append([]float64{},  quantities...)
    for i := range desiredQuantity {
    desiredQuantity[i] /= 2
    desiredQuantity[i] *= float64(portions)
	}
    return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
