package main

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type OldPrinter struct{}

func (p *OldPrinter) Print(s string) string {
	return fmt.Sprintf("Old Printer: %s", s)
}

type ModernPrinter interface {
	PrintFormatted(s string) string
}

type NewPrinter struct{}

func (p *NewPrinter) PrintFormatted(s string) string {
	return fmt.Sprintf("New Printer: %s", s)
}

type PrinterAdapter struct {
	LegacyPrinter LegacyPrinter
}

func (p *PrinterAdapter) PrintFormatted(s string) string {
	return p.LegacyPrinter.Print(s)
}

func task_21() {
	oldPrinter := &OldPrinter{}
	adapter := &PrinterAdapter{LegacyPrinter: oldPrinter}
	result := adapter.PrintFormatted("Hello, World!")
	fmt.Println(result)
}
