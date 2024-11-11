package main

import (
	"fmt"
	"math"
	"net/http"
)

const (
	width, height = 600, 320  // размер канвы в писелях
	cells = 100               // количество ячеек сетки
	xyrange = 30.0            // диапазон осей (-xyrange..+xyrange)
	xyscale = width/2/xyrange // пикселей в единице х или у
	zscale = height * 0.4     // пикселей в единице z
	angle = math.Pi / 6       // углы осей х, у (=30 градусов)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°) и cos(30°)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Запуск сервера на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml") // Устанавливаем заголовок ответа
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			// Проверка, что координаты конечны
			if isFinite(ax, ay) && isFinite(bx, by) && isFinite(cx, cy) && isFinite(dx, dy) {
				avgZ := (az + bz + cz + dz) / 4
				color := getColor(avgZ)
				// Выводим многоугольник
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Fprintf(w, "</svg>")
}

// Вычисление угла ячейки
func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

// Функция поверхности
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // расстояние от (0,0)
	return math.Sin(r) / r // функция высоты
}

// Проверка конечности координат
func isFinite(x, y float64) bool {
	return !(math.IsNaN(x) || math.IsNaN(y) || math.IsInf(x, 0) || math.IsInf(y, 0))
}

// Получение цвета на основе высоты
func getColor(z float64) string {
	minZ, maxZ := -0.5, 0.5
	normalizedZ := (z - minZ) / (maxZ - minZ)
	red := int(255 * normalizedZ)
	blue := int(255 * (1 - normalizedZ))
	return fmt.Sprintf("#%02x00%02x", red, blue) // формирование цвета в формате HEX
}
