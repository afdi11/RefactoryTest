package main

import (
	"encoding/json"
	"fmt"
)

type Inventory struct {
	Inventory_id int      `json:"inventory_id"`
	Name         string   `json:"name"`
	TypeBarang   string   `json:"type"`
	Tags         []string `json:"tags"`
	Purchased_at int      `json:"purchased_at"`
	Placement    Room     `json:"placement"`
}

type Room struct {
	Room_id int
	Name    string
}

func main() {
	//DATA INPUT
	var jsonString = `[
		{
			"inventory_id": 9382,
			"name": "Brown Chair",
			"type": "furniture",
			"tags": [
				"chair",
				"furniture",
				"brown"
				],
			"purchased_at": 1579190471,
			"placement": {
				"room_id": 3,
				"name": "Meeting Room"
			}
		},
		{
			"inventory_id": 9380,
			"name": "Big Desk",
			"type": "furniture",
			"tags": [
				"desk",
				"furniture",
				"brown"
			],
			"purchased_at": 1579190642,
			"placement": {
				"room_id": 3,
				"name": "Meeting Room"
			}
		},
		{
			"inventory_id": 2932,
			"name": "LG Monitor 50 inch",
			"type": "electronic",
			"tags": [
				"monitor"
			],
			"purchased_at": 1579017842,
			"placement": {
				"room_id": 3,
				"name": "Lavender"
			}
		},
		{
			"inventory_id": 232,
			"name": "Sharp Pendingin Ruangan 2PK",
			"type": "electronic",
			"tags": [
				"ac"
			],
			"purchased_at": 1578931442,
			"placement": {
				"room_id": 5,
				"name": "Dhanapala"
			}
		},
		{
			"inventory_id": 9382,
			"name": "Alat Makan",
			"type": "tableware",
			"tags": [
				"spoon",
				"fork",
				"tableware"
			],
			"purchased_at": 1578672242,
			"placement": {
				"room_id": 10,
				"name": "Rajawali"
			}
		}
	]`
	// fmt.Println(jsonString)
	var data []Inventory

	var err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// fmt.Println(data)
	// fmt.Println("Furnitur 1:", data[0].Name)
	// fmt.Println("Furnitur 2:", data[1].Name)
	// fmt.Println("Furnitur 2:", data[2].Name)

	//1. Find Item in Meeting Room
	fmt.Println("1. Find Item in Meeting Room")
	for _, data := range data {
		if data.Placement.Name == "Meeting Room" {
			fmt.Println(data.Name)
		}
	}
	// 2. Find all electronic devices.
	fmt.Println("2. Find all electronic devices.")
	for _, data := range data {
		if data.TypeBarang == "electronic" {
			fmt.Println(data.Name)
		}
	}
	// 3. Find all the furniture.
	fmt.Println("3. Find all the furniture.")
	for _, data := range data {
		if data.TypeBarang == "furniture" {
			fmt.Println(data.Name)
		}
	}
	// 4. Find all items were purchased on 16 Januari 2020.
	fmt.Println("4. Find all items were purchased on 16 Januari 2020.")
	fmt.Println("data purchased_at tidak dapat dikonversikan ke tanggal")
	// for _, data := range data {
	// 	if data.TypeBarang == "furniture" {
	// 		fmt.Println(data.Name)
	// 	}
	// }
	// 5. Find all items with brown color.
	fmt.Println("5. Find all items with brown color.")
	for _, data := range data {
		for _, tag := range data.Tags {
			if tag == "brown" {
				fmt.Println(data.Name)
			}
		}
	}
}
