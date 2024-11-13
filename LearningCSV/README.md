# Working with CSV FILES in GO 🦦
![1_AoO5hLnLRZvU2zSqyko-Ow](https://github.com/user-attachments/assets/181126da-1181-4c17-8acf-08d27171e4b3)

*Since I became interested in learning Go, the possibility of **creating CLIs** has really caught my attention, so I set out to create the most basic CLI in history—a **TO-DO list using Cobra.** If you think about this problem, you'll realize that **you need somewhere to store that information**, and the best option in this case is a **small CSV file** that can be downloaded with the terminal tool. So, I started exploring how to use CSV with Go, and here are my experiments.*
### Creating CSV File

So the first thing I did was look for a way to create a CSV file using Go:

```go
func main() {
	// Creating files using go and checking err
	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	//creating a writer for the file
	writer := csv.NewWriter(file)
	defer writer.Flush()
	//creating data to put in
	headers := []string{"Name", "Surname", "age"}
	data := [][]string{
		{"Jhon", "Doe", "14"},
		{"Jhon", "Doe2", "18"},
	}

	writer.Write(headers)

	for _, row := range data {
		writer.Write(row)
	}
}
```
The things I found hardest to understand were Go’s wonderfully unique way of handling errors. Often—if not almost always—Go returns errors that, if they are `nil` (empty), mean everything is fine. The issue comes when the error is not empty. For example, in `file, err := os.Create("data.csv")`, you call `os.Create`, storing the result in one variable and the error it returns in another, which you then check to see if everything went well.
### Reading CSV Files

```go
        // Reading csv files and checking err
		
	file2, err := os.Open("data2.csv")

	if err != nil {
		panic(err)
	}

	defer file2.Close()

	reader := csv.NewReader(file2)
	data2, _ := reader.Read()

	for _, row := range data2 {
		for _, column := range row {
			fmt.Printf("%s", column)
		}
	}
```
In Go, the underscore (`_`) is used to ignore return values that aren’t needed. When calling a function that returns multiple values, if not all values are required, the unused values can be assigned to `_`, which acts as a “placeholder” or “ignored value” marker.

In this example:

```go
data2, _ := reader.Read()

```

the `reader.Read()` function returns two values:

1. The first value is the data read (in this case, `data2`), representing a line or record read from the CSV file.
2. The second value is an error (`error`) in case there’s an issue while reading.

By using `_`, the developer is telling Go to ignore the error value, so the code won’t save or use it. This is helpful when it’s known that handling the error isn’t necessary, though it’s generally recommended to handle errors, especially with I/O operations (like reading files), as they are prone to issues.

In this case, they could handle the error like this:



```go
data2, err := reader.Read()
if err != nil {
    log.Fatal(err) // or any other form of error handling
}

```



This way, the error is captured and can be managed if something goes wrong.

## Reading and Creating files using GoCSV
We can create and read files using gocsv in a very similar way, but with slightly fewer lines of code. Here's how it works:

gocsv is a Go library that simplifies working with CSV files. Instead of manually managing the reading and writing process, it allows you to directly map Go structs to CSV data, making it easier to work with.
***<u>CREATING</u>***

```go
//CREATING FILES USING GOCSV
	file3, err := os.Create("data3.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	people := []*Person{
		{"Alice", 25, "Female"},
		{"Bob", 30, "Male"},
		{"Charlie", 35, "Male"},
	}

	if err := gocsv.MarshalFile(&people, file3); err != nil {
		panic(err)
	}
```

***<u>READING</u>***

```go
//READING FILES USING GOCSV
	data4, err := os.Open("data3.csv")
	if err != nil {
		panic(err)
	}

	defer data4.Close()

	var personHolder []Person
	if err := gocsv.UnmarshalFile(data4, &personHolder); err != nil {
		panic(err)
	}

	for _, person := range personHolder {
		fmt.Printf("Name: %s, Gender: %s, Age: %d", person.Name, person.Gender, person.Age)
	}
```

***<u>APPENDING</u>***

```go
//APPENDING DATA
	data5, err := os.OpenFile("data3.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer data5.Close()

	writer4 := csv.NewWriter(data5)
	defer writer4.Flush()

	person := []string{"Samuel", "7", "Male"}

	err = writer4.Write(person)
	if err != nil {
		panic(err)
	}
```

## Conclusion 👨🏻‍💻
Go is a very powerful and easy language. It just has different ways of doing things, but once you get used to it, it really has a lot of potential. I will be using gocsv for my terminal tool.

Shoutout to this amazing article: [https://earthly.dev/blog/golang-csv-files/](https://earthly.dev/blog/golang-csv-files/), which explains everything in much more detail and has been incredibly helpful to me.
