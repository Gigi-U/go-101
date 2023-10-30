WAITGROUPS


sync.WaitGroup es una estructura de sincronización proporcionada por el paquete sync en Go. Su función principal es permitir que las goroutines esperen hasta que un grupo de otras goroutines haya completado su trabajo. En esencia, sync.WaitGroup proporciona una forma de coordinar la ejecución de múltiples goroutines y garantizar que ninguna de ellas continúe hasta que todas las demás hayan terminado.

Un sync.WaitGroup mantiene un contador interno que comienza en cero. Puedes agregar goroutines al grupo utilizando el método Add(n) y luego marcar que una goroutine ha terminado su trabajo llamando al método Done(). Por último, puedes esperar a que todas las goroutines hayan terminado utilizando el método Wait().

sync.WaitGroup es una herramienta útil para coordinar goroutines y asegurarse de que todas terminen antes de continuar o realizar alguna tarea específica en el programa principal. Se utiliza comúnmente en situaciones en las que se necesita esperar a que un número desconocido de goroutines termine su trabajo, como en la realización de tareas concurrentes y su posterior sincronización.