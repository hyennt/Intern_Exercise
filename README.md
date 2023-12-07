# Offers  

### Name : Nguyen Thuy Hoang Yen

## Project's Description

**Language** : In this project I will use **Golang** to implement

**My solution** : I use Quick Sort for better implement with large JSON data and Benchmarking in Golang for evaluating executing performance

**Sample data** : I already generate sample data using Python

### Pre-condition
Plese make sure your personal laptop has Golang version !

Link for [Download Golang](https://go.dev/doc/install)


**For Window** : `go1.19.5`

**For MacOs** : `go1.19.5 darwin/arm64`

## How to install and run the project
**Note: Make sure you are in ROOT PATH**

1. Using this command-line for creating Go-binary file

```go build .```

2. To run the project

```./Exercise <CHECKIN-DATE> <PATH-TO-YOUR-JSON>```

For example of formatted params following the requirements

 <$**CHECKIN-DATE**> : ```YYYY-MM-DD```

 <$**PATH-TO-YOUR-JSON**> : ```input.json```

E.g : 
```./Exercise 2019-12-25 storages/input/input3.json```

## Output Result
- The result will be in _storages/output_ 

## Output Time Performance 
 I am using Benchmark for comparing performance, and here is my result:

| JSON File | # Input 1    | # Input 2    | # Input 3   | # Input 4   |# Input 5   |
| :---:   | :---: | :---: |:---: |:---: |:---: |
| Seconds | 2.245s   | 1.901s   |2.280s   |2.056s   |1.734s   |


Note : File _input5.json_ which 1000 offers are performed better and faster   

Using this command-line for test performance using Benchmark