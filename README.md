# Welcome to Templater

<img src="http://www.pvhc.net/img88/zxoslhbbhxtzqxydlosb.png" align="left" width="20%" />
Sometimes you don't have neither time nor resources to install a whole
ecosystem just to do something trivial. Sometimes all that is needed is
just a simple tool to get the job done.

With the rise of 12factor popularity one has to have a way of getting
environment variables into the text configuration files somehow. And
that's where Templater steps in to fill the gap for simple tooling.

All you need is ~~love~~ just a single binary with no dependencies and
some knowledge of pretty widespread Go templating language.

## Example

Let's pretend we're studying a new programming language and do a hello
world thing:

    #!/bin/bash

    NAME=John \
    SURNAME=Connor \
    ./templater -t - -o /tmp/hello.txt <<EOF
    Hello, {{env "NAME"}} {{env "SURNAME"}}!
    EOF

... And that's it. We set the vars, feed stdin with the template and
boom, we have /tmp/hello.txt ready in no time.

## Available Functions

`env _env_var_name_`

Takes the environment variable name and returns its value.

`parseInt _string_`

Parses a string into a 32 bit integer value.

`seq _until_`

Generates an array filled with integer values of the counter, i.e.
`[0, 1, 2, 3, ..., until - 1]`. Very helpful when you need to do a for
loop with a counter.

`include _filename_`

Pastes file contents in place.

## Binaries

You can download binaries for Windows, Linux and Mac on the Releases
page.
