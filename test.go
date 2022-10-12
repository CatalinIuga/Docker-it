package main

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	a := app.New()
	w := a.NewWindow("GoDocker Containers")
	w.Resize(fyne.NewSize(600, 500))
	var data = [][]string{{"Containter ID", "Container IMAGE", "Container STATUS & UPTIME"}}
	for _, container := range containers {
		temp := make([]string, 0)
		temp = append(temp, container.ID[:10])
		temp = append(temp, container.Image)
		temp = append(temp, container.Status)
		data = append(data, temp)
		fmt.Println(container.ID[:10] + " " + container.Image)
	}

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			tab := widget.NewLabel("wide content")
			return tab
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})
	for i := 0; i < len(data); i++ {
		list.SetColumnWidth(i, 200)
	}

	w.SetContent(list)
	w.ShowAndRun()
}
