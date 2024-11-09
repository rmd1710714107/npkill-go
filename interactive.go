package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/schollz/progressbar/v3"
	"log"
	"os"
)

func interactive(dirList []FileInfo) {
	var (
		selectedDir []string
	)
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().Title(fmt.Sprintf("total size is %s", transferUnit(TotalSize))),
			// Let the user select multiple toppings.
			huh.NewMultiSelect[string]().
				Title("Dir List").
				Limit(4). // there’s a 4 topping limit!
				Value(&selectedDir).
				OptionsFunc(func() []huh.Option[string] {
					var options []huh.Option[string]
					for i := 0; i < len(dirList); i++ {
						dir := dirList[i]
						key := fmt.Sprintf("%s\t\t\t\t%s", dir.Path, dir.SizeLabel)
						options = append(options, huh.NewOption(key, dir.Path))
					}
					return options
				}, &dirList),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	bar := progressbar.NewOptions(len(selectedDir), progressbar.OptionSetPredictTime(true))
	for _, dir := range selectedDir {
		go deleteDir(dir, bar)
	}

}
func deleteDir(path string, bar *progressbar.ProgressBar) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatalln(err)
	}
	barErr := bar.Add(1)
	if barErr != nil {
		log.Fatalln(barErr)
	}
}
