package main

import (
	"fmt"
	"path/filepath"
)

func previewHandler(data *Data) error {
	for _, dest := range data.cfg.Destinations {
		fmt.Printf("Destination name: %s\n", dest.Name)
		fmt.Printf("Destination path: %s\n", dest.Path)
		fmt.Println("Preview:")
		for _, link := range dest.Links {
			destFullPath := filepath.Join(dest.Path, link)
			fmt.Printf("%s -> %s\n", link, destFullPath)
		}
		fmt.Println()
	}
	return nil
}
