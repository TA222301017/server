package setup

import "server/models"

type GlobalChannel struct {
	SSETokens map[uint32]bool

	RSSIMessage       chan *models.RSSILog
	NewRSSIClients    chan chan *models.RSSILog
	ClosedRSSIClients chan chan *models.RSSILog
	TotalRSSIClients  map[chan *models.RSSILog]bool

	AccessMessage       chan *models.AccessLog
	NewAccessClients    chan chan *models.AccessLog
	ClosedAccessClients chan chan *models.AccessLog
	TotalAccessClients  map[chan *models.AccessLog]bool
}

var Channel *GlobalChannel

func (stream *GlobalChannel) Listen() {
	for {
		select {
		case client := <-stream.NewRSSIClients:
			stream.TotalRSSIClients[client] = true

		case client := <-stream.ClosedRSSIClients:
			delete(stream.TotalRSSIClients, client)
			close(client)

		case message := <-stream.RSSIMessage:
			for clientMessageChan := range stream.TotalRSSIClients {
				clientMessageChan <- message
			}

		case client := <-stream.NewAccessClients:
			stream.TotalAccessClients[client] = true

		case client := <-stream.ClosedAccessClients:
			delete(stream.TotalAccessClients, client)
			close(client)

		case message := <-stream.AccessMessage:
			for clientMessageChan := range stream.TotalAccessClients {
				clientMessageChan <- message
			}
		}
	}
}

func ChannelServer() {
	Channel = &GlobalChannel{
		SSETokens: make(map[uint32]bool),

		RSSIMessage:       make(chan *models.RSSILog),
		NewRSSIClients:    make(chan chan *models.RSSILog),
		ClosedRSSIClients: make(chan chan *models.RSSILog),
		TotalRSSIClients:  make(map[chan *models.RSSILog]bool),

		AccessMessage:       make(chan *models.AccessLog),
		NewAccessClients:    make(chan chan *models.AccessLog),
		ClosedAccessClients: make(chan chan *models.AccessLog),
		TotalAccessClients:  make(map[chan *models.AccessLog]bool),
	}

	go Channel.Listen()
}
