package display

import (
	"fmt"
	"os"
)

func Help(str string, ret int) {
	fmt.Println(str)
	os.Exit(ret)
}

func PrintInterval(s []float64) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//def print(self):
//print("Average return time: %0.0fm %0.02ds" % divmod(round(self.esp * 60), 60))
//print("Standard deviation: %.3f" % self.std_dev)
//print("Time after which 50%% of the ducks are back: %dm %02ds" % divmod(self.time_back(self.a, 50) * 60, 60))
//print("Time after which 99%% of the ducks are back: %dm %02ds" % divmod(self.time_back(self.a, 99) * 60, 60))
//print("Percentage of ducks back after 1 minute: %.1f%%" % (self.percent_back(self.a, 1)))
//print("Percentage of ducks back after 2 minutes: %.1f%%" % self.percent_back(self.a, 2))
