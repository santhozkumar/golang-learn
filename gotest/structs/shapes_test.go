package shape

import "testing"




// func TestPerimeter(t *testing.T) {
//     rectangle := Rectangle{5.0, 5.0}
//     got := Permiter(rectangle)
//     want := 20.0
//
//     if got != want {
//         t.Errorf("expected %.2f but got %.2f",want, got)
//
//     }
//
// }



func TestArea(t *testing.T) {

    // checkArea := func(t testing.TB, shape Shape, want float64){
    //     t.Helper()
    //     got := shape.Area()
    //     if got != want {
    //         t.Errorf("expected %g but got %g",want, got)
    //         }
    //     }
    // 
    // t.Run("Test Area for Rectangle", func(t *testing.T) {
    //     rectangle := Rectangle{5.0, 5.0}
    //     want := 25.0
    //     checkArea(t, rectangle, want)
    //     })
    //
    //
    // t.Run("Test Area for Circle", func(t *testing.T) {
    //     circle := Circle{10.0}
    //     want := 314.1592653589793
    //     checkArea(t, circle, want)
    //     })


        // Table Driven Test

    areaTests := []struct{
        name string
        shape Shape
        want float64
    }{
        {"Rectangle", Rectangle{5.0, 5.0}, 25.0},
        {"Circle", Circle{10.0}, 314.1592653589793},
        {"Triangle", Triangle{10.0, 2.0}, 10},
    }


    for _, tt := range(areaTests) {
        t.Run(tt.name, func(t *testing.T){

            got := tt.shape.Area()
            want := tt.want

            if got != want {
                t.Errorf("for %v expected %g but got %g",tt.shape, want, got)
                }

        })
    }
}


