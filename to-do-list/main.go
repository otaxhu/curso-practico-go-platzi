package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type task struct {
	title       string
	description string
	isCompleted bool
}

type taskList struct {
	tasks []*task
}

func (t *taskList) newTask(nt *task) {
	t.tasks = append(t.tasks, nt)
}

func (t *taskList) deleteTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *task) completed() {
	t.isCompleted = true
}

func (t *taskList) printTasks() {
	for i, tarea := range t.tasks {
		fmt.Println("(", i, ")", "Nombre:", tarea.title)
		fmt.Println("Descripcion:", tarea.description)
	}
}

func (t *taskList) printTasksCompleted() {
	for _, tarea := range t.tasks {
		if tarea.isCompleted {
			fmt.Println("Tareas completadas")
			fmt.Println("Nombre:", tarea.title)
			fmt.Println("Descripcion:", tarea.description)
		}
	}
}

func (t *task) updateDescription(newDescription string) {
	t.description = newDescription
}

func (t *task) updateTitle(newTitle string) {
	t.title = newTitle
}
func scanInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()

}

func main() {
	fmt.Println("Ingrese uno de los siguientes comandos\n\"nt\" para nueva tarea\n\n\"dt\" para borrar tarea\n\n\"ut\" para actualizar titulo\n\n\"ud\" para actualizar descripcion\n\n\"c\" para marcar completada una tarea\n\n\"pt\" para imprimir tareas\n\n\"ptc\" para imprimir tareas completadas")
	lista := &taskList{}
	command := scanInput()
	switch command {
	case "nt":
		fmt.Println("Ingrese el titulo de la nueva tarea")
		titulo := scanInput()
		fmt.Println("Ingrese la descripcion")
		descripcion := scanInput()
		tarea := &task{title: titulo,
			description: descripcion}
		lista.newTask(tarea)
		lista.printTasks()
	case "dt":
		lista.printTasks()
		fmt.Println("elije una de las tareas para eliminar segun su indice")
		index := scanInput()
		indexLimpio, _ := strconv.Atoi(index)
		lista.deleteTask(indexLimpio)
	case "ut":
		lista.printTasks()
		fmt.Println("elije una de las tareas para actualizar su titulo ingrese el indice")
		index := scanInput()
		indexLimpio, _ := strconv.Atoi(index)
		fmt.Println("escriba ahora el titulo nuevo")
		newTitle := scanInput()
		lista.tasks[indexLimpio].updateTitle(newTitle)
	case "ud":
		lista.printTasks()
		fmt.Println("elije una de las tareas para actualizar su descripcion ingrese el indice")
		index := scanInput()
		indexLimpio, _ := strconv.Atoi(index)
		fmt.Println("ingrese la descripcion nueva")
		newDescription := scanInput()
		lista.tasks[indexLimpio].updateDescription(newDescription)
	case "c":
		lista.printTasks()
		fmt.Println("elije una de las tareas para marcarlas como completadas ingrese el indice")
		index := scanInput()
		indexLimpio, _ := strconv.Atoi(index)
		lista.tasks[indexLimpio].completed()
		lista.printTasksCompleted()
	case "pt":
		lista.printTasks()
	case "ptc":
		lista.printTasksCompleted()
	default:
		fmt.Println("comando erroneo")
	}
}
