package main
import "fmt"

type taskList struct {
	tasks []*task  // Slice no Array 
}

func (t *taskList) addToList(ti *task) {
	t.tasks = append(t.tasks, ti)  //toma como parametro un slice y agrega al slice
}

func (t *taskList) deleteToList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index + 1:]...)
}

func (t *taskList)  getList() {
	for _, task := range t.tasks {
		fmt.Println("Nombre", task.name)
		fmt.Println("Descripción", task.description)
	}
}

func (t *taskList)  getListCompleted() {
	for _, task := range t.tasks {
		if task.completed {
			fmt.Println("Nombre", task.name)
			fmt.Println("Descripción", task.description)
		}
	}
}

type task struct {
	name string
	description string
	completed bool
}

func (t *task) markComplete() {
	t.completed = true
}

func (t *task) updateDescription(des string) {
	t.description = des
}

func (t *task) updateName(n string) {
	t.name = n
}

func main() {
	t1 := &task {
		name: "Completar mi curso de Go",
		description: "Completar el curso durante esta semana",
	}

	t2 := &task {
		name: "Llenar parte de actividades",
		description: "Rellenar cada actividad de la semana",
	}

	t3 := &task {
		name: "dormir",
		description: "dormir como un bebe",
	}

	list := &taskList {
		tasks: []*task {
			t1, t2,
		},
	}
	
	list.addToList(t3)
	
	list.getList()
	list.tasks[0].markComplete()
	//fmt.Println("Tareas Completadas")
	list.getListCompleted()

	mapTasks := make(map[string]*taskList)

	mapTasks["Waldo"] = list

	t4 := &task {
		name: "Completar mi curso de Python",
		description: "Completar el curso durante esta semana",
	}

	t5 := &task {
		name: "Completar mi curso de C#",
		description: "Rellenar cada actividad de la semana",
	}

	list2 := &taskList {
		tasks: []*task {
			t4, t5,
		},
	}

	mapTasks["Enzo"] = list2

	fmt.Println("Tareas de Waldo")
	mapTasks["Waldo"].getList()
	fmt.Println("Tareas de Enzo")
	mapTasks["Enzo"].getList()


	/*
	fmt.Println(list.tasks[0])
	list.addToList(t3)
	fmt.Println(len(list.tasks)) //len devuelve toda la longitud del slice
	list.deleteToList(1)
	fmt.Println(len(list.tasks))
	
	fmt.Printf("%+v\n", t)
	t.markComplete()
	t.updateName("Limpiar el baño")
	t.updateDescription("Debe limpiar el baño con todos los desinfectantes adecuados.")
	fmt.Printf("%+v\n", t)
	*/
}