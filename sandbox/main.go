package main

import (
	"fmt"
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/ctx"
	"github.com/TopoSimplify/homotopy"
)

func loads(wkt string) []geom.Point {
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

func main() {
	//homo_eg0_deformation()
	//homo_eg1_undeform()
	//homo_eg2_undeform()
	//homo_eg3_undeform()
	//paper3_homo()
	chp1_litreview_homo()
}

func chp1_litreview_homo() {
	var wkt = "Linestring(-6 2, -6 4, -4 4, -4 -1, 4 -1, 4 4, 1 4, 1 0, -2 0, -2 5, 6 5, 6 2)"
	var cwkt = []string {
		"POINT ( -1.07065 1.1577 )",
		"POINT ( 1.97068  3.1577 )",
	}

	var contexts = ctx.NewContexts()
	for _, o := range cwkt {
		contexts.Push(ctx.New(geom.ReadGeometry(o), 0, -1))
	}
	var coords = loads(wkt)
	var p = geom.NewPolygon(coords)
	fmt.Println(p.WKT())
	var bln = homotopy.Homotopy(coords, contexts)
	fmt.Println(bln)
}

func paper3_homo() {
	var wkt = "LINESTRING ( 890.9491141595076 651.5625534172342, 879.8932299895838 663.6707515220128, 885.1153871664693 681.3721679747404, 896.1398134707376 691.2561363854637, 921.9406424000867 689.8821277442557, 928.3526827257239 682.2487464042116, 931.4060352617415 671.8673477817516, 917.0552783424586 661.1806139056898, 903.46785955718 662.249287293296, 900 675, 901.6358480355694 684.0807579258222, 922.2459776536886 698.4315148451052, 943.75 687.5, 943.75 662.5 )"
	var cwkt = []string{
		"POLYGON (( 912.6279171652329 683.9280902990213, 909.7272322560161 671.8673477817516, 920.8719690124806 674.004694556964, 912.6279171652329 683.9280902990213 ))",
	}

	var contexts = ctx.NewContexts()
	for _, o := range cwkt {
		contexts.Push(ctx.New(geom.ReadGeometry(o), 0, -1))
	}
	var coords = loads(wkt)
	var p = geom.NewPolygon(coords)
	fmt.Println(p.WKT())
	var bln = homotopy.Homotopy(coords, contexts)
	fmt.Println(bln)
}

func homo_eg3_undeform() {
	var wkt = "LINESTRING ( 862.5 1043.75, 856.25 1043.75, 850 1043.75, 843.75 1037.6501699201883, 843.75 1031.380631965172, 850 1025, 856.25 1025, 868.75 1025, 875 1025, 875 1031.25, 868.75 1031.25, 856.25 1031.25, 850 1031.25, 850 1037.5, 856.25 1037.5, 868.75 1037.5, 875 1037.5, 881.25 1031.25, 881.25 1025, 875 1018.75, 868.75 1018.75, 862.5 1018.75 )"
	var cwkt = []string{
		"POLYGON (( 858.3498210290013 1033.454690559455, 858.3498210290013 1035.8652014517747, 860.8651367427265 1035.8652014517747, 860.8651367427265 1033.454690559455, 858.3498210290013 1033.454690559455 ))",
		"POLYGON (( 864.4285006705039 1026.6947795788185, 864.4285006705039 1029.419704935354, 867.2582308484448 1029.419704935354, 867.2582308484448 1026.6947795788185, 864.4285006705039 1026.6947795788185 ))",
	}

	var contexts = ctx.NewContexts()
	for _, o := range cwkt {
		contexts.Push(ctx.New(geom.ReadGeometry(o), 0, -1))
	}
	var coords = loads(wkt)
	var p = geom.NewPolygon(coords)
	fmt.Println(p.WKT())
	var bln = homotopy.Homotopy(coords, contexts)
	fmt.Println(bln)
}
func homo_eg2_undeform() {
	var wkt = "LINESTRING ( 750 943.75, 750 956.25, 756.25 968.75, 776.1464297318656 970.9213095551897, 787.5 968.75, 789.9469771861255 962.4060781472422, 787.5 956.25, 776.7336870703448 953.597218070055, 768.75 956.25, 766.8971266508191 963.2869641549609, 768.75 975, 793.75 975, 806.25 962.5, 806.25 943.75 )"
	var cwkt = []string{
		"POLYGON (( 774.0910290471886 961.3783778049036, 774.0910290471886 964.7551075011587, 778.642273420402 964.7551075011587, 778.642273420402 961.3783778049036, 774.0910290471886 961.3783778049036 ))",
	}

	var contexts = ctx.NewContexts()
	for _, o := range cwkt {
		contexts.Push(ctx.New(geom.ReadGeometry(o), 0, -1))
	}
	var coords = loads(wkt)
	var p = geom.NewPolygon(coords)
	fmt.Println(p.WKT())
	var bln = homotopy.Homotopy(coords, contexts)
	fmt.Println(bln)
}

func homo_eg1_undeform() {
	var wkt = "LINESTRING ( 743.75 1024.7311825790548, 743.75 1012.9059898726374, 756.25 1006.25, 806.25 1006.25, 806.25 1031.25, 793.75 1037.5, 793.75 1018.75, 781.25 1012.5, 781.25 1030.7150150329046, 768.75 1037.5, 768.75 1012.5, 756.25 1012.5, 756.25 1037.1262640906007, 762.6581685594776 1043.75, 806.25 1043.75, 818.75 1037.5, 818.75 1025 )"
	var cwkt = []string{
		"POLYGON (( 761.4378110757146 1016.7278721758132, 761.4378110757146 1020.9460719707306, 766.2809293587679 1020.9460719707306, 766.2809293587679 1016.7278721758132, 761.4378110757146 1016.7278721758132 ))",
		"POLYGON (( 770.3362524392878 1029.6032437932765, 770.3362524392878 1031.1925202121308, 775 1032.8392764188543, 775 1029.121459331761, 770.3362524392878 1029.6032437932765 ))",
	}

	var contexts = ctx.NewContexts()
	for _, o := range cwkt {
		contexts.Push(ctx.New(geom.ReadGeometry(o), 0, -1))
	}
	var coords = loads(wkt)
	var p = geom.NewPolygon(coords)
	fmt.Println(p.WKT())
	var bln = homotopy.Homotopy(coords, contexts)
	fmt.Println(bln)
}

func homo_eg0_deformation() {
	var wkt = "LINESTRING ( 743.75 1024.7311825790548, 743.75 1012.9059898726374, 756.25 1006.25, 806.25 1006.25, 806.25 1031.25, 793.75 1037.5, 793.75 1018.75, 781.25 1012.5, 781.25 1030.7150150329046, 768.75 1037.5, 768.75 1012.5, 756.25 1012.5, 756.25 1037.1262640906007, 762.6581685594776 1043.75, 806.25 1043.75, 818.75 1037.5, 818.75 1025 )"
	var cwkt = []string{
		"POLYGON (( 802.8682368091361 1022.986383892712, 802.8682368091361 1028.4883836252131, 808.9815698452484 1028.4883836252131, 808.9815698452484 1022.986383892712, 802.8682368091361 1022.986383892712 ))",
		"POLYGON (( 760.1433980254835 1016.7278721758132, 760.1433980254835 1020.9460719707306, 764.9865163085368 1020.9460719707306, 764.9865163085368 1016.7278721758132, 760.1433980254835 1016.7278721758132 ))",
	}

	var contexts = ctx.NewContexts()
	for _, o := range cwkt {
		contexts.Push(ctx.New(geom.ReadGeometry(o), 0, -1))
	}
	var coords = loads(wkt)
	var p = geom.NewPolygon(coords)
	fmt.Println(p.WKT())
	var bln = homotopy.Homotopy(coords, contexts)
	fmt.Println(bln)
}
