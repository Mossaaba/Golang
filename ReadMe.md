



Go is widely used for cloud services, web servers, DevOps tools, and distributed systems.


- **Simple Syntax:** Easy to learn and read.
- **Fast Compilation:** Quick build times and efficient execution, not like intermdiate java and python 
- **Concurrency:** Built-in support for concurrent programming with goroutines and channels.

- **Cross-Platform:** Compiles to multiple operating systems and architectures.
- **Strong Standard Library:** Comprehensive libraries for networking, web servers, and more.




## Key Features

Go, also known as Golang, is an open-source programming language developed by Google. 

It is designed for simplicity, efficiency, and reliability, making it ideal for building scalable and high-performance applications.

it's released by 2012 
# Introduction to Go (Golang) : 


# **Backend**
# **Systeme application**
# **Cloud Application**
# **Data processing** 
# **Automoation**  


## CRACTERSTIQUE = 

- SIPLICITY : 
- VERSATILITY (FUCTIONEL & OOP): 
- Garbage collection : manage memeory automaticly 
- Fast compilation : 
- Concurency
- Corss platforme 


🛠 Tooling
	•	gofmt → Formats code automatically.
	•	go build → Compiles code.
	•	go run → Runs code directly.
	•	go test → Unit testing.
	•	go get → Installs dependencies.


Ecosystem & Use Cases
	•	Web Development: Frameworks like Gin, Echo, Fiber.
	•	Cloud & DevOps: Kubernetes, Docker, Terraform all written in Go.
	•	Networking & APIs: Fast, concurrent servers.
	•	Microservices: Lightweight, scalable, great for distributed systems.

⸻

✅ Advantages
	•	Easy to learn, concise syntax.
	•	High performance (near C-level).
	•	Excellent concurrency model.
	•	Strong standard library.
	•	Cross-platform builds.

⚠️ Limitations
	•	No generics (introduced only in Go 1.18, still evolving).
	•	Limited GUI development.
	•	Error handling can feel verbose.
	•	Smaller ecosystem compared to Java/Python.

⸻

🔮 Future of Go
	•	Generics support (since Go 1.18) is making it more flexible.
	•	Increasing adoption in cloud-native systems, microservices, and AI infrastructure.
	•	Growing ecosystem of libraries and frameworks.

⸻

👉 In short:
Go = Simple + Fast + Concurrent + Scalable.
It’s one of the best languages for server-side development, cloud computing, and large-scale distributed systems.




## go.dev : DOCUMENTATION 

## Run 

go run main.go 
go run .

## Hierarchy : 
Workspace --> modules --> package --> Go files 

## Modue 
go mod init firstProgram

## Code rule 
- all code have to be in curly bracecs  {}, position on 1 curly braces is 
- Semi colon is optional , only if we need have many statmenet in the same line 
- case sensitive 
- Strongly typed 


## Comment 
// this is a commente 
/* This is multi line comment */ 


## Public and private function 

- Function with Capitale lettere are public 
- Function with Small letter are are private 


## Sharing package : 

we can share file / function from package to package unser the same modules 


## 1. Basic Syntax
	•	Variables: var x int = 5
	•	Short declaration: y := 10
	•	Constants: const Pi = 3.14

## 2. Data Types
	•	Basic: int, float64, string, bool
	•	Composite: array, slice (dynamic array), map (hash table), struct (custom types)

## 3. Control Structures
	•	if, for (only loop), switch, select (for channels)

## 4. Concurrency
	•	Goroutines: Lightweight threads (go functionName()).
	•	Channels: Safe communication between goroutines.

    ch := make(chan int)
    go func() { ch <- 42 }()
    fmt.Println(<-ch) // Receives 42

## 5. Error Handling
	•	No exceptions – errors are values (error type).

result, err := doSomething()
if err != nil {
    fmt.Println("Error:", err)
}
## 6. Error Handling
6. Packages & Modules
	•	Code is organized in packages.
	•	Dependency management via Go Modules (go mod).

## Slices 

