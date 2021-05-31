package strain

type Ints []int
type Lists [][]int
type Strings []string


func (i Ints) Keep(f func(int) bool) Ints{

	var ans Ints

	 for _, v := range i {
		 if f(v){
			ans = append(ans,v)
		 }
	 }
	 return ans
}
func (i Ints) Discard(f func(int) bool) Ints{

	var ans Ints

	 for _, v := range i {
		 if !f(v){
			ans = append(ans,v)
		 }
	 }
	 return ans
}
func (l Lists) Keep(f func([]int) bool) Lists{
	
	var ans Lists

	for _, v := range l {
		if f(v){
		   ans = append(ans,v)
		}
	}
	return ans

}
func (s Strings) Keep(f func(string) bool) Strings{

	var ans Strings

	for _, v := range s {
		if f(v){
		   ans = append(ans,v)
		}
	}
	return ans
}
