package main

import (
	"reflect"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)
func clearResults (resultsPanel *tview.Flex) {
	
	for i := 0; i < resultsPanel.GetItemCount(); i++ {
		if i != 0 {
			resultsPanel.RemoveItem(resultsPanel.GetItem(i))
		}
	}
}
func showResults (resultsPanel *tview.Flex,results []string) {
	clearResults(resultsPanel)
	for _,result := range results {
		resultsPanel.AddItem(tview.NewTextView().SetText(result).SetTextAlign(tview.AlignCenter), 0, 1, false)
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
	AddInputField("Name", "", 60, nil, nil).
	AddInputField("Sku", "", 60, nil, nil).
	AddInputField("Price", "", 60, nil, nil).
	AddInputField("Cost", "", 60, nil, nil).
	AddButton("Save", func (){
		showResults(resultsPanel,[]string{"Product Saved"})
	})
	return tview.NewFlex().AddItem(form, 0, 1, true)
}
func clearMain (main *tview.Flex,title string) {
	main.Clear()
	main.AddItem(tview.NewTextView().
		SetText(title).
		SetTextAlign(tview.AlignCenter), 0, 1, false)
}
func handler (selection string,main *tview.Flex,resultsPanel *tview.Flex) {
	clearMain(main,selection)
	clearResults(resultsPanel)
	switch selection {
		case "Productos":	
			main.AddItem(productForm(resultsPanel), 0, 4, false)
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
	menu.AddItem(list, 0, 4, true)
}
func saveForm (main *tview.Flex){
	for i := 0; i < main.GetItemCount(); i++ {
		item:= main.GetItem(i)
		if(reflect.TypeOf(item)==reflect.TypeOf(tview.NewFlex())) {
			item := item.(*tview.Flex)
			for j := 0; j < item.GetItemCount(); j++ {
				if(reflect.TypeOf(item.GetItem(j))==reflect.TypeOf(tview.NewForm())) {
					form := item.GetItem(j).(*tview.Form)
					//activate the save button and focus form
					form.GetButton(0).InputHandler()
				}
			}
		}
		
		
	}
}
func console () {
	newFlex:= func(text string) *tview.Flex {
		flex:= tview.NewFlex().
		SetDirection(tview.FlexRow)
		flex.AddItem(tview.NewTextView().
		SetText(text).
		SetTextAlign(tview.AlignCenter), 0, 1, false)

		return flex
	}
	menu := newFlex("Menu")
	main := newFlex("App")
	sideBar := newFlex("Results")
	configureMenu(menu,main,sideBar)
	

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
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {	
		switch event.Rune() {
			case 'm':
				app.SetFocus(menu)
			case rune(tcell.KeyCtrlS):
				saveForm(main)
			default:
				return event
		}
		return event
	})
	if err := app.Run(); err != nil {
		panic(err)
	}
	
}
