"use strict"

function myButtonClicked() {
  // This function reverses the characters of a word
  // Input
  let word = document.getElementById("word").value

  // Process
  let reversedWord = reverseWord(word)

  // Output
  document.getElementById("answer").innerHTML = reversedWord
}

function reverseWord(word) {
  // Convert the word to an array of characters
  let characters = word.split("")

  // Reverse the array of characters
  let reversed = characters.reverse()

  // Convert the array back to a string
  let reversedWord = reversed.join("")

  return reversedWord
}