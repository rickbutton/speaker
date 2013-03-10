# speaker

## Introduction

`speaker` is an application that turns a bunch of computers/embedded systems/toasters into a mesh of synchronized speaker outputs. This is the *more advanced* version of pressing PLAY on a large number of MP3 players all at the same time, hoping that they don't get out of sync, or god forbid, you didn't press them all at the same time.

## Installation 

You probably just want to know how to install it.

````
git clone http://github.com/rickbutton/speaker
cd speaker
mvn
````

## Usage

### Server

````
java -jar speaker.jar --server --input 1
````

### Client

````
java -jar speaker.jar --client
````

Go to [the options][options] to learn more about the possible command line flags.

## Challenges

The real challenge that `speaker` tries to conquer is to get near perfect audio synchronization (the human ear can detect differences less than a few dozen milliseconds) over an unreliable network. Because of network latency, it is pretty much impossible to merely play two streams of music without delay. Even on a local network, or even the same machine, there is a large non-negligble amount of latency between the sending of the audio data and the decoding step.

[options]: https://github.com/rickbutton/speaker/wiki/Options
