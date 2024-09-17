#!/bin/bash

triangles() {

echo "I will help you classify the right triangle by side. Please enter the length of all three sides."

echo -n "Enter the first number: "
read A
echo -n "Enter the second number: "
read B
echo -n "Enter the third number: "
read C

if (( A <= 0 || B <= 0 || C <= 0 ));
then
    echo "This is not a triangle. Please try again :("

elif (( A == B && B == C ));
then
    echo "This is an equilateral triangle"

elif (( A == B || B == C || C == A ));
then
    echo "This is an isosceles triangle"

else
    echo "This is a scalene triangle"
fi

elif [[ $input == "No" || $input == "no" ]];
then
echo "All right, goodbye for now :)"

else
echo "Please answer with yes or no"
read input

fi

}

triangles
