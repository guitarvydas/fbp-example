This is a quick implementation of Paul Morrison's Collate FBP example (page 91 of "Flow-Based Programming, 2nd Edition").

This example makes every func, except for main, into an FBP component.  Each component accepts parameters that consist of the input and output ports for the component.  Main "wires up" the components according to the diagram on page 91.

Each FBP component resides in its own sub-directory.


This is one of my earliest attempts at Go programming.  I could use advice on how to set up the environment, I use a Makefile at and emacs present.
