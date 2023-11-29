# naming

... is ~~hard~~ easy. 

_naming_ is a Command-Line Interface (CLI) tool designed to help you find meaningful names for various symbols within your code. This enhances code readability and contributes to the happiness of your colleagues, including your future self.

```sh
naming --lang Golang \      
       --symbol variable \ 
       --type integer \
       --unit secongs \
       --description 'time spent to come up with a good variable name' \
       --num 5

1. variableNamingTimeInHours
2. goodVariableNameTimeHours
3. namingTimeHours
4. variableNamingDurationHours
5. timeSpentOnNamingInHours
```

## Installation

### Build from source

```sh
go install github.com/fabenan-f/naming
```
You need to export your OpenAI api key.

```sh
export OPENAI_API_KEY=<YOUR_KEY>
```

## Usage

### CLI flags

flag                          |purpose                                              |default    |required
------------------------------|-----------------------------------------------------|-----------|-----------
--lang value, -l value        |programming language                                 |           |yes
--mutable, -m                 |mutable                                              |false      |no
--symbol value, -s value      |identifier for variable, function, class etc. names  |           |yes
--type value, -t value        |data type                                            |           |no
--unit value, -u value        |unit                                                 |           |no
--description value, -d value |description                                          |           |yes
--num value, -n value         |number of suggestions                                |0          |yes
--help, -h                    |show help                                            |           |no
