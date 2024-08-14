package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID     int
	Title  string
	Done   bool
}

var tasks []Task
var nextID int = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Tampilkan Semua Tugas")
		fmt.Println("3. Tandai Tugas Selesai")
		fmt.Println("4. Hapus Tugas")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih opsi: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			addTask(scanner)
		case "2":
			showTasks()
		case "3":
			markTaskDone(scanner)
		case "4":
			deleteTask(scanner)
		case "5":
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Opsi tidak valid.")
		}
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Masukkan judul tugas: ")
	scanner.Scan()
	title := scanner.Text()

	task := Task{
		ID:    nextID,
		Title: title,
		Done:  false,
	}
	nextID++
	tasks = append(tasks, task)

	fmt.Println("Tugas berhasil ditambahkan!")
}

func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("Tidak ada tugas.")
		return
	}

	fmt.Println("Daftar Tugas:")
	for _, task := range tasks {
		status := "Belum Selesai"
		if task.Done {
			status = "Selesai"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Title, status)
	}
}

func markTaskDone(scanner *bufio.Scanner) {
	fmt.Print("Masukkan ID tugas yang ingin ditandai sebagai selesai: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			fmt.Println("Tugas berhasil ditandai sebagai selesai!")
			return
		}
	}

	fmt.Println("Tugas dengan ID tersebut tidak ditemukan.")
}

func deleteTask(scanner *bufio.Scanner) {
	fmt.Print("Masukkan ID tugas yang ingin dihapus: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Tugas berhasil dihapus!")
			return
		}
	}

	fmt.Println("Tugas dengan ID tersebut tidak ditemukan.")
}
