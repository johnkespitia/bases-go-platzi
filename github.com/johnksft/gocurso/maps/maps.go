package maps

//GetMap Obtiene un mapa con llave tipo string y valores tipo enteros
func GetMap(name string) int {
	//mapTest := make(map[string]int)
	mapTest := map[string]int{
		"Juan":   18,
		"Mario":  28,
		"Freddy": 18,
	}
	mapTest["llave1"] = 1
	mapTest["llave2"] = 100
	mapTest["llave3"] = 1000
	delete(mapTest, "llave1")
	delete(mapTest, "llave2")
	delete(mapTest, "llave3")
	return mapTest[name]
}
