/*
 * @author Brayden Thistle
 * @version 1.0.0
 * @date 2025-10-07
 * @fileoverview this program outputs the volume of a sphere
 */

package main

import "fmt"

func main () {
	//this is the question
	fmt.Println("What is the volume of a sphere with a radius of 4cm?")

	//this is answering my question with an equation
	fmt.Println("The volume of a sphere with the radius of 4cm is " + fmt.Sprint(4 / 3 * 3.14 * 4 * 4 * 4) + "cm³")

	fmt.Println("\nDone")
}