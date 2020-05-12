/*
****SERVER*****
service := "0.0.0.0:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)

		var player models.Player
		decoder.Decode(&player)

		player.UpdateExp(10)

		encoder.Encode(player)
		conn.Close() // we're finished
	}
*/
package main
