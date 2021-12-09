package main

import (
	"bookstore/application"
	"bookstore/http"
	"bookstore/inmem"
	"context"
	"fmt"
	"os"
	"os/signal"
)

type WebApiMain struct {
	Config     Config
	DB         *inmem.InMemRepository
	HTTPServer *http.Server
}

func NewMain() *WebApiMain {
	config := DefaultConfig()
	inmemRepository := inmem.NewInMemoryRepository(config.DB.DSN)

	return &WebApiMain{
		Config:     config,
		DB:         inmemRepository,
		HTTPServer: http.NewServer(),
	}
}

func (m *WebApiMain) Close() error {
	if m.HTTPServer != nil {
		if err := m.HTTPServer.Close(); err != nil {
			return err
		}
	}
	if m.DB != nil {
		if err := m.DB.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (m *WebApiMain) Run(ctx context.Context) (err error) {
	if err := m.DB.Open(); err != nil {
		return fmt.Errorf("cannot open db: %w", err)
	}

	m.HTTPServer.BookService = application.NewBookService(m.DB)
	m.HTTPServer.InventoryService = application.NewInventoryService(m.DB, m.DB)

	m.HTTPServer.Addr = m.Config.HTTP.Addr
	m.HTTPServer.Domain = m.Config.HTTP.Domain

	if err := m.HTTPServer.Open(); err != nil {
		return err
	}

	return nil
}

func getContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()

	return ctx
}
