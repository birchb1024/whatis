# what-is
Make Go smarter by providing a classification module which recognises all the things.

# Preamble

This repository is a component of the Genyris project. The [Genyris Langauge](https://github.com/birchb1024/genyris) shows the concept that programming languages can classify objects and act on them the same way people do. Rather than the OO approach which is to build objects with behaviour built-in.

# Class-ification

  Classification - _the process in which ideas and objects are recognized, differentiated, and understood._ link:[Wikipedia]

* Our programming languages should be at least as good as us at
classifying basic data.
* Our operating systems should be at least as good as us at classifying
files of data

# Recognition

Our brains are amazing, they can recognise data at a glance. Here's a comparison of what we see vs what our computer sees:

|Data| How _we_ see it | How our programming languages see it|
|--|--|--|
|'42'|An integer. Also the meaning of life, the universe and everything?|string|
|[3, 4, 5, 6, 7]|A monotonically increasing integer sequence|[]int|
|[["Id", "Name"],[1, "Abe"],[2, "Alice"],[3, "Amy"]],|A Table of peoples' names|[][]interface{}|
|"http://www.example.com/"| A URL | string |
|"etoin.shrdlu@gmail.com" | An email address | string|
|"The computer is encredibly fast accurate and stupid, Stuart G. Walesh"| A quotation of Stuart G. Walesh | string|
|La computadora es increíblemente rápida, precisa y estúpida.|Spanish Text|string|

Essentially our computers do not see or recognise much. It falls to programmer to code parsers and 'input validators' which can recognise the information they want their programs to operate on.

However this work is very repetitive. Especially repetitive now that computer programs run in swarms of communicating processes, dispersed around the world. Unfortunately most computer langauges and operating systems were devised when there was only one machine, aka 'The Computer', likewise there was 'The Programmer' singular. 

Today we need to be able to write programs that communicate, above all, and we programmers almost never write the whole system.

A small step forward is to provide existing languages with the ability to recognise data with the same facility that humans have. This module (I hope) will be able to recognise most commonly-used data structures and provide interfaces for us to leverage in our programs.

# Taxonomy

Artistotle 

# API

```Go
package whatis

//
// Classify - Give an object x return the most likely classification 
func whatis.WhatIs(x interface{}) interface{}

func whatis.WhatIs(c genyris.Context, x interface{})

func whatis.WhatIs(cl []genyris.Context, x interface{})
```





