package main

import (
	"fmt"
	"sync"

	"github.com/rivo/tview"
)
func clearResults (resultsPanel *tview.Flex) {
	if(resultsPanel.GetItemCount()>1){
		resultsPanel.RemoveItem(resultsPanel.GetItem(1))
	}
}
func centerForm (form *tview.Form) *tview.Flex {
	return tview.NewFlex().AddItem(nil, 0, 1, false).   
	AddItem(tview.NewFlex().AddItem(form, 0, 1, false) , 0, 1, false).
	AddItem(nil, 0, 1, false)
}
func productForm (resultsPanel *tview.Flex) *tview.Flex {
	form := tview.NewForm()
	form.AddDropDown("Type", []string{"Simple","Virtual","Configurable"}, 0, nil).
	AddInputField("Name", "", 40, nil, nil).
	AddInputField("Sku", "", 40, nil, nil).
	AddInputField("Price", "", 40, nil, nil).
	AddInputField("Cost", "", 40, nil, nil).
	AddButton("Save", func (){
		clearResults(resultsPanel)
		resultsPanel.AddItem(tview.NewTextView().SetText("Saved").SetTextAlign(tview.AlignCenter), 0, 1, false)
	})
	return centerForm(form)
}
func handler (selection string,main *tview.Flex,resultsPanel *tview.Flex) {
	main.GetItem(0).(*tview.TextView).SetText(selection)
	if(main.GetItemCount()>1){
		main.RemoveItem(main.GetItem(1))
	}
	switch selection {
		case "Productos":
			main.AddItem(productForm(resultsPanel), 0, 3, false)
		default:
		
	}
}
func configureMenu (menu *tview.Flex,main *tview.Flex,resultsPanel *tview.Flex) {
	list:= tview.NewList().
		AddItem("Productos", "", 'p', func (){
			handler("Productos",main,resultsPanel)
		}).
		AddItem("Inventario","", 'i', func (){
			handler("Inventario",main,resultsPanel)
		}).
		AddItem("Órdenes","", 'o', func (){
			handler("Órdenes",main,resultsPanel)
		})
	menu.AddItem(list, 0, 1, true)
}
func configureMain (main *tview.Flex) {
	main.AddItem(tview.NewTextView().
		SetText("Selecione un item de la lista para comenzar").
		SetTextAlign(tview.AlignCenter), 0, 0, false)
}
func console () {
	newFlex:= func(text string) *tview.Flex {
		return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().
		SetText(text).
		SetTextAlign(tview.AlignCenter), 0, 1, false)
	}
	menu := newFlex("Menu")
	main := newFlex("App")
	sideBar := newFlex("Results")
	configureMenu(menu,main,sideBar)
	configureMain(main)
	

	grid := tview.NewGrid().
		SetRows(4, 0).
		SetColumns(40, 0, 40).
		SetBorders(true).
		AddItem(tview.NewTextView().SetText("Eworld Terminal").SetTextAlign(tview.AlignCenter), 0, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false).
		AddItem(main, 1, 0, 1, 3, 0, 0, false).
		AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
		AddItem(main, 1, 1, 1, 1, 0, 100, false).
		AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)
	app:=tview.NewApplication().SetRoot(grid, true).EnableMouse(true).SetFocus(grid)
	
	if err := app.Run(); err != nil {
		panic(err)
	}
	
}

func main() {
	// Start the terminal app in a Goroutine
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		console()
	}()
	// Wait for the terminal app to complete
	wg.Wait()
	fmt.Println("Exiting the main function.")
}
