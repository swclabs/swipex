// Author:
// - Ho Duc Hung : @kieranhoo
// - Nguyen Van Khoa: @anthony2704
// This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

package main

import (
	"fmt"
	"log"

	"swclabs/swipecore/boot"
	"swclabs/swipecore/boot/adapter"
	"swclabs/swipecore/internal/config"
)

func main() {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	server := boot.NewServer(addr)
	adapter := adapter.New(adapter.TypeBase)

	if err := server.Connect(adapter); err != nil {
		log.Fatal(err)
	}
}
