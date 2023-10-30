package waitgroups

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1) // Agregar una goroutine al grupo
        go func(i int) {
            defer wg.Done() // Marcar que la goroutine ha terminado
            fmt.Printf("Goroutine %d\n", i)
        }(i)
    }

    // Esperar a que todas las goroutines del grupo terminen
    wg.Wait()
    fmt.Println("Todas las goroutines han terminado.")
}

/* En este ejemplo, creamos un grupo sync.WaitGroup, agregamos cinco goroutines al grupo, cada una de las cuales realiza algún trabajo y llama a Done() para marcar su finalización. Luego, utilizamos Wait() para bloquear la ejecución del programa principal hasta que todas las goroutines en el grupo hayan terminado. */