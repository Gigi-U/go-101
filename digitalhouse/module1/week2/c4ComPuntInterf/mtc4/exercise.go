package mtc4
/*
 Productos
Algunas tiendas de e-commerce necesitan realizar una funcionalidad en Go para
administrar productos y RETORNAR el VALOR delPRECIO TOTAL. 
La empresa tiene tres tipos de productos: Pequeño, Mediano y Grande. Pero se espera que sean muchos más. 👉🏼(INTERFAZ)

Los costos adicionales para cada uno son:
● Pequeño: solo tiene el costo del producto. 👉🏼 return precio
● Mediano: el precio del producto + un 3% por mantenerlo en la tienda.👉🏼 precio = precio + (precio*.03)     return precio
● Grande: el precio del producto + un 6% por mantenerlo en la tienda y un adicional
de $2500 de costo de envío. 👉🏼 precio = precio + (precio*.03) + 2500   return precio

El porcentaje de mantenerlo en la tienda se calcula sobre el precio del producto.

Se requiere una función factory que reciba el tipo de producto y el precio, y retorne una
interfaz Producto que tenga el método Precio. 

Se debe poder ejecutar el método Precio y que el método devuelva el precio total
basándose en el costo del producto y los adicionales, en caso de que los tenga. 

👉🏼 INTERFAZ + METODO con switch interno? 
*/