Concurrencia:

La concurrencia en Go se refiere a la capacidad de un programa para manejar múltiples tareas o procesos de manera simultánea y eficiente.
En Go, la concurrencia se logra mediante goroutines, que son funciones ligeras que se ejecutan de forma concurrente.
Las goroutines son administradas por el runtime de Go y pueden ejecutarse en UN SOLO HILO (goroutines multiplexadas) en un proceso de manera secuencial, pero el runtime gestiona la alternancia entre ellas para lograr la concurrencia.
La concurrencia es útil cuando se deben realizar tareas que pueden bloquearse, como lectura/escritura de archivos o llamadas a API de red, ya que permite que otras goroutines continúen su ejecución mientras una está bloqueada.

Paralelismo:

Implica la ejecución simultánea real de múltiples tareas en diferentes núcleos de CPU o máquinas físicas.
En Go, el paralelismo se logra mediante la ejecución de múltiples goroutines en DIFERENTES hilos o núcleos de CPU. Esto se puede lograr utilizando la palabra clave "go" para iniciar goroutines y permitir que el runtime de Go las distribuya en paralelo.
Es beneficioso cuando se tienen tareas que pueden ejecutarse de manera independiente y se desea aprovechar completamente el potencial de múltiples núcleos de CPU.

EN CRIOLLO:
concurrencia = multitasking de tareas
paralelismo = delegar tareas a otros para que todas se realicen al mismo tiempo

------CHANNELS------
 - portal o canal de comunicación entre diferentes goroutines
 - a veces es más fácil pasar punteros y waitgroups para pasar info , en vez de canales
 - hay que evitar CONDICIONES DE CARRERA, xq son difíciles de predecir (

Una condición de carrera en un canal (channel) en Go ocurre cuando dos o más goroutines compiten para acceder y manipular los datos dentro del canal sin una sincronización adecuada. Esto puede llevar a resultados inesperados o incorrectos en la comunicación y el intercambio de datos entre goroutines. Las condiciones de carrera son un tipo común de error en la programación concurrente.

En Go, los canales son una forma segura de comunicación entre goroutines, ya que están diseñados para evitar condiciones de carrera y proporcionan una sincronización intrínseca. Sin embargo, las condiciones de carrera pueden ocurrir si no se utilizan apropiadamente.

Para evitar condiciones de carrera en canales, se pueden utilizar constructores de sincronización como sync.WaitGroup o selectores de canales (select) para coordinar y sincronizar las goroutines de manera adecuada.
)
------GOROUTINES------
- Hilo ligero administrados po Go en tiempo de ejecución
- múltiples GOROUTINES pueden ejecutarse en un mismo hilo
- al ser hilos ligeros, son más baratos que los hilos.
- La cantidad puede crecer o decrecer segun se necesite
- pueden comunicarse entre sí usando channels

time.Sleep() - para que una funcion espere x x tiempo
fmt.Scan() - al final del main para esperar a que el usuario aprete cualquier tecla para para poder terminar.
channel -  hacer que las Goroutines se comuniquen entre sí para avisar cuando realizaron su tarea.

REGLAS DE ORO:
1) siempre saber tiempo de vida de una goroutine (cuándo y cómo va a terminar)
2) mantener el ciclo de vida de la goroutine en un único bloque



waitgroup vs mutex

MUTEX. se usa cuando un recurso es usado x 2 subrutinas para q no se cree una carrera
se llama desde el paquete sync.
mutex.Lock()
mutex.Unlock()