```go
package main

import (
	"log/slog"
	"os"

	ucsm "github.com/cloud104/tks-go-ucsm-sdk/pkg"
)

func main() {
	handler := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug}))
	slog.SetDefault(handler)
	endPoint := "https://ucsm01.example.com/"
	username := "admin"
	password := "secret"
	client, err := ucsm.AaaLogin(endPoint, username, password)
	client.SetDebug(false)
	if err != nil {
		slog.Error("something failed", "err", err)
		return
	}
	defer client.AaaLogout() //nolint:errcheck
	blades, err := ucsm.GetAllComputeBlades(client)
	if err != nil {
		slog.Error("Erro ao obter blades", "err", err)
		return
	}

	if len(blades) == 0 {
		slog.Info("Nenhuma blade encontrada.")
		return
	}

	for _, blade := range blades {
		slog.Info("---------------")
		slog.Info("Location:", slog.String("Localtion", blade.Dn))
		slog.Info("Model:", slog.String("Model", blade.Model))
		slog.Info("Serial:", slog.String("Serial", blade.Serial))
		slog.Info("UUID:", slog.String("UUID", blade.UUID))
		slog.Info("PartNumber:", slog.String("PartNumber", blade.PartNumber))
		slog.Info("Name:", slog.String("PartNumber", blade.Name))
		slog.Info("UserLabel:", slog.String("PartNumber", blade.UserLabel))
		slog.Info("Number of CPUs:", slog.Int("PartNumber", blade.NumOfCpus))
		slog.Info("Number of Cores:", slog.Int("PartNumber", blade.NumOfCores))
		slog.Info("Total Memory:", slog.Int("PartNumber", blade.TotalMemory))
	}
}
```