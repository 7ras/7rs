package crafting

// Reverse reverses the input string.
func Grid(Items1 string, Items2 string,Items3 string, Items4 string) string {

	TOP := Items1 +","+ Items2
	BOT := Items1 +","+ Items2

	if TOP == "," {
		TOP = ""
	} else if BOT == "," {
		BOT = ""
	}

	Mixer:= TOP + BOT 

	if Mixer == "Cobblestone,Cobblestone,Cobblestone,Cobblestone" {
		Mixer = "bigblock"
		}

	return Mixer


}

func Hopper(Items1 string, filter string)cut string, but string {
	cut := ""
	but := ""
	for_,char := range Items1 {
		if char == filter {
		  cut = char+ ","
		}
		if char != filter {
		  but = char+ ","
		}
		
	}
	
	return cut , but
}


func Bucket(Items string)string {
	if Items == "Water"{
		return "Water Bucket"
	}

	return "err"
}
