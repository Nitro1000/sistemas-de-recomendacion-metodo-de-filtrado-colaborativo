# Grado en Ingeniería Informática
# Gestión del Conocimiento en las Organizaciones
## Sistemas de Recomendación

El programa funciona de la siguiente manera.
Tenemos una serie de parametros que le pasamos al programa entre esos estan:
- name: que corresponde con la ruta del fichero, por defecto el fichero que tomará sera ./tabla.txt.
- metric: sera la metrica elegida. Los posibles valores son:
  - CP (Correlación de Pearson).
  - DC (Distancia coseno).
  - DE (Distancia Euclídea).
  Por defecto la metrica utilizada sera CP (Correlación de Pearson).
- neighbors: es el numero de vecinos considerados, por defecto toma el valor 3.
- prediction: sera el tipo de predicción. Los posibles valores son:
  - PS (Predicción simple).
  - DM (Diferencia con la media).
  Por defecto la predicción utilizada sera PS (Predicción simple).

Una vez ejecutado el programa pasando los parametros o ejecutandolos sin los mismos, se imprimiran los valores de la matriz resultado y los valores de las metricas.

Ejemplos de ejecución del programa:
- Por defecto: 
  ```bash
  go run main.go
  ```
- Con parametros especificos:
  ```bash
  go run main.go -name tabla.txt -metric CP -neighbors 3 -prediction DM
  ```
- Para ver la ayuda:
   ```bash
  go run main.go -h
  ```