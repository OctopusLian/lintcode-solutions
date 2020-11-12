/**
 * @param A: An integer array
 * @return: An integer
 */
func singleNumber (A []int) int {
	Amap := make(map[int]int)
   for _,a := range A {
       Amap[a]++
   }
   for k,v := range Amap {
       if v == 1 {
           return k
       }
   }
   
   return -1
}