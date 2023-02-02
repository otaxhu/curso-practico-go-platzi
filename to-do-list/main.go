package main

import "fmt"

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
	t.isCompleted = !t.isCompleted
}

func (t *taskList) printTasks() {
	for _, tarea := range t.tasks {
		fmt.Println("Nombre:", tarea.title)
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

func main() {
	t1 := &task{
		title:       "Completar mi curso de Golang",
		description: "Completar el curso de Golang en Platzi en esta semana",
	}
	t2 := &task{
		title:       "Completar mi curso de Python",
		description: "Completar el curso de Python en Platzi en esta semana",
	}
	t3 := &task{
		title:       "Completar mi curso de JavaScript",
		description: "Completar el curso de JavaScript en Platzi en esta semana",
	}
	lista := taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	lista.newTask(t3)
	lista.printTasks()
	lista.tasks[1].completed()
	lista.tasks[1].completed()
	lista.printTasksCompleted()
}
