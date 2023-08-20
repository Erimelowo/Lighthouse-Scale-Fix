package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type DeviceScale struct {
	Scale       float64     `json:"scale"`
	ModelPoints [][]float64 `json:"modelPoints"`
}

func main() {
	// TODO get SteamVR install path (not hardcoded)?
	var jsonPath string = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\SteamVR\\drivers\\lighthouse\\resources\\lighthouse_scale.json"

	// Waits for the illusion of the program doing hard work
	time.Sleep(666 * time.Millisecond)

	// Load the JSON file into memory
	jsonFile, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println("Error reading lighthouse_scale.json:", err)
		time.Sleep(3 * time.Second)
		fmt.Println("Exiting...")
		time.Sleep(2 * time.Second)
		return
	}

	// Make a backup (.bkp)
	if err := os.WriteFile(jsonPath + ".bkp", jsonFile, 0644); err != nil {
		fmt.Println("Error creating backup of lighthouse_scale.json:", err)
		time.Sleep(3 * time.Second)
		fmt.Println("Exiting...")
		time.Sleep(2 * time.Second)
		return
	}

	// Unmarshal the JSON data into a map
	var jsonData map[string]DeviceScale
	if err := json.Unmarshal(jsonFile, &jsonData); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		time.Sleep(3 * time.Second)
		fmt.Println("Exiting...")
		time.Sleep(2 * time.Second)
		return
	}

	// Update the scale values
	for key := range jsonData {
		jsonData[key] = DeviceScale{Scale: 1, ModelPoints: jsonData[key].ModelPoints}
	}

	// Marshal the updated data back to JSON
	updatedJSON, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		time.Sleep(3 * time.Second)
		fmt.Println("Exiting...")
		time.Sleep(2 * time.Second)
		return
	}

	// Write the updated JSON to a file
	if err := os.WriteFile(jsonPath, updatedJSON, 0644); err != nil {
		fmt.Println("Error writing updated lighthouse_scale.json:", err)
		time.Sleep(3 * time.Second)
		fmt.Println("Exiting...")
		time.Sleep(2 * time.Second)
		return
	}

	fmt.Println("lighthouse_scale.json file updated successfully!")
	time.Sleep(2 * time.Second)
	fmt.Println("Exiting...")
	time.Sleep(1 * time.Second)
}
