package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struktur tugas
type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

var tasks []Task
const fileName = "tasks.json"

// Load data dari file JSON
func loadTasks() {
	file, err := os.ReadFile(fileName)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
}

// Simpan data ke file JSON
func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

// Menampilkan semua tugas
func listTasks() {
	fmt.Println("\n📋 To-Do List:")
	if len(tasks) == 0 {
		fmt.Println("   (Belum ada tugas)")
		return
	}
	for _, task := range tasks {
		status := "❌"
		if task.Complete {
			status = "✅"
		}
		fmt.Printf("   [%s] %d. %s\n", status, task.ID, task.Name)
	}
}

// Menambahkan tugas baru
func addTask(name string) {
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	tasks = append(tasks, Task{ID: id, Name: name, Complete: false})
	saveTasks()
	fmt.Println("✔️  Tugas berhasil ditambahkan!")
}

// Menghapus tugas berdasarkan ID
func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			fmt.Println("🗑️  Tugas berhasil dihapus!")
			return
		}
	}
	fmt.Println("❌ Tugas tidak ditemukan.")
}

// Menandai tugas sebagai selesai
func completeTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Complete = true
			saveTasks()
			fmt.Println("✅ Tugas selesai!")
			return
		}
	}
	fmt.Println("❌ Tugas tidak ditemukan.")
}

// Menjalankan CLI
func main() {
	loadTasks()
	for {
		fmt.Print("\nMasukkan perintah (add/list/done/del/exit): ")
		var input string
		fmt.Scanln(&input)

		switch strings.ToLower(input) {
		case "add":
			fmt.Print("Masukkan nama tugas: ")
			var name string
			fmt.Scanln(&name)
			addTask(name)

		case "list":
			listTasks()

		case "done":
			fmt.Print("Masukkan ID tugas: ")
			var idStr string
			fmt.Scanln(&idStr)
			id, _ := strconv.Atoi(idStr)
			completeTask(id)

		case "del":
			fmt.Print("Masukkan ID tugas: ")
			var idStr string
			fmt.Scanln(&idStr)
			id, _ := strconv.Atoi(idStr)
			deleteTask(id)

		case "exit":
			fmt.Println("👋 Bye!")
			return

		default:
			fmt.Println("❌ Perintah tidak dikenali.")
		}
	}
}
