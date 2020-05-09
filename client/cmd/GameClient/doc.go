// Package main.
/*

	****CLIENT***
	player := models.NewPlayer("Adrianhh")

	service := "0.0.0.0:1200"

	conn, err := net.Dial("tcp", service)
	checkError(err)

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	encoder.Encode(player)

	var newPlayer models.Player
	decoder.Decode(&newPlayer)

	fmt.Println(newPlayer.Exp)

	fmt.Scanln()
*/
package main
