package main

import (
	CyberCore "github.com/MrmaderatorYT/CyberCore/server"
	"github.com/Tnze/go-mc/net"
	//server "github.com/Tnze/go-mc/server"
)

func acceptConnection(conn net.Conn) {
	defer func(conn *net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(&conn)
	// Читаем пакет-рукопожатие(HandSnake)
	_, nextState, _, _, err := CyberCore.ReadHandSnake(conn)
	// Если при чтении была некая ошибка, то просто перестаём обрабатывать подключение
	if err != nil {
		return
	}

	// Обрабатываем следющее состояние(1 - пинг, 2 - игра)
	switch nextState {
	case 1:
		acceptPing(conn)
	default:
		return
	}
}
