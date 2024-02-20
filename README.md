# Go Gen(erator)

## Motivation

Allow a way to yeild in between running a function,

and the ability to resume from the same place.

// how it should be;

## Design

### How to make it work

// - use of channels
// - use of a goroutine

A simple interface/struct to allow step (to next yeild) or run (continue working) methods. These methods would drive an actual implementation of the generator.

### Two methods
// - an exec to run to yeild

// - a run that returns a chennel and calls the generator over a over
//		to allow range over the channel, but the channel should not be buffered

### Generator code

Code has to be written for each type of generator.

Some samples and test-main functions are provided.

## Current issues

Synchronization with generoator goroutine is not working properly.

Might need something else make it work, like ready signal from goroutine.